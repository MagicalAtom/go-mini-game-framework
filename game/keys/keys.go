package event

import (
	"fmt"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	W = rl.KeyW
	A = rl.KeyA
	S = rl.KeyS
	D = rl.KeyD

	Space  = rl.KeySpace
	Shift  = rl.KeyLeftShift
	Escape = rl.KeyEscape

	Up    = rl.KeyUp
	Down  = rl.KeyDown
	Left  = rl.KeyLeft
	Right = rl.KeyRight
)

var (
	i    *Events
	once sync.Once
)

func GetInstance() *Events {
	once.Do(func() {
		i = &Events{}
		fmt.Println("Hello World")
	})
	return i
}

type Event struct {
	Name string
	Key  int32
}

type Events struct {
	events []Event
}

func AddKey(Name string, key int32) Event {
	return Event{
		Name: Name,
		Key:  key,
	}
}

func AddEvent(event Event) {
	ins := GetInstance()
	ins.events = append(i.events, event)
}

func Listen() string {
	for _, v := range i.events {
		if rl.IsKeyDown(v.Key) {
			return v.Name
		}
	}
	return ""
}

func GetEventList() {
	fmt.Println(i.events)
}
