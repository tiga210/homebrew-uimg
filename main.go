package main

import (
	"fmt"
	"os"
	"uimg/api"
	"uimg/model"
)

func main() {
	args := os.Args

	if len(args) == 1 || args[1] == "-h" {
		fmt.Println("Example usage:")
		fmt.Println("  uimg -i ")
		fmt.Println("  uimg /1.png /2.jpg")
		return
	}

	if args[1] == "-i" {
		model.InitConfig()
		return
	}

	filesNames := args[1:]

	c := model.GetConfig()
	if c == nil || c.Host == "" || c.Token == "" {
		fmt.Println("please init: uimg -i")
		return
	}

	for _, f := range filesNames {
		status, imgUrl := api.UploadImg(f, c.Token, c.Host)
		if status {
			fmt.Println(imgUrl)
		}
	}

}
