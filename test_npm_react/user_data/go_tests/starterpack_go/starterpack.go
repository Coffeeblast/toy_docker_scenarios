// STARTERPACK with examples for GO language basic syntax
// run with ```go run .```
// see also Tour of GO ```https://go.dev/tour/list```

// package with entrypoint must be called "main" so that ```go run .``` knows which files to search for main function
// it must contain a function called "main()" ... TODO: find out how command line arguments are represented (main(...?))
package main

// import multiple modules in one statement, a.k.a. "factored import statement"
// functions exported from these modules (and imported here), start with capital letter (e.g. fmt.Println). 
// per go convention functions starting with lower letters are not exported
import (
	"fmt" // fmt: builtin module fmt contains functionality for printing into console
	"time"
	"io"
	"strings"
	"image"
	"sync" // mutex
)

// main function is program's entrypoint. If you remain it to something else, program breaks
func main() {
	helperWelcome()
	helperRunSection(sectionPrintingToConsole, "printing to console")
	helperRunSection(sectionVariables, "variables")
	helperRunSection(sectionDatatypes, "datatypes")
	helperRunSection(sectionFunctions, "functions")
	helperRunSection(sectionBranching, "branching")
	helperRunSection(sectionLoops, "loops")
	helperRunSection(sectionPointers, "pointer")
	helperRunSection(sectionStructures, "structures")
	helperRunSection(sectionArrays, "arrays")
	helperRunSection(sectionMaps, "maps")
	helperRunSection(sectionMethods, "methods")
	helperRunSection(sectionInterfaces, "interfaces")
	helperRunSection(sectionGenerics, "generics")
	helperRunSection(sectionConcurrency, "concurrency")
}

// SECTION: printing to console
// provides basic examples on how fmt.Println() can be used to print into console

func sectionPrintingToConsole() {
	// print single message
	fmt.Println("I am single message")

	// print multiple strings in row (notice how go inserts space in between them)
	fmt.Println("Passing multiple strings, string_0", ", and string_1", " and string_2")

	// passing non-string datatypes to print
	fmt.Println("Passing boolean", true, false)
	fmt.Println("Passing number", 1, 0.424, 1e120)
	fmt.Println("Passing time", time.Now())

	// formatted printing (TODO: examine more options)
	fmt.Printf("This is simple %v string that %v inserting values\n", "formatted", "demonstrates") // format printing values
	fmt.Printf("You can also print datatypes in formatted output %T, %T, %T, %T.\n", true, "whee", -5, -5.0) // printing datatypes
	fmt.Printf("This inserts a %q string\n", "quoted") // adding quotes when format printing strings

	// returning formatted string instead of printing to console
	s := fmt.Sprintf("I am a %v string", "formatted")
	fmt.Println(s)
}

// SECTION: variables
// provides examples on how variables are declared and assigned

func sectionVariables() {
	// variable from package scope is also visible inside functions defined in this package
	fmt.Println("Variable packageScopedVariable:", packageScopedVariable)

	// declare single variable
	var i int
	fmt.Println("i", i)

	// declare multiple variables of the same type
	// for successive variables of the same type, only last identifier has to have type explicitly given
	var j1, j2 int
	fmt.Println("j1", j1, "j2", j2)

	// !!! NOT ALLOWED
	// declare multiple variables of two mixed types
	// var k1, k2 int, l1, l2, l3 string

	// do it this way --> factored declaration
	var (
		k1, k2 int
		l1, l2, l3 string = "a", "b", "c"
	)
	fmt.Println("k1", k1, "k2", k2, "l1", l1, "l2", l2, "l3", l3)

	// variables with initiallizers
	var m1, m2 int = 1, 2
	fmt.Println("m1", m1, "m2", m2)

	// !!! NOT ALLOWED
	// providing initializtion value only for part of variables defined in one statement
	// var n1, n2 int = 1
	// fmt.Println("n1", n1, "n2", n2)

	// short variable definition
	o1 := "whee"
	p1, p2 := "glee", "bree"
	fmt.Println("o1", o1, "p1", p1, "p2", p2)

	// constants

	const q1 = 3.14
	const q2 float32 = 3.14
	const q3, q4 int = 9, 2
	fmt.Printf("Constants: (%v, %T), (%v, %T), (%v, %T), (%v, %T)\n", q1, q1, q2, q2, q3, q3, q4, q4)

	// q3 = 5 // not working, constant cannot be reassigned
	// const q5 := 9 // not working, constants do not support short assignment syntax
}

// variable declaration visible on the scope of whole package
var packageScopedVariable bool

// short variable definition does not work in package scope, only inside functions
// packageScopedVariable2 := true

// SECTION: data types
// discusses primitive datatypes available in go
func sectionDatatypes() {
	// bools
	var a1, a2 bool = true, false
	var a3 bool // default value is false
	fmt.Printf("Booloeans (bool): %v %v; default value: %v; data type: %T\n", a1, a2, a3, a3)

	a4 := true // short assignment produces correct datatype
	fmt.Printf("%v %t\n", a4, a4)

	// Strings
	var b1 = "I am a string"
	fmt.Printf("%v >> %T\n", b1, b1)

	b2 := "I am short assigned string" // short assignment produces correct datatype
	fmt.Printf("%v >> %T\n", b2, b2)
	
	// escape sequences
	b3 := "Printing string with \"quotes\", tab \"\t\", newline \"\n\" and backslash \\ characters"
	fmt.Println(b3)

	// NOT ALLOWED
	// strings with single quotes
	// b4 := 'I am single quoted string'
	// fmt.Println(b4)

	// single quotes work with one character only, but this produces type int32 with Unicode number of character
	b5 := 'A'
	fmt.Printf("%v %T\n", b5, b5)

	// NOT ALLOWED
	// single character with '' and datatype string
	// var b6 string = 'A'
	// fmt.Printf("%v %T\n", b6, b6)

	// integral datatypes
	// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.

	var c1 int = 9
	var c2 int8 = 9
	var c3 int16 = 9
	var c4 int32 = 9
	var c5 int64 = 9
	var c6 uint = 9
	var c7 uint8 = 9
	var c8 uint16 = 9
	var c9 uint32 = 9
	var c10 uint64 = 9
	var c11 uintptr = 9
	var c12 byte = 9 // alias for uint8
	var c13 rune = 9 // alias for int32, represents a Unicode code point
	
	fmt.Printf("%v: %v; type: %T\n", "int", c1, c1)
	fmt.Printf("%v: %v; type: %T\n", "int8", c2, c2)
	fmt.Printf("%v: %v; type: %T\n", "int16", c3, c3)
	fmt.Printf("%v: %v; type: %T\n", "int32", c4, c4)
	fmt.Printf("%v: %v; type: %T\n", "int64", c5, c5)
	fmt.Printf("%v: %v; type: %T\n", "uint", c6, c6)
	fmt.Printf("%v: %v; type: %T\n", "uint8", c7, c7)
	fmt.Printf("%v: %v; type: %T\n", "uint16", c8, c8)
	fmt.Printf("%v: %v; type: %T\n", "uint32", c9, c9)
	fmt.Printf("%v: %v; type: %T\n", "uint64", c10, c10)
	fmt.Printf("%v: %v; type: %T\n", "uintptr", c11, c11)
	fmt.Printf("%v: %v; type: %T\n", "byte", c12, c12)
	fmt.Printf("%v: %v; type: %T\n", "rune", c13, c13)

	// short assignment of literal produces int
	c14 := 9
	fmt.Printf("%v: %v; type: %T\n", "Short assignment of a literal", c14, c14)

	// short assignment of variable produces the same type as the variable
	c15 := c10 
	fmt.Printf("%v: %v; type: %T\n", "Short assignment of a variable", c15, c15)

	// floating point numbers
	var d1 float32 = 16.1
	var d2 float64 = 16.1
	d3 := 16.1 // this produces float64

	fmt.Printf("%v: %v; type: %T\n", "float32", d1, d1)
	fmt.Printf("%v: %v; type: %T\n", "float64", d2, d2)
	fmt.Printf("%v: %v; type: %T\n", "Short assignment of a literal", d3, d3)

	// NOT ALLOWED
	// float is undefined
	// var d4 float = 16.1
	// fmt.Printf("%v: %v; type: %T\n", "float", d4, d4)

	// other formats for declaration + assignment
	var d5, d6, d7 float32 = 9, 9e0, .5
	fmt.Printf("%v: %v; type: %T\n", "float32 (9)", d5, d5)
	fmt.Printf("%v: %v; type: %T\n", "float32 (9e0)", d6, d6)
	fmt.Printf("%v: %v; type: %T\n", "float32 (.5)", d7, d7)

	// complex numbers
	var e1 complex64 = 1 + 2i
	var e2 complex128 = 1 + 2i
	fmt.Printf("%v: %v; type: %T\n", "complex64", e1, e1)
	fmt.Printf("%v: %v; type: %T\n", "complex128", e2, e2)

	// short assignment
	e3 := 1 + 2i // this produces complex128
	fmt.Printf("%v: %v; type: %T\n", "Short assignment literal", e3, e3)

	// type conversion
	f1 := 3
	fmt.Printf("Variable %v of type %T can be converted to %v of type %T\n", f1, f1, float64(f1), float64(f1))
	f2 := float32(3)
	fmt.Printf("%v: %v; type: %T\n", "Converting to float32 by combination of short assignment and type conversion", f2, f2)
	
	// NOT ALLOWED
	// most operations do not work when types do not match properly (no implicit conversion)
	// fmt.Printf("Adding two different types: %v", float32(2) + int(3))

	var f3 float32
	f3 = 3 // this gets implicitly interpreted as float32
	// f3 = float64(3) // NOT ALLOWED - here types do not explicitly match
	fmt.Printf("%v: %v; type: %T\n", "Implicit type conversion on assignment", f3, f3)
}

// SECTION: functions
// provides examples on how functions are declared in go
func sectionFunctions() {
	// call function
	fmt.Println("Calling add(2, 3):", add(2, 3))
	fmt.Println("Calling add(2, 3):", add2(2, 3))
	aSwapped, bSwapped := swap("a", "b")
	fmt.Println("Calling swap(\"a\",\"b\")", aSwapped, bSwapped)
	three_plus_one, three_plus_two := addMany(3)
	fmt.Println("Calling addMany(3)", three_plus_one, three_plus_two)

	helperEmptyLine()
	// defer -- immediately evaluates function's arguments and executes it after current function returns
	deferOuterFunction()

	helperEmptyLine()
	// multiple defers are evaluated in LIFO order (stack)
	deferOuterFunctionTwoDefers()

	helperEmptyLine()
	// functions as values
	var square1 func(int, int)(int) = func(x, y int)(int){return x * x + y * y}
	fmt.Printf("square1(3, 4) = %v, type = %T\n", square1(3, 4), square1)

	square2 := func(x, y int)(int){return x * x + y * y}
	fmt.Printf("square2(3, 4) = %v, type = %T\n", square2(3, 4), square2)

	var square3 func(int, int)(int) // if we assign nothing, the value is nil
	fmt.Printf("square3 %v, type = %T, is nil? %v\n", square3, square3, square3 == nil)

	helperEmptyLine()
	// closures - functions that hold contex from the scope of function, in which they are defined
	counter := 0
	useCounter := func()(){
		fmt.Println("Count:", counter)
		counter += 1
	}
	useCounter()
	useCounter()
	useCounter()

	// closure can be returned as a result of a function. In such case, each closure holds its separate context
	counterMaker := func(counterName string)(func()()) {
		innerCounter := 0
		fn := func()(){
			fmt.Printf("[Counter %v] count is: %v\n", counterName, innerCounter)
			innerCounter += 1
		}
		return fn 
	}
	counterBob := counterMaker("Bob")
	counterAlice := counterMaker("Alice")
	counterBob()
	counterBob()
	counterAlice()
	counterBob()
	counterAlice()
	counterAlice()
}

// basic declaration of function with 2 parameters returning 1 result
func add(x int, y int) int {
	return x + y
}

// when declaring successive arguments of the same type, only last argument of 
// that type has to have explicitly written type
func add2(x, y int) int {
	return x + y
}

// function can have multiple return types
func swap(x, y string) (string, string) {
	return y, x
}

// named return values. In this case, the rule about skipping datatype declaration for multiple 
// successive return types of the same datatype holds as well
func addMany(value int) (plus_one, plus_two int) {
	plus_one = value + 1
	plus_two = value + 2
	return // this is called "naked return"
}

func deferArgument(x int) (int) {
	fmt.Println("Evaluating argument", x)
	return x
}

func deferInnerFunction(x int) {
	fmt.Println("Evaluating inner function with argument", x)
}

func deferOuterFunction() {
	fmt.Println("Before defering inner function")
	defer deferInnerFunction(deferArgument(9))
	fmt.Println("After defering inner function")
}

func deferOuterFunctionTwoDefers() {
	fmt.Println("Before defering inner function")
	defer deferInnerFunction(deferArgument(0))
	defer deferInnerFunction(deferArgument(1))
	fmt.Println("After defering inner function")
}

// SECTION: branching
// if statement
func sectionBranching() {
	// if statement
	
	// basic syntax
	// if condition  {
	//	statements_executed_if_true
	// } [else {
	// statements_executed_if_false
	// }]

	condition := true
	if condition {
		fmt.Println("Statement 1: Condition is true")
	} else { // note: "else" must be on the same row as }, otherwise go will not detect it as continuation of if statement
		fmt.Println("Statement 1: Condition is false")
	}

	condition = false
	if condition {
		fmt.Println("Statement 2: Condition is true")
	} else {
		fmt.Println("Statement 2: Condition is false")
	}

	// else statement is optional
	condition = true
	if condition {fmt.Println("Statement 3: Condition is true")}

	// works also with brackets around condition
	if (condition) {fmt.Println("Statement 4: Condition is true")}

	// NOT ALLOWED: braces are not optional
	// if (condition) fmt.Println("Statement 5: Condition is true")

	// if with short initialization statement

	// syntax:
	// if init_statement; condition {
	//	statements_executed_if_true
	// } [else {
	// statements_executed_if_false
	// }]

	if someBool := true; someBool {
		fmt.Println("Statement 6: Condition is true")
	} else {
		fmt.Println("Statement 6: Condition is false")
	}

	// fmt.Println(someBool) // NOT ALLOWED: someBool is scoped to the if statement
	
	// NOT ALLOWED: if (init_statment, condition) ... syntax not working with ()
	// if (someBool := true; someBool) {fmt.Println("Statement 7: Condition is true")}

	// (if, else if,  ... else) construct
	const someInt = 1
	if (someInt == 0) {
		fmt.Println("someInt is 0");
	} else if (someInt == 1) {
		fmt.Println("someInt is 1");
	} else {
		fmt.Println("someInt is neither 0 nor 1");
	}

	// switch statement
	// note: in go "switch" statement breaks automatically with first matched condition, it does not "fall through"

	switch someString := "b"; someString {
	case "a":
		fmt.Println("someString is a")
	case "b":
		fmt.Println("someString is b")
	default:
		fmt.Println("someString is neither a, nor b")
	}

	// fmt.Println("Somestring is:", someString) // NOT ALLOWED: someString is scoped to Switch Statement

	// alternative to if, else if ... else construction 
	switch someString := "b"; true {
	case someString == "a":
		fmt.Println("someString is a")
	case someString == "b" || someString == "c":
		fmt.Println("someString is b or c")
	case someString == "b" || someString == "d":
		fmt.Println("someString is b or d") // does not get called, switch breaks on the previous case
	// default: // default statement is optional
	// 	fmt.Println("someString is something else")
	}

	// switch true statement can also be writen like this
	someString := "b"
	switch {
	case someString == "a":
		fmt.Println("someString is a")
	case someString == "b" || someString == "c":
		fmt.Println("someString is b or c")
	case someString == "b" || someString == "d":
		fmt.Println("someString is b or d") // does not get called, switch breaks on the previous case
	}

}

// SECTION: loops
// for loop - the only loop in go
func sectionLoops() {
	// for loop - the only loop in go
	
	// syntax:
	// for init_statement; end_condition; post_statement {loop_body_Statements}
	
	for i := 0; i < 5; i++ {println("For loop 1 (standard for loop) index: ", i)}
	helperEmptyLine()

	// println(i) // NOT ALLOWED - scope of "i" is limited to the loop

	// for (i := 0; i < 5; i++) {println("For loop 2 index: ", i)} // NOT ALLOWED - syntax disallows for (...) {...}
	
	// break statement works as usual
	for i := 0; i < 5; i++ {
		if (i == 2) {break}
		fmt.Println("For loop 3 (for loop with break) index: ", i)
	}
	helperEmptyLine()

	// continue statement works as usual
	for i := 0; i < 5; i++ {
		if (i == 2) {continue}
		fmt.Println("For loop 4 (for loop with continue) index: ", i)	
	}
	helperEmptyLine()

	// while loop - made as for loop with skipped init_statement and post_statement
	j := 0
	for ; j < 3 ; {
		fmt.Println("For loop 5 (while loop with semicolons) index: ", j)
		j++
	}
	helperEmptyLine()

	// also for this usecase, the semicolons may be skipped
	k := 0
	for k < 3 {
		fmt.Println("For loop 6 (while loop without semicolons) index: ", k)
		k++
	}
	helperEmptyLine()

	// infinite loop - skipping all 3 statements in for loop. Must be broken out of using "break"
	l := 0
	for {
		if (l == 3){break}
		fmt.Println("For loop 7 (infinite loop) index: ", l)
		l++
	}

}

// SECTION: pointers
func sectionPointers() {
	// pointer declaration
	var val1 = 9 // normal value
	var a1 *int // declaration without definition
	var a2 *int = &val1 // declaration with definition
	a3 := &val1 // short syntax
	fmt.Printf(
		"(a1, %v, %v, %T), (a2, %v, %v, %T), (a3, %v, %v, %T)\n", 
		a1, "dereference not allowed for null pointer", a1, 
		a2, *a2, a2, 
		a3, *a3, a3)

	// dereference (indirection) operator (*)
	val2 := *a2
	fmt.Printf("val2 is dereference of a2 %v, but they have different addresses %v, %v\n", val2, &val2, a2)
	val2 = -5
	fmt.Printf("After changing val2 is %v, but *a2 is still %v, the addresses are still the same %v, %v\n", val2, *a2, &val2, a2)

	// address of operator (&)
	fmt.Printf("Pointer stores memory address of variable, which can be accessed as &val2 %v\n", &val2)

	// nil - default value for pointers, can also be assigned explicitly
	var b1 *float64
	var b2 *float64 = nil
	// b3 := nil // NOT ALLOWED - go does not know what type of pointer this should be
	fmt.Printf("Nil pointers: %v, %T, %v, %T\n", b1, b1, b2, b2)

	// TODO: pointer assignment
	// TODO: constant pointers
}

// SECTION: structures
func sectionStructures() {
	// structure declaration without initializer
	var a1 MyStruct
	fmt.Printf("a1: %v, type: %T\n", a1, a1) // type is "main.MyStruct"

	// structure declaration with initializer
	var a2 MyStruct = MyStruct{0.3, -0.2, "whee"}
	fmt.Println("a2", a2)

	// member variable access - standard dot syntax
	fmt.Printf("a2.x: %v, a2.name: %v\n", a2.x, a2.name)

	a3 := MyStruct{y : 1.3} // initialize only part of structure, other parts have default values
	fmt.Println("a3", a3)

	// pointer to structure
	var b1 *MyStruct = &a2
	b2 := &a2
	fmt.Printf("b1 %v %v %T\n", b1, *b1, b1)
	fmt.Printf("b2 %v %v %T\n", b2, *b2, b2)
	fmt.Printf("For pointer to structure member access (*b1).x %v may be shortened as b1.x %v\n", (*b1).x, b1.x)

	// initializing pointer on literal is also allowed
	b3 := &MyStruct{x: 5.0, name: "glee"}
	fmt.Println("b3", b3, "b3.name", b3.name)
}

// structure definition
type MyStruct struct {
	x float64
	y float64
	name  string
}

// SECTION: Arrays
func sectionArrays() {
	// static arrays - size of array is fixed
	var a1 [2]int
	var a2 [2]int = [2]int{2, 5}
	a3 := [2]int{3, 6}
	fmt.Printf("a1 %v %T, a2 %v %T, a3 %v %T\n", a1, a1, a2, a2, a3, a3)

	// assignment of element
	a1[1] = 5
	fmt.Println("a1", a1)

	// access to variables
	fmt.Println("!!! indexing is zero-based: a1[1]", a1[1])

	// access array length as
	fmt.Println("len(a1)", len(a1))

	// slicing
	b1 := [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Original array b1", b1)
	fmt.Printf("Standard slicing. It is left-inclusive right-exclusive: b1[2:6]: %v %T\n", b1[2:6], b1[2:6])
	fmt.Printf("Slice without first index. It takes slice from start till given index(right-exclusive): b1[:3]: %v %T\n", b1[:3], b1[:3])
	fmt.Printf("Slice without second index. It takes slice from first index (left-inclusive) till the end: b1[6:]: %v %T\n", b1[6:], b1[6:])
	fmt.Printf("Slice without both indices. It takes slice from the start till the end: b1[:]: %v %T\n", b1[:], b1[:])

	// type of slice is []T (without filled size)
	var b2 []int
	fmt.Printf("%v %T\n", b2, b2)

	// slice references array. Modifying slice modifies the original array
	b2 = b1[2:6]
	b2[0] = 999
	fmt.Println("b2", b2, "b1", b1)

	// slices have: 
	// * length (number of elements in slice)
	// * capacity (number of elements in underlying array, counting from the first element of slice)
	fmt.Printf("b2 = %v len(b2) = %v, cap(b2) = %v\n", b2, len(b2), cap(b2))

	// you can extend or shorten the slice as long as the underlying array allows it
	// extend
	// b2 = b2[:16] // NOT ALLOWED: underlying array is not large enough
	b2 = b2[:6]
	fmt.Printf("b2 = %v len(b2) = %v, cap(b2) = %v\n", b2, len(b2), cap(b2))

	// shorten
	b2 = b2[1:]
	fmt.Printf("b2 = %v len(b2) = %v, cap(b2) = %v\n", b2, len(b2), cap(b2))

	// slice can be also defined using slice literal. This in background creates an array that references it
	b3 := []int {0, -1, -2, -3}
	fmt.Printf("b3: %v %T, length = %v, capacity = %v\n", b3, b3, len(b3), cap(b3))

	// uninitialized slice is nil
	var b4 []int
	if (b4 == nil) {fmt.Println("b4 is nil")}
	fmt.Printf("b4 %v length = %v capacity = %v\n", b4, len(b4), cap(b4))
	var b5 []int = nil 
	if (b5 == nil) {fmt.Println("b5 is nil")}

	// slice of slices
	b6 := [][]int{
		[]int {1, 2, 3},
		[]int {4, 5, 6},
	}
	fmt.Printf("b6: %v length = %v capacity = %v type = %T\n", b6, len(b6), cap(b6), b6)

	// dynamic arrays
	// slices can be created with function make
	c1 := make([]int, 0, 5) // make(type, size, capacity) ... can also be (type, size), then capacity = size
	fmt.Printf("c1: %v length = %v capacity = %v type = %T\n", c1, len(c1), cap(c1), c1)

	// appending into arrays
	c2 := append(c1, 0, 10, 20, 30)
	fmt.Printf("c1: %v length = %v capacity = %v type = %T address = %v\n", c1, len(c1), cap(c1), c1, &c1)
	fmt.Printf("c2: %v length = %v capacity = %v type = %T address = %v\n", c2, len(c2), cap(c2), c2, &c2)

	// here appending causes increase of capacity, since the underlying array is not enough to hold the new data anymore
	c3 := append(c2, 40, 50)
	fmt.Printf("c1: %v length = %v capacity = %v type = %T address = %v\n", c1, len(c1), cap(c1), c1, &c1)
	fmt.Printf("c2: %v length = %v capacity = %v type = %T address = %v\n", c2, len(c2), cap(c2), c2, &c2)
	fmt.Printf("c3: %v length = %v capacity = %v type = %T address = %v\n", c3, len(c3), cap(c3), c3, &c3)

	// appending makes a copy of the array, mutating result does not mutate the original array, but only if resize occurs
	c3[1] = 999
	fmt.Printf("c2: %v length = %v capacity = %v type = %T address = %v\n", c2, len(c2), cap(c2), c2, &c2)
	fmt.Printf("c3: %v length = %v capacity = %v type = %T address = %v\n", c3, len(c3), cap(c3), c3, &c3)

	// if resize does not occur, original slice and appended slice share the underlying array
	c4 := append(c3, 60, 70)
	c4[1] = 555
	fmt.Printf("c3: %v length = %v capacity = %v type = %T address = %v\n", c3, len(c3), cap(c3), c3, &c3)
	fmt.Printf("c4: %v length = %v capacity = %v type = %T address = %v\n", c4, len(c4), cap(c4), c4, &c4)

	// looping over arrays with range construct
	d4 := []string{"a", "b", "c"}
	for idx, val := range d4 {
		fmt.Printf("Range loop 1 (index + value): d4[%v] = %v\n", idx, val)
	}

	for idx, _ := range d4 {
		fmt.Printf("Range loop 2 (index only): d4[%v] = ?\n", idx)
	}

	for idx := range d4 {
		fmt.Printf("Range loop 3 (index only, concise syntax): d4[%v] = ?\n", idx)
	}

	for _, val := range d4 {
		fmt.Printf("Range loop 4 (value only): d4[?] = %v\n", val)
	}
}

// SECTION: maps
func sectionMaps() {
	// map[keyType]valueType
	// default value is nil
	var a1 map[string]float64
	if a1 == nil {fmt.Println("a1 is nil")}
	var a2 map[string]float64 = nil
	if a2 == nil {fmt.Println("a2 is nil")}

	// you cannot assign to nil map
	// a1["pi"] = 3.14 // NOT ALLOWED: panic

	// to create an empty map we can use make function
	a1 = make(map[string]float64)
	// assign value to key
	a1["pi1"] = 3.14
	a1["e1"] = 2.7
	fmt.Println("a1", a1)

	// or we can assign a map literal
	a1 = map[string]float64{
		"pi2" : 3.14,
		"e2" : 2.7,
	}
	fmt.Println("a1", a1)

	// or we can assign empty map literal and then assign values
	a1 = map[string]float64{}
	// assign value to key
	a1["pi3"] = 3.14
	a1["e3"] = 2.7
	fmt.Println("a1", a1)

	// for maps with structure element, it is enough to use {} inside literals, we do not have to repeat name of structure
	a3 := map[string]MapElem{
		"a" : {-1, -2},
		"b" : {-3, -4},
	}
	fmt.Println("a3", a3)

	// element retrieval
	elem := a3["b"]
	fmt.Printf("Value of elem for b = %v\n", elem)

	// retrieving non-existing element will create it with default value, but does not modify original map
	elem = a3["c"]
	fmt.Printf("Value of elem for c = %v\n", elem)
	// now map still does not contain "c"
	fmt.Println("a3", a3)

	// you can add check parameter to test if elem exists in map
	var ok bool
	elem, ok = a3["b"]
	fmt.Printf("Value of elem for b = %v, exists? %v\n", elem, ok)
	elem, ok = a3["c"]
	fmt.Printf("Value of elem for c = %v, exists? %v\n", elem, ok)

	// upserting element into map
	a3["b"] = MapElem{999, 888} // update
	a3["c"] = MapElem{-5, -6} // insert
	fmt.Println("a3", a3)

	// deleting key from a map
	delete(a3, "a")
	delete(a3, "d") // deleting a non-existing key works without panic
	fmt.Println("a3", a3)

}

type MapElem struct {
	x int
	y int
}

// SECTION: methods
func sectionMethods(){
	// method is function with "receiver" argument that is defined using "type" keyword
	// receiver and method must be defined in the same package

	// syntax
	// type typeAlias typeDefinition // first define a type
	// func (receiverAlias receiverType) functionAlias(otherFunctionArgsAsUsual) returnType {...}

	// custom type can be alias for existing type
	a1 := MyFloat(3.14)
	// calling methods using dot syntax
	fmt.Printf("a1: %v type: %T, adding -0.5 1.1 %v\n", a1, a1, a1.addTwoNumbers(-0.5, 1.1))

	// custom type can also be structure
	a2 := Point2{3, 4}
	fmt.Printf("a2: %v type:%T, square = %v\n", a2, a2, a2.square())

	// when method has non-pointer receiver, then receiver is passed by value (a copy)
	// when method has a pointer receiver, then receiver is passed by reference
	fmt.Printf("a2 original: %v\n", a2)
	a2.shiftXBad(3) // this does not work as intended
	fmt.Printf("a2 after calling a2.shiftXBad(3) ... not working: %v\n", a2)
	
	(&a2).shiftXGood(3) // this works as intended
	fmt.Printf("a2 after calling (&a2).shiftXGood(3) ... working: %v\n", a2)

	// with pointer receivers, also simplified syntax of calling is allowed
	a2.shiftXGood(3) // this works as intended
	fmt.Printf("a2 after calling a2.shiftXGood(3) ... still working: %v\n", a2)

	// also, methods with value receivers can be applied to pointers without using indirection operator
	a3 := &a2
	fmt.Printf("a3: %v type:%T, square = %v\n", a3, a3, a3.square())
}

type MyFloat float64

func (f MyFloat) addTwoNumbers(x1 MyFloat, x2 MyFloat) MyFloat {return f + x1 + x2}

type Point2 struct {
	x int
	y int
}

func (p Point2) square() int {return p.x * p.x + p.y * p.y}


func (p Point2) shiftXBad(amount int)() {p.x += amount}
func (p *Point2) shiftXGood(amount int)() {p.x += amount}

// SECTIONL: interfaces
func sectionInterfaces() {
	// type T i said to have interface I, if it implements set of methods required by that interface
	// interfaces are (value, type) pairs under the hood
	a1 := CustomString("Whee")
	var a2 MyInterface = a1
	fmt.Printf("a2: %v, %T\n", a2, a2) // type is actually still main.CustomString when printed
	a2.MyMethod1()
	a3 := a2.MyMethod2(11, "Glee")
	fmt.Printf("Result of method 2 was %v, %T\n", a3, a3)

	// NOT ALLOWED: CustomString2 has one method with value receiver and the other with pointer receiver
	// thus it does not conform to the MyInterface (all methods have to either have value receivers or all have pointer ones)
	// a4 := CustomString2("Bree")
	// a2 = a4 // throws error

	// this works again
	a5 := CustomString3("Jaja")
	// a2 = a5 // NOT ALLOWED: must assign address of a5 (pointer)
	a2 = &a5
	fmt.Printf("a2: %v, %T\n", a2, a2)
	a2.MyMethod1()
	a6 := a2.MyMethod2(11, "Glee")
	fmt.Printf("Result of method 2 was %v, %T\n", a6, a6)

	// nil values: 
	// interface without assigned value
	var a7 MyInterface;
	fmt.Printf("a7: %v, %T\n", a7, a7) // nil of type nil
	// a7.MyMethod1() // NOT ALLOWED: panic, go does not know which method to call
	a7 = nil // nil is actually default value for interface
	fmt.Printf("a7: %v, %T\n", a7, a7) // nil of type nil

	// interface with nil value and not-nil type is allowed, but receiver must support handling it (cannot go into panic)
	var a8 *CustomString3
	a7 = a8
	a7.MyMethod1()

	helperEmptyLine()
	// empty interface - requires implementation of 0 methods
	// can be used to hold object of any value
	var a9 AnyValue
	fmt.Printf("a9 -- value:%v, type: %T\n", a9, a9)
	a9 = 1
	fmt.Printf("a9 -- value:%v, type: %T\n", a9, a9)
	a9 = 1.4
	fmt.Printf("a9 -- value:%v, type: %T\n", a9, a9)
	a9 = "Whee"
	fmt.Printf("a9 -- value:%v, type: %T\n", a9, a9)
	// you can create lists or maps (values not keys) with non-homogeneous data types
	a9 = []AnyValue {
		9,
		1.4,
		"Whee",
		map[string]AnyValue {
			"x" : 9,
			"y" : "Glee",
		},
	}
	fmt.Printf("a9 -- value:%v, type: %T\n", a9, a9)
	
	// type switches
	switch v := a9.(type) {
	case float64:
		fmt.Printf("Interface has type float64, value is %v, type is %T\n", v, v)
	case int:
		fmt.Printf("Interface has type int, value is %v, type is %T\n", v, v)
	case []AnyValue:
		fmt.Printf("Interface has type []AnyValue, value is %v, type is %T\n", v, v)
	default:
		fmt.Printf("Interface has some other type, value is %v, type is %T\n", v, v)
	}

	// interface type
	// iType := a9.(type) // NOT ALLOWED: this syntax does not work outside of type switch

	// interface{} keyword in function argument - makes function accept empty interface (any argument)
	functionOnInterfaces := func (i interface{})(){fmt.Printf("functionOnInterfaces: We got %v of type %T\n", i, i)}
	functionOnInterfaces(a9)
	var a10 int = 3
	functionOnInterfaces(a10)
	functionOnInterfaces(10)

	// well-known interfaces 

	// fmt.Stringer interface
	// type Stringer interface {
	// 	String() string
	// }
	p1 := Person{"Harry", "Potter", 23, "M"}
	fmt.Println(p1) // prints normally (not a pointer)
	fmt.Println(&p1) // prints using custom function

	//p2 := fmt.Stringer(p1) // NOT ALLOWED: *Person is stringer, not Person
	p3 := fmt.Stringer(&p1)
	fmt.Println(p3)

	// error interface
	// type error interface {
	// 	Error() string
	// }
	e1 := MyError{-1, "Something went wrong"}
	fmt.Println(e1)
	fmt.Println(&e1) // this prints Error() as well ... weird
	e2 := error(e1)
	fmt.Println(e2)
	e3 := error(&e1) // this works too ... weird
	fmt.Println(e3)

	// io.Reader interface
	// func (T) Read(b []byte) (n int, err error)
	// populates slice "b" with data
	// returns "n" number of bytes populated
	// when stream ends, returns io.EOF error
	reader := strings.NewReader("ABC")

	// note, how buffer is not resized even if nBytes < 2
	buffer := make([]byte, 2)
	for {
		nBytes, err := reader.Read(buffer)
		fmt.Printf("nBytes = %v err = %v buffer = %v, len(buffer) = %v, cap(buffer) = %v\n", nBytes, err, buffer, len(buffer), cap(buffer))
		fmt.Printf("b[:nBytes] = %q\n", buffer[:nBytes])
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("io.EOF: value = %v, type = %T\n", io.EOF, io.EOF)

	helperEmptyLine()

	// image.Image interface
	// type Image interface {
	// 	ColorModel() color.Model
	// 	Bounds() Rectangle
	// 	At(x, y int) color.Color
	// }

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0, 0).RGBA())
}

type MyInterface interface {
	MyMethod1() ()
	MyMethod2(int, string) float64
}

// implements MyInterface -- both receivers are value-based
type CustomString string
func (s CustomString) MyMethod1() () {fmt.Printf("Calling MyMethod1 on %v, %T\n", s, s)}
func (s CustomString) MyMethod2(x int, name string) float64 {
	result := float64(x)
	fmt.Printf("Calling MyMethod2 on %v, %T with args (%v, %v) and returning %v\n", s, s, x, name, result)
	return result
} 

// does not implement MyInterface - one receiver is value-based and one pointer-based
type CustomString2 string
func (s CustomString2) MyMethod1() () {fmt.Printf("Calling MyMethod1 on %v, %T\n", s, s)}
func (s *CustomString2) MyMethod2(x int, name string) float64 {
	result := float64(x)
	fmt.Printf("Calling MyMethod2 on %v, %T with args (%v, %v) and returning %v\n", s, s, x, name, result)
	return result
} 

// implements MyInterface -- both receivers are pointer-based
type CustomString3 string
func (s *CustomString3) MyMethod1() () {fmt.Printf("Calling MyMethod1 on %v, %T\n", s, s)}
func (s *CustomString3) MyMethod2(x int, name string) float64 {
	result := float64(x)
	fmt.Printf("Calling MyMethod2 on %v, %T with args (%v, %v) and returning %v\n", s, s, x, name, result)
	return result
} 

// empty interface
type AnyValue interface {}

// structure that implements Stringer
type Person struct {
	name string
	surname string
	age int
	sex string
}

func (p *Person) String() string {
	var title string
	switch p.sex{
	case "M":
		title = "Mr."
	case "F":
		title = "Ms."
	default:
		title = "Oth."
	} 
	return fmt.Sprintf("%v %v %v (%v years old)", title, p.name, p.surname, p.age)
}

// structure that implements error
type MyError struct {
	code int
	description string
}

func (e MyError) Error() string {return fmt.Sprintf("ERROR [%v]: %v", e.code , e.description)}

// SECTION: generics
func sectionGenerics(){
	// generic functions
	
	// syntax: 
	// func functionIdentifier[typeIdentifier constraint, ...](arguments) returnType
	// arguments and return type can use typeIdentifier as generic type
	// constraint argument is required. Use "any" to indicate "no constraint"

	a1 := []string{"a", "b", "c"}
	a2 := []int{1, 2 , 3}

	fmt.Printf("Index of c in a1 is: %v\n", Index(a1, "c"))
	fmt.Printf("Index of 2 in a2 is: %v\n", Index(a2, 2))

	// NOT ALLOWED: generics cannot be defined inside other functions like this

	// InsideIndex := func [T comparable](s []T, x T) int {
	// 	for i, v := range s {
	// 		// v and x are type T, which has the comparable
	// 		// constraint, so we can use == here.
	// 		if v == x {
	// 			return i
	// 		}
	// 	}
	// 	return -1
	// }

	// fmt.Printf("InsideIndex of c in a1 is: %v\n", InsideIndex(a1, "c"))
	// fmt.Printf("InsideIndex of 2 in a2 is: %v\n", InsideIndex(a2, 2))

	a3 := MakeEmpty[string]() // must specifically use [...] syntax, otherwise go cannot infer [T]
	fmt.Println("MakeEmpty[string]()", a3)

	fmt.Printf("%v\n", MakeCopy("bla"))
	fmt.Printf("%v\n", MakeCopy(9))

	// generic structures

	// syntax
	// type identifier[typeIdentifier constraint] struct {structure definition}
	// structure definiton may use type declared in []

	b1 := ValueWrapper[int]{1, "This holds int"}
	b2 := ValueWrapper[float64]{9.35, "This holds float64"}
	fmt.Println(b1, b2)
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func MakeEmpty[T comparable]() T {
	var result T
	return result
}


// this works
func MakeCopy[T any](val T) T {
	return val
}

type ValueWrapper[T any] struct {
	val  T
	description string
}

// SECTION: concurrency
func sectionConcurrency() {
	// goroutine - lightweight thread managed by go
	// arguments evaluation happens in current goroutine
	// body evaluation happens in new goroutine
	// call goroutine by putting "go" before function call, like: "go functionName(...)"

	createName := func (name string) string {
		fmt.Printf("Evaluating createName %q\n", name)
		return name
	}

	someFunction := func (name string) () {
		for i := 0; i < 5; i++ {
			fmt.Printf("someFunction[%v] i = %v\n", name, i)
		}
	}

	fmt.Println("Before starting someFunction")
	go someFunction(createName("Bob")) // call first goroutine ... opened in another thread independent from main thread
	go someFunction(createName("Alice")) // call second grorutine ... opened in another thread independent from main thread or Bob
	
	waitMilliseconds := int64(200)
	fmt.Printf("After starting someFunction. Now wait %v seconds so that goroutines complete\n", float32(waitMilliseconds) / 1000)

	// wait a while so that goroutines finish ... this is dirty way to do this just for sake of example
	time.Sleep(time.Duration(waitMilliseconds) * time.Millisecond)
	fmt.Printf("After %v second\n", float32(waitMilliseconds) / 1000)

	// channels are queue - like structures that allow sending data between threads
	// by default sends and receives block, until the other side is ready
	// channel type is "chan T", where T is type of data that will be sent through channels

	helperEmptyLine()

	producerGoroutine := func (channel chan int, nValues int, name string) () {
		for i:= 0; i < nValues; i++ {
			val := 10 * i
			fmt.Printf("[%v] Start pushing %v\n", name, val)
			channel <- (10 * i) // operator of writing into channel
			fmt.Printf("[%v] End pushing %v\n", name, val)
		}
	}

	channel := make(chan int)
	fmt.Printf("Channel value: %v, type: %T\n", channel, channel)

	nValues := 5
	// trigger goroutine
	// note, here output will be kinda mixed, we would have to make concurrency lock to include 
	// print statements as well, but that would deny purpose of the channel
	go producerGoroutine(channel, nValues, "producer 1")
	for i:= 0; i < nValues; i++ {
		fmt.Println("Pop next value")
		val := <-channel
		fmt.Printf("Value popped %v\n", val)
	}

	// reads from empty channel block until something fills it
	// go will detect this as deadlock (set if statement to "true" if you want to see it)
	wantSeeEmptyChannelDeadlock := false
	if wantSeeEmptyChannelDeadlock {
		emptyChannel := make(chan int)
		fmt.Println("Now we will wait indefinitely (channel is empty)")
		val := <-emptyChannel
		fmt.Println("This line will never execute, empty channel blocks, waiting for val", val)
	}

	// channel can be buffered to accept only maxN values
	// writes to full channel will block until something is popped
	// code below would block indefinitely when writing "5" into buffered channel and go would detect deadlock
	// set if statement to true, if you want to see this

	wantSeeBufferedChannelDeadlock := false
	if wantSeeBufferedChannelDeadlock {
		bufferedChannel := make(chan int, 5)
		for i:= 0; i < 10; i++ {
			fmt.Printf("Start writing %v to bufferedChannel\n", i)
			bufferedChannel <- i
			fmt.Printf("End writing %v to bufferedChannel\n", i)
		}
	}
	
	helperEmptyLine()

	// channels can be closed by sender to indicate no more values are coming
	producerGoroutineWithClosing := func (channel chan int, nValues int, name string) () {
		for i:= 0; i < nValues; i++ {
			val := 10 * i
			fmt.Printf("[%v] Start pushing %v\n", name, val)
			channel <- (10 * i) // operator of writing into channel
			fmt.Printf("[%v] End pushing %v\n", name, val)
		}
		close(channel) // close the channel
	}

	channel2 := make(chan int)
	// run goroutine that will close the channel in the end
	go producerGoroutineWithClosing(channel2, nValues, "producer 2")
	
	// on read side we read in an infinite loop and check channel for closing until the channel is closed
	for {
		fmt.Println("Pop next value")
		val, ok := <-channel2
		fmt.Printf("Value popped %v, ok = %v\n", val, ok)
		if !ok {
			// break out of infinite loop on channel closing
			fmt.Println("Channel2 is closed, terminating read operations.")
			break
		}
	}

	helperEmptyLine()
	// also, same thing can be more compactly writen using range loop construct
	
	channel3 := make(chan int)
	// run goroutine that will close the channel in the end
	go producerGoroutineWithClosing(channel3, nValues, "producer 3")
	
	for val := range channel3 {
		fmt.Printf("Value popped %v\n", val)
	}
	fmt.Println("Channel3 was closed.")

	helperEmptyLine()
	// alternately, "select" statement can be used
	// this is like "switch", but it blocks, until one of its cases can run

	producerGoroutineWithQuitChannel := func (numberChannel, quitChannel chan int, nValues int, name string) () {
		for i:= 0; i < nValues; i++ {
			val := 10 * i
			fmt.Printf("[%v] Start pushing %v\n", name, val)
			numberChannel <- (10 * i) // normal output is writen to numberChannel
			fmt.Printf("[%v] End pushing %v\n", name, val)
		}
		quitChannel <- 0 // on finishing, 0 is writen to quitChannel
	}

	numberChannel := make(chan int)
	quitChannel := make(chan int)

	// run producer goroutine
	go producerGoroutineWithQuitChannel(numberChannel, quitChannel, nValues, "producer 4")

	// now reader code will loop until quitChannel receives input
	productionFinished := false
	for !productionFinished {
		select {
		case val := <- numberChannel:
			fmt.Printf("Value popped %v\n", val)
		case <- quitChannel:
			fmt.Println("Quit signal received. Terminating loop")
			productionFinished = true
		// default: // if default clause is present, "select" statement runs it instead of blocking, if no case is ready
		// 	fmt.Println("No other case was ready at this moment")
		}
	}

	helperEmptyLine()
	// note sending into unbuffered channel when 1 thread only is available results in blocking on first iteration, because for 
	// unbuffered channels, reader and writer must be ready at the same time
	
	// singleThreadChannel := make(chan int) // this would block due to unbuffered channel
	// singleThreadChannel := make(chan int, 3) // this would block due to overflowing channel capacity
	singleThreadChannel := make(chan int, 10) // this does not block
	for i:= 0; i < nValues; i++ {
		val := 10 * i
		fmt.Printf("[%v] Start pushing %v\n", "producer 5", val)
		singleThreadChannel <- (10 * i) // normal output is writen to numberChannel
		fmt.Printf("[%v] End pushing %v\n", "producer 5", val)
	}

	helperEmptyLine()
	// sync.Mutex -- mutually exclusive locks

	q1 := MyQueue{sync.Mutex{}, make([]int, 0), "MyQueue 1"}
	fmt.Printf("q1 %v, type: %T\n", q1, q1)

	producerGorutineWithMyQueue := func (queue *MyQueue, nValues int) () {
		for i:= 0; i < nValues; i++ {
			queue.push(10 * i)
		}
	}
	
	// start producer goroutine in another thread
	go producerGorutineWithMyQueue(&q1, nValues)
	
	counter := 0
	for counter < nValues {
		val, ok := q1.pop()
		if ok {
			fmt.Printf("Processing val %v", val)
			counter++
		} else {
			fmt.Printf("Queue is empty ... try again later\n")
		}
	}

}

type MyQueue struct {
	// simple implementation of queue that uses mutex to secure against concurent read, write operations
	mu sync.Mutex
	queue []int
	name string
}

func (q *MyQueue) push(val int) {
	// if any other function comes into mu.lock() that thread blocks until mu.unlock() is called
	q.mu.Lock()
	fmt.Printf("[%v] Start pushing %v\n", q.name, val)
	q.queue = append(q.queue, val)
	fmt.Printf("[%v] End pushing %v\n", q.name, val)
	q.mu.Unlock()
}

func (q *MyQueue) pop() (val int, success bool) {
	// if any other function comes into mu.lock() that thread blocks until mu.unlock() is called
	q.mu.Lock()
	fmt.Printf("[%v] Start popping\n", q.name)
	defer q.mu.Unlock() // unlock happens only after returning
	if len(q.queue) > 0 {
		val = q.queue[0]
		q.queue = q.queue[1:]
		fmt.Printf("[%v] End popping success, val=%v \n", q.name, val)
		return val, true // return value and indication that pop() was successful
	} else {
		fmt.Printf("[%v] End popping failure. Queue is empty \n", q.name)
		return 0, false // return default value and indication that pop() was unsuccesful
	}
}

// APPENDIX: helper functions 

func helperRunSection(section func (), sectionTitle string) {
	fmt.Println("__________________________________________________________________\n")
	fmt.Println("***** SECTION:", sectionTitle, "*******************************\n")
	section()
	fmt.Println("")
}

func helperWelcome() {
	helperEmptyLine()
	fmt.Println("*********************************************************************")
	fmt.Println("************** Welcome to go basic syntax starterpack ***************")
	fmt.Println("*********************************************************************")
	helperEmptyLine()
}

func helperEmptyLine() {fmt.Println()}