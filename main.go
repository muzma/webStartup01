package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "muzma:P@ssw0rd@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//input := user.LoginInput{
	//	Email:    "muzma.niar@gmail.com",
	//	Password: "password",
	//}
	//user, err := userService.Login(input)
	//if err != nil {
	//	log.Println("Terjadi kesalahan")
	//	log.Println(err.Error())
	//}
	//
	//log.Println(user.Email)
	//log.Println(user.Name)

	//userByEmail, err := userRepository.FindByEmail("muzma.niar@gmail.com")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if userByEmail.ID == 0 {
	//	fmt.Println("User tidak ditemukan")
	//} else {
	//	fmt.Println(userByEmail.Name)
	//}

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/session", userHandler.Login)

	router.Run()

	//input dari user
	//handler , mapping input dari user -> input struct input
	//service : melakukan mapping dari struct input ke struct
	//repository
	//db

}
