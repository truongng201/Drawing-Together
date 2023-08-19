package controller

import (
	database "server/pkg/lib/database/sqlc"

	"fmt"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GetUsersController struct {}

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}


type GetUsersResponse struct {
	StatusCode 	int32 					`json:"status_code"`
	Message 	string 					`json:"message"`
	Users 		[]database.GetUsersRow 	`json:"users"`
}

func (controller GetUsersController) handleParams(params url.Values) GetUsersParams {
	var limit, offset string
	if limit := params.Get("limit"); limit == "" {
		limit = "10"
	}

	if offset := params.Get("offset"); offset == "" {
		offset = "0"
	}

	Limit, _  := strconv.Atoi(limit)

	Offset, _ := strconv.Atoi(offset)

	return GetUsersParams{
		Limit: int32(Limit),
		Offset: int32(Offset),
	}
}


func (controller *GetUsersController) Execute(c echo.Context) error {
	params := controller.handleParams(c.QueryParams())
	
	arg := database.GetUsersParams{
		Limit: params.Limit,
		Offset: params.Offset,
	}
	
	return c.JSON(200, "GetUsersController")
}