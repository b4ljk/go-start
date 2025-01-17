package admin

import (
	databas "back/db"
	"back/services"
	"back/structs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type AdminController struct {
	services.Controller
}

type GetPaginatedUsersParams struct {
	structs.PaginationInput
}

func (co AdminController) Register(router *gin.RouterGroup) {
	{
		router.GET("/", co.Getusers)
		router.GET("/create", co.CreateUsers)
	}
}

func (co AdminController) CreateUsers(c *gin.Context) {
	db := co.DB

	userData := databas.CreateUserParams{
		Email:    "bataa",
		Password: "boldoo",
		UserName: "gg",
		Gender:   pgtype.Text{String: "MALE", Valid: true},
	}

	result, err := db.Queries.CreateUser(c, userData)

	if err != nil {
		co.SetError(500, "Error while creating user")
		c.JSON(500, co.Response.Body)
		return
	}

	co.SetBody(result)
	c.JSON(co.GetBody())
}

func (co AdminController) Getusers(c *gin.Context) {
	db := co.DB

	users, err := db.Queries.ListAllUsers(c)

	if err != nil {
		co.SetError(500, "Error while fetching users")
		c.JSON(500, co.Response.Body)
		return
	}

	co.SetBody(users)
	c.JSON(co.GetBody())
}

func (co AdminController) GetPaginatedUsers(c *gin.Context) {
	db := co.DB

	var params GetPaginatedUsersParams
	if err := c.ShouldBindQuery(&params); err != nil {
		co.SetError(400, "Invalid request")
		c.JSON(400, co.Response.Body)
		return
	}

	users, err := db.Queries.ListUsers(c, struct {
		Limit  int32
		Offset int32
	}{
		Limit:  params.Limit,
		Offset: params.Offset,
	})

	if err != nil {
		co.SetError(500, "Error while fetching users")
		c.JSON(500, co.Response.Body)
		return
	}

	co.SetBody(users)
	c.JSON(co.GetBody())
}
