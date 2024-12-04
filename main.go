// // // 

// // package main

// // import "fmt"

// // func main() {
// // 	fmt.Println("Data Types in Go")

// // 	var number1 int
// // 	var number2 float64
// // 	var complex complex64
// // 	var name string
// // 	var IsTrue bool

// // 	fmt.Println(number1, number2, complex, IsTrue, name)
// // 	fmt.Printf("An empty string %q \n", name) //null values of different types

// // 	fmt.Println("Ways to define variable or constants")
// // 	var age int                  //declaration without initialisation
// // 	var user = "Tommy"           //declaration with initialisation
// // 	email := "email@example.com" //shorthand for declaration with initialisation
// // 	fmt.Println(age, user, email)
// // 	const distance = 25
// // 	fmt.Printf("Type of distance: %T \n",distance)

// // }




// package main

// import "fmt"

// func main() {
// 	fmt.Println("Arrays")
// 	var fruits [10]string
// 	fruits[3] = "Mango"
// 	fruits[9] = "Apple"
// 	fmt.Println(fruits, fruits[1])
// 	numbers := [10]int{1, 4, 7, 5, 8, 9, 0}
// 	fmt.Println(numbers)

// 	fmt.Println("Slices")
// 	random := numbers[2:6]
// 	fmt.Println(random)
// 	var ages []int
// 	// ages[0] = 25
// 	ages = append(ages, 25, 75, 66)
// 	fmt.Println(ages)
// 	ages = append(ages, 12,57,87,98)
// 	names := []string{"Tommy", "Danny", "Robin"}
// 	fmt.Println(names)

// 	fmt.Println("Capacity and Length of slices")
// 	fmt.Println( "capacity of numbers array: ",cap(numbers))
// 	fmt.Println( "length of numbers array: ", len(numbers))
// 	fmt.Println( "capacity of slice 'random': ",cap(random))
// 	fmt.Println( "length of slice 'random': ", len(random))
// 	// fmt.Println( "capacity of slice 'ages': ",cap(ages))

// }


// package main

// import "fmt"

// func main() {
// 	// phonebook := make(map[string]string)
// 	// var phonebook = map[string]string{"Alice":"43356"}
// 	phonebook := map[string]string{"Tony": "5764784"}
// 	phonebook["Tommy"] = "123435"
// 	phonebook["Danny"] = "676789"
// 	fmt.Println(phonebook)
// 	fmt.Println(phonebook["Tommy"])

// 	delete(phonebook, "Tommy")
// 	fmt.Println(phonebook)

// 	elem, ok := phonebook["Danny"]
// 	fmt.Println(elem, ok)

// }


package main

import "fmt"

func main() {
    fmt.Println("Addition:", Add(5, 3))
    fmt.Println("Subtraction:", Subtract(5, 3))
	
	fmt.Println(x,y)
	a,b:=getVariable()
	fmt.Println("get variables are: \n", a, b)
	fmt.Println("identifier is: \n", name1)
	fmt.Println("Value of Pi: \n", pi)
	fmt.Println("Constant with type: \n", n)
	area := length * width
	fmt.Println("area is: \n", area)
	fmt.Println(zero, one, two)
	fmt.Println(isActive)
	// fmt.Println("array is", array(arr))
}
