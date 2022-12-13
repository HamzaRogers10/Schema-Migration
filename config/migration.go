package config

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	n, err := migrate.Exec(DB.DB(), "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
