package main

import (
	"flag"
	"mathalgorithms"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = flag.String("port", ":8080", "Port to use")

	flag.Parse()

	r := gin.Default()
	r.GET("/fib", func(c *gin.Context) {
		fibResult := []int32{}
		fib := mathalgorithms.NewFibonacciIterator()
		for fib.Next() {
			fibResult = append(fibResult, fib.Value())
			if fib.Value() == 34 {
				break
			}
		}

		c.JSON(200, fibResult)
	})
	r.Run(*port)
}
