package main
import(
"fmt"
)

//struct definitions
type Triangle struct{
	base float64
	height float64
}

type Rectangle struct{
	length float64
	breadth float64
}

type Square struct{
	length float64
}

//function definitions
func(tri Triangle) Area() float64{
	return 0.5 * tri.base * tri.height
}

func(rect Rectangle) Area() float64{
	return  rect.length * rect.breadth
}

func(sqr Square) Area() float64{
	return sqr.length * sqr.length
}

type AreaInterface interface{
	Area() float64		
}

func main(){
rectObj := Rectangle{length:10 , breadth:9} 
triObj := Triangle{base:10 , height:9} 
sqrObj := Square{length:10} 

var Areainterfaceobj AreaInterface ;
Areainterfaceobj = rectObj
fmt.Printf("Area of rectangle = %f\n", Areainterfaceobj.Area())

Areainterfaceobj = triObj
fmt.Printf("Area of triangle = %f\n", Areainterfaceobj.Area())

Areainterfaceobj = sqrObj
fmt.Printf("Area of square = %f\n", Areainterfaceobj.Area())

}
