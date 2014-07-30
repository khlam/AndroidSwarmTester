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
a.Var1,(*((((((<- make(chan func(float64,uint) [][0][][2]*map[int16]rune ))(a.Func81(false,1),(uint(1)) + (uint(1))))[((a.Var150)[(int)(a.Var151 ,)]) - (1)])[(int)((*((*(Var152)))) ,)])[((func(complex64) int)(nil))(complex64(1i))])[(int)(Var262 )])),a.Var520 = uintptr(0),make(map[int16]rune ,1),byte(0)
return 
}
var Var262 = ((func(Param312 byte, Param313 chan *float32, Param314 struct { }, Param315 interface{})[0]int {
_ = Param312
_ = Param313
_ = Param314
_ = Param315
return func(Param316 error, Param317 complex128)func(float64,uintptr,uintptr) [0]int {
Var486 := 1
_, _, Var318 := a.Var319,float32(1.0),(*((((((*(((*((*(a.Var366)))))[func(Param408 interface { }, Param409 struct { Field406 float32
Field407 []error
}, Param410 []chan int)func(complex64,func([2]byte,map[bool]interface{},interface{}) error) int {
make(chan map[uint]int16 ,(([0]map[struct { Field436 bool
}]int{})[(<- make(chan func(struct { Field437 *bool
Field438 float32
},struct { Field439 byte
Field440 complex128
Field441 string
}) int ))((struct { Field437 *bool
Field438 float32
}{}),(struct { Field439 byte
Field440 complex128
Field441 string
}{}))])[(struct { Field436 bool
}{})]) <- map[uint]int16{}
if false {
var Var443 float64 = 1.0
var Var444 **byte = (**byte)(nil)
make(chan func(complex128) (error,bool,string,int16) ) <- ((func(complex128) (error,bool,string,int16))(nil))
Var451 := (*float64)(nil)
switch (*(Var451)) --;  {
case false:
SINK = Var452
fallthrough
default:
}
_ = Var443
_ = Var444
_ = Var451
} else {
}
_ = Param408
_ = Param409
_ = Param410
return ((func(complex64,func([2]byte,map[bool]interface{},interface{}) error) int)(nil))
}(interface { }(nil),(struct { Field406 float32
Field407 []error
}{}),[]chan int{})(complex64(1i),((func([2]byte,map[bool]interface{},interface{}) error)(nil)))])))[(<- make(chan int ,1))])[(*((([][0]*int{})[b.Var485])[(Var486) + (1)]))])[(<- make(chan int ))])[Var486]))
_ = Var486
_ = Var318
_ = Param316
_ = Param317
return ((func(float64,uintptr,uintptr) [0]int)(nil))
}(error(nil),1i)(1.0,uintptr(0),uintptr(0))
}(byte(0),make(chan *float32 ,1),(struct { }{}),interface{}(nil)))[((func([]int,uintptr,struct { Field519 chan error
}) int)(nil))([]int{},uintptr(0),(struct { Field519 chan error
}{}))]) + (1)
