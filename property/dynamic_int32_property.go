package property

//DynamicInt32Property ...
type DynamicInt32Property struct {
    prop                *DynamicProperty
    defaultValue        int32
}

//NewDynamicInt32Property ...
func NewDynamicInt32Property(propName string, defaultValue int32) *DynamicInt32Property {
    return &DynamicInt32Property{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
    }
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicInt32Property)Get() int32 {
    return p.prop.GetInt32(p.defaultValue)
}

//GetValue ...
func (p *DynamicInt32Property)GetValue() int32 {
    return p.Get()
}

//GetName ...
func (p *DynamicInt32Property)GetName() string {
    return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicInt32Property)GetDefaultValue() int32 {
    return p.defaultValue
}
