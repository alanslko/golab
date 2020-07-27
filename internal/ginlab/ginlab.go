package ginlab

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/gin-gonic/gin"
)

func currentContext() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
	fmt.Println("Username: " + user.Username)
	fmt.Println("Home Dir: " + user.HomeDir)

	// Get "Real" User under sudo.
	// More Info: https://stackoverflow.com/q/29733575/402585
	fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
}

func GinLab() {
	currentContext()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		t := time.Now().String()

		c.JSON(200, gin.H{
			"message": "pong " + t,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
