package group

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type LongPollResponse struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
	ErrCode int      `json:"failed"`
}

type Update struct {
	Type   string      `json:"type"`
	Object interface{} `json:"object"`
}

// TODO: ЯВНО можно сделать что-то получше двойного парсинга, это оверхед мне кажется
func (u *Update) UnmarshalJSON(d []byte) error {
	preUnmarshaled := make(map[string]interface{})

	err := json.Unmarshal(d, &preUnmarshaled)
	if err != nil {
		return err
	}

	t, ok := UpdateTypes[preUnmarshaled["type"].(string)]
	if !ok {
		panic(fmt.Sprintf("Some unknown type was founded: %v\nData is: %v", preUnmarshaled["type"], string(d)))
	}

	u.Type = preUnmarshaled["type"].(string)
	u.Object = reflect.New(t).Interface()

	return ToObject(preUnmarshaled["object"].(map[string]interface{}), u.Object)
}

func ToObject(in map[string]interface{}, out interface{}) error {
	if reflect.ValueOf(out).Elem().Kind() != reflect.Struct {
		panic("out type is not struct. karamba!")
	}

	config := &mapstructure.DecoderConfig{
		Result:     out,
		ZeroFields: true,
		TagName:    "json",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}
	return decoder.Decode(in)
}
