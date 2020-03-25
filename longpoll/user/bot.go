package user

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"github.com/xelaj/vk"
)

type Bot struct {
	client *vk.Client
	User   string
	key    string
	ts     uint
	server string // может отличаться, поэтому не константа
	funcs  map[int]interface{}
}

func New(c *vk.Client, user string) (*Bot, error) {
	res, err := c.By(user).MessagesGetLongPollServer(false, 0)
	if err != nil {
		return nil, errors.Wrap(err, "getting longpoll server")
	}

	u := &Bot{
		client: c,
		User:   user,
		key:    res.Key,
		ts:     uint(res.Ts),
		server: res.Server,
		funcs:  make(map[int]interface{}),
	}
	return u, nil
}

func (u *Bot) Start(ctx context.Context) chan error {
	errchan := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				endpoint := u.botUserEndpoint()
				res, err := http.Get(endpoint)
				if err != nil {
					errchan <- errors.Wrap(err, "requesting")
					continue
				}

				buf, _ := ioutil.ReadAll(res.Body)
				res.Body.Close()
				resp := new(LongPollResponse)
				err = json.Unmarshal(buf, resp)
				if err != nil {
					errchan <- errors.Wrap(err, "json parsing")
					continue
				}

				u.ts = uint(resp.Ts)

				if resp.ErrCode == 1 {
					continue
				}
				if resp.ErrCode == 2 && resp.ErrCode == 3 {
					res, err := u.client.By(u.User).MessagesGetLongPollServer(false, 0)
					if err != nil {
						errchan <- errors.Wrap(err, "updating server key")
					}
					u.key = res.Key
					continue
				}

				for _, update := range resp.Updates {
					f, ok := u.funcs[update.Code()]
					if ok {
						err := callFunc(f, update)
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
	id, ok := userUpdateTypesStrings[actionType]
	if !ok {
		panic("action " + actionType + " is not supported")
	}
	inputType := userUpdateTypes[id]

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

	b.funcs[id] = f
}

func (b *Bot) botUserEndpoint() string {
	return "https://" + b.server + "?act=a_check&key=" + b.key + "&ts=" + strconv.Itoa(int(b.ts)) + "&wait=25&mode=2&version=2 "
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
