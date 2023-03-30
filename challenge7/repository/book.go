package repository

import (
	"challenge/challenge7/model"
	"challenge/challenge7/repository/query"
	"database/sql"
	"errors"
)

type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(int, model.Book) error
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(data model.Book) (res model.Book, err error) {
	res = model.Book{}

	err = r.db.QueryRow(
		query.AddBook,
		data.Title,
		data.Author,
		data.Desc,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Desc,
	)

	if err != nil {
		return model.Book{}, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	res = []model.Book{}

	rows, err := r.db.Query(query.GetAllBook)
	if err != nil {
		return []model.Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return []model.Book{}, err
		}
		res = append(res, book)

	}

	return res, nil
}

func (r Repo) GetBookByID(id int64) (res model.Book, err error) {
	res = model.Book{}

	rows := r.db.QueryRow(query.GetBookById, id)

	err = rows.Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Book{}, errors.New("DATA NOT FOUND")
		}
		return model.Book{}, err
	}

	return res, nil
}

func (r Repo) UpdateBook(id int, data model.Book) error {
	res, err := r.db.Exec(query.UpdateBook, id, data.Title, data.Author, data.Desc)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("DATA NOT FOUND")
	}
	return nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	res, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("DATA NOT FOUND")
	}
	return nil
}
