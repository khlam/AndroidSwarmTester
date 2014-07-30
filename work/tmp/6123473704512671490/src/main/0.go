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
SINK = b.Var1
var Var2 interface{} = b.Var3(Var4)
_ = Var2
return 
}
var Var4 = 1i
func init() {
for ((<- Var83))[(int)((len)(a.Var158))] <- interface { }(nil); (*(((*((*(((((*(((*(Var436)))[(int)(((b.Var815)[((*(((*(((a.Var877)[(*((*(Var911))))])[(<- make(chan int ,1))])))[(((*((*((*(a.Var1668))))))).Field1667) - (1)]))) - (1)])[((*(Var1691)))[((*((Var1692)[((func(rune,struct { },rune,[]byte) int)(nil))(rune(0),(struct { }{}),rune(0),[]byte{})]))) - (1)]] ,)])))[(int)((*(((*((*(((*(((*(b.Var1827)))[(<- make(chan int ))])))[(make(map[interface{}]int ,1))[interface{}(nil)]])))))[(map[int]int{})[1]])) )])[((func(chan float64,rune,error,[0]struct { }) int)(nil))(make(chan float64 ,1),rune(0),error(nil),[0]struct { }{})])[(<- make(chan int ,1))])))))[a.Var2165]));  {
continue
make(chan float32 ) <- float32(1.0)
select {
default:
SINK = Var2166
make(chan complex64 ) <- complex64(1i)
Label2167:
goto Label2167
switch "foo" {
case "foo":
fallthrough
default:
}
}
}
return 
}
var Var83 = make(chan [1]chan interface { } )
var Var436 = (*[2]*[0][2][0]**[1]*bool)(&(Var525))
var Var919 = (<- (Var955(((func(struct { Field949 uintptr
Field950 rune
Field951 byte
},float64) ([]byte,int16,[]float32))(nil)),[]struct { Field952 string
}{},interface {  Method953 (func(complex128,uintptr,int16,string) float64,error,interface { }) (map[rune]int,[]error)
 Method954 ([]interface{}) *bool
}(nil)))[(<- make(chan int ,1))])
var Var1207 = [1]int{}
var Var1561 = (**[1]int)(nil)
var Var1692 = [2]*int{}
