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
func main() {
type Type1 float32
Var1145 := [0]float32{}
Var337 := (<- (<- (<- ((func(complex64,*struct { Field383 float32
}) chan chan chan [0][2][2]*func(rune,map[bool]int) (chan byte,uint))(nil))(complex64(1i),(*struct { Field383 float32
})(nil)))))
Var1200 := make(chan float32 ,1)
Var192 := (<- (<- make(chan chan chan map[complex64]string ,((func(struct { },[]uintptr) int)(nil))((struct { }{}),[]uintptr{}))))
Var1199 := "foo"
Var153 := b.Var154((<- (((func(complex64) []chan int)(nil))(complex64(1i)))[func(Param155 float32)int {
Var305, Var306 := <-Var192
Var1140 := make(chan complex64 ,1)
Var414 := Var480
Var1142 := rune(0)
go (*((((Var337)[(Var414) + (1)])[(int)((len)(Var1140))])[(int)((b.Var1141)[(int)(Var1142 ,)] )]))(rune(0),make(map[bool]int ))
_ = Var305
_ = Var306
_ = Var1140
_ = Var414
_ = Var1142
_ = Param155
return 1
}(float32(1.0))]),(*chan uint)(nil),complex64(1i),[]int16{})
Var24 := ((*(((((func(chan []bool,float32) [][]*[0][0]int)(nil))(make(chan []bool ,1),float32(1.0)))[(int)(((Var153)[(int)((len)(([]string{})[((func(error,chan [0]complex64,*Type1) int)(nil))(error(nil),make(chan [0]complex64 ),(*Type1)(nil))]))])[(int)((Var1145)[(<- make(chan int ))] ,)] ,)])[(int)((len)(([]map[byte]uint{})[(int)(((*(([]*[2]rune{})[(int)(([]uintptr{})[(int)((len)(Var1199))] ,)])))[(int)((cap)(Var1200))] ,)]))])))[(<- make(chan int ))]
Var3 := (((([2][1][1][][0]Type1{})[((func(string,*[0]Type1) int)(nil))("foo",(*[0]Type1)(nil))])[b.Var23])[((Var24)[((func(uint,bool) int)(nil))(uint(1),false)]) + (1)])[(int)((*(([]*byte{})[((func(float32,interface {  Method1318 (uintptr,chan bool,interface{},[]bool) (struct { Field1319 rune
Field1320 Type1
},*byte,Type1,int)
 Method1321 ([]float64,rune) func(complex128) rune
 Method1322 (uintptr,[]uintptr,uintptr,float64) byte
},float32,byte) int)(nil))(float32(1.0),interface {  Method1318 (uintptr,chan bool,interface{},[]bool) (struct { Field1319 rune
Field1320 Type1
},*byte,Type1,int)
 Method1321 ([]float64,rune) func(complex128) rune
 Method1322 (uintptr,[]uintptr,uintptr,float64) byte
}(nil),float32(1.0),byte(0))])) )]
Var1323 := 1.0
Var2 := Var3
(Var2)[((func(*struct { },map[interface { }]float32,interface{}) int)(nil))((*struct { })(nil),map[interface { }]float32{},interface{}(nil))],Var1323 = Type1(float32(1.0)),1.0
_ = Var1145
_ = Var337
_ = Var1200
_ = Var192
_ = Var1199
_ = Var153
_ = Var24
_ = Var3
_ = Var1323
_ = Var2
return 
}
var Var858 = [1]int{}
var Var1074 = [0]int{}
func init() {
return 
}
