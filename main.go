package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {

	// ARRAYS - Fixed length. When assigned to another variable it makes an entirely new copy of the array

	// Defines array of 4 int values - [0 0 0 0]
	var arr [4]int

	arr1 := [4]int{0, 1, 2, 3}

	arr[1] = 10

	fmt.Println(arr, arr1)

	var twoDimArr [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoDimArr[i][j] = i + j
		}
	}

	fmt.Println(twoDimArr)

	// SLICES - Passed by reference not a full copy as in an array

	var sl []int
	fmt.Println(sl, sl == nil)

	//make(Type, len, capacity(optional))
	sl1 := make([]int, 4)
	fmt.Println(sl1, sl1 == nil)

	sl2 := []int{0, 1, 2, 3}
	sl2[1] = 10
	fmt.Println(sl2, len(sl2))

	// append([]T, element1, element2)
	// adding one slice to another slice. Use of ellipses required when adding a slice
	// append([]T, []T...)
	sl2 = append(sl2, 4, 5, 6)
	fmt.Println(sl2, len(sl2))

	slAppend := []int{7, 8, 9}
	sl2 = append(sl2, slAppend...)
	fmt.Println(sl2, len(sl2))

	// copy(dst, src []T) int
	// copies from source slice to destiation slice returns an int, which is the number of elements copied
	slCopy := make([]int, len(sl2))
	el := copy(slCopy, sl2)
	fmt.Println(len(slCopy))
	fmt.Println(slCopy)
	fmt.Println(el)

	// []T[low:high]
	a := sl2[2:4] // gives elements at index 2 and 3
	b := sl2[:5]  // gives elements at index 0 to 4
	c := sl2[3:]  // gives elements at index 3 onwards
	fmt.Println(a, b, c)

	// MAPS - Unordered key-value pairs
	// var name map[type for key]type for value
	var prodPrice map[string]int
	// declare and initialise with make
	tempPrice := make(map[string]int)
	tempPrice["convertible widget"] = 150
	prodPrice = tempPrice
	prodPrice["widget"] = 100
	fmt.Println(prodPrice)

	// declae and initialise using map literal
	empPrice := map[string]int{
		"widget": 75,
	}
	empPrice["turbo widget"] = 100
	fmt.Println(empPrice)

	// check that a key is present
	el, ok := empPrice["widget"]
	fmt.Println(el, ok)

	//delete
	delete(empPrice, "widget")

	// POINTERS - Stores the memory address of a value rather than the value itself
	// address of operaror = &
	// & creates a pointer
	// * (pointer operator/dereferencing operator) used to define a pointer variable and to reveal a pointer's underlying value

	// var <name> *T
	val := 123
	var ptr *int = &val // address of val
	fmt.Println(ptr, *ptr)

	// var <name> *T = new(T)
	var ptr1 *int = new(int)
	fmt.Println(ptr1, *ptr1)
	*ptr1 = 345
	fmt.Println(ptr1, *ptr1)

	ptr2 := &val
	fmt.Println(ptr2, *ptr2)

	var ptr3 = &val
	fmt.Println(ptr3, *ptr3)

	// FLOW CONTROL
	// For loop - init statement, condition expression, post statement
	sum := 0
	for i := 0; i < 5; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// The init and post statement are both optional
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Can have break or continue statements within for loops
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// goto statements
	// Transfers control to a labeled statement, within same function

	// infinite for loop
	var j int
	for {
		if j == 3 {
			break
		}
		fmt.Println(j)
		j++
	}

	// If statements
	// can inialise temp within if statement - only in scope until end of if statement
	if temp := -10; temp < 0 {
		fmt.Println("Below freezing!")
	} else if temp == 0 {
		fmt.Println("At freezing point")
	} else {
		fmt.Println("Above freezing")
	}

	// Switch statement
	// Don't need break clauses
	workday := 3
	switch workday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other")
	}

	temp := -10
	switch {
	case temp < 0:
		fmt.Println("Below freezing!")
	case temp == 0:
		fmt.Println("At freezing!")
	default:
		fmt.Println("Above freezing!")
	}

	// Defer statement
	// Defers the execution of a function until the surrounding function, in which it was called, returns
	// Arguments are evaluated immediately but the function is only called after the parent function returns
	defer fmt.Println("second")
	fmt.Println("first")

	// looping over collections
	slice := []int{10, 20, 30, 40}
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}

	// easier to use range to iterate over collection
	for i, value := range slice {
		fmt.Println(i, value)
	}

	prodPrice1 := map[string]int{
		"widget":       75,
		"turbo widget": 100,
		"smol widget":  150,
	}
	for key, value := range prodPrice1 {
		fmt.Println(key, value)
	}

	for key := range prodPrice1 {
		fmt.Println(key)
	}

	for _, value := range prodPrice1 {
		fmt.Println(value)
	}

	fmt.Println(strings.Contains("Working with string functions", "functions"))
	// if less than 0 no limit on number of times to replace
	fmt.Println(strings.Replace("Working with string functions", "functions", "variables", -1))

	fmt.Println(strings.Title("Working with string functions"))
	fmt.Println(strings.Trim("___Working with string functions___", "_"))

	// can include the field names to use struct literal to specify which field we're setting, or don't use but have to input in the correct order
	cube := Cube{d: 4, w: 4, h: 4}
	cube.d = 6
	fmt.Println("volume: ", cube.volume())

	t := Tire{Part{Manufacturer: "Brocadero"}}
	fmt.Println("Manufactured by", t.Part.Mfc())

	t1 := Tire1{Part{Manufacturer: "Brocadero"}}
	fmt.Println("Manufactured by", t1.Mfc())

	sp := Sphere{radius: 4}

	fmt.Println(totalVolume(&cube, &sp))

	x := 20
	y := 10
	fmt.Println(add(x, y))

	// For addVals - Pass arguments as variadic parameters
	// Can pass slice followed by ... as the unpack operator, to unpack the contents of a slice
	vals := []int{1, 2, 3, 4}
	addVals(vals...)

	addVals(1, 2, 3, 4, 5, 6)

	radius := 1.0
	vol, err := volume(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Volume of sphere is %0.2f\n", vol)

	// Defer is a process that places a function call onto a stack. Each deferred function is
	// executed in reverse order when the calling function completes whether a panic is called or not
	// defer r()
	// p("runtime error: enter panic state", 3)

	// FILE OPERATIONS

	// Create file
	// func Create(name string) (*File, error)
	// f, err := os.Create("create.txt")
	// defer f.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(f)

	// // Open and close file
	// // custom closer() function
	// // func Open(name string) (*File, error)
	// f1, err := os.Open("file.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("File successfully opened:", f1.Name())
	// closer(f1)

	// // Delete file
	// // func Remove(name string) error
	// // err1 := os.Remove("del.txt")
	// // if err1 != nil {
	// // 	log.Fatal(err1)
	// // }
	// // fmt.Println("file removed")

	// // Copy file
	// // First open source file
	// src, err := os.Open("src.txt")
	// defer src.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // the flag allows the program to create it, if the file does not exist
	// // func OpenFile(name string, flag int, perm FileMode) (*File, error)
	// dst, err := os.OpenFile("dst.txt", os.O_RDWR|os.O_CREATE, 0755)
	// defer dst.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // func Copy(dst Writer, src Reader) (written int64, err error)
	// w, err := io.Copy(dst, src)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(reflect.TypeOf(w))
	// fmt.Println(w)

	// // Rename a file
	// oldPath := "file.txt"
	// newPath := "new.txt"
	// // func Rename(oldpath, newpath string) error
	// err2 := os.Rename(oldPath, newPath)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// Truncate a file
	// func Truncate(name string, size int64) error

	// File info
	// os.Stat() returns a FileInfo interface that describes a file
	// type FileInfo interface {
	// 	Name() string
	// 	Size() int64   		// length in bytes
	// 	Mode() FileMpdea	// file mode bits, permissions
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool 		// Mode().IsDir()
	// 	Sys() interface{}	// underlying data source
	// }

	// READING FILES INTO MEMORY

	// func ReadFile(name string) ([]byte, error)
	contents, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(contents))
	fmt.Println(contents)
	// type cast slice of uint8 values (the slice of bytes) in "contents" to string type
	fmt.Println(string(contents))

	// READINF FILES LINE BY LINE
	f, err := os.Open("file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Reader implements buffering for a Reader object
	// Scanner implements an interface for reading data such lines of text in a file with line breaks

	// func NewScanner(r io.Reader) *Scanner
	s := bufio.NewScanner(f)
	// func (s *Scanner) Scan() bool -- returns false when finished scanning
	for s.Scan() {
		// calling the Text method on the scanner - scan scans through the text, accessing each token and makes it available via the Text method
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// PERFORMING DIRECTORY I/O OPERATIONS
	// MkdirALL Create directory
	// Stat() returns FileInfo which implements methods that describes a file
	d, err := os.Stat("subdir")
	fmt.Println("error returned by os.Stat() is:", err)
	// if err is nil then os.Stat() found a file/directory with the same name
	if err == nil {
		fmt.Println(d.Mode(), d.IsDir())
		// log.Fatal("file/directory name already exists")
	}
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("subdir", 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("subdir directory created")
	}

	// MkdirAll Nest directories
	// func Join(elem ...string) string
	p := filepath.Join("./test", "subdir1", "subdir2")
	err1 := os.MkdirAll(p, 0777)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(p, "nested directory created")

	// ReadDir List Files
	// func ReadDir(name string) ([]Directory, error)
	ls, err := os.ReadDir("../go-basics")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range ls {
		fmt.Println(f.Name(), f.IsDir())
	}

}

// FUNCTIONS

// METHODS - Special type of function that includes a receiver
// Allows us to use the . operator when calling the function to access the attributes from the cube object

func (c *Cube) volume() float64 {
	return float64(c.d * c.w * c.h)
}

type Cube struct {
	d float64
	w float64
	h float64
}

// Can have multiple return values and these can be named
// Don't need to include the return values in the return statement but can if we want
func add(x, y int) (a, b int, c bool) {
	a = x + y
	b = x - y
	c = x > y
	return
}

// Variadic functions
// ... is the pack operator - Used to declare variadic function, because the arguments are packed into a slice when they are passed to the function
func addVals(vals ...int) {
	total := 0
	for _, val := range vals {
		total += val
	}
	fmt.Println(vals, total)
}

// Inheritance - Need to use embedded types to achieve this
// Objects defined by other objects
// The tire is a part so inherits the fields from the Part struct
// However this is a problems because it implies the tire has a Part
type Part struct {
	Manufacturer string
}

func (p *Part) Mfc() string {
	return p.Manufacturer
}

type Tire struct {
	Part Part
}

// Instead can use anonymous fields
// Provide the type Part but don't supply a name. So can call on any Part methods directly on the Tire
type Tire1 struct {
	Part // anonymous field
}

// INTERFACES
// Define a method set, a list of methods that a type needs in order to implement the interface
// Both cube and sphere have volume methods that return float64 values
// Therefore both types are said to implement the shape interface
type Shape interface {
	volume() float64
}

// Can how specify interface types as function arguments
func totalVolume(shapes ...Shape) float64 {
	// Here the volume method defined for each respective shape is called by ranging over the shapes in a for loop
	var volume float64
	for _, s := range shapes {
		volume += s.volume()
	}
	return volume
}

type Sphere struct {
	radius float64
}

func (sp *Sphere) volume() float64 {
	return ((4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius))
}

// ERROR HANDLING
// Compile time errors = Lexical, syntactical, semantical, logical
// Runtime errors = Signal abort, non-zero exit code, segmentation fault, floating point execptions

// error is an interface type with one method
// type error interface {
//Error() string
// }

// Custom errors
// errors.New() function which takes a string value as its only parameter

func volume(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("Radius negative")
	}
	return ((4.0 / 3.0) * math.Pi * r * r * r), nil
}

// Defer, panic, recover only used in exceptional cases
// Defer places a function call onto a defer stack
// Panic terminates the natural flow of execution
// Recover is a mechanism to recover the execution from terminating

// func panic (v interface{}) // empty interface can basically be any value
// func recover() interface{}

func p(s string, i int) {
	// x := [3]int{1, 2, 3}
	// x[i] = 11
	if i > 2 {
		panic(s)
	}
}

func r() {
	// the value returned by a recover function call is the value passed when the panic function call was made or if a real runtime error, the error condition
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("Recovered from panic")
	}
}

// errors.Is(err, target error) bool
// checks if error produced matches a known error
// if _, err := os.Open("myFile.txt"); err != nil {
// if errors.Is(err, os.ErrNotExist) {
// log.Println("file does not exist")
// 	} else {
// 	log.Println(err)
// }
// return
// }

// FILE OPERATIONS

func closer(f *os.File) error {
	// func (f *file) Close() error
	f.Close()
	fmt.Println(f.Name(), "successfully closed")
	return nil
}
