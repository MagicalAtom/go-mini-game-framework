package trigger

import "sync"

type CreateTrigger struct {
	EventName string
	Handler   func()
}

type Trigger struct {
	event   string
	handler func()
}

type TriggerStore struct {
	Store []Trigger
}

var (
	T    *TriggerStore
	once sync.Once
)

func GetInstance() *TriggerStore {
	once.Do(func() {
		T = &TriggerStore{
			Store: []Trigger{},
		}
	})
	return T
}
func (t *TriggerStore) AddTrigger(createTrigger CreateTrigger) {
	GetInstance().Store = append(GetInstance().Store, Trigger{
		event:   createTrigger.EventName,
		handler: createTrigger.Handler,
	})
}
func (t *TriggerStore) Listen(event string) {
	for _, v := range GetInstance().Store {
		if v.event == event {
			v.handler()
		}
	}
}
