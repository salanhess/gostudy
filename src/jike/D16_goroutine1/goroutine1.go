package main

import (
	"fmt"
	"time"
)

/*
 Go 语言不但有着独特的并发编程模型，以及用户级线程 goroutine，还拥
 有强大的用于调度 goroutine、对接系统级线程的调度器。
 这个调度器是 Go 语言运行时系统的重要组成部分，它主要负责统筹调配 Go 并发编程
 模型中的三个主要元素，即： G（goroutine 的缩写）、 P（processor 的缩写）和 M
 （machine 的缩写）。
 其中的 M 指代的就是系统级线程。而 P 指的是一种可以承载若干个 G，且能够使这些
 G 适时地与 M 进行对接，并得到真正运行的中介。
 从宏观上说， G 和 M 由于 P 的存在可以呈现出多对多的关系。当一个正在与某个 M 对
 接并运行着的 G，需要因某个事件（比如等待 I/O 或锁的解除）而暂停运行的时候，调
 度器总会及时地发现，并把这个 G 与那个 M 分离开，以释放计算资源供那些等待运行
 的 G 使用。
*/

/*
在拿到了一个空闲的 G 之后， Go 语言运行时系统会用这个 G 去包装当前的那个go函数
（或者说该函数中的那些代码），然后再把这个 G 追加到某个存放可运行的 G 的队列
中。
这类队列中的 G 总是会按照先入先出的顺序，很快地由运行时系统内部的调度器安排运
行。虽然这会很快，但是由于上面所说的那些准备工作还是不可避免的，所以耗时还是存
在的。
因此， go函数的执行时间总是会明显滞后于它所属的go语句的执行时间。当然了，这里
所说的“明显滞后”是对于计算机的 CPU 时钟和 Go 程序来说的。我们在大多数时候都不
会有明显的感觉。
在说明了原理之后，我们再来看这种原理下的表象。请记住，只要go语句本身执行完
毕， Go 程序完全不会等待go函数的执行，它会立刻去执行后边的语句。这就是所谓的异
步并发地执行。这里“后边的语句”指的一般是for语句中的下一个迭代。然而，当最后一个迭代运行的时
候，这个“后边的语句”是不存在的。
在 demo38.go 中的那条for语句会以很快的速度执行完毕。当它执行完毕时，那 10 个
包装了go函数的 goroutine 往往还没有获得运行的机会。
*/
func Start_goroutine1() {
	for j := 0; j < 10; j++ {
		go func() {
			fmt.Println(j)
			//fmt.Println("vim-go ", j)
		}()
	}
}

func Start_goroutine3() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("vim-go ", i)
		}(i)
	}
}

//With sleep
func Start_goroutine2() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("vim-go ", i)
		}()
	}
	time.Sleep(1 * time.Microsecond)
}
func main() {
	Start_goroutine1()
	time.Sleep(2 * time.Second)
	fmt.Println("=========Add sleep========")
	Start_goroutine2()
	fmt.Println("=========Add para========")
	Start_goroutine3()
}
