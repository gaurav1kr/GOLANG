//An example of Goroutine synchronization issue
package main

import(
"fmt"
"sync"
)

var counter int

func increment(wg *sync.WaitGroup){
	for i :=0;i<1000;i++{			
		counter++
	}
	wg.Done()
}

func main(){
	var wg sync.WaitGroup 
	for i :=0;i<10;i++{				
		go increment(&wg) ;	
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("Counter = %d\n", counter)
}
//Output
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 8591
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 7392
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine.go
//Counter = 7721
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$
