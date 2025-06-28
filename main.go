package main

import (
	"fmt"
	"os"

	"github.com/MegeKaplan/gobox/cmd"
	"github.com/MegeKaplan/gobox/internal/storage"
)

func main() {
    if err := storage.Init(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
