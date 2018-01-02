package manyfaced

import (
	"time"

	"github.com/nienie/manyfaced/configuration"
	"github.com/nienie/manyfaced/property"
)

var (
	propertyFactory *dynamicPropertyFactory
)

type dynamicPropertyFactory struct{}

func newDynamicPropertyFactory() *dynamicPropertyFactory {
	return &dynamicPropertyFactory{}
}

func (f *dynamicPropertyFactory) initWithConfigurationSource(config configuration.Configuration) {
	dynamicPropertySupport := property.NewDynamicPropertySupportImpl(config)
	property.RegisterWithDynamicPropertySupport(dynamicPropertySupport)
}

//GetStringProperty ...
func GetStringProperty(propName string, defaultValue string) *property.DynamicStringProperty {
	return property.NewDynamicStringProperty(propName, defaultValue)
}

//GetBoolProperty ...
func GetBoolProperty(propName string, defaultValue bool) *property.DynamicBoolProperty {
	return property.NewDynamicBoolProperty(propName, defaultValue)
}

//GetIntProperty ...
func GetIntProperty(propName string, defaultValue int) *property.DynamicIntProperty {
	return property.NewDynamicIntProperty(propName, defaultValue)
}

//GetUIntProperty ...
func GetUIntProperty(propName string, defaultValue uint) *property.DynamicUIntProperty {
	return property.NewDynamicUIntProperty(propName, defaultValue)
}

//GetInt32Property ...
func GetInt32Property(propName string, defaultValue int32) *property.DynamicInt32Property {
	return property.NewDynamicInt32Property(propName, defaultValue)
}

//GetUInt32Property ...
func GetUInt32Property(propName string, defaultValue uint32) *property.DynamicUInt32Property {
	return property.NewDynamicUInt32Property(propName, defaultValue)
}

//GetInt64Property ...
func GetInt64Property(propName string, defaultValue int64) *property.DynamicInt64Property {
	return property.NewDynamicInt64Property(propName, defaultValue)
}

//GetUInt64Property ...
func GetUInt64Property(propName string, defaultValue uint64) *property.DynamicUInt64Property {
	return property.NewDynamicUInt64Property(propName, defaultValue)
}

//GetFloat32Property ...
func GetFloat32Property(propName string, defaultValue float32) *property.DynamicFloat32Property {
	return property.NewDynamicFloat32Property(propName, defaultValue)
}

//GetFloat64Property ...
func GetFloat64Property(propName string, defaultValue float64) *property.DynamicFloat64Property {
	return property.NewDynamicFloat64Property(propName, defaultValue)
}

//GetInterfaceProperty ...
func GetInterfaceProperty(propName string, defaultValue interface{}) *property.DynamicInterfaceProperty {
	return property.NewDynamicInterfaceProperty(propName, defaultValue)
}

//GetTimeDurationProperty ...
func GetTimeDurationProperty(propName string, defaultValue time.Duration) *property.DynamicDurationProperty {
	return property.NewDynamicDurationProperty(propName, defaultValue)
}
