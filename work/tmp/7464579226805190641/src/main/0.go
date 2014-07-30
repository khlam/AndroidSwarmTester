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
var _ = a.UsePackage
var _ = b.UsePackage
var UsePackage = 0
var SINK interface{}
func main() {
select {
case ((*(([]*struct { Field979 int16
Field980 func(float32) (byte,int,complex64)
Field981 [2]string
Field982 error
Field983 bool
}{})[((func(interface { },complex128,chan *int16) int)(nil))(interface { }(nil),1i,make(chan *int16 ,1))]))).Field983, ([]bool{})[(int)((len)((Var986)[(make(map[int16]int ))[int16(1)]]))] = <-make(chan bool ,a.Func1(Func2(int16(1)))):
}
return 
}
var Var157 = (*((Var229)[((func(struct { Field631 []int16
}) int)(nil))((struct { Field631 []int16
}{}))]))
var Var986 = [2]string{}
func init() {
return 
}
