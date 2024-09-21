package main
import "fmt"

func main(){
	messages := make(chan string)
	go func(){
		messages <- "gaurav"
	}()

	reciever := <- messages
	fmt.Printf("%s",reciever)
}
