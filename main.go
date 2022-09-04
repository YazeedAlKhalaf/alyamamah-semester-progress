package main

import (
	"fmt"
	"os"

	"github.com/yazeedalkhalaf/alyamamah-sp/internal"
)

func main() {
	config, err := internal.LoadConfig(".")
	if err != nil {
		fmt.Println("error happenned:", err)
		os.Exit(1)
	}

	fmt.Println("config:", config)
}
