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
return 
}
var Var130 = ((<- (([2][1]chan [][1]**string{})[(b.Var170) - (1)])[(<- Var171)]))[((struct { Field510 int
}{})).Field510]
var Var200 = (<- Var201)
var Var203 = ([...]func([0][1]complex128,int,uintptr,interface {  Method204 (*error,string,func(complex128,complex64,uintptr,float32) float64) (rune,interface {  Method205 (uint,byte,error,complex64) (interface{},rune)
},interface{},bool)
 Method206 (interface { },struct { },complex128) (struct { },float64,uintptr,float32)
}) func(chan string,[1]func(float64) (uint,complex128,int16),int16,error) chan chan func(int16,uint,*bool,chan complex64) chan []chan int{((func([0][1]complex128,int,uintptr,interface {  Method204 (*error,string,func(complex128,complex64,uintptr,float32) float64) (rune,interface {  Method205 (uint,byte,error,complex64) (interface{},rune)
},interface{},bool)
 Method206 (interface { },struct { },complex128) (struct { },float64,uintptr,float32)
}) func(chan string,[1]func(float64) (uint,complex128,int16),int16,error) chan chan func(int16,uint,*bool,chan complex64) chan []chan int)(nil))})[(<- ([]chan int{})[((func(chan uint,interface { },interface {  Method273 (complex64,[]int16,complex128) float64
},[]*byte) int)(nil))(make(chan uint ,1),interface { }(nil),interface {  Method273 (complex64,[]int16,complex128) float64
}(nil),[]*byte{})])]([0][1]complex128{},1,uintptr(0),interface {  Method204 (*error,string,func(complex128,complex64,uintptr,float32) float64) (rune,interface {  Method205 (uint,byte,error,complex64) (interface{},rune)
},interface{},bool)
 Method206 (interface { },struct { },complex128) (struct { },float64,uintptr,float32)
}(nil))(make(chan string ),[1]func(float64) (uint,complex128,int16){},int16(1),error(nil))
var Var659 = (**map[error]chan uint)(nil)
