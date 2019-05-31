package main

import (
	"fmt"

	//"time"
	"time"
)

func notify(id int, channel chan int){
	<-channel//接收到数据或通道关闭时退出阻塞
	fmt.Printf("%d receive a message.\n", id)
}

func broadcast(channel chan int){
	fmt.Printf("Broadcast:\n")
	close(channel)//关闭通道
}

//超时设置
func timeout()  {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()

	go func() {
		for i:=0;i<10000000;i++{
			c <- i
		}
	}()
	<- o

}

func main()  {

	 //fmt.Println("测试数据")
	 //var wg sync.WaitGroup //计数信号量
	 //wg.Add(2)
	 //
	 //go func(){
	 //	defer wg.Done()
	 //	for i:=0;i < 1000;i++{
	 //		fmt.Println("Hello,Go.This is %d\n",i)
		//}
	 //}()
	 //
	 //go func() {
	 //	defer wg.Done()
	 //	for i:=0;i < 1000; i ++{
	 //		fmt.Println("Hello,World.This is %d\n",i)
		//}
	 //}()
	 //wg.Wait()





		//channel := make(chan int,1)
		//
		//for i:=0;i<10 ;i++  {
		//	go notify(i,channel)
		//}
		//go broadcast(channel)
		//time.Sleep(time.Second)

	//ch := make(chan int)
	//go func() {
	//	var sum int = 0
	//	for i := 0; i < 10; i++ {
	//		sum += i
	//	}
	//	//发送数据到通道
	//	ch <- sum
	//}()
	////从通道接收数据
	//fmt.Println(<-ch)





}
