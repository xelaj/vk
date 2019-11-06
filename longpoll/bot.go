package longpoll

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/xelaj/vk"
	"github.com/xelaj/vk/response"
)

// пакет longpoll является оберткой над основной библиотекой
// это значит, что longpoll использует vk, но не наоборот.

type RespFunc func(payload Update) error

type GroupBot struct {
	client  *vk.Client
	GroupId int
	key     string
	ts      uint
	funcs   map[string]RespFunc
}

func NewGroupBot(c *vk.Client, groupId int) (*GroupBot, error) {
	key, ts, err := getServerData(groupId, c)

	if err != nil {
		vkerr := errors.Cause(err).(*response.ResponseError)

		if vkerr.Code == response.ErrAccessDenied {
			return nil, errors.New("user is not admin of group")
		}
		return nil, errors.Wrap(err, "bot setup")
	}

	b := new(GroupBot)
	b.client = c
	b.GroupId = groupId
	b.key = key
	b.ts = uint(ts)
	b.funcs = make(map[string]RespFunc)

	return b, nil
}

func (b *GroupBot) Start(done *bool) error {
	for !*done {
		endpoint := botEndpoint(b.GroupId, b.key, b.ts)
		res, err := http.Get(endpoint)
		if err != nil {
			errors.Wrap(err, "requesting")
		}
		buf, _ := ioutil.ReadAll(res.Body)
		resp := new(LongPollResponse)
		err = json.Unmarshal(buf, resp)
		if err != nil {
			errors.Wrap(err, "json parsing")
		}

		b.ts = func() uint {
			i, _ := strconv.Atoi(resp.Ts)
			return uint(i)
		}()

		if resp.ErrCode == 1 {
			continue
		}
		if resp.ErrCode == 2 && resp.ErrCode == 3 {
			key, ts, err := getServerData(b.GroupId, b.client)
			if err != nil {
				return errors.Wrap(err, "updating server key")
			}
			b.key = key
			b.ts = uint(ts)
			continue
		}

		for _, update := range resp.Updates {
			err := b.runFunc(update)
			if err != nil {
				return errors.Wrap(err, "running user function")
			}
		}
	}
	return nil
}

func (b *GroupBot) On(actionType string, f RespFunc) {
	if !validAction(actionType) {
		panic("action '" + actionType + "' is invalid")
	}
	b.funcs[actionType] = f
}

func (b *GroupBot) runFunc(u Update) error {
	f, ok := b.funcs[u.Type]
	if !ok {
		return nil
	}
	return f(u)
}

//// DEPRECATED: поскольку нет никакой возможности тестировать
////             работу user long poll api, этот тип не поддерживается
//type UserBot struct {
//}

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

func getServerData(gid int, c *vk.Client) (key string, ts int, err error) {
	res, err := c.By(-gid).GroupsGetLongPollServer(gid)
	if err != nil {
		return "", 0, err // т.к. функция — частое действие, обертку ошибки можно опустить
	}

	ts = func() int {
		i, _ := strconv.Atoi(res.Ts)
		return i
	}() // или как еще элегантней?

	return res.Key, ts, nil
}
