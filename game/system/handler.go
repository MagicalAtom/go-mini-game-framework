package system

import "fmt"

type Hook[T Property] struct {
	Model        *Model
	PropertyName string
}
type Handler[T Property] struct {
	model    *Model
	Property T
}

func Make[T Property](hook Hook[T]) (*Handler[T], error) {
	property, err := hook.Model.GetProperty(hook.PropertyName)

	if err != nil {
		fmt.Println(err)
	}

	final := property.(T)

	return &Handler[T]{
		model:    hook.Model,
		Property: final,
	}, err
}
