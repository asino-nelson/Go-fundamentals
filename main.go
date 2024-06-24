package main
import "fmt"


func main (){
  // fmt.Println("Hello world")

  // var intNum int32 = 100
  // fmt.Println(intNum)

  // var floatNum float64 = 100.0
  // fmt.Println(floatNum)

  // var intNum1 int = 3
  // var intNum2 int = 2
  // fmt.Println(intNum1 + intNum2)
  // fmt.Println(intNum1 - intNum2)
  // fmt.Println(intNum1 * intNum2)
  // fmt.Println(intNum1 / intNum2)
  // fmt.Println(intNum1 % intNum2)


  // var myString string = "Hello Nelson"
  // fmt.Println(myString)

  // var myBool bool = true
  // fmt.Println(myBool)

  // myVar := "Milk"
  // fmt.Println(myVar)

  sentence := "This is the print value"

  PrintMe(sentence)
  answer, remainder := intDivision(100,9)
  fmt.Println(answer)
  fmt.Printf("The result is %v with remainder %v", answer, remainder)

}

func PrintMe(printValue string){
  fmt.Println(printValue)
}

func intDivision(numerator int, denominator int)(int, int){
  result := numerator/denominator
  remainder := numerator % denominator
  return result, remainder
}