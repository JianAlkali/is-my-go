// 包名
package main

// 导入包
import (
	"fmt"
	"math/rand" // 下级的导入用'/'
)

// 命名：驼峰 包名用名词全小写 接口加er 对外私有小写开头，公有大写开头
// 变量 常量
const Pi = 3.14

var cir_h = 3.
var cir_r float64 // 未指定初始值时必须指定类型
var (
	a, b int
	name = "Alkali"
)

// 定义结构体
type Person struct {
	age, idx int
	name     string
}

// 接口
type Reneamable interface {
	Rename()
}

// var const 变量 常量
func varConst() {
	var v1 int = 1
	var v2 = "2"
	var v3 string // 未指定初始值要指定类型
	v4 := "3"     // := 声明，自动推断类型，不用var

	var (
		v5 int = 4
		v6     = "5"
		v7 string
		// v8 := 6 那么这样自然不行
	)

	const C1 int = 6
	const C2 = "7"
	// const C3 string // const需初始值
	const (
		C4 int = 8
		C5     = "9"
	)
	fmt.Println("var and const usage:")
	fmt.Println("It can:", v1, v2, v3, v4, v5, v6, v7, C1, C2, C4, C5)
	fmt.Println("")
}

func main() {
	varConst()
	var r_num = rand.Int()
	fmt.Println(r_num)
}
