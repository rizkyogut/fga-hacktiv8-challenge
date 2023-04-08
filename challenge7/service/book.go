package service

import "challenge/challenge7/model"

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(id int64, data model.Book) (err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
	// return s.repo.CreateEmployee(in)
}

func (s *Service) GetAllBook() (res []model.Book, err error) {
	return s.repo.GetAllBook()
}

func (s *Service) GetBookByID(id int64) (res model.Book, err error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) UpdateBook(id int64, data model.Book) (err error) {
	return s.repo.UpdateBook(id, data)
}

func (s *Service) DeleteBook(id int64) (err error) {
	err = s.repo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
