package group

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/pkg/errors"

	"github.com/xelaj/vk"
	"github.com/xelaj/vk/response"
)

const (
	BOTS_LONGPOLL_HOST = "lp.vk.com"
	LONGPOLL_TIMEOUT   = 25
)

type Bot struct {
	client *vk.Client
	Group  string
	key    string
	ts     uint
	funcs  map[string]interface{}
}

func New(c *vk.Client, group string) (*Bot, error) {
	b := &Bot{
		client: c.By(group),
		Group:  group,
		funcs:  make(map[string]interface{}),
	}

	key, ts, err := b.getServerData()
	if err != nil {
		vkerr := errors.Cause(err).(*response.ResponseError)

		if vkerr.Code == response.ErrAccessDenied {
			return nil, errors.New("user is not admin of group")
		}
		return nil, errors.Wrap(err, "bot setup")
	}

	b.key = key
	b.ts = uint(ts)

	return b, nil
}

func (b *Bot) Start(ctx context.Context) chan error {
	errchan := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				id, _ := b.client.ConvertToID(b.Group)
				endpoint := botEndpoint(id, b.key, b.ts)

				res, err := http.Get(endpoint)
				if err != nil {
					errchan <- errors.Wrap(err, "requesting")
				}
				buf, _ := ioutil.ReadAll(res.Body)
				resp := new(LongPollResponse)
				err = json.Unmarshal(buf, resp)
				if err != nil {
					errchan <- errors.Wrap(err, "json parsing")
				}

				b.ts = func() uint {
					i, _ := strconv.Atoi(resp.Ts)
					return uint(i)
				}()

				if resp.ErrCode == 1 {
					continue
				}
				if resp.ErrCode == 2 && resp.ErrCode == 3 {
					key, ts, err := b.getServerData()
					if err != nil {
						errchan <- errors.Wrap(err, "updating server key")
					}
					b.key = key
					b.ts = uint(ts)
					continue
				}

				for _, update := range resp.Updates {
					f, ok := b.funcs[update.Type]
					if ok {
						err := callFunc(f, update.Object)
						if err != nil {
							errchan <- errors.Wrap(err, "running user function")
						}
					}

				}
			}
		}
	}()
	return errchan
}

func (b *Bot) On(actionType string, f interface{}) {
	inputType, ok := UpdateTypes[actionType]
	if !ok {
		panic("action " + actionType + " is not supported")
	}

	funcType := reflect.TypeOf(f)
	if funcType.Kind() != reflect.Func {
		panic("not a function")
	}
	if funcType.NumOut() != 1 {
		panic("function output arguments is not single")
	}
	if !funcType.Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic("function output is not error")
	}

	if funcType.NumIn() != 1 {
		panic("function input arguments is not single")
	}

	if funcType.In(0) != reflect.PtrTo(inputType) {
		panic("function output is not *" + inputType.String())
	}

	b.funcs[actionType] = f
}

func (b *Bot) getServerData() (key string, ts int, err error) {
	groupId, err := b.client.ConvertToID(b.Group)
	if err != nil {
		return "", 0, err
	}

	res, err := b.client.GroupsGetLongPollServer(-groupId)
	if err != nil {
		return "", 0, err // т.к. функция — частое действие, обертку ошибки можно опустить
	}

	ts, _ = strconv.Atoi(res.Ts)

	return res.Key, ts, nil
}

func botEndpoint(chatId int, key string, ts uint) string {
	u := url.URL{
		Scheme: "https",
		Host:   BOTS_LONGPOLL_HOST,
		Path:   fmt.Sprintf("wh%v", chatId),
	}

	params := map[string]interface{}{
		"act":  "a_check",
		"key":  key,
		"ts":   ts,
		"wait": LONGPOLL_TIMEOUT,
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, fmt.Sprint(v))
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func callFunc(f, arg interface{}) error {
	value := reflect.ValueOf(f)
	argument := reflect.ValueOf(arg)
	res := value.Call([]reflect.Value{argument})
	if res[0].IsNil() {
		return nil
	}
	return res[0].Interface().(error)
}
