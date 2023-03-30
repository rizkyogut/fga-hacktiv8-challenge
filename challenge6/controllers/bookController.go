package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	BookID int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	Year   string `json:"year"`
}

var BookDatas []Book

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if len(BookDatas) > 0 {
		newBook.BookID = BookDatas[len(BookDatas)-1].BookID + 1
	} else {
		newBook.BookID = 1
	}
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book":    newBook,
		"message": fmt.Sprintf("Book has been sucessfully to Created"),
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))

	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].BookID = bookID
			return
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully updated", bookID),
	})
}

func DeletedBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("the book you're trying to delete with id %v is not exist", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been sucessfully deleted", bookID),
	})
}
