# go-data-chain Easy Navigation of complex interface{} data types

## Description
Package to help navigate complex interface{} data types with auto casting of various types where possible.


## When to use go-data-chain
- Is you need to read data from complex interface{} for example when Unmarshal data into an interface.
- Read data returned from an REST API call and you don't Unmarshal the data into a type but an interface.


## Requirements
* go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install go-data-chain


## Project folders
Below is a description go-data-chain project folders and what they contain
|   Folder        | Description  | 
|-----------|---|
| examples    | yaml payload used to test the package  |
| /      | Main code |


## Installation and Basic usage
This will take you through the steps to install and get go-data-chain up and running.
<details>
<summary>1. Install</summary>

Once you have installed golang you can run the following command to get go-data-chain
```bash
go get github.com/Mrpye/go-data-chain
```
</details>

<details>
<summary>2. Include in your project</summary>

```go
    include github.com/Mrpye/go-data-chain
```
</details>


## Examples

<details>
<summary>1. Read yaml and navigate the data</summary>

### Example yaml data:

```yaml

data:
  arrays:
    - "array_string_1"
    - "array_string_2"
    - "array_string_3"
  maps:
    map_string_1: "map_string_1"
    map_string_2: "map_string_2"
    map_string_3: "map_string_3"
  convert_bool:
    string_bool_true: "true"
    string_bool_false: "false"
    string_bool_y: "y"
    string_bool_n: "n"
    string_bool_1: "1"
    string_bool_0: "0"
    string_bool_yes: "yes"
    string_bool_no: "no"
    string_bool_pass: "pass"
    string_bool_fail: "fail"
  convert_int:
    float_int: 1.56
    string_int: "2"
    int_int: 3
    bool_int_true: true
    bool_int_false: false
  convert_float:
    float_float: 1.56
    string_float: "1.56"
    int_float: 5
    bool_float_true: true
    bool_float_false: false
  convert_string:
    float_string: 1.56
    string_string: "string"
    int_string: 5
    bool_string_true: true
    bool_string_false: false
  convert_map:
  

```
### code example:

```go
package main

import (
	"fmt"
	"io/ioutil"

	go_data_chain "github.com/Mrpye/go-data-chain"
	"gopkg.in/yaml.v3"
)

func main() {
	//*******************
	// Store example data
	//*******************
	var test_data interface{}

	//******************************
	//Read the example data from a file
	//******************************
	file, err := ioutil.ReadFile("example_data.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(file), &test_data)
	if err != nil {
		panic(err)
	}
	//***********************************
	//Create an instance if go_data_chain
	//***********************************
	chain := go_data_chain.CreateDataChain(test_data)

	//*****************
	//Navigate the data
	//*****************

	//get string data from a map map_string_1
	chain_item := chain.GetMapItem("data").GetMapItem("maps")     //get a map item
	fmt.Println(chain_item.GetMapItem("map_string_1").ToString()) //cast as a string
	//Loop through an array
	chain_item = chain.GetMapItem("data").GetMapItem("arrays") //get a map item
	for i := 0; i < chain_item.GetArrayCount(); i++ {
		fmt.Println(chain_item.GetArrayItem(i).ToString())
		fmt.Println(chain_item.GetArray()[i].ToString())
	}

	//get bool data from a map convert_bool
	chain_item = chain.GetMapItem("data").GetMapItem("convert_bool") //get a map item
	fmt.Println(chain_item.GetMapItem("string_bool_true").ToBool())  //convert "true" to true
	fmt.Println(chain_item.GetMapItem("string_bool_no").ToBool())    //convert "no" to false

	//get int data from a map convert_int
	chain_item = chain.GetMapItem("data").GetMapItem("convert_int") //get a map item
	fmt.Println(chain_item.GetMapItem("float_int").ToInt64())       //convert 1.56 to 1
	fmt.Println(chain_item.GetMapItem("bool_int_true").ToInt8())    //convert true to 1

	//get float data from a map convert_float
	chain_item = chain.GetMapItem("data").GetMapItem("convert_float") //get a map item
	fmt.Println(chain_item.GetMapItem("float_float").ToFloat64())     //convert 1.56 to 1
	fmt.Println(chain_item.GetMapItem("float_float").ToFloat32())     //convert true to 1

	//get string data from a map convert_string
	chain_item = chain.GetMapItem("data").GetMapItem("convert_string") //get a map item
	fmt.Println(chain_item.GetMapItem("int_string").ToString())        //convert 1.56 to 1
	fmt.Println(chain_item.GetMapItem("string_string").ToString())     //convert true to 1

}

```

### Result:

```bash
map_string_1
array_string_1
array_string_1
array_string_2
array_string_2
array_string_3
array_string_3
true
false
1
1
1.56
1.56
5
string
```

</details>

---

## go-data-chain documentation

```go
//Data type 
type Data

// CreateDynamicData creates a new Data object
// - value: the value to set
CreateDataChain(value interface{}) *Data

// GetArrayItem returns an item from the array
// - index: the index of the item to get
GetArrayItem(index int) *Data

// GetArrayCount returns the number of items in the array
GetArrayCount() int

// GetArray returns an array of Data objects
GetArray() []Data

// GetMap returns a map of Data objects
GetMap() map[string]Data

// GetMapItem a Data object
// - Key: the key to get
GetMapItem(key string) *Data

// GetInterface returns the data as an interface{}
GetInterface() interface{}

// GetType returns the type of the data as a string
GetType() string

// ToString returns the data as a string
ToString() string

// ToBool returns the data as a bool
ToBool() bool 

// ToFloat32 returns the data as a float32
ToFloat32() float32

// ToFloat64 returns the data as a float64
ToFloat64() float64

// ToInt returns the data as an int
ToInt() int 

// ToInt8 returns the data as an int8
ToInt8() int8

// ToInt32 returns the data as an int32
ToInt32() int32

// ToInt64 returns the data as an int64
ToInt64() int64

```

---

## Change Log
### v0.1.0
- First build

## license
go-data-chain is Apache 2.0 licensed.
