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
func Func100(Param101 string, Param102 float64, Param103 rune)complex128 {
select {
case make(chan func(*byte,chan uint) map[uintptr]complex128 ) <- Var104:
case (<- make(chan func(float64,rune) chan byte ))((*((*((*(((*((([][1]*[]***float64{})[Var374([]*interface{}{},(*uint)(nil),[][2]rune{},[2]byte{})])[(<- make(chan int ))])))[(int)(((b.Var479)[([]int{})[((((Var607)[Var608])[(*(((*((*(Var609)))))[(int)(([]uint{})[((func(chan complex64,string) int)(nil))(make(chan complex64 ,1),"foo")] ,)]))])[(Var653)[((func(complex64) int)(nil))(complex64(1i))]]) - (1)]])[(int)(([]byte{})[(<- make(chan int ))] ,)] ,)])))))),rune(0)) <- byte(0):
defer ((func(*complex128,func(uint,int16) (int16,error,int,int16),struct { },int) (struct { },rune,uint,chan string))(nil))((*complex128)(nil),((func(uint,int16) (int16,error,int,int16))(nil)),(struct { }{}),1)
var Var727 struct { Field728 struct { }
} = (struct { Field728 struct { }
}{})
_ = Var727
case ([][1]rune{})[(((*((*(Var755)))))[((func([]struct { },float64) int)(nil))([]struct { }{},1.0)]) - (1)], (((Var780)[(([]int{})[(<- make(chan int ,1))]) + (1)])[(((((*(Var781)))[(Var608) + (1)])[(<- make(chan int ))])[([]int{})[(([][1]int{})[(make(map[interface { }]int ))[interface { }(nil)]])[(int)((len)((*(([]*[]complex64{})[Var608]))))]]]) - (1)])[(int)((*(((([][0][1]*int{})[(int)((cap)((*((*((*((*((((*(Var986)))[(int)(b.Var1010 ,)])[(int)((len)(([]string{})[(int)((a.Var1048)[(*(([]*int{})[([]int{})[(((Var1169)[(make(map[bool]int ))[false]])[(([]int{})[(int)(([]float64{})[(<- make(chan int ,1))] )]) - (1)]) - (1)]]))] )]))]))))))))))])[(Var608) + (1)])[((func([0][2]int16,float32,struct { Field1237 func(rune,int16) (int,interface{})
}) int)(nil))([0][2]int16{},float32(1.0),(struct { Field1237 func(rune,int16) (int,interface{})
}{}))])) ,)] = <-make(chan [1]rune ):
var Var1238 []interface {  Method1239 (complex64) (complex128,uintptr)
} = []interface {  Method1239 (complex64) (complex128,uintptr)
}{}
Var1549 := (*int16)(nil)
Var1291 := [2]*rune{}
Var1284 := [2]uint{}
switch COND := (interface{})(([]string{})[(([][2]int{})[((([][1][1]int{})[(<- make(chan int ,1))])[(<- make(chan int ,1))])[(<- make(chan int ,1))]])[(int)((Var1284)[((func(struct { },uintptr,bool) int)(nil))((struct { }{}),uintptr(0),false)] )]]); COND.(type) {
}
for (*((Var1291)[(<- make(chan int ))])) ++; false;  {
if Var1293, Var1294, _, _ := false,[]struct { Field1292 complex128
}{},int16(1),map[chan interface{}]*byte{}; false {
(((*((*(([]**[0]struct { Field1324 uintptr
Field1325 rune
Field1326 struct { }
Field1327 uintptr
}{})[(*(([]*int{})[(*(([]*int{})[((func(byte,uint) int)(nil))(byte(0),uint(1))]))]))])))))[(int)((*(Var1549)) )]).Field1325 --
_ = Var1293
_ = Var1294
}
}
_ = Var1238
_ = Var1549
_ = Var1291
_ = Var1284
}
_ = Param101
_ = Param102
_ = Param103
return 1i
}
var Var608 = 1
var Var1169 = [1][2]int{}
func init() {
return 
}
