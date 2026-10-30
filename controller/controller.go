package controller

import "github.com/hci-group-15/backend/model"

type Controller struct {
	db *model.DBConnector
}

func NewController(dbcon *model.DBConnector) *Controller {
	return &Controller{
		db: dbcon,
	}
}
