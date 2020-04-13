// 优雅的关闭channel的方法
package main

import (
	"math/rand"
	"time"
	"log"
	"sync"
	"strconv"
)

// 多个消费者，单个生产者
func test1() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumRecevers = 100

	var wg sync.WaitGroup
	wg.Add(NumRecevers)

	datach := make(chan int, 100)
	go func () {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				close(datach)
				return
			} else {
				datach <- value
			}
		}
	}()
	for i := 0; i < NumRecevers; i++ {
		go func () {
			defer wg.Done()
			for value := range datach {
				log.Println(value)
			}
		}()
	}
	wg.Wait()
}

//多个生产者，单个消费者
func test2() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumSenders = 01000

	var wg sync.WaitGroup
	wg.Add(1)

	dataCh := make(chan int, 1000)
	// 在消费端添加一个附加信号来通知发送端停止生产数据
	stopCh := make(chan struct{})

	// 生产者
	for i := 0; i < NumSenders; i++ {
		go func () {
			for {
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(MaxRandomNumber):
				}
			}
		}()
	}

	// 消费者
	go func () {
		defer wg.Done()
		for value := range dataCh {
			if value == MaxRandomNumber - 1 {
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()
	wg.Wait()
}


//多个生产者，多个消费者
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumRecevers = 10
	const NumSenders = 1000

	var wg sync.WaitGroup
	wg.Add(NumRecevers)

	dataCh := make(chan int, 100)
	// 在消费端添加一个附加信号来通知发送端停止生产数据
	stopCh := make(chan struct{})
	// 引入一个chan来关闭stopCh
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func () {
		stoppedBy = <- toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumRecevers; i++ {
		go func (id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}
				//the first select here is to try to exit the goroutine as early as possible.
				// select {
				// case <- stopCh:
				// 	return
				// default:
				// }
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(MaxRandomNumber):
				}
			}	
		}(strconv.Itoa(i))
	}

	// 接收者
	for i := 0; i < NumRecevers; i++ {
		go func (id string) {
			defer wg.Done()
			for {
				// same as senders, the first select here is to 
				// try to exit the goroutine as early as possible.
				select {
					case <- stopCh:
						return
					default:
				}

				select {
					case <- stopCh:
						return
					case value := <-dataCh:
						if value == MaxRandomNumber - 1 {
							// the same trick is used to notify the moderator 
                        	// to close the additional signal channel.
							select {
							case toStop <- "receiver#" + id:
							default:
							}
							return
						}
						log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
	wg.Wait()
	log.Println("stopped by", stoppedBy)
}