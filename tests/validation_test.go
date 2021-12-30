package main

import (
	"testing"
	"os"
	"log"
	"io/ioutil"
	"../valid"
)

func readFile(path string)([]byte){
	file, err := os.Open(path)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	out, err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err)
	}
	return out
}

func TestValidateMyJSON(t *testing.T){
	//Arrange test data
	byteValValid1 := readFile("valid_test_data/valid1.json")
	byteValValid2 := readFile("valid_test_data/valid2.json")
	byteInvalid1 := readFile("valid_test_data/invalid1.json")
	byteInvalid2 := readFile("valid_test_data/invalid2")
	byteInvalid3 := readFile("valid_test_data/invalid3.json")
	byteInvalid4 := readFile("valid_test_data/invalid4.json")
	byteInvalid5 := readFile("valid_test_data/invalid5.json")
	byteInvalid6 := readFile("valid_test_data/invalid6.json")

	//Act
	resValid1 := valid.ValidateMyJSON(byteValValid1)
	resValid2 := valid.ValidateMyJSON(byteValValid2)
	resInvalid1 := valid.ValidateMyJSON(byteInvalid1)
	resInvalid2 :=valid.ValidateMyJSON(byteInvalid2)
	resInvalid3 :=valid.ValidateMyJSON(byteInvalid3)
	resInvalid4 := valid.ValidateMyJSON(byteInvalid4)
	resInvalid5 :=valid.ValidateMyJSON(byteInvalid5)
	resInvalid6 := valid.ValidateMyJSON(byteInvalid6)

	//Assert
	if resValid1 != 1{
		t.Error("validFunc did not pass valid1.json file")
	}
	if resValid2 != 1{
		t.Error("validFunc did not pass valid2.json file")
	}
	if resInvalid1 != 0{
		t.Error("validFunc passed invalid1.json file")
	}
	if resInvalid2 != 0{
		t.Error("validFunc passed invalid2.json file")
	}
	if resInvalid3 != 0{
		t.Error("validFunc passed invalid3.json file")
	}
	if resInvalid4 != 0{
		t.Error("validFunc passed invalid4.json file")
	}
	if resInvalid5 != 0{
		t.Error("validFunc passed invalid5.json file")
	}
	if resInvalid6 != 0{
		t.Error("validFunc passed invalid6.json file")
	}
}