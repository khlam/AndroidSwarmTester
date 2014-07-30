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
var _ = a.UsePackage
var _ = b.UsePackage
var UsePackage = 0
var SINK interface{}
var Var41 = Var42
var Var82 = len(Var83)
var Var150 = ((*((*((((([][0][1]**func(uint,interface{},bool,struct { }) []uint{})[([]int{})[([]int{})[(int)((len)(Var189))]]:((([][2]struct { Field755 func(float64,int,byte) (byte,bool,int,byte)
Field756 int
}{})[((Var757)[(<- make(chan int ,1))]) - (1)])[(make(map[float64]int ,1))[1.0]]).Field756])[(<- make(chan int ))])[(a.Var836) - (1)])[(<- make(chan int ,1))]))))(uint(1),interface{}(nil),false,(struct { }{})))[(<- make(chan int ))]
var Var757 = [0]int{}
func init() {
return 
}
