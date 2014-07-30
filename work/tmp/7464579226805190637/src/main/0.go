package main
import "a"
import "b"
import "runtime"
func init() {
	go func() {
		for {
			runtime.GC()
			runtime.Gosched()
		}
	}()
}
var _ = b.UsePackage
var _ = a.UsePackage
var UsePackage = 0
var SINK interface{}
func main() {
select {
case ((*(a.Var85))(a.Func86("foo",1.0,"foo",float32(1.0)),[][1]bool{}))[(<- make(chan int ))](uintptr(0),complex64(1i),"foo") <- make(map[chan float64]chan byte ):
(([]chan float64{})[((*(b.Var89)))[(int)((len)(((([][1][0][]byte{(<- make(chan [1][0][]byte ,1)),[1][0][]byte{0: [0][]byte{}}})[(([]int{})[b.Var182]) + (1)])[((func([]uintptr,complex128,string) int)(nil))([]uintptr{},1i,"foo")])[(*((*((([][1]**int{})[(([]int{})[(int)(Var223 )]) + (1)])[(([][2]int{})[(<- make(chan int ,1))])[b.Var966]]))))]))]:])[(([][0]int{})[(int)((*(Var994)) ,)])[((Var995)[((func(map[complex64][]uintptr,[]uint,*[1]int16,struct { Field1025 *int16
}) int)(nil))(make(map[complex64][]uintptr ,1),[]uint{},(*[1]int16)(nil),(struct { Field1025 *int16
}{}))]) + (1)]] <- 1.0
case make(chan complex128 ,1) <- 1i:
case make(chan uint ,1) <- uint(1):
case ([]int{})[((func(map[*int16]map[interface{}]int,map[struct { }]bool,int16,[2]chan error) int)(nil))(make(map[*int16]map[interface{}]int ),make(map[struct { }]bool ),int16(1),[2]chan error{})], ([]bool{})[(int)(Var1052 )] = <-make(chan int ,1):
case make(chan interface { } ,1) <- interface { }(nil):
}
return 
}
func init() {
return 
}
