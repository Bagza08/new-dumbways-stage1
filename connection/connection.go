package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConn() {
	databaseurl := "postgres://postgres:123@localhost:5432/db_personal_web" // ini adalah connection string

	var err error

	Conn, err = pgx.Connect(context.Background(), databaseurl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "DATABASE GAGAL CONNECT !: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n\n\tSELAMAT DATABASE BERHASIL CONNECT!")

}
