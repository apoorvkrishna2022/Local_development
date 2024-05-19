package route

import (
	myuitls "github.com/apoorvkrishna22/prometheus_local_testing/uitls"
	"time"
)

//func randomInt(min, max int) int {
//	rand.Seed(time.Now().UnixNano())
//	return min + rand.Intn(max-min)
//}

func ReturnResponse() string {
	switch result := myuitls.RandomInt(1, 100) % 3; result {
	case 1:
		return "200"
	case 2:
		return "500"
	case 0:
		return "400"
	default:
		return "300"
	}
}

func SleepRouteInMS() time.Duration {
	switch result := myuitls.RandomInt(1, 100) % 14; result {
	case 0:
		return 1 * time.Millisecond
	case 1:
		return 3 * time.Millisecond
	case 2:
		return 5 * time.Millisecond
	case 3:
		return 10 * time.Millisecond
	case 4:
		return 15 * time.Millisecond
	case 5:
		return 30 * time.Millisecond
	case 6:
		return 50 * time.Millisecond
	case 7:
		return 100 * time.Millisecond
	case 8:
		return 200 * time.Millisecond
	case 9:
		return 25 * time.Millisecond
	case 10:
		return 10 * time.Millisecond
	case 11:
		return 200 * time.Millisecond
	case 12:
		return 40 * time.Millisecond
	case 13:
		return 75 * time.Millisecond
	default:
		return 70 * time.Millisecond
	}
}

func GetTicket() string {
	time.Sleep(SleepRouteInMS())
	return ReturnResponse()
}

func AddTicket() string {
	time.Sleep(SleepRouteInMS())
	return ReturnResponse()
}

func UpdateTicket() string {
	time.Sleep(SleepRouteInMS())
	return ReturnResponse()
}

func DelTicket() string {
	time.Sleep(SleepRouteInMS())
	return ReturnResponse()
}
