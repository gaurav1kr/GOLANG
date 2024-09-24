package main

import (
"fmt"
"slices"
)

func main(){
	str := []string {"gaurav" ,"baba" , "apple" , "father"} 
	slices.Sort(str)	
	fmt.Println(str)
}
