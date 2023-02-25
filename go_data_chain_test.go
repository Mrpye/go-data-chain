// go-data-chain

// Created: 24/02/2023

// Written by: Andrew Pye

//Package to allow navigation of inteface{} data types in a generic way

//with auto conversion to the correct type where possible

package go_data_chain

import (
	"io/ioutil"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGetArrayItem(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("arrays")
	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "array_string_1",
			args: 0,
			want: "array_string_1",
		},
		{
			name: "array_string_2",
			args: 1,
			want: "array_string_2",
		},
		{
			name: "array_string_3",
			args: 2,
			want: "array_string_3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetArrayItem(tt.args).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArray(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("arrays")
	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "array_string_1",
			args: 0,
			want: "array_string_1",
		},
		{
			name: "array_string_2",
			args: 1,
			want: "array_string_2",
		},
		{
			name: "array_string_3",
			args: 2,
			want: "array_string_3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.ToArray()[tt.args].ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetArrayCount(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("arrays")
	tests := []struct {
		name string
		want int
	}{
		{
			name: "array_count",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetArrayCount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapItem(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("maps")
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "map_string_1",
			args: "map_string_1",
			want: "map_string_1",
		},
		{
			name: "map_string_2",
			args: "map_string_2",
			want: "map_string_2",
		},
		{
			name: "map_string_3",
			args: "map_string_3",
			want: "map_string_3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_bool")
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "string_bool_true",
			args: "string_bool_true",
			want: true,
		},
		{
			name: "string_bool_false",
			args: "string_bool_false",
			want: false,
		},
		{
			name: "string_bool_y",
			args: "string_bool_y",
			want: true,
		},
		{
			name: "string_bool_n",
			args: "string_bool_n",
			want: false,
		},
		{
			name: "string_bool_1",
			args: "string_bool_1",
			want: true,
		},
		{
			name: "string_bool_0",
			args: "string_bool_0",
			want: false,
		},
		{
			name: "string_bool_yes",
			args: "string_bool_yes",
			want: true,
		},
		{
			name: "string_bool_no",
			args: "string_bool_no",
			want: false,
		},
		{
			name: "string_bool_pass",
			args: "string_bool_pass",
			want: true,
		},
		{
			name: "string_bool_fail",
			args: "string_bool_fail",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToBool(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_int")
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "float_int",
			args: "float_int",
			want: 1,
		},
		{
			name: "string_int",
			args: "string_int",
			want: 2,
		},
		{
			name: "int_int",
			args: "int_int",
			want: 3,
		},
		{
			name: "bool_int_true",
			args: "bool_int_true",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_int")
	tests := []struct {
		name string
		args string
		want int8
	}{
		{
			name: "float_int",
			args: "float_int",
			want: 1,
		},
		{
			name: "string_int",
			args: "string_int",
			want: 2,
		},
		{
			name: "int_int",
			args: "int_int",
			want: 3,
		},
		{
			name: "bool_int_true",
			args: "bool_int_true",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToInt8(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}

}
func TestToInt32(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_int")
	tests := []struct {
		name string
		args string
		want int32
	}{
		{
			name: "float_int",
			args: "float_int",
			want: 1,
		},
		{
			name: "string_int",
			args: "string_int",
			want: 2,
		},
		{
			name: "int_int",
			args: "int_int",
			want: 3,
		},
		{
			name: "bool_int_true",
			args: "bool_int_true",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToInt32(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}

}
func TestToInt64(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_int")
	tests := []struct {
		name string
		args string
		want int64
	}{
		{
			name: "float_int",
			args: "float_int",
			want: 1,
		},
		{
			name: "string_int",
			args: "string_int",
			want: 2,
		},
		{
			name: "int_int",
			args: "int_int",
			want: 3,
		},
		{
			name: "bool_int_true",
			args: "bool_int_true",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToInt64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestToFloat32(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_float")
	tests := []struct {
		name string
		args string
		want float32
	}{
		{
			name: "float_float",
			args: "float_float",
			want: 1.56,
		},
		{
			name: "string_float",
			args: "string_float",
			want: 1.56,
		},
		{
			name: "int_float",
			args: "int_float",
			want: 5,
		},
		{
			name: "bool_float_true",
			args: "bool_float_true",
			want: 1,
		},
		{
			name: "bool_float_false",
			args: "bool_float_false",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToFloat32(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestToFloat64(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_float")
	tests := []struct {
		name string
		args string
		want float64
	}{
		{
			name: "float_float",
			args: "float_float",
			want: 1.56,
		},
		{
			name: "string_float",
			args: "string_float",
			want: 1.56,
		},
		{
			name: "int_float",
			args: "int_float",
			want: 5,
		},
		{
			name: "bool_float_true",
			args: "bool_float_true",
			want: 1,
		},
		{
			name: "bool_float_false",
			args: "bool_float_false",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToFloat64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestToString(t *testing.T) {
	var test_data interface{}

	file, _ := ioutil.ReadFile("examples/example_data.yaml")
	err := yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		t.Error(err)
	}
	chain := CreateDataChain(test_data, false)
	chain_item := chain.GetMapItem("data").GetMapItem("convert_string")
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "float_string",
			args: "float_string",
			want: "1.56",
		},
		{
			name: "string_string",
			args: "string_string",
			want: "string",
		},
		{
			name: "int_string",
			args: "int_string",
			want: "5",
		},
		{
			name: "bool_string_true",
			args: "bool_string_true",
			want: "true",
		},
		{
			name: "bool_string_false",
			args: "bool_string_false",
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chain_item.GetMapItem(tt.args).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDataChain() = %v, want %v", got, tt.want)
			}
		})
	}

}
