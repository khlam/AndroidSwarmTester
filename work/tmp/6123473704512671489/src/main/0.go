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
var Var737 = (((([2][2][0][][1][2]**[2]**[1][2][0][0]**[0]*[0][2][1]**int{})[((func([]string,complex64) int)(nil))([]string{},complex64(1i))])[a.Var920])[((*((*((a.Var921)[((b.Var922)[(<- ([]chan int{})[((func([0]map[byte]string,int16,error,int16) int)(nil))([0]map[byte]string{},int16(1),error(nil),int16(1))])]) - (1)]))))) - (1)])[(int)((cap)((*((*((((((Var1048)[(map[complex128]int{})[1i]])[(int)((len)(((*((a.Var2452)[(<- make(chan int ))])))[(int)((((*(Var2492)))[((*(b.Var2504))) - (1)])[((*((*((((*((*(((Var2551)[(Var1804) - (1)])[((struct { Field2561 float64
Field2562 int
Field2563 func(uint) (complex64,int)
}{})).Field2562])))))[(int)(Var2564 ,)])[(int)((len)((*((*(b.Var2742))))))]))))) + (1)] )]))])[(int)((len)((((*(((*((((*((*((((*(a.Var3116)))[(((a.Var3117)[(int)(((*((Var3118)[(int)((len)(((Var3284)[(int)((cap)(b.Var3346))])[(((*((*((((*(((*((*(((*((*((*(Var3493)))))))[(*((*(((*(((*(a.Var3529)))[(Var1804) + (1)])))[(Var1804) - (1)]))))])))))[(<- make(chan int ,1))])))[((func(byte,*int16) int)(nil))(byte(0),(*int16)(nil))])[(int)((cap)(b.Var3542))])))))[(map[float64]int{})[1.0]]) - (1)]))])))[(int)(((*(((*(a.Var3604)))[(<- make(chan int ))])))[(Var3706).Field3674] )] ,)])[((b.Var3744).Field3743) + (1)]) + (1)])[(int)((*(b.Var3745)) ,)])))))[(int)(Var3746 ,)])[(<- make(chan int ))])))[(int)(((Var3840)[(<- make(chan int ,1))]).Field3839 ,)])))[((*(Var3972)))[(b.Var3982).Field3981]])[(Var1804) + (1)]))])[(map[chan error]int{})[make(chan error ,1)]])[(<- make(chan int ))]))))))]
var Var1323 = (<- (*(Var1486))(float32(1.0),complex64(1i))(map[chan complex64]func(uint,interface{}) (complex64,float32,float32,bool){}))
var Var1486 = (*func(float32,complex64) func(map[chan complex64]func(uint,interface{}) (complex64,float32,float32,bool)) chan int)(&((*((((*(a.Var1544)))[((*((b.Var1741)[(int)((Var1765)[(<- make(chan int ,1))] )])))[(Var1804) - (1)]])[(*((b.Var1853)[Var1804]))]))))
var Var1804 = 1
var Var2042 = (*[2]int)(nil)
var Var2564 = 1.0
var Var4075 = [1]**[0]**[2][0]int{}
var Var4383 = [0]*int{}
var Var5091 = (*****[2][0][1]*[0]*int16)(nil)
