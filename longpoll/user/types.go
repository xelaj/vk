package user

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"
)

type UpdateItem interface {
	Code() int
}

type LongPollResponse struct {
	Ts      int          `json:"ts"`
	Updates []UpdateItem `json:"updates"`
	ErrCode int          `json:"failed"`
}

func (r *LongPollResponse) UnmarshalJSON(b []byte) error {
	type raw struct {
		Ts      int             `json:"ts"`
		Updates [][]interface{} `json:"updates"`
		ErrCode int             `json:"failed"`
	}
	res := new(raw)
	err := json.Unmarshal(b, res)
	if err != nil {
		return err
	}
	r.Ts = res.Ts
	r.ErrCode = res.ErrCode
	r.Updates = make([]UpdateItem, 0)
	for _, update := range res.Updates {
		new, err := parseUpdate(update)
		if err != nil {
			return err
		}
		r.Updates = append(r.Updates, new)
	}
	return nil
}

func parseUpdate(in []interface{}) (UpdateItem, error) {
	_type, ok := userUpdateTypes[int(in[0].(float64))]
	if !ok {
		return nil, errors.New("unknown type: " + strconv.Itoa(int(in[0].(float64))))
	}

	if len(in)-1 != _type.NumField() {
		return nil, errors.New("object invalid")
	}

	resType := reflect.New(_type)
	for i := 0; i < resType.Elem().NumField(); i++ {
		switch resType.Elem().Field(i).Type().String() { // Kind() нельзя, потому что есть например time.Time
		case "int", "int8", "int16", "int32", "int64":
			if string(_type.Field(i).Tag) == "invert" {
				resType.Elem().Field(i).SetInt(-int64(in[i+1].(float64)))
			} else {
				resType.Elem().Field(i).SetInt(int64(in[i+1].(float64)))
			}
		case "user.Flags":
			value := Flags(in[i+1].(float64))
			resType.Elem().Field(i).Set(reflect.ValueOf(value))
		case "time.Time":
			value := time.Unix(int64(in[i+1].(float64)), 0)
			resType.Elem().Field(i).Set(reflect.ValueOf(value))
		case "string":
			resType.Elem().Field(i).SetString(in[i+1].(string))
		case "map[string]interface {}":
			resType.Elem().Field(i).Set(reflect.ValueOf(in[i+1]))
		default:
			panic("что? " + resType.Elem().Field(i).Type().String())
		}

	}

	return resType.Interface().(UpdateItem), nil
}
