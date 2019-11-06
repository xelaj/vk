package methods

import (
	"reflect"

	"github.com/xelaj/vk/types"
)

func ExecuteMethod(c types.Client, method string, params map[string]interface{}, resType reflect.Type) (interface{}, error) {
	res, err := c.RawMethod("execute."+method, params, resType)
	if err != nil {
		return nil, err
	}

	return res.Object(), nil
}
