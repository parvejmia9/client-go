package main

import "client-go/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

}
