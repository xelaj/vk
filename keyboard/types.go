package keyboard

import (
	"encoding/json"
)

type Keyboard struct {
	OneTime bool        `json:"one_time"`
	Buttons [][]*Button `json:"buttons"`
}

type Button struct {
	Action Action `json:"action"`
	Color  Color  `json:"color,omitempty"`
}

type Color string

const (
	ColorPrimary   Color = "primary"
	ColorSecondary Color = "secondary"
	ColorNegative  Color = "negative"
	ColorPositive  Color = "positive"
)

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Hash    string `json:"hash,omitempty"`
	Label   string `json:"label,omitempty"`
}

func Text(label, payload string, color Color) *Button {
	b := new(Button)
	b.Color = color
	b.Action = Action{
		Type:    "text",
		Payload: payload,
		Label:   label,
	}
	return b
}

func Geo(payload string) *Button {
	b := new(Button)
	b.Action = Action{
		Type:    "location",
		Payload: payload,
	}
	return b
}

func Pay(payload, hash string) *Button {
	b := new(Button)
	b.Action = Action{
		Type:    "vkpay",
		Payload: payload,
		Hash:    hash,
	}
	return b
}

func New(onetime bool, rows ...[]*Button) *Keyboard {
	if len(rows) > 10 && len(rows) < 1 {
		panic("keyboard can contain rows in count [1:10]")
	}

	k := new(Keyboard)
	k.OneTime = onetime
	k.Buttons = rows
	return k
}

func Col(b ...*Button) []*Button {
	if len(b) > 4 && len(b) < 1 {
		panic("keyboard can contain columns in count [1:4]")
	}
	return b
}

func (k *Keyboard) String() string {
	buf, err := json.Marshal(k)
	if err != nil {
		panic(err) // DELETE ME
	}
	return string(buf)
}
