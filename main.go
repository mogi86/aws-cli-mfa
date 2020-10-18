package main

import (
	"fmt"

	"github.com/mogi86/mogi-awscli/cmd"
)

func main()  {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("cmd failed: %v", err)
	}
}
