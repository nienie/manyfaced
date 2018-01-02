package property

import (
	"time"
)

//DynamicDurationProperty ...
type DynamicDurationProperty struct {
	prop         *DynamicProperty
	defaultValue time.Duration
}

//NewDynamicDurationProperty ...
func NewDynamicDurationProperty(propName string, defaultValue time.Duration) *DynamicDurationProperty {
	return &DynamicDurationProperty{
		prop:         GetDynamicProperty(propName),
		defaultValue: defaultValue,
	}
}

//Get ...
func (p *DynamicDurationProperty) Get() time.Duration {
	return p.prop.GetTimeDuration(p.defaultValue)
}

//GetValue ...
func (p *DynamicDurationProperty) GetValue() time.Duration {
	return p.Get()
}

//GetDefaultValue ...
func (p *DynamicDurationProperty) GetDefaultValue() time.Duration {
	return p.defaultValue
}
