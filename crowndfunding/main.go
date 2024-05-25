package main

import (
	"crowndfunding/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection to databases is good")
	var users []user.User
	db.Find(&users)
	length := len(users)
	fmt.Print(length)

	for _, user := range users {
		fmt.Println("======mengakses semua nilai")
		fmt.Println(user)
		fmt.Println("======mengakses nilai satu persatu")
		fmt.Println(user.ID)
		fmt.Println(user.AvatarFileName)
		fmt.Println(user.CreatedAt)
		fmt.Println(user.Email)
		fmt.Println(user.Name)

	}
}
