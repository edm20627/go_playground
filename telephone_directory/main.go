package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type Record struct {
	ID    int64
	Name  string
	Phone string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func run() error {
	db, err := sql.Open("sqlite", "mydata.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	if err := createTable(db); err != nil {
		return err
	}

	for {
		if err := showRecords(db); err != nil {
			return err
		}

		fmt.Println("INSERT(i) or UPDATE(u) or FIND(f)")
		var q string
		fmt.Scan(&q)

		switch q {
		case "i", "insert", "INSERT":
			if err := inputRecord(db); err != nil {
				return err
			}
		case "u", "update", "UPDATE":
			if err := updateRecord(db); err != nil {
				return err
			}
		case "f", "find", "FIND":
			if err := showRecord(db); err != nil {
				return err
			}
		default:
			fmt.Println("INSERT か UPDATE を選んでください")
			break
		}
	}

	return nil
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS addressbook (
			id    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name  TEXT NOT NULL,
			phone TEXT NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("全件表示")
	rows, err := db.Query("SELECT * FROM addressbook")
	if err != nil {
		return err
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Phone); err != nil {
			return err
		}
		fmt.Printf("[%d] Name:%s TEL:%s\n", r.ID, r.Name, r.Phone)
	}

	fmt.Println("------")

	return nil
}

func showRecord(db *sql.DB) error {
	fmt.Println("1件表示")

	var r Record

	fmt.Println("ID >")
	_, err := fmt.Scan(&r.ID)
	if err != nil {
		return err
	}

	row := db.QueryRow("SELECT * FROM addressbook WHERE id = ?", r.ID)
	if err := row.Scan(&r.ID, &r.Name, &r.Phone); err != nil {
		return err
	}
	fmt.Printf("[%d] Name:%s TEL:%s\n", r.ID, r.Name, r.Phone)

	fmt.Println("------")

	return nil
}

func inputRecord(db *sql.DB) error {
	fmt.Println("--- INSERT ---")

	var r Record

	fmt.Println("Name >")
	fmt.Scan(&r.Name)

	fmt.Println("TEL >")
	fmt.Scan(&r.Phone)

	const sql = "INSERT INTO addressbook(name, phone) values (?, ?)"
	_, err := db.Exec(sql, r.Name, r.Phone)
	if err != nil {
		return err
	}

	fmt.Println("------")

	return nil
}

func updateRecord(db *sql.DB) error {
	fmt.Println("--- UPDATE ---")

	var r Record

	fmt.Println("ID >")
	_, err := fmt.Scan(&r.ID)
	if err != nil {
		return err
	}

	fmt.Println("Name >")
	fmt.Scan(&r.Name)

	fmt.Println("TEL >")
	fmt.Scan(&r.Phone)

	const sql = "UPDATE addressbook SET name = ?, phone = ? WHERE id = ?"
	result, err := db.Exec(sql, r.Name, r.Phone, r.ID)

	if err != nil {
		return err
	}

	cnt, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(cnt, "件更新しました")
	fmt.Println("------")

	return nil
}
