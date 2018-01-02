package property

//DynamicStringProperty ...
type DynamicStringProperty struct {
	prop         *DynamicProperty
	defaultValue string
}

//NewDynamicStringProperty ...
func NewDynamicStringProperty(propName string, defaultValue string) *DynamicStringProperty {
	return &DynamicStringProperty{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicStringProperty) Get() string {
	return p.prop.GetString(p.defaultValue)
}

//GetValue ...
func (p *DynamicStringProperty) GetValue() string {
	return p.Get()
}

//GetName ...
func (p *DynamicStringProperty) GetName() string {
	return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicStringProperty) GetDefaultValue() string {
	return p.defaultValue
}
