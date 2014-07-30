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
var Var991 = (*[1]**[2]int)(nil)
var Var1785 = (*[2]float64)(nil)
var Var1888 = (struct { Field1887 int
}{})
var Var2200 = [0][0]*[2][1]*[1]*[2]*[1]int{}
var Var2789 = (*[2]struct { Field2718 *uintptr
Field2719 string
Field2720 rune
Field2721 map[string]complex64
})(nil)
