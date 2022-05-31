package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"math/rand"
)

type ticks struct {
	Time time.Time `json:"time"`
	Symbol string `json:"symbol"`
	Open float32 `json:"open"`
	High float32 `json:"high"`
	Low float32 `json:"low"`
	Close float32 `json:"close"`
	Volume int `json:"volume"`
}

var arr = []ticks{} // global variable for the ticks cache
var isRunning = false // to check whether server is running or not


// main function to create intial 10 ticks and then changing random ticks
func updateTicks() {
	if len(arr) < 10{
		for i := 0; i < 10; i++{
			tick := ticks{Open: 100.00, High: 100.00, Low: 100.00, Close: 100.00, Volume: 10000, Symbol: fmt.Sprintf("%x", rand.Intn(1000))}
			arr = append(arr, tick)
		}
	}

	num := rand.Intn(9)
	close := arr[num].Close + (rand.Float32() * ( (arr[num].Close + 10/arr[num].Close*100) - (arr[num].Close - 10/arr[num].Close*100) ))
	arr[num].Close = close
	if close > arr[num].High{
		arr[num].High = close
	}
	if close < arr[num].Low{
		arr[num].Low = close
	}

	arr[num].Volume = rand.Intn(1000)
	arr[num].Time = time.Now()

	time.Sleep(time.Duration(100000000)) // for server sleep for 100 ms before changing a tick

	if isRunning {
			updateTicks()
	}

}

func HomePage(context *gin.Context) {
    context.JSON(200, arr)
}

func main() {
    isRunning = true
    go updateTicks() // go routine to update ticks by calling the function
    r := gin.Default()
    r.GET("/", HomePage)
    r.Run(":9000")
}
