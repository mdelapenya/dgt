package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var pwd = os.Getenv("MYSQL_ROOT_PASSWORD")
var dbServer = os.Getenv("MYSQL_SERVER")
var dsn = "root:" + pwd + "@tcp(" + dbServer + ":3306)/dgt"

const insertSQL = "INSERT INTO plates(plate, sticker_id, counts) VALUES(?,?,1)"
const selectSQL = "SELECT COUNT(1) as count FROM plates WHERE plate=?"
const updateSQL = "UPDATE plates SET counts=counts+1, sticker_id=? WHERE plate=?"

// InsertPlate inserts plate into the database, checking if it previously exists
func InsertPlate(plate string, stickerID int) {
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
	var arg1 any
	var arg2 any

	if count == 0 {
		// new plate
		saveSQL = insertSQL
		arg1 = plate
		arg2 = stickerID
	} else {
		// the plate already exists
		saveSQL = updateSQL
		arg1 = stickerID
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
