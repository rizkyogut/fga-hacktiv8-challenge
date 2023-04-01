package repository

import (
	"challenge/project1/model"
)

type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r Repo) GetBookByID(id int64) (res model.Book, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).First(&res).Updates(model.Book{
		Title:  in.Title,
		Author: in.Author,
		Desc:   in.Desc,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	book := model.Book{}
	err = r.gorm.Where("id = ?", id).First(&book).Delete(&book).Error

	if err != nil {
		return err
	}

	return err
}
