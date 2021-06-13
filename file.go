package main

import (
	"io"
	"os"
)

func main()  {
	const filepath = "./message.txt"

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		panic(err)
	}

}
