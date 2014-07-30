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
Label1:
goto Label1
var Var2 error = (*((*((a.Var27)[func(Param28 map[bool]func(complex64,rune) (uintptr,float32))int {
SINK = a.Var29
Var235 := (<- ([]chan chan chan chan int{})[(<- (<- ([1]chan chan int{})[(b.Var275) - (1)]))])
Label30:
goto Label30
Var234 := (<- (<- (<- Var235)))
Var233 := Var234
type Type31 int16
Var32 := [2]*int{0: (*int)(nil),1: ((<- make(chan [1]*int ,1)))[((((<- make(chan func(chan chan complex128,chan float32,bool) [][0]int ,1))(make(chan chan complex128 ),make(chan float32 ,1),false))[(<- make(chan int ,1))])[((*((*(([]**int{})[(([][2]int{})[((func(**string,*error,string,chan []int16) int)(nil))((**string)(nil),(*error)(nil),"foo",make(chan []int16 ,1))])[(Var233) + (1)]]))))) - (1)]) - (1)]}
Var349 := (**int)(nil)
Var299 := ((func(interface{}) struct { Field297 int
Field298 map[uint]interface{}
})(nil))(interface{}(nil))
Var300 := [2]int{}
(([1][]int{0: (<- ((func(int16) chan []int)(nil))(int16(1)))})[(*((Var32)[(Var299).Field297]))])[((Var300)[((struct { Field307 uint
Field308 int
Field309 chan error
Field310 float64
}{})).Field308]) - (1)] ++
([]Type31{})[((*((*(Var349))))) + (1)],([]int{})[(<- make(chan int ))] = Type31(int16(1)),1
type Type350 complex128
(([]struct { Field360 int
}{})[((func(float32,interface {  Method361 (interface {  Method362 (error) (string,string,bool)
 Method363 (complex64,complex64,bool,uint) (complex64,Type350)
},[1]byte,*bool) (struct { },float32,map[int]int,bool)
},float32) int)(nil))(float32(1.0),interface {  Method361 (interface {  Method362 (error) (string,string,bool)
 Method363 (complex64,complex64,bool,uint) (complex64,Type350)
},[1]byte,*bool) (struct { },float32,map[int]int,bool)
}(nil),float32(1.0))]).Field360 ++
make(chan interface {  Method370 (error) int
} ) <- interface {  Method370 (error) int
}(nil)
_ = Var235
_ = Var234
_ = Var233
_ = Var32
_ = Var349
_ = Var299
_ = Var300
_ = Param28
return 1
}(map[bool]func(complex64,rune) (uintptr,float32){})]))))
_ = Var2
return 
}
func init() {
return 
}
