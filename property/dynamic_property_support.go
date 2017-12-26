package property

import (
    "fmt"
    "github.com/nienie/manyfaced/configuration"
)

//DynamicPropertySupport the interface that defines the contract between DynamicProperty and its underlying support system.
type DynamicPropertySupport interface {

    //GetString get the string value of a given property. The string value will be futher
    //cached and parsed into specific type for DynamicProperty.
    //propName: the name of the property
    //return: the string value of the property
    GetString(propName string) string

    //AddConfigurationListener add the property change listener. This is necessary for the DynamicProperty to
    //receive callback once a property is updated in the underlying DynamicPropertySupport.
    AddConfigurationListener(expandedPropertyListener PropertyListener)
}

type DynamicPropertySupportImpl struct {
    Config    configuration.Configuration
}

//NewDynamicPropertySupportImpl ...
func NewDynamicPropertySupportImpl(config configuration.Configuration) *DynamicPropertySupportImpl {
    return &DynamicPropertySupportImpl{
        Config:         config,
    }
}

//GetString ...
func (o *DynamicPropertySupportImpl)GetString(key string) string {
    value := o.Config.GetProperty(key)
    if value == nil {
        return ""
    }
    return fmt.Sprint(value)
}

//AddConfigurationListener ...
func (o *DynamicPropertySupportImpl)AddConfigurationListener(propertyListener PropertyListener) {
    adapter := NewConfigurationListenerAdapter(propertyListener)
    o.Config.AddConfigurationListener(adapter)
}

