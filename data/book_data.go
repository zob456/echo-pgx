package data

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/zob456/echo-pgx/models"
	"github.com/zob456/echo-pgx/utils"
	"log"
)

var dbConn *pgx.Conn

func init() {
	dbConn = utils.ConnDB()
}

func SelectBook(id string, ctx context.Context) (*models.Book, error) {
	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const statement = `SELECT
       "ID",
       "Title",
       "AuthorID"
	FROM "Library"."Book"
	WHERE "ID" = $1;`
	book := &models.Book{}

	row := dbConn.QueryRow(ctx, statement, id)

	err := row.Scan(&book.ID, &book.Title, &book.AuthorID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return book, nil
}
