package property

//Listener listener that handles property event notification.
//It handles events to add a property, set property, remove property, load and clear of the configuration source.
//DynamicPropertySupport registers this type listener with a DynamicPropertySupport to receive callbacks on
//changes so that it can dynamically change a value of a DynamicProperty.
type Listener interface {

	//ConfigurationSourceLoaded notifies the listener about a new source of configuration being added.
	ConfigurationSourceLoaded(source interface{})

	//AddProperty notifies this listener about a new value for the given property.
	AddProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error

	//SetProperty notifies this listener about a changed value for the given property.
	SetProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error

	//ClearProperty notifies this listener about a cleared property, which now has no value.
	ClearProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error

	//Clear notifies this listener that all properties have been cleared.
	Clear(source interface{}, beforeUpdate bool)
}
