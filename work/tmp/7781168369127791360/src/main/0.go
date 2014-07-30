package main
import "runtime"
func init() {
	go func() {
		for {
			runtime.GC()
			runtime.Gosched()
		}
	}()
}
var UsePackage = 0
var SINK interface{}
func main() {
Label1:
Var26 := (*int)(nil)
Var65 := (<- ([]chan [0][][][0]float32{(<- make(chan chan [0][][][0]float32 ))})[(((([...][]struct { Field182 uint
Field183 float32
Field184 interface {  Method185 (float64,interface{}) byte
 Method186 (string) uint
}
Field187 [0]uint
Field188 int
}{})[(<- make(chan func(map[complex64]map[rune]uintptr,interface { },complex64) int ))(make(map[complex64]map[rune]uintptr ,1),interface { }(nil),complex64(1i))])[((struct { Field319 chan string
Field320 int
Field321 float32
Field322 float32
Field323 struct { }
Field324 complex64
}{})).Field320]).Field188) - (1)])
goto Label1
Var325 := (<- (*(((<- ([]chan []*chan [1]int{})[(<- make(chan int ,1))]))[(<- (<- ([2]chan chan int{})[((func([][]rune,rune,interface {  Method596 (*byte) (float64,interface { },map[float64]float64,*float32)
}) int)(nil))([][]rune{},rune(0),interface {  Method596 (*byte) (float64,interface { },map[float64]float64,*float32)
}(nil))]))])))
select {
case (<- (((func(map[int16][]int16) [2]chan func(interface{}) chan [1]func(float64,int16,bool) (complex64,int,int16,rune))(nil))(map[int16][]int16{}))[((*(Var26))) - (len("foo"))])((<- make(chan interface{} ,1))) <- [1]func(float64,int16,bool) (complex64,int,int16,rune){}:
return 
Var850 := 1
Var797 := [0]*string{}
Var27, Var28, Var29 := float32(1.0),(((Var65)[((Var325)[(<- ((func(rune) chan int)(nil))(rune(0)))]) + (1)])[(([]struct { Field598 int
}{})[(int)((len)((*((Var797)[(Var850) + (1)]))))]).Field598])[((func(int,interface { }) int)(nil))(1,interface { }(nil))],interface { }(nil)
Var1308 := (**int)(nil)
Var1030 := [2][0]struct { }{}
make(chan func(byte,bool) error ) <- ((func(byte,bool) error)(nil))
Var1029 := (*map[float64]chan float32)(nil)
Var1307 := [1]int{}
_, Var964, _ := ((func([]complex64,bool,func(bool,uintptr,int16) (interface{},bool,int),uint) (struct { },map[rune]float64,struct { Field963 interface{}
},int16))(nil)),make(chan func(byte) string ),float32(1.0)
([]struct { }{})[(int)((len)((*(Var1029))))],((Var1030)[(int)((len)(((([][0]struct { Field1044 string
Field1045 []complex128
}{})[(([]int{})[((func(map[byte]func(complex64,byte) string,complex128,uint) int)(nil))(make(map[byte]func(complex64,byte) string ,1),1i,uint(1))]) - (1)])[Var850]).Field1044))])[((((*(((*(([]*[2]*[1][1]int{})[(([]int{})[((func(func(byte,struct { Field1272 rune
}) ([1]float32,int,[0]byte,chan error),interface{},map[chan int]struct { Field1273 bool
}) int)(nil))(((func(byte,struct { Field1272 rune
}) ([1]float32,int,[0]byte,chan error))(nil)),interface{}(nil),map[chan int]struct { Field1273 bool
}{})]) - (1)])))[((func(chan struct { Field1274 error
},[]interface { }) int)(nil))(make(chan struct { Field1274 error
} ),[]interface { }{})])))[((*((*((*(([]***int{})[(*(([]*int{})[(<- make(chan int ,1))]))]))))))) - (1)])[(Var1307)[((func([1]float64) int)(nil))([1]float64{})]]) - (1)],(([][2]bool{})[(*((*(Var1308))))])[([]int{})[((func(interface {  Method1336 (interface {  Method1337 (error) (error,int16)
},complex64) (int,chan float64,byte,map[complex128]float64)
 Method1338 (*int16,[]float64,interface { },func(interface{}) (int16,interface{},rune)) ([0]float32,func(error,byte,int16,int16) (string,float32,byte),interface {  Method1339 (uint) (byte,uintptr,error,byte)
},struct { })
},map[[0]uint]interface {  Method1340 (byte) (interface{},int16,rune,uintptr)
},interface { }) int)(nil))(interface {  Method1336 (interface {  Method1337 (error) (error,int16)
},complex64) (int,chan float64,byte,map[complex128]float64)
 Method1338 (*int16,[]float64,interface { },func(interface{}) (int16,interface{},rune)) ([0]float32,func(error,byte,int16,int16) (string,float32,byte),interface {  Method1339 (uint) (byte,uintptr,error,byte)
},struct { })
}(nil),map[[0]uint]interface {  Method1340 (byte) (interface{},int16,rune,uintptr)
}{},interface { }(nil))]] = (struct { }{}),(struct { }{}),false
Var1344, Var1345 := <-make(chan struct { Field1343 rune
} ,1)
_ = Var850
_ = Var797
_ = Var27
_ = Var28
_ = Var29
_ = Var1308
_ = Var1030
_ = Var1029
_ = Var1307
_ = Var964
_ = Var1344
_ = Var1345
}
_ = Var26
_ = Var65
_ = Var325
return 
}
func init() {
return 
}
