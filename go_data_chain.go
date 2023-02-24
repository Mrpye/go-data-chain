// go-data-chain
// Created: 24/02/2023
// Written by: Andrew Pye
//Package to allow navigation of inteface{} data types in a generic way
//with auto conversion to the correct type where possible
package go_data_chain

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//Data is a struct that can hold any type of data
type Data struct {
	value interface{}
}

// GetType returns the type of the data as a string
func (m *Data) GetType() string {
	//Return the type as a string
	return reflect.TypeOf(m.value).Kind().String()
}

// ToString returns the data as a string
func (m *Data) ToString() string {
	//check if the value is a string
	if m.value != nil && reflect.TypeOf(m.value).Kind() == reflect.String {
		return m.value.(string)
	}
	return fmt.Sprintf("%v", m.value)
}

// ToInt returns the data as an int
func (m *Data) ToInt() int {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return int(i)
	case int:
		return val
	case float32:
		return int(val)
	case float64:
		return int(val)
	case int64:
		return int(val)
	case int8:
		return int(val)
	default:
		return 0
	}
}

// ToInt8 returns the data as an int8
func (m *Data) ToInt8() int8 {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return int8(i)
	case int:
		return int8(val)
	case float32:
		return int8(val)
	case float64:
		return int8(val)
	case int64:
		return int8(val)
	case int8:
		return val
	default:
		return 0
	}
}

// ToInt32 returns the data as an int32
func (m *Data) ToInt32() int32 {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0
		}
		return int32(i)
	case int:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case int64:
		return int32(val)
	case int32:
		return val
	default:
		return 0
	}
}

// ToInt64 returns the data as an int64
func (m *Data) ToInt64() int64 {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a int64
		if val == "" {
			return 0
		}
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return i
	case int:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case int64:
		return val
	default:
		return 0
	}
}

// ToFloat32 returns the data as a float32
func (m *Data) ToFloat32() float32 {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a float32
		if val == "" {
			return 0
		}
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return 0
		}
		return float32(f)
	case int:
		return float32(val)
	case float32:
		return val
	case float64:
		return float32(val)
	default:
		return 0
	}
}

// ToFloat64 returns the data as a float64
func (m *Data) ToFloat64() float64 {
	switch val := m.value.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		//convert the string to a float64
		if val == "" {
			return 0
		}
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return f
	case int:
		return float64(val)
	case float64:
		return val
	default:
		return 0
	}
}

// ToBool returns the data as a bool
func (m *Data) ToBool() bool {
	switch val := m.value.(type) {
	case bool:
		return val
	case string:
		if strings.ToLower(val) == "t" {
			return true
		} else if strings.ToLower(val) == "f" {
			return false
		} else if strings.ToLower(val) == "yes" {
			return true
		} else if strings.ToLower(val) == "no" {
			return false
		} else if strings.ToLower(val) == "y" {
			return true
		} else if strings.ToLower(val) == "n" {
			return false
		} else if strings.ToLower(val) == "1" {
			return true
		} else if strings.ToLower(val) == "pass" {
			return true
		} else if strings.ToLower(val) == "fail" {
			return false
		} else if strings.ToLower(val) == "0" {
			return false
		}
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val > 0
	case int8:
		return val > 0
	case int32:
		return val > 0
	case int64:
		return val > 0
	case float32:
		return val > 0
	case float64:
		return val > 0
	default:
		return false
	}
}

// GetInterface returns the data as an interface{}
func (m *Data) GetInterface() interface{} {
	return m.value
}

// GetMapItem gets a map item by key
// - Key: the key to get
// returns a Data object if the key exists or nil if it does not
func (m *Data) GetMapItem(key string) *Data {
	if m.value != nil && reflect.TypeOf(m.value).Kind() == reflect.Map {
		if m.value.(map[string]interface{})[key] != nil {
			return &Data{value: m.value.(map[string]interface{})[key]}
		}
	}
	return nil
}

// GetMap returns the data as a map
// returns a map of Data objects
func (m *Data) GetMap() map[string]Data {
	//check if the value is a map
	if m.value != nil && reflect.TypeOf(m.value).Kind() == reflect.Map {
		items := make(map[string]Data)
		for k, o := range m.value.(map[string]interface{}) {
			items[k] = Data{value: o}
		}
		return items
	}
	return nil
}

// GetArray returns the data as an array
// returns an array of Data objects
func (m *Data) GetArray() []Data {
	var items []Data
	k := reflect.TypeOf(m.value).Kind()
	//check if the value is an array
	if m.value != nil && (k == reflect.Slice || k == reflect.Array) {
		var items []Data
		for _, o := range m.value.([]interface{}) {
			items = append(items, Data{value: o})
		}
		return items
	}
	return items
}

// GetArrayCount returns the number of items in the array
func (m *Data) GetArrayCount() int {
	//check if the value is an array
	if m.value != nil && reflect.TypeOf(m.value).Kind() == reflect.Slice {
		return len(m.value.([]interface{}))
	}
	return 0
}

// GetArrayItem returns an item from the array
// - index: the index of the item to get
func (m *Data) GetArrayItem(index int) *Data {
	//check if the value is an array
	k := reflect.TypeOf(m.value).Kind()
	if m.value != nil && (k == reflect.Slice || k == reflect.Array) {
		if len(m.value.([]interface{})) > index {
			return &Data{value: m.value.([]interface{})[index]}
		}
	}
	return m
}

// CreateDynamicData creates a new Data object
// - value: the value to set
func CreateDataChain(value interface{}) *Data {
	data := Data{value: value}
	return &data
}
