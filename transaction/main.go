package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type Record struct {
	ID     int64
	Name   string
	Amount int64
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

		fmt.Println("i or u")
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
		default:
			fmt.Println("INSERT か UPDATE を選んでください")
			break
		}
	}
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS bank (
			id    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name  TEXT NOT NULL,
			amount TEXT NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("全件表示")
	rows, err := db.Query("SELECT * FROM bank")
	if err != nil {
		return err
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
			return err
		}
		fmt.Printf("[%d] Name:%s Amount:%d\n", r.ID, r.Name, r.Amount)
	}

	fmt.Println("------")

	return nil
}

func inputRecord(db *sql.DB) error {
	fmt.Println("--- INSERT ---")

	var r Record

	fmt.Println("Name >")
	fmt.Scan(&r.Name)

	fmt.Println("Amount >")
	fmt.Scan(&r.Amount)

	const sql = "INSERT INTO bank(name, amount) values (?, ?)"
	_, err := db.Exec(sql, r.Name, r.Amount)
	if err != nil {
		return err
	}

	fmt.Println("------")

	return nil
}

func updateRecord(db *sql.DB) error {
	fmt.Println("--- UPDATE ---")

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	fmt.Println("from ID >")
	var fromID int64
	_, err = fmt.Scan(&fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("to ID >")
	var toID int64
	_, err = fmt.Scan(&toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("Amount >")
	var amount int64
	_, err = fmt.Scan(&amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	const selectSQL = "SELECT * FROM bank WHERE id = ?"
	const updateSQL = "UPDATE bank SET amount = ? WHERE id = ?"
	var r Record

	// fromID
	row := tx.QueryRow(selectSQL, fromID)
	if err := row.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
		tx.Rollback()
		return err
	}

	result, err := tx.Exec(updateSQL, r.Amount-amount, fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	cnt, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	} else if cnt == 0 {
		tx.Rollback()
		return errors.New("error: 更新件数0")
	}

	// toID
	row = tx.QueryRow(selectSQL, toID)
	if err := row.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
		tx.Rollback()
		return err
	}

	result, err = tx.Exec(updateSQL, r.Amount+amount, toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	cnt, err = result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	} else if cnt == 0 {
		tx.Rollback()
		return errors.New("error: 更新件数0")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Println("送金完了")
	fmt.Println("------")

	return nil
}
