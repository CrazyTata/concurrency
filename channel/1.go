package main

import "fmt"

func main() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

//第一次循环：0被发送到通道ch中，此时i=0。
//
//第二次循环：1被发送到通道ch中，此时i=1。
//
//第三次循环：通道ch中已经有两个元素（0和1），所以执行case v := <-ch:，v的值被赋为0，然后输出0，此时i=2。
//
//第四次循环：通道ch中已经有一个元素（1），所以执行case v := <-ch:，v的值被赋为1，然后输出1，此时i=3。
//
//第五次循环：4被发送到通道ch中，此时i=4。
//
//第六次循环：5被发送到通道ch中，此时i=5。
//
//第七次循环：通道ch中已经有两个元素（4和5），所以执行case v := <-ch:，v的值被赋为4，然后输出4，此时i=6。
//
//第八次循环：通道ch中已经有一个元素（5），所以执行case v := <-ch:，v的值被赋为5，然后输出5，此时i=7。
//
//第九次循环：8被发送到通道ch中，此时i=8。
//
//第十次循环：通道ch中已经有一个元素（8），所以执行case v := <-ch:，v的值被赋为8，然后输出8，此时i=9。
//
//循环结束，没有其他发送操作。
//因此，最终的输出结果为：
//
//0
//1
//4
//5
//8
