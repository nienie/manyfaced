package property

//DynamicUInt64Property ...
type DynamicUInt64Property struct {
    prop                *DynamicProperty
    defaultValue        uint64
}

//NewDynamicUInt64Property ...
func NewDynamicUInt64Property(propName string, defaultValue uint64) *DynamicUInt64Property {
    return &DynamicUInt64Property{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
    }
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicUInt64Property)Get() uint64 {
    return p.prop.GetUInt64(p.defaultValue)
}

//GetValue ...
func (p *DynamicUInt64Property)GetValue() uint64 {
    return p.Get()
}

//GetName ...
func (p *DynamicUInt64Property)GetName() string {
    return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicUInt64Property)GetDefaultValue() uint64 {
    return p.defaultValue
}

