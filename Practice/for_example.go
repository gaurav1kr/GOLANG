package main

import (
	"fmt"
	"strings"
)

func main() {
	// Correct slice declaration
	str := "Thales is waiting for you"
	str_array := strings.Split(str," ")
	for idx,strs := range str_array{
		if(strs == "waiting"){
			str_array[idx] = "passing"	
		}
	}
	fmt.Printf("Replaced string is --> %s\n", strings.Join([]string(str_array)," "))
}

