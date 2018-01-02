package configuration

import (
	"fmt"
	"strings"
	"sync"

	"github.com/nienie/manyfaced/parser"
)

//MapConfiguration ...
type MapConfiguration struct {
	sync.RWMutex
	configurationAccessor  Accessor
	storage                map[string]interface{}
	configurationListeners []Listener
}

//NewMapConfiguration ...
func NewMapConfiguration(m map[string]interface{}) *MapConfiguration {
	cfg := &MapConfiguration{
		storage:                make(map[string]interface{}),
		configurationListeners: make([]Listener, 0),
	}
	for key, val := range m {
		cfg.storage[key] = val
	}
	cfg.SetConfigurationAccessor(cfg)
	return cfg
}

//AddConfigurationListener ...
func (c *MapConfiguration) AddConfigurationListener(listener Listener) {
	if listener != nil {
		c.configurationListeners = append(c.configurationListeners, listener)
	}
}

//RemoveConfigurationListener ...
func (c *MapConfiguration) RemoveConfigurationListener(listener Listener) {
	if listener != nil {
		for i, l := range c.configurationListeners {
			if l == listener {
				c.configurationListeners = append(c.configurationListeners[:i], c.configurationListeners[i+1:]...)
				break
			}
		}
	}
}

//GetConfigurationListeners ...
func (c *MapConfiguration) GetConfigurationListeners() []Listener {
	return c.configurationListeners
}

//SetConfigurationAccessor ...
func (c *MapConfiguration) SetConfigurationAccessor(ca Accessor) {
	c.configurationAccessor = ca
}

//Subset ...
func (c *MapConfiguration) Subset(prefix string) Configuration {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	subConfiguration := NewMapConfiguration(nil)
	for key, value := range c.storage {
		if strings.HasPrefix(key, prefix) {
			subKey := key[0:len(prefix)]
			subConfiguration.AddProperty(subKey, value)
		}
	}
	return subConfiguration
}

//IsEmpty ...
func (c *MapConfiguration) IsEmpty() bool {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	return len(c.storage) == 0
}

//ContainsKey ...
func (c *MapConfiguration) ContainsKey(key string) bool {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	_, ok := c.storage[key]
	return ok
}

//AddProperty ...
func (c *MapConfiguration) AddProperty(key string, value interface{}) {
	c.fireEvent(EventAddProperty, key, value, true)
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.storage[key] = value
	c.fireEvent(EventAddProperty, key, value, false)
}

//SetProperty ...
func (c *MapConfiguration) SetProperty(key string, value interface{}) {
	c.fireEvent(EventSetProperty, key, value, true)
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.storage[key] = value
	c.fireEvent(EventSetProperty, key, value, false)
}

//ClearProperty ...
func (c *MapConfiguration) ClearProperty(key string) {
	c.fireEvent(EventClearProperty, key, nil, true)
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	delete(c.storage, key)
	c.fireEvent(EventClearProperty, key, nil, false)
}

//Clear ...
func (c *MapConfiguration) Clear() {
	c.fireEvent(EventClear, "", nil, true)
	for key := range c.storage {
		c.ClearProperty(key)
	}
	c.storage = map[string]interface{}{}
	c.fireEvent(EventClear, "", nil, false)
}

//GetProperty ...
func (c *MapConfiguration) GetProperty(key string) interface{} {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	return c.storage[key]
}

//GetKeys ...
func (c *MapConfiguration) GetKeys() []string {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	keys := make([]string, 0, len(c.storage))
	for key := range c.storage {
		keys = append(keys, key)
	}
	return keys
}

//GetBool ...
func (c *MapConfiguration) GetBool(key string, defaultValue bool) bool {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseBool(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetBool gets a bool associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetBool(key string) bool {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseBool(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as bool type", value))
	}
	return val
}

//GetInt gets a int associated with the given configuration key.
func (c *MapConfiguration) GetInt(key string, defaultValue int) int {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetInt gets a int associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetInt(key string) int {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as int type", value))
	}
	return val
}

//GetUInt gets an uint associated with the given configuration key.
func (c *MapConfiguration) GetUInt(key string, defaultValue uint) uint {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetUInt gets an uint associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetUInt(key string) uint {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as uint type", value))
	}
	return val
}

//GetInt32 gets an int32 associated with the given configuration key.
func (c *MapConfiguration) GetInt32(key string, defaultValue int32) int32 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt32(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetInt32 gets an int32 associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetInt32(key string) int32 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt32(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as int32 type", value))
	}
	return val
}

//GetUInt32 gets an uint32 associated with the given configuration key.
func (c *MapConfiguration) GetUInt32(key string, defaultValue uint32) uint32 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt32(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetUInt32 gets an int32 associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetUInt32(key string) uint32 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt32(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as uint32 type", value))
	}
	return val
}

//GetInt64 gets a int64 associated with the given configuration key.
func (c *MapConfiguration) GetInt64(key string, defaultValue int64) int64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt64(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetInt64 gets a int64 associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetInt64(key string) int64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseInt64(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as int64 type", value))
	}
	return val
}

//GetUInt64 gets an int64 associated with the given configuration key.
func (c *MapConfiguration) GetUInt64(key string, defaultValue uint64) uint64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt64(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetUInt64 gets an uint64 associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetUInt64(key string) uint64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseUInt64(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as uint64 type", value))
	}
	return val
}

//GetString gets a string associated with the given configuration key.
func (c *MapConfiguration) GetString(key string, defaultValue string) string {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseString(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetString gets a string associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetString(key string) string {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseString(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as string type", value))
	}
	return val
}

//GetFloat64 gets a float64 associated with the given configuration key.
func (c *MapConfiguration) GetFloat64(key string, defaultValue float64) float64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		return defaultValue
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseFloat64(stringValue)
	if err != nil {
		return defaultValue
	}
	return val
}

//MustGetFloat64 gets a float64 associated with the given configuration key. If failed, a panic is thrown.
func (c *MapConfiguration) MustGetFloat64(key string) float64 {
	value := c.configurationAccessor.GetProperty(key)
	if value == nil {
		panic(fmt.Errorf("configuratioin key=%s is empty", key))
	}
	stringValue := fmt.Sprint(value)
	val, err := parser.ParseFloat64(stringValue)
	if err != nil {
		panic(fmt.Errorf("invalid configuration value=%+v, it can not be parsed as float64 type", value))
	}
	return val
}

func (c *MapConfiguration) fireEvent(eventType EventType,
	propName string, propValue interface{}, before bool) {
	event := c.createEvent(eventType, propName, propValue, before)
	for _, listener := range c.configurationListeners {
		listener.ConfigurationChanged(event)
	}
}

func (c *MapConfiguration) createEvent(eventType EventType, propName string,
	propValue interface{}, before bool) *Event {
	return NewConfigurationEvent(c, eventType, propName, propValue, before)
}
