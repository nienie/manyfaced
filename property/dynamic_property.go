package property

import (
    "time"
    "fmt"
    "sync"
    "github.com/nienie/manyfaced/parser"
)

var (
    dynamicPropertiesLock       sync.RWMutex
    allDynamicProperties        map[string]*DynamicProperty
    dynamicPropertySupportImpl  DynamicPropertySupport
)

func init() {
    dynamicPropertiesLock = sync.RWMutex{}
    allDynamicProperties = make(map[string]*DynamicProperty)
}

//ChangedCallback ...
type ChangedCallback func()

type cachedValue struct {
    isCached  bool
    value     interface{}
    exception error
    parse     parser.ParseValue
}

func (c cachedValue)getValue(stringValue string) (interface{}, error) {
    if !c.isCached {
        c.value, c.exception = c.parse(stringValue)
        c.isCached = true
    }
    return c.value, c.exception
}

//DynamicProperty a cached configuration property value that is automatically updated when the config
// is changed.
//The struct is intended for those situations where the value of a property is fetched many times, and
//the value may be change on-the-fly.
//If the property is being read only once, "normal" access methods should be used.
//If the property is fixed, consider just caching the value in a variable.
type DynamicProperty struct {
    propName          string
    stringValue       string
    changedTime       time.Duration
    callbacks         []ChangedCallback
    validators        []ChangedValidator
    boolValue         cachedValue
    cachedStringValue cachedValue
    intValue          cachedValue
    uintValue         cachedValue
    int32Value        cachedValue
    uint32Value       cachedValue
    int64Value        cachedValue
    uint64Value       cachedValue
    float32Value      cachedValue
    float64Value      cachedValue
    durationValue     cachedValue
}

//GetDynamicProperty ...
func GetDynamicProperty(propName string) *DynamicProperty {
    dynamicPropertiesLock.Lock()
    defer dynamicPropertiesLock.Unlock()
    dynamicProperty, ok := allDynamicProperties[propName]
    if !ok {
        dynamicProperty = newDynamicProperty(propName)
        allDynamicProperties[propName] = dynamicProperty
    }
    return dynamicProperty
}

//InitializeDynamicProperty ...
func InitializeDynamicProperty(config DynamicPropertySupport) {
    dynamicPropertySupportImpl = config
    listener := &DynamicPropertyListener{}
    config.AddConfigurationListener(listener)
    updateAllProperties()
}

//RegisterWithDynamicPropertySupport ...
func RegisterWithDynamicPropertySupport(support DynamicPropertySupport) {
    InitializeDynamicProperty(support)
}

func newDynamicProperty(propName string) *DynamicProperty {
    dynamicProperty := &DynamicProperty{
        propName:           propName,
        callbacks:          make([]ChangedCallback, 0),
        validators:         make([]ChangedValidator, 0),
        boolValue:          cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseBool(stringValue)
            }),
        },
        cachedStringValue:  cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseString(stringValue)
            }),
        },
        intValue:           cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseInt(stringValue)
            }),
        },
        uintValue:          cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseUInt(stringValue)
            }),
        },
        int32Value:         cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseInt32(stringValue)
            }),
        },
        uint32Value:        cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseUInt32(stringValue)
            }),
        },
        int64Value:         cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseInt64(stringValue)
            }),
        },
        uint64Value:        cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseUInt64(stringValue)
            }),
        },
        float32Value:       cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseFloat32(stringValue)
            }),
        },
        float64Value:       cachedValue{
            parse: parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseFloat64(stringValue)
            }),
        },
        durationValue:      cachedValue{
            parse:  parser.ParseValue(func(stringValue string)(interface{}, error) {
                return parser.ParseTimeDuration(stringValue)
            }),
        },
    }
    dynamicProperty.initValue()
    return dynamicProperty
}

func (p *DynamicProperty)initValue() {
    newValue := dynamicPropertySupportImpl.GetString(p.propName)
    p.updateValue(newValue)
}

func (p *DynamicProperty)updateValue(newValue interface{}) {
    p.stringValue = fmt.Sprint(newValue)
    p.setStatusForValues()
    p.changedTime = time.Duration(time.Now().UnixNano())
}

func (p *DynamicProperty)setStatusForValues() {
    p.cachedStringValue.isCached = false
    p.boolValue.isCached = false
    p.intValue.isCached = false
    p.uintValue.isCached = false
    p.int32Value.isCached = false
    p.uint32Value.isCached = false
    p.int64Value.isCached = false
    p.uint64Value.isCached = false
    p.float32Value.isCached = false
    p.float64Value.isCached = false
    p.durationValue.isCached = false
}

func (p *DynamicProperty)validate(value interface{}) error {
    var err error
    newValue := fmt.Sprint(value)
    for _, validator := range p.validators {
        err = validator.Validate(newValue)
        if err != nil {
            return err
        }
    }
    return nil
}

func updateProperty(propName string, value interface{}) error {
    dynamicPropertiesLock.RLock()
    defer dynamicPropertiesLock.RUnlock()
    prop, ok := allDynamicProperties[propName]
    if !ok {
        return fmt.Errorf("these is no configuration key{%s} in the configuration", propName)
    }
    if err := prop.validate(value); err != nil {
        return err
    }
    prop.updateValue(value)
    return nil
}

func updateAllProperties() bool {
    changed := false
    dynamicPropertiesLock.RLock()
    defer dynamicPropertiesLock.RUnlock()
    for _, prop := range allDynamicProperties {
        prop.initValue()
        prop.notifyCallbacks()
        changed = true
    }
    return changed
}

//GetName ...
func (p *DynamicProperty) GetName() string {
    return p.propName
}

//GetChangedTimestamp ...
func (p *DynamicProperty) GetChangedTimestamp() time.Duration {
    return p.changedTime
}

//AddCallback ...
func (p *DynamicProperty)AddCallback(callback ChangedCallback) {
    if callback != nil {
        p.callbacks = append(p.callbacks, callback)
    }
}

//AddValidator ...
func (p *DynamicProperty)AddValidator(validator ChangedValidator) {
    if validator != nil {
        p.validators = append(p.validators, validator)
    }
}

func (p *DynamicProperty)notifyCallbacks() {
    for _, callback := range p.callbacks {
        go callback()
    }
}

//GetString ...
func (p *DynamicProperty)GetString(defaultValue string) string {
    value, err := p.cachedStringValue.getValue(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(string)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetString ...
func (p *DynamicProperty)MustGetString() string {
    value, err := p.cachedStringValue.getValue(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(string)
    if !ok {
        panic(fmt.Errorf("value type is %T not string", val))
    }
    return val
}

//GetBool ...
func (p *DynamicProperty)GetBool(defaultValue bool) bool {
    value, err := p.boolValue.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(bool)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetBool ...
func (p *DynamicProperty)MustGetBool() bool {
    value, err := p.boolValue.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(bool)
    if !ok {
        panic(fmt.Errorf("value type is %T not bool", val))
    }
    return val
}

//GetInt ...
func (p *DynamicProperty)GetInt(defaultValue int) int {
    value, err := p.intValue.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(int)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetInt ...
func (p *DynamicProperty)MustGetInt() int {
    value, err := p.intValue.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(int)
    if !ok {
        panic(fmt.Errorf("value type is %T not int", val))
    }
    return val
}


//GetUInt ...
func (p *DynamicProperty)GetUInt(defaultValue uint) uint {
    value, err := p.uintValue.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(uint)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetUInt ...
func (p *DynamicProperty)MustGetUInt() uint {
    value, err := p.uintValue.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(uint)
    if !ok {
        panic(fmt.Errorf("value type is %T not uint", val))
    }
    return val
}

//GetInt32 ...
func (p *DynamicProperty)GetInt32(defaultValue int32) int32 {
    value, err := p.int32Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(int32)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetInt32 ...
func (p *DynamicProperty)MustGetInt32() int32 {
    value, err := p.int32Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(int32)
    if !ok {
        panic(fmt.Errorf("value type is %T not int32", val))
    }
    return val
}

//GetUInt32 ...
func (p *DynamicProperty)GetUInt32(defaultValue uint32) uint32 {
    value, err := p.uint32Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(uint32)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetUInt32 ...
func (p *DynamicProperty)MustGetUInt32() uint32 {
    value, err := p.uint32Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(uint32)
    if !ok {
        panic(fmt.Errorf("value type is %T not uint32", val))
    }
    return val
}

//GetInt64 ...
func (p *DynamicProperty)GetInt64(defaultValue int64) int64 {
    value, err := p.int64Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(int64)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetInt64 ...
func (p *DynamicProperty)MustGetInt64() int64 {
    value, err := p.int64Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(int64)
    if !ok {
        panic(fmt.Errorf("value type is %T not int64", val))
    }
    return val
}

//GetUInt64 ...
func (p *DynamicProperty)GetUInt64(defaultValue uint64) uint64 {
    value, err := p.uint64Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(uint64)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetUInt64 ...
func (p *DynamicProperty)MustGetUInt64() uint64 {
    value, err := p.uint64Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(uint64)
    if !ok {
        panic(fmt.Errorf("value type is %T not uint64", val))
    }
    return val
}

//GetFloat32 ...
func (p *DynamicProperty)GetFloat32(defaultValue float32) float32 {
    value, err := p.float32Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(float32)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetFloat32 ...
func (p *DynamicProperty)MustGetFloat32() float32 {
    value, err := p.float32Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(float32)
    if !ok {
        panic(fmt.Errorf("value type is %T not float32", val))
    }
    return val
}

//GetFloat64 ...
func (p *DynamicProperty)GetFloat64(defaultValue float64) float64 {
    value, err := p.float64Value.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(float64)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetFloat64 ...
func (p *DynamicProperty)MustGetFloat64() float64 {
    value, err := p.float64Value.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(float64)
    if !ok {
        panic(fmt.Errorf("value type is %T not float64", val))
    }
    return val
}

//GetTimeDuration ...
func (p *DynamicProperty)GetTimeDuration(defaultValue time.Duration) time.Duration {
    value, err := p.durationValue.parse(p.stringValue)
    if err != nil {
        return defaultValue
    }
    val, ok := value.(time.Duration)
    if !ok {
        return defaultValue
    }
    return val
}

//MustGetTimeDuration ...
func (p *DynamicProperty)MustGetTimeDuration() time.Duration {
    value, err := p.durationValue.parse(p.stringValue)
    if err != nil {
        panic(err)
    }
    val, ok := value.(time.Duration)
    if !ok {
        panic(fmt.Errorf("value type is %T not float64", val))
    }
    return val
}
