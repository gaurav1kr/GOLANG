package main
import(
	"fmt"
	"time"
      )

func gosample1(){
	fmt.Println("we are in child thread-->1")
}

func gosample2(){
	fmt.Println("we are in child thread-->2")
}

func main(){
	a,b,c := 10,"gaurav",12.23
	go gosample1()
	go gosample2()
	fmt.Println("we are starting main thread")
	fmt.Printf("a=%d b=%s c=%f\n",a,b,c)
	time.Sleep(1*time.Second)
	fmt.Println("we are ending main thread")
}

//output
//we are starting main thread
//a=10 b=gaurav c=12.230000
//we are in child thread-->2
//we are in child thread-->1
//we are ending main thread
