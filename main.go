package main

import (
	"fmt"
	"md/initialization"
)

func main() {
	gin := initialization.Gin
	fmt.Println(gin.GetString())
}
