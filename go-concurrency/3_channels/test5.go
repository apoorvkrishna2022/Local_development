//package main
//
//import (
//	"fmt"
//	"go-concurrency/3_channels/Channel"
//	"math/rand"
//	"sync"
//)
//
//func generateRandomNumber() int {
//	maxRange := 1000
//	return rand.Intn(maxRange)
//}
//
//func transform(val interface{}) string {
//	temp := fmt.Sprintf("The value is %d", val)
//	return temp
//
//}
//
//func main() {
//
//	length := 10
//	maxNumberOfGoroutines := 8
//
//	myChannel := Channel.NewSafeChannel(length + 1)
//	for i := 0; i < length; i++ {
//		myChannel.Push(generateRandomNumber())
//	}
//	myChannel.Close()
//
//	myResultChannel := Channel.NewSafeChannel(length + 1)
//
//	var wg sync.WaitGroup
//
//	for maxNumberOfGoroutines > 0 {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for {
//				data, ok := myChannel.GetData()
//				if !ok {
//					return
//				}
//				myResultChannel.Push(transform(data))
//			}
//		}()
//		maxNumberOfGoroutines--
//	}
//
//	wg.Wait()
//	myResultChannel.Close()
//	myresultChannelData, err := myResultChannel.GetChannelArray()
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	for _, ans := range *myresultChannelData {
//		fmt.Println(ans)
//	}
//
//}
