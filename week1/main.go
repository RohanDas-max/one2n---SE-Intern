package main

import "fmt"

func main() {

	// var (
	// 	name   string
	// 	age    int
	// 	ph     int
	// 	income float32
	// )
	// name = "rohan"
	// age = 12
	// ph = 1212312312312
	// income = 1200.2
	// fmt.Println(name, age, ph, income)

	// storing a string to a int value
	// val := "100"
	// newVal := 0
	// fmt.Sscan(val, &newVal)
	// fmt.Printf("%d and type is %T\n", newVal, newVal)

	//converting string to float
	// fl := "-3.143434"

	// f, err := strconv.ParseFloat(fl, 32)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%T\n", f)

	//infinte for loop
	// i := 1
	// for {
	// 	fmt.Println(i, "Hello")
	// 	if i == 10 {
	// 		break
	// 	}
	// 	i++
	// }

	//anonymous function
	// fmt.Printf(
	// 	"100 (°F) = %.3f (°C)\n", //%.3f means to round float to 3 decimal places
	// 	func(f float64) float64 {
	// 		return (f - 32.0) * (5.0 / 9.0)
	// 	}(100),
	// )
	// fmt.Printf("sum of two number %d\n",
	// 	func(a, b int) int {
	// 		return a + b
	// 	}(10, 5),
	// )

	//running Higher Order function
	p := PartialSum(10)
	ps := p(10)
	fmt.Println(ps)

	//

}

//can use func as a type
type fn func(int) int

//Higher Order Function (type of which return/recives a function)
func sum(x, y int) int {
	return x + y
}
func PartialSum(x int) fn {
	return func(i int) int {
		return sum(x, i)
	}
}
