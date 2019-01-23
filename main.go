package main

import "time"

//当打印完100个数的时候让程序停掉
var channel chan bool

func main() {
	//创建信道(默认无缓冲)
	channel = make(chan bool)

	go test1(0)

	println("--1--")

	//取不出时阻塞
	<-channel

	println("--2--")

	go test(1)

	println("--3--")

	<-channel

	println("--4-")

}

func test(index int) {
	time.Sleep(time.Second)
	//没被取出时阻塞
	channel <- true
	println(index, " ")
}

func test1(index int) {
	time.Sleep(2*time.Second)
	channel <- true
	println(index, " ")
}
