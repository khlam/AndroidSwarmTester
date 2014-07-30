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
var Var84 = (*((*((*((*(([]****[2]int{Var85})[((*(([]*int{})[(int)((len)(Var1367))]))) - (1)]))))))))
var Var85 = Var86
var Var86 = ((<- (((func(func(interface { },byte,float32,string) (struct { Field162 complex128
Field163 error
},bool,chan uintptr),map[int16]struct { },[0]interface { }) [0]chan []****[2]int)(nil))(((func(interface { },byte,float32,string) (struct { Field162 complex128
Field163 error
},bool,chan uintptr))(nil)),map[int16]struct { }{},[0]interface { }{}))[(<- Var164)]))[((func(string,map[[1]int][]int,map[complex128]interface {  Method1322 (interface{},error,int,uint) (error,complex128,float64,complex64)
},[1]chan complex64) int)(nil))("foo",make(map[[1]int][]int ),map[complex128]interface {  Method1322 (interface{},error,int,uint) (error,complex128,float64,complex64)
}{},[1]chan complex64{})]
var Var164 = (b.Var165)[((*((*((Var210)[(*(Var1321))]))))) - (1)]
var Var210 = func(Param289 string, Param290 int)[1]**int {
Var291 := (b.Var292)[(<- ([]func(func(map[uintptr]float32,error) (complex128,interface { },rune),interface{},byte) func(rune,func(rune,uint,string,complex128) (complex64,complex128,string),float32,[]uintptr) chan int{})[(((*(([]*[1]int{})[(<- make(chan int ))])))[(([]int{})[((func(interface { },rune,interface {  Method423 (error,[1]uintptr,error,interface{}) (bool,rune,interface { },uint)
 Method424 (float32,int16) (chan complex128,error)
 Method425 (interface { },*bool,interface {  Method426 (int,error) interface{}
}) (int16,map[uintptr]float64)
 Method427 (interface {  Method428 (interface{},error) (string,string,byte)
},struct { Field429 complex128
},func(uint) uintptr) (chan complex64,bool)
},int16) int)(nil))(interface { }(nil),rune(0),interface {  Method423 (error,[1]uintptr,error,interface{}) (bool,rune,interface { },uint)
 Method424 (float32,int16) (chan complex128,error)
 Method425 (interface { },*bool,interface {  Method426 (int,error) interface{}
}) (int16,map[uintptr]float64)
 Method427 (interface {  Method428 (interface{},error) (string,string,byte)
},struct { Field429 complex128
},func(uint) uintptr) (chan complex64,bool)
}(nil),int16(1))]) - (1)]) - (1)](((func(map[uintptr]float32,error) (complex128,interface { },rune))(nil)),interface{}(nil),byte(0))(rune(0),((func(rune,uint,string,complex128) (complex64,complex128,string))(nil)),float32(1.0),[]uintptr{}))]
Var1315 := [2]bool{}
Var722 := [1]***float32{}
Var1006 := uint(1)
Var485 := [0][2]int{}
Var1049 := (struct { Field1010 bool
}{})
select {
case (*((*((*((Var722)[((struct { Field736 rune
Field737 float32
Field738 int
Field739 int
}{})).Field738])))))), ((*(([]*[1]bool{})[((func(int) int)(nil))(1)])))[((func([]float32,map[complex64]bool,float64,*error) int)(nil))([]float32{},make(map[complex64]bool ),1.0,(*error)(nil))] = <-([]chan float32{})[(int)((cap)((*(([]*chan int{})[((*(([]*[0]int{})[([]int{})[((([][2]int{})[(((Var485)[(<- make(chan int ,1))])[((*(a.Var486))) - (1)]) + (1)])[(([]int{})[(<- make(chan int ,1))]) + (1)]) + (1)]])))[((struct { Field591 int
}{})).Field591]]))))]:
type Type1005 float64
Var1260 := (*[1][0][]rune)(nil)
Var1316 := (*int)(nil)
SINK = Var1006
Var1051 := [2][0]int{}
select {
case Var1007 := <-make(chan complex128 ):
Var1006, (Var1049).Field1010 = <-make(chan uint ,1)
switch ;  {
default:
}
_ = Var1007
case make(chan byte ) <- byte(0):
case make(chan float64 ,1) <- 1.0:
case Var1050 := <-make(chan [0]complex64 ):
_ = Var1050
case ([]func(error,float64,complex64) func(complex128,uintptr,uintptr,uint) (Type1005,rune,float64){})[(((Var1051)[(([]int{})[((func(error) int)(nil))(error(nil))]) + (1)])[(([]int{})[((func(*map[int16]Type1005,struct { },[]float64) int)(nil))((*map[int16]Type1005)(nil),(struct { }{}),[]float64{})]) - (1)]) - (1)], (([][1]bool{})[Param290])[((struct { Field1136 int
Field1137 rune
Field1138 chan interface{}
Field1139 bool
}{})).Field1136] = <-make(chan func(error,float64,complex64) func(complex128,uintptr,uintptr,uint) (Type1005,rune,float64) ,1):
case ((*(Var1260)))[((func(chan complex64) int)(nil))(make(chan complex64 ))], (Var1315)[(*(Var1316))] = <-make(chan [0][]rune ):
case make(chan complex64 ,1) <- complex64(1i):
}
_ = Var1260
_ = Var1316
_ = Var1051
case make(chan *float64 ) <- (*float64)(nil):
case make(chan struct { Field1317 string
} ) <- (struct { Field1317 string
}{}):
case Var1318 := <-make(chan *int ,1):
_ = Var1318
case Var1319, Var1320 := <-make(chan chan []int ,1):
_ = Var1319
_ = Var1320
default:
}
_ = Var291
_ = Var1315
_ = Var722
_ = Var1006
_ = Var485
_ = Var1049
_ = Param289
_ = Param290
return [1]**int{}
}("foo",1)
