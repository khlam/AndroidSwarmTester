package main
import "b"
import "a"
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
func init() {
return 
}
var Var285 = (a.Func286(Func287(1i))) + (1)
func Func332(Param333 interface{}, Param334 float32)(string,uint,float32,bool) {
var Var335 byte = b.Var336
_ = Var335
_ = Param333
_ = Param334
return ([]string{})[((([]struct { Field411 int
}{})[:((*(((*(([]*[]*struct { Field499 int
Field500 error
}{})[((func(complex64) int)(nil))(complex64(1i))])))[(*(([]*int{})[((*((([][0]*int{})[(<- make(chan int ))])[(([]struct { Field678 int
Field679 *int
Field680 uintptr
Field681 float32
Field682 float64
Field683 error
}{})[(int)((len)(([]string{})[((struct { Field691 int
Field692 float64
Field693 uintptr
}{})).Field691]))]).Field678]))) - (1)]))]))).Field499])[((func(func(int16,map[int16]string) map[rune]complex64,int16) int)(nil))(((func(int16,map[int16]string) map[rune]complex64)(nil)),int16(1))]).Field411],uint(1),float32(1.0),false
}
