package scanner

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"fmt"
	"runtime/debug"
)

//Rows defines methods that scanner needs, which database/sql.Rows already implements
type Rows interface {
	Close() error

	Columns() ([]string, error)

	Next() bool

	Scan(dest ...interface{}) error
}

const (
	// DefaultTagName is the default struct tag name
	DefaultTagName = "ddb"
	cTimeFormat    = "2006-01-02 15:04:05"
)

var (
	userDefinedTagName string
	//ErrTargetNotSettable means the second param of Bind is not settable
	ErrTargetNotSettable = errors.New("[scanner]: target is not settable! a pointer is required")
	//ErrNilRows means the first param can't be a nil
	ErrNilRows = errors.New("[scanner]: rows can't be nil")
	//ErrSliceToString means only []uint8 can be transmuted into string
	ErrSliceToString = errors.New("[scanner]: can't transmute a non-uint8 slice to string")
)

//SetTagName can be set only once
func SetTagName(name string) {
	if userDefinedTagName != "" {
		return
	}
	userDefinedTagName = name
}

//ScanErr will be returned if an underlying type couldn't be AssignableTo type of target field
type ScanErr struct {
	structName, fieldName string
	from, to              reflect.Type
}

func (s ScanErr) Error() string {
	return fmt.Sprintf("[scanner]: %s.%s is %s which is not AssignableBy %s", s.structName, s.fieldName, s.to, s.from)
}

func newScanErr(structName, fieldName string, from, to reflect.Type) ScanErr {
	return ScanErr{structName, fieldName, from, to}
}

// Scan scans data from rows to target
// Don't forget to close the rows
func Scan(rows Rows, target interface{}) error {
	if reflect.TypeOf(target).Kind() != reflect.Ptr || nil == target || reflect.ValueOf(target).IsNil() {
		return ErrTargetNotSettable
	}

	data, err := resolveDataFromRows(rows)
	if nil != err || nil == data {
		return err
	}

	switch reflect.TypeOf(target).Elem().Kind() {
	case reflect.Slice:
		err = bindSlice(data, target)
	default:
		err = bind(data[0], target)
	}

	return err
}

// ScanMap returns the result in the form of []map[string]interface{}
func ScanMap(rows Rows) ([]map[string]interface{}, error) {
	return resolveDataFromRows(rows)
}

// ScanMapClose is the same as ScanMap and close the rows
func ScanMapClose(rows Rows) ([]map[string]interface{}, error) {
	result, err := ScanMap(rows)
	if nil != err {
		return nil, err
	}
	err = rows.Close()
	return result, err
}

// ScanClose is the same as Scan and helps you Close the rows
// Don't exec the rows.Close after calling this
func ScanClose(rows Rows, target interface{}) error {
	err := Scan(rows, target)
	if nil != rows {
		if nil == err {
			err = rows.Close()
		}
	}
	return err
}

//caller must guarantee to pass a &slice as the second param
func bindSlice(arr []map[string]interface{}, target interface{}) error {
	targetObj := reflect.ValueOf(target)
	if !targetObj.Elem().CanSet() {
		return ErrTargetNotSettable
	}
	length := len(arr)
	valueArrObj := reflect.MakeSlice(targetObj.Elem().Type(), 0, length)
	typeObj := valueArrObj.Type().Elem()
	var err error
	for i := 0; i < length; i++ {
		newObj := reflect.New(typeObj)
		newObjInterface := newObj.Interface()
		err = bind(arr[i], newObjInterface)
		if nil != err {
			return err
		}
		valueArrObj = reflect.Append(valueArrObj, newObj.Elem())
	}
	targetObj.Elem().Set(valueArrObj)
	return nil
}

func bind(result map[string]interface{}, target interface{}) (resp error) {
	if nil != resp {
		return
	}
	defer func() {
		if r := recover(); nil != r {
			resp = fmt.Errorf("error:[%v], stack:[%s]", r, string(debug.Stack()))
		}
	}()
	valueObj := reflect.ValueOf(target).Elem()
	if !valueObj.CanSet() {
		return ErrTargetNotSettable
	}
	typeObj := valueObj.Type()
	if typeObj.Kind() == reflect.Ptr {
		ptrType := typeObj.Elem()
		newObj := reflect.New(ptrType)
		newObjInterface := newObj.Interface()
		err := bind(result, newObjInterface)
		if nil == err {
			valueObj.Set(newObj)
		}
		return err
	}
	typeObjName := typeObj.Name()

	for i := 0; i < valueObj.NumField(); i++ {
		fieldTypeI := typeObj.Field(i)
		fieldName := fieldTypeI.Name

		//for convenience
		wrapErr := func(from, to reflect.Type) ScanErr {
			return newScanErr(typeObjName, fieldName, from, to)
		}

		valuei := valueObj.Field(i)
		if !valuei.CanSet() {
			continue
		}
		tagName, ok := lookUpTagName(fieldTypeI)
		if !ok || "" == tagName {
			continue
		}
		mapValue, ok := result[tagName]
		if !ok {
			continue
		}
		err := convert(mapValue, valuei, wrapErr)
		if nil != err {
			return err
		}
	}
	return nil
}

func isIntSeriesType(k reflect.Kind) bool {
	return k >= reflect.Int && k <= reflect.Int64
}

func isUintSeriesType(k reflect.Kind) bool {
	return k >= reflect.Uint && k <= reflect.Uint64
}

func isFloatSeriesType(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func resolveDataFromRows(rows Rows) ([]map[string]interface{}, error) {
	if nil == rows {
		return nil, ErrNilRows
	}
	columns, err := rows.Columns()
	if nil != err {
		return nil, err
	}
	length := len(columns)
	var result []map[string]interface{}
	//unnecessary to put below into rows.Next loop,reduce allocating
	values := make([]interface{}, length)
	for i := 0; i < length; i++ {
		values[i] = new(interface{})
	}
	for rows.Next() {
		err = rows.Scan(values...)
		if nil != err {
			return nil, err
		}
		mp := make(map[string]interface{})
		for idx, name := range columns {
			//mp[name] = reflect.ValueOf(values[idx]).Elem().Interface()
			mp[name] = *(values[idx].(*interface{}))
		}
		result = append(result, mp)
	}
	return result, nil
}

func lookUpTagName(typeObj reflect.StructField) (string, bool) {
	var tName string
	if "" != userDefinedTagName {
		tName = userDefinedTagName
	} else {
		tName = DefaultTagName
	}
	name, ok := typeObj.Tag.Lookup(tName)
	if !ok {
		return "", false
	}
	name = resolveTagName(name)
	return name, ok
}

func convert(mapValue interface{}, valuei reflect.Value, wrapErr func(from, to reflect.Type) ScanErr) error {
	//vit: ValueI Type
	vit := valuei.Type()
	//mvt: MapValue Type
	mvt := reflect.TypeOf(mapValue)
	if nil == mvt {
		return nil
	}
	//[]byte tp []byte && time.Time to time.Time
	if mvt.AssignableTo(vit) {
		valuei.Set(reflect.ValueOf(mapValue))
		return nil
	}
	//time.Time to string
	switch assertT := mapValue.(type) {
	case time.Time:
		if vit.Kind() == reflect.String {
			sTime := assertT.Format(cTimeFormat)
			valuei.SetString(sTime)
			return nil
		}
		return wrapErr(mvt, vit)
	default:
	}

	//according to go-mysql-driver/mysql, driver.Value type can only be:
	//int64 or []byte(> maxInt64)
	//float32/float64
	//[]byte
	//time.Time if parseTime=true or DATE type will be converted into []byte
	switch mvt.Kind() {
	case reflect.Int64:
		if isIntSeriesType(vit.Kind()) {
			valuei.SetInt(mapValue.(int64))
		} else if isUintSeriesType(vit.Kind()) {
			valuei.SetUint(uint64(mapValue.(int64)))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Float32:
		if isFloatSeriesType(vit.Kind()) {
			valuei.SetFloat(float64(mapValue.(float32)))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Float64:
		if isFloatSeriesType(vit.Kind()) {
			valuei.SetFloat(mapValue.(float64))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Slice:
		mapValueSlice, ok := mapValue.([]byte)
		if !ok {
			return ErrSliceToString
		}
		mapValueStr := string(mapValueSlice)
		vitKind := vit.Kind()
		switch {
		case vitKind == reflect.String:
			valuei.SetString(mapValueStr)
		case isIntSeriesType(vitKind):
			intVal, err := strconv.ParseInt(mapValueStr, 10, 64)
			if nil != err {
				return wrapErr(mvt, vit)
			}
			valuei.SetInt(intVal)
		case isUintSeriesType(vitKind):
			uintVal, err := strconv.ParseUint(mapValueStr, 10, 64)
			if nil != err {
				return wrapErr(mvt, vit)
			}
			valuei.SetUint(uintVal)
		case isFloatSeriesType(vitKind):
			floatVal, err := strconv.ParseFloat(mapValueStr, 64)
			if nil != err {
				return wrapErr(mvt, vit)
			}
			valuei.SetFloat(floatVal)
		default:
			return wrapErr(mvt, vit)
		}
	default:
		return wrapErr(mvt, vit)
	}
	return nil
}
