package response

import "reflect"

type Basic struct {
	Response interface{}    `json:"response"`
	Error    *ResponseError `json:"error"`
}

func New(object reflect.Type) *Basic {
	r := new(Basic)
	r.Response = reflect.New(object).Interface()
	return r
}

func (r *Basic) Err() error {
	if r.Error == nil { // это потому что если просто возвращать ошибку, то случится паника, из-за того что невозможно напечатать текст ПАНИКИ
		return nil //      МЫ УСТАНОВИЛИ ТЕБЕ ПАНИКУ В ПАНИКЕ, ЧТО БЫ МОЖНО БЫЛО ЗАПАНИКОВАТЬ ПОКА ПАНИКУЕШЬ
	}
	return r.Error
}

func (r *Basic) Object() interface{} {
	if r.Response == nil {
		return nil
	}
	return r.Response
}
