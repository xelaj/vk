package streaming

import (
	"testing"

	"github.com/xelaj/vk"
)

func TestAddingRules(t *testing.T) {
	c, err := vk.LoadClient("/home/r0ck3t/go/src/github.com/xelaj/vk/client.yaml")
	if err != nil {
		panic(err)
	}

	stream, err := New(c)
	if err != nil {
		panic(err)
	}

	//	err = stream.NewRule("кот", "cat tag")
	//	if err != nil {
	//		panic(err)
	//	}

	storage := make(chan string)

	_, err = stream.Subscribe(storage)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case data := <-storage:
			t.Log(data)
		}
	}
}
