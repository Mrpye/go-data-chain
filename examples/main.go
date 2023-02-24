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
