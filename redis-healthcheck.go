package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var ctx = context.Background()
var arguments = os.Args

func rClient() *redis.Client {
	redisAddr := arguments[1]
	redisPassword := arguments[2]

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		PoolSize: 1,
	})

	return client
}

func ping(client *redis.Client) error {
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)

	return nil
}

func healthcheck(c echo.Context) error {
	// creates a client
	client := rClient()

	// check connection status
	err := ping(client)
	client.Close()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Redis ERROR\n")
	}

	return c.String(http.StatusOK, "Redis OK\n")
}

func main() {
	if len(arguments) == 2 {
		fmt.Println("Please provide host:port <password>.")
		os.Exit(1)
	}

	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Route => handler
	e.GET("/healthcheck", healthcheck)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
