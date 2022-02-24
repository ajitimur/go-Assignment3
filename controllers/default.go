package controllers

import (
	"assignment3/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

var result models.Status

func (c *MainController) Get() {
	result = models.GetStatus()

	go Looper()

	c.Data["json"] = result
	c.ServeJSON()
}

func Looper() {

	ch := time.Tick(15 * time.Second)
	for next := range ch {
		fmt.Println(next)
		result = models.GetStatus()

		fmt.Println(result)
	}
}
