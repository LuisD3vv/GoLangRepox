package main

import "fmt"

func default_values() {

	// en go, todos las variables tienen valor asignado, ya sea 0 o false
	var (
		defaultInt    int     //0
		defaultUint   uint    //0
		defaultFloat  float32 //0
		defaultbool   bool    //false
		defaultString string  // 0
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultbool, defaultString)
}
