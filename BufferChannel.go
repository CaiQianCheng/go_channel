package main

var bufferedChannel chan bool

func main() {

	//缓存5个的通道
	bufferedChannel = make(chan bool, 5)


}



