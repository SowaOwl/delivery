package main

import (
	"crypto/rand"
	"delivery/cmd"
	"delivery/internal/auth"
	"delivery/internal/cache"
	"delivery/internal/repository"
	"encoding/base64"
	"github.com/redis/go-redis/v9"
	"log"
)

func generateHMACKey() (string, error) {
	key := make([]byte, 32) // 256-битный ключ
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func main() {
	//db := cmd.InitDataBase()
	//
	//r := gin.Default()
	//
	//api.RoutesDefine(r, db)
	//
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//}

	db := cmd.InitDataBase()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "172.19.0.1:6379",
		Password: "root_pass",
	})

	jwtService := auth.NewJWTService("KRW+eik+tIYsM/48NElXUXz3DcGAtTV4WJYxhHDhJz0=")
	redisCache := cache.NewRedisCache(redisClient)
	userRepo := repository.NewGormUserRepository(db)

	service := auth.NewAuthService(jwtService, redisCache, userRepo)

	token, err := service.Login("User", "123")
	if err != nil {
		log.Fatal(err)
	}

	println(token)

	err = service.Authenticate(token)
	if err != nil {
		log.Fatal(err)
	}
}
