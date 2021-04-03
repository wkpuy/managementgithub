package helper

import (
	"database/sql"
	"log"
)

func InsertGetId(db *sql.DB, table string, params map[string]interface{}, fields []string) (int64, error) {
	sql, value := GetInsertStatement(table, params, fields)
	stmt, err := db.Prepare(sql)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(value...)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}


func UpdateData(db *sql.DB, table string, params map[string]interface{}, ignoredKey []string, condKey []string) error {
	sql, value := GetEditStatement(table, params, ignoredKey, condKey)

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()
	if _, err := stmt.Exec(value...); err != nil {
		return err
	}

	return err
}


func CountRows(db *sql.DB, table string, cond string, bindValue []interface{}) int {
	ID := 0
	if err := db.QueryRow("SELECT COUNT(id) FROM "+table+" WHERE "+cond, bindValue...).Scan(&ID); err != nil {
		log.Println("err CountRows -> ", err)
		return ID
	}
	return ID
}
