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
var Var30 = (<- (<- (<- make(chan func(float64) chan chan *chan uint ))(([]float64{})[((func([2]complex128) int)(nil))([2]complex128{})])))
