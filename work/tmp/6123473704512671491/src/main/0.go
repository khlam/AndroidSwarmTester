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
func main() {
go ((*(Var1)))[(int)((*(Var176)) ,)](((((<- make(chan [0]complex128 ,1)))[(*((((*((*((*((*((Var359)[((func(float64,float32,[]uint,[]func(complex64,error,bool,float32) (interface{},complex128)) int)(nil))(1.0,float32(1.0),[]uint{},[]func(complex64,error,bool,float32) (interface{},complex128){})])))))))))[((*((a.Var3417)[((struct { Field3429 chan bool
Field3430 int
Field3431 interface { }
}{})).Field3430]))) - (1)])[((struct { Field3452 int
}{})).Field3452]))]) + (1i)) + (1i),complex64(1i))
return 
}
var Var2 = func(Param3 byte)*[2]func(complex128,complex64) (error,rune,uint) {
_ = Param3
return (*[2]func(complex128,complex64) (error,rune,uint))(nil)
}((byte)((float64)(((struct { Field74 byte
Field75 int
Field76 float32
Field77 float32
Field78 string
Field79 struct { Field80 complex128
Field81 error
}
Field82 interface {  Method83 (complex64,float64) error
 Method84 (float64) (string,float64,complex64)
}
}{})).Field74 ,) ,))
var Var359 = [...]****[0][2]*int{(*(((*((*(((*(Var425)))[((((*((*(((*((*(Var616)))))[((*(((*(b.Var1776)))[((struct { Field1799 int
}{})).Field1799])))[(Var1284) + (1)]])))))[(int)(((((Var2042)[(*((*((*((((*((*((*(b.Var2226)))))))[(*(Var2227))])[(((a.Var2228)[(*((*(((b.Var2264)[(<- make(chan int ))])[(make(map[*byte]int ))[(*byte)(nil)]]))))])[((func(bool,rune) int)(nil))(false,rune(0))]) - (1)]))))))])[((func(chan int16) int)(nil))(make(chan int16 ,1))])[(make(map[interface{}]int ))[interface{}(nil)]])[((*((*((*(b.Var2427))))))) + (1)] )])[(int)((((*(a.Var2509)))[(int)(((((*(a.Var2573)))[(Var1284) + (1)])[((*(((*((((Var2755)[((*(b.Var2810))) + (1)])[((func(struct { },float32) int)(nil))((struct { }{}),float32(1.0))])[(Var2902)[(int)((len)((b.Var2983)[(*(Var2984))]))]])))[((*(((((*((Var3108)[(int)(Var3109 )])))[((struct { Field3133 int
}{})).Field3133])[Var1284])[(Var1284) - (1)]))) + (1)])))[((*(((*((*(a.Var3391)))))[((func(rune,map[complex128]float64) int)(nil))(rune(0),map[complex128]float64{})])))[(int)(Var3392 )]]]).Field2515 ,)]).Field2434 ,)])[(int)((len)(Var3395))]])))))[(<- make(chan int ,1))]))}
var Var616 = (<- make(chan **[1]**[0][2][1]int ,(<- (*((((*(Var696)))[((*(Var1616))) + (1)])[(int)(((*(((*(a.Var1745)))[((func(uintptr,rune) int)(nil))(uintptr(0),rune(0))])))[(<- make(chan int ))] )])))))
var Var697 = (*((((*(b.Var769)))[(<- (<- ((func(bool,struct { },func(byte,rune,float64,uint) (uintptr,map[int16]rune,error)) chan func(bool,byte,float32) chan int)(nil))(false,(struct { }{}),((func(byte,rune,float64,uint) (uintptr,map[int16]rune,error))(nil))))(Func839(Func1412(int16(1))),byte(0),float32(1.0)))])[(make(map[interface {  Method1417 (byte,interface{},interface{}) (error,rune)
 Method1418 (complex128,float64,float32) (uintptr,error,complex128)
 Method1419 (uintptr) (interface{},int16,int16,interface{})
 Method1420 (bool) (string,rune,error)
}]int ))[interface {  Method1417 (byte,interface{},interface{}) (error,rune)
 Method1418 (complex128,float64,float32) (uintptr,error,complex128)
 Method1419 (uintptr) (interface{},int16,int16,interface{})
 Method1420 (bool) (string,rune,error)
}(nil)]]))
func Func839(Param840 interface{})bool {
switch COND := (interface{})((*((*(a.Var969))))); COND.(type) {
case *chan string:
Var970, Var971, Var972, _ := (<- b.Var1023)(func(Param1028 interface{}, Param1029 func(interface {  Method1024 (complex64,bool,complex128) (rune,error,uint,uintptr)
 Method1025 (string,complex128,rune) (uintptr,int,int)
 Method1026 (complex64,uintptr,int,error) (float32,uint,uint)
 Method1027 (float64,rune) (uint,string,string,int16)
}) (*uint,struct { },complex64), Param1030 chan error, Param1031 complex64)[]int {
for ; (a.Func1034(Func1035(interface{}(nil),rune(0),false))) == (rune(0));  {
}
_ = Param1028
_ = Param1029
_ = Param1030
_ = Param1031
return []int{}
}(interface{}(nil),((func(interface {  Method1024 (complex64,bool,complex128) (rune,error,uint,uintptr)
 Method1025 (string,complex128,rune) (uintptr,int,int)
 Method1026 (complex64,uintptr,int,error) (float32,uint,uint)
 Method1027 (float64,rune) (uint,string,string,int16)
}) (*uint,struct { },complex64))(nil)),make(chan error ),complex64(1i)),(*struct { Field973 float64
})(nil),rune(0)),(*rune)(nil),[1]interface{}{},false
_ = Var970
_ = Var971
_ = Var972
default:
}
_ = Param840
return false
}
var Var1238 = (struct { Field1230 float32
Field1231 complex64
Field1232 chan rune
Field1233 struct { Field1234 error
Field1235 uintptr
}
Field1236 float64
Field1237 uintptr
}{})
var Var1616 = (*int)(nil)
var Var2227 = (*int)(nil)
var Var2902 = [1]int{}
func init() {
return 
}
