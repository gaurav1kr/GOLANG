package main

import (
	"fmt"
	"strings"
)

func ReplaceAll(text string, searchword string, replaceword string) string {
	split_text := strings.Split(text, " ")
	for index, elem := range split_text {
		if string(elem) == searchword {

			split_text[index] = replaceword
		}
	}
	return (strings.Join([]string(split_text), " "))
}
func main() {
	var str string = "gaurav is learning golang in golang tool"
	str = ReplaceAll(str, "golang", "cplusplus")
	fmt.Println(str)
}
