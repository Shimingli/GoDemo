package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("面向并发的内存模型")
}

func main() {
	fmt.Println("")

	//在早期，CPU都是以单核的形式顺序执行机器指令。Go语言的祖先C语言正是这种顺序编程语言的代表。顺序编程语言中的顺序是指:所有的指令都是以串行的方式执行，在相同的时刻有且仅有一个CPU在顺序执行程序的指令。

	//随着处理器技术的发展，单核时代以提升处理器频率来提高运行效率的方式遇到了瓶颈，目前各种主流的CPU频率基本被锁定在了3GHZ附近。单核CPU的发展的停滞，给多核CPU的发展带来了机遇。相应地，编程语言也开始逐步向并行化的方向发展。Go语言正是在多核和网络化的时代背景下诞生的原生支持并发的编程语言。

	//常见的并行编程有多种模型，主要有多线程、消息传递等。从理论上来看，多线程和基于消息的并发编程是等价的。由于多线程并发模型可以自然对应到多核的处理器，主流的操作系统因此也都提供了系统级的多线程支持，同时从概念上讲多线程似乎也更直观，因此多线程编程模型逐步被吸纳到主流的编程语言特性或语言扩展库中。而主流编程语言对基于消息的并发编程模型支持则相比较少，Erlang语言是支持基于消息传递并发编程模型的代表者，它的并发体之间不共享内存


	//Go语言是基于消息并发模型的集大成者，它将基于CSP模型的并发编程内置到了语言中，通过一个go关键字就可以轻易地启动一个Goroutine，与Erlang不同的是Go语言的Goroutine之间是共享内存的


	goroutine()



}

func goroutine() {
	//Goroutine是Go语言特有的并发体，是一种轻量级的线程，由go关键字启动。在真实的Go语言的实现中，goroutine和系统线程也不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了Go语言并发编程质的飞跃

  //首先，每个系统级线程都会有一个固定大小的栈（一般默认可能是2MB），这个栈主要用来保存函数递归调用时参数和局部变量。固定了栈的大小这导致了两个问题：一是对于很多只需要很小的栈空间的线程来说是一个巨大的浪费，二是对于少数需要巨大栈空间的线程来说又面临栈溢出的风险。针对这两个问题的解决方案是：要么降低固定的栈大小，提升空间的利用率;要么增大栈的深度以允许更深的函数递归调用，但这两者是没法同时兼得的。相反，一个Goroutine会以一个很小的栈启动（可能是2KB或4KB），当遇到深度递归导致当前栈空间不足时，Goroutine会根据需要动态地伸缩栈的大小（主流实现中栈的最大值可达到1GB）。因为启动的代价很小，所以我们可以轻易地启动成千上万个Goroutine。



	//Go的运行时还包含了其自己的调度器，这个调度器使用了一些技术手段，可以在n个操作系统线程上多工调度m个Goroutine。Go调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的Goroutine。Goroutine采用的是半抢占式的协作调度，只有在当前Goroutine发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得多。运行时有一个runtime.GOMAXPROCS变量，用于控制当前运行正常非阻塞Goroutine的系统线程数目。
	// todo 运行时有一个runtime.GOMAXPROCS变量，用于控制当前运行正常非阻塞Goroutine的系统线程数目。
	runtime.GOMAXPROCS(10)

   //在Go语言中启动一个Goroutine不仅和调用函数一样简单，而且Goroutine之间调度代价也很低，这些因素极大地促进了并发编程的流行和发展




}