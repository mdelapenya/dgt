package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:passw0rd@tcp(db:3306)/dgt"
const insertSQL = "INSERT INTO plates(plate, sticker, counts) VALUES(?,?,1)"
const selectSQL = "SELECT COUNT(1) as count FROM plates WHERE plate=?"
const updateSQL = "UPDATE plates SET counts=counts+1, sticker=? WHERE plate=?"

// InsertPlate inserts plate into the database, checking if it previously exists
func InsertPlate(plate string, sticker string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error opening the database connection: %s", err.Error())
		return
	}
	defer db.Close()

	// check if exists
	stmtGet, err := db.Prepare(selectSQL)
	if err != nil {
		log.Printf("Error checking if the plate exists (1): %s", err.Error())
		return
	}
	defer stmtGet.Close()

	var count int
	err = stmtGet.QueryRow(plate).Scan(&count)
	if err != nil {
		log.Printf("Error checking if the plate exists (2): %s", err.Error())
		return
	}

	var saveSQL string
	var arg1 string
	var arg2 string

	if count == 0 {
		// new plate
		saveSQL = insertSQL
		arg1 = plate
		arg2 = sticker
	} else {
		// the plate already exists
		saveSQL = updateSQL
		arg1 = sticker
		arg2 = plate
	}

	stmtSave, errSave := db.Prepare(saveSQL)
	if errSave != nil {
		log.Printf("Error saving the plate in the database: %s", err.Error())
		return
	}
	defer stmtSave.Close()
	_, err = stmtSave.Exec(arg1, arg2)
}
