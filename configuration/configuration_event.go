package configuration

type ConfigurationEventType int

const (
    EventAddProperty ConfigurationEventType = iota
    EventClearProperty
    EventSetProperty
    EventClear
	EventConfigurationSourceChanged
)

//ConfigurationEvent ...
type ConfigurationEvent struct {
	EventType     ConfigurationEventType
	PropertyName  string
	PropertyValue interface{}
	BeforeUpdate  bool
	Source        interface{}
}

//NewConfigurationEvent ...
func NewConfigurationEvent(source interface{}, eventType ConfigurationEventType, propertyName string,
	propertyValue interface{}, beforeUpdate bool) *ConfigurationEvent {
	return &ConfigurationEvent{
		Source:        source,
		EventType:     eventType,
		PropertyName:  propertyName,
		PropertyValue: propertyValue,
		BeforeUpdate:  beforeUpdate,
	}
}
