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
_, Var2, _ := (*uint)(&(((((<- (<- a.Var274)))[((<- (([][0]chan struct { Field314 int
}{})[([]int{})[((func(complex64,func(interface { },*int,*interface{}) (uint,float32,error,chan float32),uint) int)(nil))(complex64(1i),((func(interface { },*int,*interface{}) (uint,float32,error,chan float32))(nil)),uint(1))]])[(*(Var414))])).Field314])[((func(byte,error) int)(nil))(byte(0),error(nil))])[(<- make(chan int ,1))])),1.0,((func([]interface{}) (interface {  Method1 (int,byte,int,bool) (bool,complex128)
},uintptr,byte,*interface{}))(nil))
_ = Var2
return 
}
