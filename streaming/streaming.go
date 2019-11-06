package streaming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/xelaj/vk"
)

// пакет streaming является оберткой над основной библиотекой
// это значит, что streaming использует vk, но не наоборот.

type StreamingClient struct {
	vk   *vk.Client
	host string
	key  string
	http *http.Client
}

func New(c *vk.Client) (*StreamingClient, error) {
	res := new(StreamingClient)
	res.vk = c

	keyserver, err := c.StreamingGetServerUrl()
	if err != nil {
		return nil, errors.Wrap(err, "getting streaming server error")
	}
	res.host = keyserver.Endpoint
	res.key = keyserver.Key

	res.http = &http.Client{}

	return res, nil
}

func (s *StreamingClient) GetRules() ([]*StreamingRule, error) {
	res, err := s.http.Get(s.ruleUrl())
	if err != nil {
		return nil, errors.Wrap(err, "http requesting")
	}

	type responseType struct {
		Code  int              `json:"code"`
		Rules []*StreamingRule `json:"rules"`
		Err   struct {
			Msg  string `json:"message"`
			Code int    `json:"error_code"`
		} `json:"error"`
	}

	response := new(responseType)

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(response)
	if err != nil {
		return nil, errors.Wrap(err, "response parsing error")
	}

	fmt.Println(response)

	if response.Code == 400 {
		return nil, errors.New(response.Err.Msg)
	}

	return response.Rules, nil
}

func (s *StreamingClient) NewRule(value, tag string) error {
	rule := new(StreamingRule)
	rule.Value = value
	rule.Tag = tag

	data := map[string]*StreamingRule{
		"rule": rule, //  соответствии с документацией: https://vk.com/dev/streaming_api_docs?f=5.%2B%D0%94%D0%BE%D0%B1%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%2B%D0%BF%D1%80%D0%B0%D0%B2%D0%B8%D0%BB
	}

	b, _ := json.Marshal(data) // ошибку опускаем, ее тут априори не может быть

	res, err := s.http.Post(s.ruleUrl(), "application/json", bytes.NewReader(b))
	if err != nil {
		return errors.Wrap(err, "http requesting")
	}

	type responseType struct {
		Code int `json:"code"`
		Err  struct {
			Msg  string `json:"message"`
			Code int    `json:"error_code"`
		} `json:"error"`
	}

	response := new(responseType)

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(response)
	if err != nil {
		return errors.Wrap(err, "response parsing error")
	}

	if response.Code == 400 {
		return errors.New(response.Err.Msg)
	}

	return nil
}

func (s *StreamingClient) DeleteRules() {
	panic("not implemented")
}

func (s *StreamingClient) Subscribe(storeTo chan string) (chan struct{}, error) {
	conn, _, err := websocket.DefaultDialer.Dial("wss://"+s.host+"/stream?key="+s.key, nil)
	if err != nil {
		return nil, errors.Wrap(err, "dial error")
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				done <- struct{}{}
				return
			}
			println("recieved!!! " + string(message))
			storeTo <- string(message)
		}
	}()
	println("goroutine is running")
	return done, nil

}

func (s *StreamingClient) ruleUrl() string {
	// здесь сделано таким образом, потому что streaming api катастрофически маленький
	// так что смысла подтягивать url.Encode смысла нет
	return "https://" + s.host + "/rules?key=" + s.key
}

type StreamingRule struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}
