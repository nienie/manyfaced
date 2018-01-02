package property

//DynamicInt64Property ...
type DynamicInt64Property struct {
	prop         *DynamicProperty
	defaultValue int64
}

//NewDynamicInt64Property ...
func NewDynamicInt64Property(propName string, defaultValue int64) *DynamicInt64Property {
	return &DynamicInt64Property{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicInt64Property) Get() int64 {
	return p.prop.GetInt64(p.defaultValue)
}

//GetValue ...
func (p *DynamicInt64Property) GetValue() int64 {
	return p.Get()
}

//GetName ...
func (p *DynamicInt64Property) GetName() string {
	return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicInt64Property) GetDefaultValue() int64 {
	return p.defaultValue
}
