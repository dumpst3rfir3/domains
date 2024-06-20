package main

import (
	"domains"
	"fmt"
)

func main() {
	for _, domain := range domains.List() {
		fmt.Println(domain)
	}

}
