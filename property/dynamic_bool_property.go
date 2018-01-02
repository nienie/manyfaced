package property

//DynamicBoolProperty ...
type DynamicBoolProperty struct {
	prop         *DynamicProperty
	defaultValue bool
}

//NewDynamicBoolProperty ...
func NewDynamicBoolProperty(propName string, defaultValue bool) *DynamicBoolProperty {
	return &DynamicBoolProperty{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicBoolProperty) Get() bool {
	return p.prop.GetBool(p.defaultValue)
}

//GetValue ...
func (p *DynamicBoolProperty) GetValue() bool {
	return p.Get()
}

//GetName ...
func (p *DynamicBoolProperty) GetName() string {
	return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicBoolProperty) GetDefaultValue() bool {
	return p.defaultValue
}
