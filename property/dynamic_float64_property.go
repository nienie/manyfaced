package property

//DynamicFloat64Property ...
type DynamicFloat64Property struct {
    prop                *DynamicProperty
    defaultValue        float64
}

//NewDynamicFloat64Property ...
func NewDynamicFloat64Property(propName string, defaultValue float64) *DynamicFloat64Property {
    return &DynamicFloat64Property{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
    }
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicFloat64Property)Get() float64 {
    return p.prop.GetFloat64(p.defaultValue)
}

//GetValue ...
func (p *DynamicFloat64Property)GetValue() float64 {
    return p.Get()
}

//GetName ...
func (p *DynamicFloat64Property)GetName() string {
    return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicFloat64Property)GetDefaultValue() float64 {
    return p.defaultValue
}

