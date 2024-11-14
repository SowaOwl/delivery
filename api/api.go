package api

import (
	"delivery/internal/auth"
	"delivery/internal/cache"
	"delivery/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RoutesDefine(r *gin.Engine, db *gorm.DB) {
	r.GET("/test", func(c *gin.Context) {

	})
}

func testRoute(c *gin.Context, db *gorm.DB) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "172.19.0.1:6379",
		Password: "root_pass",
	})

	jwtService := auth.NewJWTService("")
	redisCache := cache.NewRedisCache(redisClient)
	userRepo := repository.NewGormUserRepository(db)

	service := auth.NewAuthService(jwtService, redisCache, userRepo)

	token, err := service.Login("User", "123")
	if err != nil {
		c.JSON(200, err)
	}

	c.JSON(200, token)
}
