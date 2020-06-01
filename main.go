package main

import (
	// "github.com/AdrianMendez1199/simple-crud/handlers"
	"fmt"

	"github.com/AdrianMendez1199/simple-crud/models"
)

func main() {
	user := new(models.User)
	user.CreateUser("Adrian", "Mendez", "mendezadrian149@gamil.com", "123456")
	fmt.Println(user)

}
