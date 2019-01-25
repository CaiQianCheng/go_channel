package main
import "time"

var bufferedChannel chan int
func main() {
	//缓存5个的通道
	bufferedChannel = make(chan int, 5)
	go testFun()
	//读不到会阻塞
	println(<-bufferedChannel)
	println(<-bufferedChannel)
	println(<-bufferedChannel)
	println(<-bufferedChannel)
	println(<-bufferedChannel)
}

func testFun() {
	//写满5个才阻塞
	for i := 0; i < 5; i++ {
		bufferedChannel <- i
		time.Sleep(time.Second)
	}
}
