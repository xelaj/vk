package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/ololosha228/keystorage"
	"github.com/pkg/errors"
	"github.com/xelaj/vk/response"
	"github.com/xelaj/vk/types"
)

type Client struct {
	id           int
	accessKey    string
	secureKey    string
	storage      keystorage.Storage
	lastRequest  time.Time
	counters     map[string]int // https://vk.com/dev/data_limits
	parentClient *Client        // при вызове client.By() создается новый инстанс клиента, который возвращает родительский клиент самого приложения
	err          error
}

const (
	HOST              = "api.vk.com"
	DEFAULT_VERSION   = "5.103"
	MAX_COUNT_FRIENDS = 5000
	MAX_COUNT_GROUPS  = 1000
	MAX_COUNT_POSTS   = 100
)

type ClientConfig struct {
	ID         int
	SecureKey  string
	ServiceKey string
	KeyStorage keystorage.Storage
}

func NewClient(config ClientConfig) (*Client, error) {
	if !config.KeyStorage.HasService("vk") {
		return nil, errors.New("storage can't handle vk service")
	}

	app := &Client{
		id:          config.ID,
		accessKey:   config.ServiceKey,
		secureKey:   config.SecureKey,
		storage:     config.KeyStorage,
		lastRequest: time.Unix(0, 0),
	}

	res, err := app.AppsGet(nil)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if len(res.Items) == 0 {
		return nil, errors.New("unknown server-side error")
	}

	if config.ID != res.Items[0].Id {
		return nil, errors.New("app is invalid")
	}
	return app, nil
}

func (c *Client) RootClient() *Client {
	if c.parentClient != nil {
		return c.parentClient.RootClient()
	}
	return c
}

// согласно документации vk api любые положительные id считаются пользователями,
// а отрицательные — группами. здесь все реализовано идентично, поскольку нет иной
// возможности сделать как-то более удобно и очевидно
func (c *Client) By(user string) *Client {
	root := c.RootClient()

	key, err := root.storage.UserKey(user, "vk")
	if err != nil {
		return &Client{
			err: errors.Wrap(err, "finding user"),
		}
	}

	return &Client{
		parentClient: root,
		accessKey:    key,
	}
}

func (c *Client) ConvertToID(user string) (int, error) {
	idStr, err := c.RootClient().storage.ServiceID("vk", user)
	if err != nil {
		return 0, errors.Wrap(err, "getting service user id")
	}
	return strconv.Atoi(idStr)
}

func (c *Client) RawMethod(method string, params map[string]interface{}, storeTo interface{}) (*response.Basic, error) {
	if c.err != nil {
		return nil, c.err
	}

	if params == nil {
		params = make(map[string]interface{})
	}

	if v, ok := params["access_token"]; !ok || v == "" {
		if c.accessKey != "" {
			params["access_token"] = c.accessKey
		}
	}
	c.wait() // чтобы не словить таймаут

	body, err := rawMethod(method, params)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	data := response.New(reflect.TypeOf(storeTo))

	err = json.NewDecoder(body).Decode(data)
	body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "json parsing error")
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

func (c *Client) Client() types.Client {
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
		resp, err = http.Get(u.String())
	}

	// TODO: иногда выкидывает net/http: TLS handshake timeout.
	// придумать, как обработать, т.к. стандартного метода похоже нет
	if err != nil {
		return nil, errors.Wrap(err, "request error")
	}

	return resp.Body, nil
}
