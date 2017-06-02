package controllers

import (
	"github.com/h-tko/revel-base/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	userModel := models.NewUser()

	revel.ERROR.Printf("%d\n\n", userModel.AllJoinSample()[0].UserParam.Stamina)

	return c.Render()
}
