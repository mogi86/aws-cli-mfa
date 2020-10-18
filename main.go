package main

import (
	"flag"
	"fmt"
)

func main()  {
	var greeting string

	//f := flag.String("greeting", "", "please set the greeting.")
	flag.StringVar(&greeting, "greeting", "", "please set the greeting.")
	flag.Parse()

	fmt.Println(greeting)
}
