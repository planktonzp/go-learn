package main //方法就是函数的一个使用方式的变换 并没有什么不一样的地方
import (
	"fmt"
	"math"
)

type FVertex struct {
	X, Y float64
}

func Abs2(v FVertex) float64 {	//Abs2函数没有接收者 但是内置声明了一个变量v类型是FVertex，在调用Abs2时初始化v也可以将运算结果返回
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := FVertex{3, 4}
	fmt.Println(Abs2(v))
}
