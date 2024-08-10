package main
import "fmt"
import "errors"
import "strings"
// import "unicode/utf8"

func main() {
	// fmt.Println("Hello World!")

	// var intNum int // int can be followed by memory amount int16, int32, etc.
	// fmt.Println(intNum)

	// // uint can be used in order to store ints twice as large in same memory space
	// // uint is ONLY positives


	// var floatNum float32 // Must be type float32 or float64
	// fmt.Println(floatNum)

	// // Think about the datatype being used for all vars


	// // Operations between mixed types cannot be done, must cast to same type
	// var floatNum32 float32 = 10.1
	// var intNum32 int32 = 2
	// var result float32 = floatNum32 * float32(intNum32)
	// fmt.Println(result)

	// var intNum1 int = 3
	// var intNum2 int = 2

	// fmt.Println(intNum1/intNum2) // Answer will be rounded down
	// fmt.Println(intNum1%intNum2) // Will answer with the remainder

	// var someString string = "Hello" + " " + "World" + "\nNew Lined"
	// fmt.Println(someString)

	// fmt.Println(len("A")) // Returns not the length but the # of bytes in a string
	// fmt.Println(utf8.RuneCountInString("A")) // Returns the length of the string (Runes represent a character)
	
	// var char rune = 'a'
	// fmt.Println(char)

	// var boolean bool = false
	// fmt.Println(boolean)

	// /**
	// 	Default values exist for most things....
	// 	All int + floats and rune 0
	// 	Strings ''
	// 	Booleans false
	// */

	// //Type can be inferred

	// var typeless = "String"
	// fmt.Println(typeless)
	
	// noDeclaration := "missing a var"
	// fmt.Println(noDeclaration)

	// // You can apply the above as well!
	// var var1, var2 int = 1, 2
	// fmt.Println(var1 + var2)

	// // Constants must be declared
	// const aConst string = "This value can no longer be changed"
	// fmt.Println(aConst)
	printMe("Okay")

	var result, remainder, err = intDivision(3, 0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

	// Some other things

	// [length of the array]type of the array
	var intArr [3]int32 // Assigns 12 bytes of contiguous memory for this

	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0]) // Prints memory location
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var newIntArr [3]int32 = [3]int32{1, 2, 3}
	// newIntArr := [3]int32{1, 2, 3}
	// newIntArr := [...]int32{1, 2, 3}  can infer the size of the array w/ ...
	fmt.Println(newIntArr)

	// Slicing acts similar to lists, backed by arrays
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))


	// Can append multiple items to slice using ... (the spread operator)
	var intSlice2 []int32 = []int32{10, 11, 12}
	intSlice2 = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Make can create a slice of fixed length, useful for not having to reallocate memory
	// Type, length, capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var mapping map[string]uint8 = make(map[string]uint8)
	fmt.Println(mapping)

	var mapping2 = map[string]uint8 {
		"Sam": 1,
		"John": 3,
	}

	fmt.Println(mapping2["Sam"])
	fmt.Println(mapping2["John"])
	// If item is not found in mapping, the default value for the value type is returned
	fmt.Println(mapping2["NOT IN MAPPING"]) // Prints 0 (default for type uint8)

	// Maps in go return an optional secondary value that is a boolean reporting if the item was in the mapping
	var age, exists = mapping2["NOT HERE"]
	fmt.Printf("The age was %v, but the value found status is: %v", age, exists)

	// To delete from map
	delete(mapping2, "Sam")

	// Iterations in Go
	// Order is not preserved in Go mappings, multiple iterations will return a different ordering
	for name, age := range mapping2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for index, value := range intSlice2 {
		fmt.Printf("Index: %v, Value: %v \n", index, value)
	}

	// No while loop in go, must use for to achieve
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	stringsNotes()
	structsAndTypes()
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// Can optionally return tuples of results
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func stringsNotes() {
	var myString = []rune("resume") // Important to note that if this was type string, then iterating over would reveal ascii and count bytes
	for index, value := range myString {
		fmt.Println(index, value)
	}

	var myRune = 'a'
	fmt.Println(myRune)

	var strSlice = []string{"h", "e", "l", "l", "o"}
	var catStr = ""
	for index := range strSlice {
		catStr += strSlice[index]
	}
	fmt.Printf("\n%v", catStr)

	// Or use string builder

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var newStr = strBuilder.String()
	fmt.Println(newStr)
}

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

type engine interface {
	milesLeft() uint8
}

type owner struct {
	name string
}


func structsAndTypes() {
	var engine gasEngine = gasEngine{25, 15, owner{"Tenchi"}}

	//Anon struct
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{25, 15} // Must be defined and initialized in one go
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Println(engine.mpg, engine.gallons, engine.name)

	var myElectricEngine electricEngine = electricEngine{100, 100}

	canMakeIt(myElectricEngine, 200)
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You got it!")
	} else {
		fmt.Println("Get some gas!")
	}
}