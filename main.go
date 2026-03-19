// 包名
package main

// 导入包
import (
	"fmt"
	"math/rand" // 下级的导入用'/'
)

// 命名：驼峰 包名用名词全7小写 接口加er 对外私有小写开头，公有大写开头

// var const 变量 常量
func varConst() {
	fmt.Println("## var and const:")
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
	fmt.Println("These var and const are legal:")
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, C1, C2, C4, C5)
	fmt.Println("")
}

// 控制流
func stream() {
	// if-else if-else
	fmt.Println("## if, else if, else:")
	a := 1
	if a <= 0 {
		fmt.Println("a is neg or 0")
	} else if a%2 == 0 {
		fmt.Println("a is pos even")
	} else {
		fmt.Println("a is pos odd")
	}
	// goto
	fmt.Println("## goto:")
	if true {
		goto alkali
	}
	fmt.Println("goto will jmp over this")
alkali:
	fmt.Println("goto finish")

	// for
	fmt.Println("## for:")
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println("")

	// while
	fmt.Println("## while & while(true):")
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println("")

	// while(true)
	i = 0
	for {
		fmt.Print(i)
		i++
		if i >= 10 {
			goto outwhile
		}
	}
outwhile:
	fmt.Println("")

	// continue break
	fmt.Println("## continue & break:")
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Print(i)
	}
	fmt.Println("")

	// select case default 仅用于通道
	// b := 3
	// select{
	// case b==1:
	// }
}

// 类型
// struct
type Point struct {
	x, y float64
}
type Person struct {
	name string
	age  int
	pos  Point // 普通字段（区别于匿名字段
}
type Student struct {
	Person // 匿名字段 类似继承，可直接调用Person的成员
	id     uint64
}

// struct 方法
// 成员函数定义方式：func后加类型
func (p Point) cross(p2 Point) float64 {
	return p.x*p2.y - p2.x*p.y
}

// 只有加*传指针，才能修改（默认值传递好）
func (p *Point) addBy(p2 Point) {
	p.x += p2.x
	p.y += p2.y
}
func types() {
	fmt.Println("## types:")
	// 声明的结构同变量一致
	pt1 := Point{1, 5}
	var pt2 = Point{x: -2, y: 1} // 指定字段（顺序时可省略）
	var pt3 Point
	pt4 := new(Point)
	pt3 = pt1
	fmt.Printf("pt1: %v\n", pt1) // 输入.print!自动补全 调试输出
	fmt.Printf("pt2: %v\n", pt2)
	fmt.Printf("pt3: %v\n", pt3)

	// 调用成员
	pt4.x, pt4.y = pt1.y, pt1.x
	fmt.Printf("pt4: %v\n", pt4)
	pt4.addBy(pt1)
	fmt.Printf("pt4: %v\n", pt4)
	x := pt1.cross(pt2)
	fmt.Println("crossx = ", x)
}

func main() {
	varConst()
	stream()
	types()
	var r_num = rand.Int()
	fmt.Println(r_num)
}
