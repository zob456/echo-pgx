package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func ConnDB() *pgx.Conn {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"localhost", "5432", "postgres", "postgres", "postgres", "disable")
	conn, err  := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}
