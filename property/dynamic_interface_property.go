package property

import (
    "reflect"
    "encoding/json"
)

//DynamicInterfaceProperty ...
type DynamicInterfaceProperty struct {
    prop            *DynamicProperty
    defaultValue    interface{}
    isCached        bool
    cachedValue     interface{}
    stringValue     string
}

//NewDynamicInterfaceProperty ...
func NewDynamicInterfaceProperty(propName string, defaultValue interface{}) *DynamicInterfaceProperty {
    return &DynamicInterfaceProperty{
        prop:           GetDynamicProperty(propName),
        defaultValue:   defaultValue,
        isCached:       false,
        cachedValue:    nil,
        stringValue:    "",
    }
}

//Get ...
func (p *DynamicInterfaceProperty)Get() interface{} {
    if p.isCached && p.cachedValue != nil && p.stringValue == p.prop.GetString("") {
        return p.cachedValue
    }
    p.stringValue = p.prop.GetString("")
    t := reflect.TypeOf(p.defaultValue)
    var isPtr = false
    if t.Kind() == reflect.Ptr {
        isPtr = true
        t = t.Elem()
    }
    ptr := reflect.New(t)
    it := ptr.Interface()
    err := json.Unmarshal([]byte(p.stringValue), it)
    if err != nil {
        return p.defaultValue
    }
    if isPtr == false {
        i := ptr.Elem().Interface()
        p.cachedValue = i
    } else {
        p.cachedValue = it
    }
    p.isCached = true
    return p.cachedValue
}

//GetValue ...
func (p *DynamicInterfaceProperty)GetValue() interface{} {
    return p.Get()
}

//GetDefaultValue ...
func (p *DynamicInterfaceProperty)GetDefaultValue() interface{} {
    return p.defaultValue
}