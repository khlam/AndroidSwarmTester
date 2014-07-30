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
var Var59 = ((<- make(chan []int )))[(int)((len)((*(((*(((*(((*(b.Var141)))[(int)((*(Var143)) )])))[Var1207])))[(([]int{})[(int)(((([][2][1]byte{})[((((b.Var1404)[(make(map[[2]bool]int ,1))[[2]bool{}]])[(int)(([]int16{})[(map[float32]int{})[float32(1.0)]] )])[(int)(([]float32{})[((func(int,[0]map[complex64]string,rune) int)(nil))(1,[0]map[complex64]string{},rune(0))] )])[(int)((*(([]*int{})[(int)(([]float32{})[((func(string,*struct { Field1516 uint
Field1517 byte
Field1518 complex64
},interface {  Method1519 (complex128,byte,[]uint,[]float64) func(rune,uintptr,uintptr) (interface{},complex128,int16,float32)
},int) int)(nil))("foo",(*struct { Field1516 uint
Field1517 byte
Field1518 complex64
})(nil),interface {  Method1519 (complex128,byte,[]uint,[]float64) func(rune,uintptr,uintptr) (interface{},complex128,int16,float32)
}(nil),1)] )])) ,)]])[(Var1207) - (1)])[(int)((len)(([]map[interface {  Method1520 (interface{},rune,float32,uint) (bool,float64,uintptr,byte)
}]chan float64{})[((func(string,struct { },[0][0]float32) int)(nil))("foo",(struct { }{}),[0][0]float32{})]))] )]) + (1)]))))]
var Var143 = ((<- b.Var144))[(int)((([]struct { Field153 struct { Field154 string
Field155 float64
Field156 error
}
Field157 float32
}{7: ((*((*((*((*((*((Var387)[((([][0]int{})[(Var1207) - (1)])[(([]int{})[((func([]*complex128) int)(nil))([]*complex128{})]) - (1)]) + (1)])))))))))))[(([]int{})[Var1207]) + (1)]})[((func(chan error,error,float64,complex128) int)(nil))(make(chan error ,1),error(nil),1.0,1i)]).Field157 ,)]
var Var389 = [1][0]chan [1]*****[2]struct { Field153 struct { Field154 string
Field155 float64
Field156 error
}
Field157 float32
}{[0]chan [1]*****[2]struct { Field153 struct { Field154 string
Field155 float64
Field156 error
}
Field157 float32
}{}}
var Var937 = ([][0]*float64{(<- (<- make(chan chan [0]*float64 ,1)))})[(a.Var955) - ((*((*((b.Var1023)[(<- make(chan int ,1))])))))]
var Var1207 = 1
var Var1546 = float32(1.0)
func init() {
return 
}
