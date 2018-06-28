package main

import (
	"fmt"
	"math"
)

func init() {

	//那你是否想过函数当作struct的字段一样来处理呢？今天我们就讲解一下函数的另一种形态，带有接收者的函数，我们称为method
	fmt.Println("object-Oriented  start    面向对象  ")



	//  todo    这个是关键的意思  非常的有意思
	//也就是说：
	//如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
	//类似的
	//如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
	//所以，你不用担心你是调用的指针的method还是不是指针的method，Go知道你要做的一切，这对于有多年C/C++编程经验的同学来说，真是解决了一个很大的痛苦。
	methodDemo()


	//method继承
	methodExtendsDemo()


	//method重写 如果Employee想要实现自己的SayHi,怎么办？简单，和匿名字段冲突一样的道理，我们可以在Employee上面定义一个method，重写了匿名字段的方法
	methodDemoC()


}
func methodDemoC() {

	//设计出基本的面向对象的程序了，但是Go里面的面向对象是如此的简单，没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则。
	mark := Student1{Human11{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee1{Human11{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
type Human11 struct {
	name string
	age int
	phone string
}

type Student1 struct {
	Human11 //匿名字段
	school string
}

type Employee1 struct {
	Human11 //匿名字段
	company string
}

//Human定义method
func (h *Human11) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee1) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}








//学习了字段的继承，那么你也会发现Go的一个神奇之处，method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
func methodExtendsDemo() {
	mark := StudentExtends{HumanExtends{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := EmployeeExtends{HumanExtends{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	//Hi, I am Mark you can call me on 222-222-YYYY
	//Hi, I am Sam you can call me on 111-888-XXXX
	mark.SayHi()
	sam.SayHi()
}
//在human上面定义了一个method
func (h *HumanExtends) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func (h *StudentExtends) SayHi()  {
	fmt.Println("我是重写的方法",h.age,h.name)
}
type HumanExtends struct {
	name string
	age int
	phone string
}

type StudentExtends struct {
	HumanExtends //匿名字段
	school string
}

type EmployeeExtends struct {
	HumanExtends //匿名字段
	company string
}




func methodDemo() {
	//定义了一个struct叫做长方形，你现在想要计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
   //  todo   上面的实现没有问题  这样实现当然没有问题咯，但是当需要增加圆形、正方形、五边形甚至其它多边形的时候，你想计算他们的面积的时候怎么办啊？那就只能增加新的函数咯，但是函数名你就必须要跟着换了，变成area_rectangle, area_circle, area_triangle...

   //很显然，这样的实现并不优雅，并且从概念上来说"面积"是"形状"的一个属性，它是属于这个特定的形状的，就像长方形的长和宽一样。

    //  todo  基于上面的原因所以就有了method的概念，method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。
	//用上面提到的形状的例子来说，method area() 是依赖于某个形状(比如说Rectangle)来发生作用的。Rectangle.area()的发出者是Rectangle， area()是属于Rectangle的方法，而非一个外围函数。
	//更具体地说，Rectangle存在字段 height 和 width, 同时存在方法area(), 这些字段和方法都属于Rectangle。


	//用Rob Pike的话来说就是：
	//"A method is a function with an implicit first argument, called a receiver."

	//method的语法如下：
	//func (r ReceiverType) funcName(parameters) (results)
	r11 := Rectangle{12, 2}
	r22 := Rectangle{9, 4}
	c11 := Circle{10}
	c22 := Circle{25}

	fmt.Println("Area of r1 is: ", r11.area())
	fmt.Println("Area of r2 is: ", r22.area())
	fmt.Println("Area of c1 is: ", c11.area())
	fmt.Println("Area of c2 is: ", c22.area())


	//  todo   在使用method的时候重要注意几点

	//虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
	//method里面可以访问接收者的字段
	//调用method通过.访问，就像struct里面访问字段一样

	//method area() 分别属于Rectangle和Circle， 于是他们的 Receiver 就变成了Rectangle 和 Circle, 或者说，这个area()方法 是由 Rectangle/Circle 发出的。

	//那是不是method只能作用在struct上面呢？当然不是咯，他可以定义在任何你自定义的类型、内置类型、struct等各种类型上面。这里你是不是有点迷糊了，什么叫自定义类型，自定义类型不就是struct嘛，不是这样的哦，struct只是自定义类型里面一种比较特殊的类型而已，还有其他自定义类型申明，可以通过如下这样的申明来实现。

    //下面这个申明自定义类型的代码
	type ages int
	type money float32
	type months map[string]int
	m := months {
		"January":31,
		"February":28,
		"December":31,
	}
	fmt.Println(m)
	//看到了吗？简单的很吧，这样你就可以在自己的代码里面定义有意义的类型了，实际上只是一个定义了一个别名,有点类似于c中的typedef，例如上面ages替代了int

	//  todo   你可以在任何的自定义类型中定义任意多的method


     //定义了一个slice:BoxList，含有Box
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	//意思 表面积最大的颜色是什么  -----
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())


}

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)
//Color作为byte的别名
type Color byte
//定义了一个struct:Box，含有三个长宽高字段和一个颜色属性
type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box //a slice of boxes
//体积计算公式  Volume()定义了接收者为Box，返回Box的容量
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
//SetColor(c Color)，把Box的颜色改为c
/*
现在让我们回过头来看看SetColor这个method，它的receiver是一个指向Box的指针，是的，你可以使用*Box。想想为啥要使用指针而不是Box本身呢？
我们定义SetColor的真正目的是想改变这个Box的颜色，如果不传Box的指针，那么SetColor接受的其实是Box的一个copy，也就是说method内对于颜色值的修改，其实只作用于Box的copy，而不是真正的Box。所以我们需要传入指针
 */
func (b *Box) SetColor(c Color) {
	/*
	这里你也许会问了那SetColor函数里面应该这样定义*b.Color=c,而不是b.Color=c,因为我们需要读取到指针相应的值。
你是对的，其实Go里面这两种方式都是正确的，当你用指针去访问相应的字段时(虽然指针没有任何的字段)，Go知道你要通过指针去获取这个值，看到了吧，Go的设计是不是越来越吸引你了。
	 */
	b.color = c
}
//BiggestColor()定在在BoxList上面，返回list里面容量最大的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	fmt.Println("WHITE iota===",WHITE) //WHITE iota=== 0
	k := Color(WHITE)//WHITE 是枚举
	fmt.Println("k===",k.String())//k=== WHITE
	fmt.Println("b1的全部的值",bl)
	//这是两次的循环  第二次循环才会 走到这样
	//       todo  b1的全部的值 [{4 4 4 1} {10 10 1 1} {1 1 20 1} {10 10 1 1} {10 30 1 1} {20 20 20 1}]
	for _, b := range bl {
		fmt.Println("每次循环的时候 b=",b)  //其实b呢 就是一个 Box
		if bv := b.Volume(); bv > v {
			fmt.Println("每次循环的时候 最大的表面积bv==",bv)
			fmt.Println("每次循环的时候 最大的表面积的颜色b.color==",b.color)
			v = bv
			k = b.color
		}
	}
	return k
}
//PaintItBlack()把BoxList里面所有Box的颜色全部变成黑色
func (bl BoxList) PaintItBlack() {
	//如果省略掉的话，这里就是这个意思 代表的是角标
	for i := range bl {
		fmt.Println("每次变成黑色的时候，这个角标是多少------》",i)
		/*
		也许细心的读者会问这样的问题，PaintItBlack里面调用SetColor的时候是不是应该写成(&bl[i]).SetColor(BLACK)，因为SetColor的receiver是*Box，而不是Box。你又说对了，这两种方式都可以，因为Go知道receiver是指针，他自动帮你转了。
		*/
		bl[i].SetColor(BLACK)
	}


	//也就是说：
	//如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
	//类似的
	//如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
	//所以，你不用担心你是调用的指针的method还是不是指针的method，Go知道你要做的一切，这对于有多年C/C++编程经验的同学来说，真是解决了一个很大的痛苦。


	//后面代表的是value的值
	for _,i :=range bl {
		fmt.Println("这个i值应该是 value值 我觉得=-----》》》",i)
	}
}
//String()定义在Color上面，返回Color的具体颜色(字符串格式)
//  todo  为啥打印出来是  字符串的格式   原因 就是在这里
func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}





type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

//func (c Circle) area() float64 {
//	return c.radius * c.radius * math.Pi
//}
//func (r ReceiverType) funcName(parameters) (results)
func (c Circle)area() float64 {
	return c.radius*c.radius*math.Pi
}

//定义了一个struct叫做长方形，你现在想要计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现
type Rectangle struct {
	width, height float64
}
//这段代码可以计算出来长方形的面积，但是area()不是作为Rectangle的方法实现的（类似面向对象里面的方法），而是将Rectangle的对象（如r1,r2）作为参数传入函数计算面积的。
func area(r Rectangle) float64 {
	return r.width*r.height
}


func main() {

}
