package handler

//route, controller

import (
	"challenge/challenge7/handler/helper"
	"challenge/challenge7/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	book := model.Book{}

	err := c.BindJSON(&book)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.service.CreateBook(book)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.service.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	if len(res) == 0 {
		helper.NoContent(c)
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.service.GetBookByID(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in.ID = idInt
	// call service
	err = h.service.UpdateBook(int64(idInt), in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, in)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.service.DeleteBook(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Successfully deleted Book")
}
