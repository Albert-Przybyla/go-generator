package main

import (
	"fmt"

	"github.com/Albert-Przybyla/go-generator/config"
)

func main() {
	res, err := config.LoadConfig(".conf.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("Config:", res.Paths.Models)
	fmt.Println("Config:", res.Paths.Handlers)
	fmt.Println("Config:", res.Paths.Repositories)
	fmt.Println("Config:", res.Paths.Services)
}
