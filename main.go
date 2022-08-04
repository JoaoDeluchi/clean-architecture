package main

import (
	"fmt"

	"github.com/JoaoDeluchi/clean-architecture/src/entity"
)

func main() {
	u := entity.User{
		Name: "someName",
		Age: 20, 
		LastName: "",
	}

	fmt.Println(u.IsValid())
}
