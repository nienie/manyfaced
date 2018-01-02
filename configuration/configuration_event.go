package configuration

//EventType ...
type EventType int

const (
	//EventAddProperty ...
	EventAddProperty EventType = iota
	//EventClearProperty ...
	EventClearProperty
	//EventSetProperty ...
	EventSetProperty
	//EventClear ...
	EventClear
	//EventConfigurationSourceChanged ...
	EventConfigurationSourceChanged
)

//Event ...
type Event struct {
	EventType     EventType
	PropertyName  string
	PropertyValue interface{}
	BeforeUpdate  bool
	Source        interface{}
}

//NewConfigurationEvent ...
func NewConfigurationEvent(source interface{}, eventType EventType, propertyName string,
	propertyValue interface{}, beforeUpdate bool) *Event {
	return &Event{
		Source:        source,
		EventType:     eventType,
		PropertyName:  propertyName,
		PropertyValue: propertyValue,
		BeforeUpdate:  beforeUpdate,
	}
}
