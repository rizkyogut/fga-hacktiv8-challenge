package query

const (
	AddBook = `
	INSERT INTO db_go_sql.books (title, author, "description")
	VALUES ($1, $2, $3) RETURNING *;
	`

	GetAllBook = `
	SELECT * FROM db_go_sql.books ;
	`
	GetBookById = `
	SELECT * FROM db_go_sql.books
	WHERE id = $1;
	`
	UpdateBook = `
	UPDATE db_go_sql.books
	SET title = $2, author = $3 , "description" = $4
	WHERE id = $1 ; 
	`
	DeleteBook = `
	DELETE FROM db_go_sql.books
	WHERE id = $1 ; 
	`
)
