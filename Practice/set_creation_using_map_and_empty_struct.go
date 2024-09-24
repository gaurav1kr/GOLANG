package main

import(
"fmt"
)

func main(){
	set := make(map[string]struct{})

	set["car"] = struct{}{}
	set["bus"] = struct{}{}
	set["train"] = struct{}{}

	for set_loop := range(set){
		fmt.Println(set_loop)
	}
	var str string = "van"

	if _,exists := set[str]; exists{
		fmt.Println(str," exists")		
	}else{
		fmt.Println(str," doesnt exist")		
	}
}
