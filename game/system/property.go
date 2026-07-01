package system

type Property interface {
	GetName() string
	GetModel() string
}

type CommonProperty struct {
	Name      string
	ModelName string
}

func (c *CommonProperty) GetName() string {
	return c.Name
}

func (c *CommonProperty) GetModel() string {
	return c.ModelName
}
