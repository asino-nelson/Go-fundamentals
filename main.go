package main
import "fmt"


func main (){

  // DATA TYPES AND VARIABLES

  /*
  fmt.Println("Hello world")

  var intNum int32 = 100
  fmt.Println(intNum)

  var floatNum float64 = 100.0
  fmt.Println(floatNum)

  var intNum1 int = 3
  var intNum2 int = 2
  fmt.Println(intNum1 + intNum2)
  fmt.Println(intNum1 - intNum2)
  fmt.Println(intNum1 * intNum2)
  fmt.Println(intNum1 / intNum2)
  fmt.Println(intNum1 % intNum2)

  var myString string = "Hello Nelson"
  fmt.Println(myString)

  var myBool bool = true
  fmt.Println(myBool)

  myVar := "Milk"
  fmt.Println(myVar)
  */

  // FUNCTIONS

  /*
  sentence := "This is the print value"

  PrintMe(sentence)
  answer, remainder := intDivision(100,9)
  fmt.Println(answer)
  fmt.Printf("The result is %v with remainder %v", answer, remainder)
*/

// ARRAYS
/*
  var intArr = [3] int {1,2,3}
  fmt.Println(intArr[0])
  fmt.Println(intArr[1:3])
  fmt.Println(intArr)

  fruits := [...]string{"Apple","Mango","Kiwi"}
  fmt.Println(fruits)
*/

// SLICES
/*
  var intSlice []int = []int{4,5,6}
  fmt.Println(intSlice)
  fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))

  intSlice = append(intSlice,7)
  fmt.Println(intSlice)
  fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))

  var intSlice2 []int = []int{8,9}
  intSlice = append(intSlice, intSlice2 ...)
  fmt.Println(intSlice)

  var intSlice3 []int = make([]int, 3,8) // Weve specified length 3 and capcity 8
  fmt.Println(intSlice3)
*/
//go run .\main.go

// MAPS
var myMap map[string]string = make(map[string]string)
fmt.Println(myMap)

var myMap2 = map[string]uint{"Nelson":21, "Ruth":44}
fmt.Println(myMap2)
fmt.Println(myMap2["Nelson"])





}


// FUNCTIONS

/*
func PrintMe(printValue string){
  fmt.Println(printValue)
}

func intDivision(numerator int, denominator int)(int, int){
  result := numerator/denominator
  remainder := numerator % denominator
  return result, remainder
}
*/