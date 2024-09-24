package main
import(
"fmt"
"time"
)

func worker(chd chan struct{}){ 
	fmt.Println("start....")						
	time.Sleep(2 * time.Second)
	fmt.Println("end....")						

	chd <- struct{}{}
}
func main(){
	chd := make(chan struct{})   					
	go worker(chd)			

	<-chd
	fmt.Println("Main: worker completed")
}
