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
	chain := go_data_chain.CreateDataChain(test_data, false)

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
		fmt.Println(chain_item.ToArray()[i].ToString())
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

	//*****************************************************************************************
	// Create an instance if go_data_chain but with safe parameter to true
	// This will always return a chain item even if the data does not exist
	// For this to work you must use the GetMapItem(), GetArrayItem() and GetArrayCount methods
	// And check the Err property of the chain item for errors
	// This is a workaround so that if data does not exist the program does not crash
	// And give you the option to handle the error
	//*****************************************************************************************
	chain_error := go_data_chain.CreateDataChain(test_data, true)

	//***********************************************
	//get string data from a map maps_does_not_exists
	//***********************************************
	chain_item = chain_error.GetMapItem("data").GetMapItem("maps_does_not_exists") //get a map item
	fmt.Println(chain_item.GetMapItem("map_string_1").ToString())                  //cast as a string
	if chain_error.Err != nil {
		fmt.Printf("there was an error for GetMapItem:%v \n", chain_error.Err)
		chain_error.Err = nil //reset the error
	}

	//********************************
	//Get an array that does not exist
	//********************************
	chain_item = chain_error.GetMapItem("data").GetMapItem("arrays_does_not_exists") //get a map item
	if chain_item.Err != nil {
		fmt.Printf("there was an error for GetMapItem:%v \n", chain_error.Err)
	} else {
		fmt.Print(chain_item.ToString())
	}

	//****************************************
	//Because the array does not exist
	//The GetArrayCount() method will return 0
	// And return an error
	//****************************************
	for i := 0; i < chain_item.GetArrayCount(); i++ {
		fmt.Println(chain_item.GetArrayItem(i).ToString())
		fmt.Println(chain_item.ToArray()[i].ToString())
	}
	if chain_error.Err != nil {
		fmt.Printf("there was an error GetArrayCount:%v \n", chain_error.Err)
	}

}
