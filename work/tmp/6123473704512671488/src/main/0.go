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
var Var188 = (([][0]*[2][2]*interface { }{})[(*((a.Var189)[(<- (<- Var190))])):((*((*(((a.Var282)[([1]func(interface{},float64,error) func(uint) int{})[(<- make(chan int ,1))](interface{}(nil),1.0,error(nil))(uint(1))])[(<- make(chan int ))])))))[(int)(a.Var283 ,)]])[((*((*((*(((((*((*(((Var669)[(b.Var670) - (1)])[(map[*byte]int{})[(*byte)(nil)]])))))[((struct { Field698 map[float64]int
Field699 int16
Field700 int
}{})).Field700])[((func(complex128,[]interface {  Method701 (error) (complex64,complex128,complex128,complex128)
}) int)(nil))(1i,[]interface {  Method701 (error) (complex64,complex128,complex128,complex128)
}{})])[(int)((*((((*(Var752)))[(*(((*(b.Var957)))[((*(Var1102))) + (1)]))])[(int)((((*((*(b.Var1103)))))[((*((((b.Var1232)[(make(map[float64]int ))[1.0]])[((*(((*(((*(((*(((*(Var1423)))[(a.Var1424) + (1)])))[(b.Var1425) - (1)])))[(map[[2]complex128]int{})[[2]complex128{}]])))[(<- make(chan int ,1))]))) - (1)])[(int)((len)((*(((*(((*((*(Var1470)))))[(<- make(chan int ))])))[((func(func(int,[]float32,chan complex128,func(int,complex128) complex128) (interface {  Method1471 (complex128,error) (uint,byte)
 Method1472 (bool,int16,int16,int16) int16
 Method1473 (rune) (rune,complex64,int)
 Method1474 (byte,error,interface{},int16) (int,float32)
},float32,int)) int)(nil))(((func(int,[]float32,chan complex128,func(int,complex128) complex128) (interface {  Method1471 (complex128,error) (uint,byte)
 Method1472 (bool,int16,int16,int16) int16
 Method1473 (rune) (rune,complex64,int)
 Method1474 (byte,error,interface{},int16) (int,float32)
},float32,int))(nil)))]))))]))) - (1)])[(map[[2]int16]int{})[[2]int16{}]] )])) )])))))))[(<- make(chan int ))]]
var Var223 = (<- make(chan float32 ))
var Var669 = [0][1]**[1][0][2]***[2]int{}
var Var752 = (*[1][1]*int)(nil)
var Var1470 = (**[2]*[1]*string)(nil)
func init() {
return 
}
