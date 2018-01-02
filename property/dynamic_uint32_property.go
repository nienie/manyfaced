package property

//DynamicUInt32Property ...
type DynamicUInt32Property struct {
	prop         *DynamicProperty
	defaultValue uint32
}

//NewDynamicUInt32Property ...
func NewDynamicUInt32Property(propName string, defaultValue uint32) *DynamicUInt32Property {
	return &DynamicUInt32Property{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicUInt32Property) Get() uint32 {
	return p.prop.GetUInt32(p.defaultValue)
}

//GetValue ...
func (p *DynamicUInt32Property) GetValue() uint32 {
	return p.Get()
}

//GetName ...
func (p *DynamicUInt32Property) GetName() string {
	return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicUInt32Property) GetDefaultValue() uint32 {
	return p.defaultValue
}
