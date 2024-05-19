package Channel

import (
	"errors"
	"sync"
)

type SafeChannel struct {
	ch     chan interface{}
	closed bool
	mutex  sync.Mutex
}

func NewSafeChannel(bufferSize int) *SafeChannel {
	return &SafeChannel{
		ch: make(chan interface{}, bufferSize),
	}
}

func (sc *SafeChannel) Push(data interface{}) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	if !sc.closed {
		sc.ch <- data
	}
}

func (sc *SafeChannel) Close() {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if sc.closed {
		return
	}

	close(sc.ch)
	sc.closed = true
}

func (sc *SafeChannel) IsChannelClosed() bool {
	return sc.closed
}

func (sc *SafeChannel) GetData() (interface{}, bool) {

	data, ok := <-sc.ch
	return data, ok

}

func (sc *SafeChannel) GetChannelArray() (*[]interface{}, error) {
	if !sc.IsChannelClosed() {
		return nil, errors.New("Channel has not been Closed")
	}
	var channelArray []interface{}
	for item := range sc.ch {
		channelArray = append(channelArray, item)
	}
	return &channelArray, nil
}

//func main() {
//	sc := NewSafeChannel(3)
//
//	for i := 1; i <= 5; i++ {
//		sc.Push(i)
//	}
//
//	sc.Close()
//
//	// Try pushing more data after closing
//	sc.Push(6)
//
//	for {
//		data, ok := <-sc.ch
//		if !ok {
//			break
//		}
//		fmt.Println("Received:", data)
//	}
//}
