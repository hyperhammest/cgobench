package cgobench

import "C"
//export MyFunction
func MyFunction(a, b C.int) C.int {
   return a+b
}
