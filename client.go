package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
	"github.com/xelaj/vk/response"
	"github.com/xelaj/vk/types"
)

type Client struct {
	id               int
	secureKey        string
	serviceKey       string
	clientTokens     map[int]string
	groupTokens      map[int]string
	savedPath        string
	NeedToSaveTokens bool
	lastRequest      time.Time
	tempKey          string         // используется для метода client.By(). после первого запроса сразу же сбрасывается
	deleteTempKey    bool           // нужно ли заблокировать удаление временного ключа сразу же. по умолчанию true
	Counters         map[string]int // https://vk.com/dev/data_limits
}

type File struct {
	Client struct {
		Id      int    `yaml:"id"`
		Seckey  string `yaml:"secure key"`
		Servkey string `yaml:"service key"`
	} `yaml:"client"`
	Tokens struct {
		Clients map[int]string `yaml:"clients"`
		Groups  map[int]string `yaml:"groups"`
	} `yaml:"tokens"`
}

const (
	HOST              = "api.vk.com"
	DEFAULT_VERSION   = "5.103"
	MAX_COUNT_FRIENDS = 5000
	MAX_COUNT_GROUPS  = 1000
	MAX_COUNT_POSTS   = 100
)

func NewClient(id int, seckey, servkey string) (*Client, error) {
	app := &Client{
		id:            id,
		secureKey:     seckey,
		serviceKey:    servkey,
		clientTokens:  make(map[int]string),
		groupTokens:   make(map[int]string),
		lastRequest:   time.Unix(0, 0),
		deleteTempKey: true,
	}

	res, err := app.AppsGet(nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if len(res.Items) == 0 {
		return nil, errors.New("unknown server-side error")
	}

	if id != res.Items[0].Id {
		return nil, errors.New("app is invalid")
	}

	return app, nil
}

func LoadClient(path string) (*Client, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "reading file")
	}

	var f File
	if err = yaml.Unmarshal(buf, &f); err != nil {
		return nil, errors.Wrap(err, "parsing file")
	}

	app := &Client{
		id:           f.Client.Id,
		secureKey:    f.Client.Seckey,
		serviceKey:   f.Client.Servkey,
		clientTokens: f.Tokens.Clients,
		groupTokens:  f.Tokens.Groups,
		savedPath:    path,
		lastRequest:  time.Unix(0, 0),
	}
	res, err := app.AppsGet(nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if len(res.Items) == 0 {
		return nil, errors.New("unknown server-side error")
	}

	if f.Client.Id != res.Items[0].Id {
		return nil, errors.New("app is invalid")
	}
	return app, nil
}

func (c *Client) SaveTo(path string) error {
	if path == "" {
		path = c.savedPath
	}

	var f File
	f.Client.Id = c.id
	f.Client.Seckey = c.secureKey
	f.Client.Servkey = c.serviceKey
	f.Tokens.Clients = c.clientTokens
	f.Tokens.Groups = c.groupTokens

	res, _ := yaml.Marshal(&f)
	return ioutil.WriteFile(path, res, 0666)
}

func (c *Client) ByClient(userId int) types.Client {
	return types.Client(c.By(userId))
}

func (c *Client) DisableTempTokenDeleting() {
	c.deleteTempKey = false
}

func (c *Client) EnableTempTokenDeleting() {
	c.deleteTempKey = false
}

func (c *Client) ForceDeleteTempToken() {
	c.tempKey = ""
}

// согласно документации vk api любые положительные id считаются пользователями,
// а отрицательные — группами. здесь все реализовано идентично, поскольку нет иной
// возможности сделать как-то более удобно и очевидно
func (c *Client) By(userId int) *Client {
	if userId >= 0 {
		token, ok := c.clientTokens[userId]
		if !ok {
			token = ""
			fmt.Printf("WARNING! You don't have %v user token!\n", userId)
		}
		c.tempKey = token
		return c
	} else {
		userId = -userId
		token, ok := c.groupTokens[userId]
		if !ok {
			token = ""
			fmt.Printf("WARNING! You don't have %v group token!\n", userId)
		}
		c.tempKey = token
		return c
	}
}

func (c *Client) RawMethod(method string, params map[string]interface{}, storeTo interface{}) (*response.Basic, error) {
	methodName := strings.Split(method, "#")[0]

	if params == nil {
		params = make(map[string]interface{})
	}

	if v, ok := params["access_token"]; !ok || v == "" {
		if c.tempKey != "" {
			params["access_token"] = c.tempKey
			if c.deleteTempKey {
				c.tempKey = ""
			}
		} else {
			params["access_token"] = c.serviceKey
		}
	}

	c.wait()

	i, e := rawMethod(methodName, params)
	if e != nil {
		return nil, errors.Wrap(e, "request failed")
	}

	data := response.New(reflect.TypeOf(storeTo))

	e = json.NewDecoder(i).Decode(data)
	i.Close()
	if e != nil {
		return nil, errors.Wrap(e, "json parsing error")
	}
	return data, data.Err()
}

// wait дожидается, пока не истечет треть секунды для следующего
// запроса (т.к. вк позволяет отправлять лишь 3 запроса в секунду)
func (c *Client) wait() {
	dur := time.Now().Sub(c.lastRequest).Nanoseconds() / 1000000 // millis
	if dur <= 330 {                                              // 1/3 second
		time.Sleep(time.Duration((330 - dur)) * time.Millisecond)
	}
}

func (c *Client) client() types.Client {
	return types.Client(c)
}

func rawMethod(methodName string, params map[string]interface{}) (io.ReadCloser, error) {
	if _, ok := params["access_token"]; !ok {
		return nil, errors.New("access token is required. sad, but true")
	}

	if _, ok := params["v"]; !ok {
		params["v"] = DEFAULT_VERSION
	}

	u := url.URL{
		Scheme: "https",
		Host:   HOST,
		Path:   filepath.Join("method", methodName),
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, fmt.Sprint(v))
	}
	var (
		resp *http.Response
		err  error
	)

	encoded := q.Encode()
	if len(encoded) > 500 {
		resp, err = http.PostForm(u.String(), q)
	} else {
		u.RawQuery = encoded
		// println(u.String())
		resp, err = http.Get(u.String())
	}

	// TODO: иногда выкидывает net/http: TLS handshake timeout.
	// придумать, как обработать, т.к. стандартного метода похоже нет
	if err != nil {
		return nil, errors.Wrap(err, "request error")
	}

	return resp.Body, nil
}
