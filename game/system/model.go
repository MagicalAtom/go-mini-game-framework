package system

import "fmt"

type ModelName struct {
	Name string
}

type Model struct {
	Name       string
	properties map[string]Property
}

func New(Name ModelName) *Model {
	return &Model{
		Name:       Name.Name,
		properties: make(map[string]Property),
	}
}

// if you walk like a duck or if you talk like a duck , you are a duck you don't say this class extends or implements this interface . only things you do is talk like a interface and after that you type like a interface
func (model *Model) AddProperty(property Property) {
	model.properties[property.GetName()] = property
}

func (model *Model) GetProperty(Name string) (Property, error) {
	if model.properties[Name] != nil {
		return model.properties[Name], nil
	}
	return nil, fmt.Errorf("Can't Find Property . First Please set This Properties")
}
