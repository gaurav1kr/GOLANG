package main
import(
	"fmt"
)

func main(){
	var value interface{} = "gaurav" //Declaring an interface having string value 
	t := value.(string) //retrieving the string stored in variable value	

//	v := value.(int) //retrieving the int stored in variable value and thus it will give error as orig val is string stored in the interface 
	fmt.Println(t)
}
