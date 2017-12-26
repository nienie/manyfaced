package property

//DynamicFloat32Property ...
type DynamicFloat32Property struct {
    prop                *DynamicProperty
    defaultValue        float32
}

//NewDynamicFloat32Property ...
func NewDynamicFloat32Property(propName string, defaultValue float32) *DynamicFloat32Property {
    return &DynamicFloat32Property{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
    }
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicFloat32Property)Get() float32 {
    return p.prop.GetFloat32(p.defaultValue)
}

//GetValue ...
func (p *DynamicFloat32Property)GetValue() float32 {
    return p.Get()
}

//GetName ...
func (p *DynamicFloat32Property)GetName() string {
    return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicFloat32Property)GetDefaultValue() float32 {
    return p.defaultValue
}

