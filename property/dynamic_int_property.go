package property

//DynamicIntProperty ...
type DynamicIntProperty struct {
    prop                *DynamicProperty
    defaultValue        int
}

//NewDynamicIntProperty ...
func NewDynamicIntProperty(propName string, defaultValue int) *DynamicIntProperty {
    return &DynamicIntProperty{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
    }
}

//Get get the current value from the underlying DynamicProperty
func (p *DynamicIntProperty)Get() int {
    return p.prop.GetInt(p.defaultValue)
}

//GetValue ...
func (p *DynamicIntProperty)GetValue() int {
    return p.Get()
}

//GetName ...
func (p *DynamicIntProperty)GetName() string {
    return p.prop.GetName()
}

//GetDefaultValue ...
func (p *DynamicIntProperty)GetDefaultValue() int {
    return p.defaultValue
}