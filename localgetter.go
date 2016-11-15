package main

import "io/ioutil"

func GetLocalJson(_ string, _ string) []byte {
	result, err := ioutil.ReadFile("sample.json")

	if err != nil {
		panic(err.Error())
	}

	return result
}
