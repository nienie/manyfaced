package parser

import (
    "fmt"
    "math"
    "strings"
    "strconv"
    "time"
)

//ParseValue ...
type ParseValue func(string) (interface{}, error)

// make this a var to overwrite it in a test
var is32Bit = ^uint(0) == math.MaxUint32

//ParseBool ...
func ParseBool(stringValue string) (bool, error) {
    for _, t := range []string{"true", "t", "yes", "y", "on", "1"} {
        if strings.ToLower(stringValue) == t {
            return true, nil
        }
    }
    for _, f := range []string{"false", "f", "no", "n", "off", "0"} {
        if strings.ToLower(stringValue) == f {
            return false, nil
        }
    }
    return false, fmt.Errorf("%s can not be parsed as bool", stringValue)
}

//ParseInt ...
func ParseInt(stringValue string) (int, error) {
    integer, err := strconv.ParseInt(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return intRangeCheck(integer)
}

//ParseUInt ...
func ParseUInt(stringValue string) (uint, error) {
    integer, err := strconv.ParseUint(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return uintRangeCheck(integer)
}

//ParseInt32 ...
func ParseInt32(stringValue string) (int32, error) {
    integer, err := strconv.ParseInt(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return int32(integer), nil
}

//ParseUInt32 ...
func ParseUInt32(stringValue string) (uint32, error) {
    integer, err := strconv.ParseUint(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return uint32(integer), err
}

//ParseInt64 ...
func ParseInt64(stringValue string) (int64, error) {
    integer, err := strconv.ParseInt(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return integer, nil
}

//ParseUInt64 ...
func ParseUInt64(stringValue string) (uint64, error) {
    integer, err := strconv.ParseUint(stringValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return integer, nil
}

//ParseFloat32 ...
func ParseFloat32(stringValue string)(float32, error) {
    float, err := strconv.ParseFloat(stringValue, 64)
    if err != nil {
        return 0.0, err
    }
    return float32(float), nil
}

//ParseFloat64 ...
func ParseFloat64(stringValue string)(float64, error) {
    float, err := strconv.ParseFloat(stringValue, 64)
    if err != nil {
        return 0.0, err
    }
    return float, nil
}

//ParseString ...
func ParseString(stringValue string)(string, error) {
    return stringValue, nil
}

//ParseTimeDuration ...
func ParseTimeDuration(stringValue string)(time.Duration, error) {
    return time.ParseDuration(stringValue)
}

// intRangeCheck checks if the value fits into the int type
func intRangeCheck(v int64) (int, error) {
    if is32Bit && (v < math.MinInt32 || v > math.MaxInt32) {
        return 0, fmt.Errorf("Value %d out of int range", v)
    }
    return int(v), nil
}

// uintRangeCheck checks if the value fits into the uint type
func uintRangeCheck(v uint64) (uint,error) {
    if is32Bit && v > math.MaxUint32 {
        return 0, fmt.Errorf("Value %d out of uint range", v)
    }
    return uint(v), nil
}
