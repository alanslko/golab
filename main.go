package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/alanslko/golab/internal/database"
	"github.com/alanslko/golab/pkg/formatter"

	"github.com/gin-gonic/gin"
)

// nullable(YES), editable(YES), unique(NO), autofill(NO), refreshable(NO), displayed(YES)

type parent struct {
	name  string
	child `isNullable:"NO" isKey:"YES"`
}

type child struct {
	name string
	age  int
}

func init() {
	fmt.Println(database.GetName())
	fmt.Println(formatter.GetName())
}

func main() {
	p := parent{"cindy", child{"stefan", 7}}
	fmt.Println(p.name)
	var t = reflect.TypeOf(p)

	var tag = t.Field(1).Tag
	fmt.Printf("reflect: %s %s", tag.Get("isNullable"), tag.Get("isKey"))
	fmt.Printf("parent.age %d", p.age)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		now := time.Now().String()
		c.Header("Cache-Control", "no-cache")
		c.JSON(200, gin.H{
			"message": now,
		})
	})
	r.Run()
}
