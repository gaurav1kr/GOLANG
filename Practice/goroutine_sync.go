//An example of Goroutine synchronization issue
package main

import(
"fmt"
"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup){
	for i :=0;i<1000;i++{			
		mutex.Lock()
		counter++
		mutex.Unlock()
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
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
//gakumar@DESKTOP-CTCEABV:~/GOLANG/Practice$ go run goroutine_sync.go
//Counter = 10000
