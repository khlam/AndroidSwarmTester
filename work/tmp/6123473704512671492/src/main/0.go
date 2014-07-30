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
var Var285 = ((*((*(Var314)))))[((func(struct { Field2821 interface { }
}) int)(nil))((struct { Field2821 interface { }
}{}))]
var Var314 = (<- (<- (([]chan chan **[][0]**float64{})[a.Var355:(((*((*((*((*(((*(a.Var397)))[((struct { Field409 struct { Field410 rune
Field411 float32
}
Field412 map[complex64]uint
Field413 int
}{})).Field413])))))))))[(<- make(chan int ))])[((func(map[[2]rune]rune,int,[1]interface { }) int)(nil))(make(map[[2]rune]rune ),1,[1]interface { }{})]])[((Var472)[((((*(b.Var2615)))[((struct { Field2679 int
}{})).Field2679])[(b.Var2769)[(<- make(chan int ,1))]]) + (1)]) - (1)]))
var Var473 = (<- b.Var474)
var Var612 = (*(b.Var613))
var Var1820 = [2]float64{}
var Var2209 = (**[2]string)(nil)
var Var2257 = (struct { Field2256 int16
}{})
var Var2344 = [2][1]int{}
var Var2468 = (*[2]int)(nil)
var Var2951 = (struct { Field2945 bool
Field2946 complex128
Field2947 string
Field2948 complex64
Field2949 *uint
Field2950 float32
}{})
func init() {
return 
}
