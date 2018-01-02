package property

//DynamicUIntProperty ...
type DynamicUIntProperty struct {
	prop         *DynamicProperty
	defaultValue uint
}

//NewDynamicUIntProperty ...
func NewDynamicUIntProperty(propName string, defaultValue uint) *DynamicUIntProperty {
	return &DynamicUIntProperty{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicUIntProperty) Get() uint {
	return p.prop.GetUInt(p.defaultValue)
}

//GetValue ...
func (p *DynamicUIntProperty) GetValue() uint {
	return p.Get()
}

//GetName ...
func (p *DynamicUIntProperty) GetName() string {
	return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicUIntProperty) GetDefaultValue() uint {
	return p.defaultValue
}
