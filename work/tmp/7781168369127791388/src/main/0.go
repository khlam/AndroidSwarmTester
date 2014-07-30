package main
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
var _ = a.UsePackage
var UsePackage = 0
var SINK interface{}
var Var414 = (*(a.Var415))
var Var780 = 1
func Func782(Param793 uint, Param794 float32, Param795 byte, Param796 uint)(interface{},rune) {
_ = Param793
_ = Param794
_ = Param795
_ = Param796
return interface{}(nil),rune(0)
}
func init() {
return 
}
