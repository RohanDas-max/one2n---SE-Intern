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
	// p := PartialSum(10)
	// ps := p(10)
	// fmt.Println(ps)

	//variadic function (a function that accepts variable numbers of arguments)
	// sum(1, 2, 3, 4, 4, 5)

	//panic and recover
	// var action int
	// fmt.Println("Enter 1 for Student and 2 for Professional")
	// fmt.Scanln(&action)
	// /*  Use of Switch Case in Golang */
	// switch action {
	// case 1:
	// 	fmt.Printf("I am a  Student")
	// case 2:
	// 	fmt.Printf("I am a  Professional")
	// default:
	// 	panic(fmt.Sprintf("I am a  %d", action))
	// }
	// defer func() {
	// 	action := recover()
	// 	fmt.Println(action)
	// }()

	//array with variadic operator
	// var arr = [...]int{1, 2, 3, 4, 5}
	// // fmt.Println("\n", len(arr))
	// j := 0
	// for range arr {
	// 	if arr[j] == 1 {
	// 		fmt.Print(true, "\n")
	// 	} else {
	// 		fmt.Println(false)
	// 	}
	// 	j++
	// }

	//interface
	vi := value{1, 2}
	vf := valueF{1.2, 1.2}
	addvalue(vi)
	addvalue(vf)

}

func addvalue(s Sum) {

	fmt.Println(s)
	fmt.Println(s.add())
	fmt.Println(s.add())
	fmt.Println(s.addf())
}

//Interface
type Sum interface {
	add() int
	addf() float64
}
type value struct {
	a int
	b int
}
type valueF struct {
	a float64
	b float64
}

func (v value) add() int {
	return v.a + v.b
}
func (v valueF) add() int {
	return int(v.a + v.b)
}

func (v valueF) addf() float64 {
	return v.a + v.b
}
func (v value) addf() float64 {
	return float64(v.a + v.b)
}

// type Emp int
// func (e Emp) PrintEm(name string) string {
// 	return name
// }

//can use func as a type
// type fn func(int) int

//Higher Order Function (type of which return/recives a function)
// func sum(x, y int) int {
// 	return x + y
// }
// func PartialSum(x int) fn {
// 	return func(i int) int {
// 		return sum(x, i)
// 	}
// }

//variadic function
// func sum(x ...int) {
// 	var res int
// 	for i := range x {
// 		res += x[i]
// 	}
// 	fmt.Println(res)
// }
