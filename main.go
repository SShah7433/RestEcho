package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func printDivider() {
	print(colorBlue)

	for i := 1; i < 45; i++ {
		fmt.Printf("-")
	}

	println(colorReset)
}

func main() {
	// Hides all debug info and warning messages
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {

		printDivider()

		// Print URL Parameters (Query Parameters)
		println("--- URL PARAMETERS ---", colorYellow)
		for param, value := range c.Request.URL.Query() {
			fmt.Printf("%s: %s\n", param, value[0])
		}

		println(colorReset)

		// Print Request Body (does not get pretty formatted on purpose)
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("--- REQUEST BODY ---\n%v%v%v\n--- REQUEST BODY END ---\n", colorGreen, string(data), colorReset)

		printDivider()

		// Reply with 200 (always)
		c.Status(http.StatusOK)
	})
	println(colorPurple, "Starting Server...you should assume the server is up", colorReset)
	r.Run() // listen and serve on 0.0.0.0:8080
}
