package property

import (
    "github.com/nienie/manyfaced/configuration"
)

//ConfigurationListenerAdapter convert PropertyListener to ConfigurationListener
type ConfigurationListenerAdapter struct {
    propertyListener    PropertyListener
}

//NewConfigurationListenerAdapter ...
func NewConfigurationListenerAdapter(propertyListener PropertyListener) *ConfigurationListenerAdapter {
    return &ConfigurationListenerAdapter{
        propertyListener:           propertyListener,
    }
}

//ConfigurationChanged ...
func (adapter *ConfigurationListenerAdapter)ConfigurationChanged(event *configuration.ConfigurationEvent) {
    switch event.EventType {
    case configuration.EventAddProperty:
        adapter.propertyListener.AddProperty(event.Source, event.PropertyName, event.PropertyValue, event.BeforeUpdate)
    case configuration.EventSetProperty:
        adapter.propertyListener.SetProperty(event.Source, event.PropertyName, event.PropertyValue, event.BeforeUpdate)
    case configuration.EventClearProperty:
        adapter.propertyListener.ClearProperty(event.Source, event.PropertyName, event.PropertyValue, event.BeforeUpdate)
    case configuration.EventClear:
        adapter.propertyListener.Clear(event.Source, event.BeforeUpdate)
    default:

    }
}

//GetPropertyListener ...
func (adapter *ConfigurationListenerAdapter) GetPropertyListener() PropertyListener {
    return adapter.propertyListener
}