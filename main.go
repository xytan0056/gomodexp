package main

import (
 "fmt"
 a "github.com/xytan0056/depwithgen/api"
 c "github.com/xytan0056/depwithgen/.gen/client"
)

func main() {
	a.API()
	fmt.Println(c.Client("production"))
	fmt.Println("well...calling" )
}
