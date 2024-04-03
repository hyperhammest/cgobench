package cgobench

/*
#include <time.h>
int trivial_add(int a, int b) {
  return a+b;
}
extern int MyFunction(int a, int b);
int trivial_add2(int a, int b) {
  return MyFunction(a, b);
}
*/
import "C"
func Call() {
	// I don't do much. Yet.
}
// wow this is easy
// import "C"
func CgoCall() {
	C.trivial_add(1,2)
}
func CgoCall2() {
	C.trivial_add2(1,2)
}
