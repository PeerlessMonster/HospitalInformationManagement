package conn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func GetInstance() *gorm.DB {
	if db == nil {
		dsn := "root:12345678@tcp(127.0.0.1)/hospital"

		var err error
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(err)
		}
	}
	return db
}

func Select_doctorNo_by_doctorName(doctorName string) (result string) {
	sql := "SELECT employeeNo FROM doctor WHERE doctorName = ?"
	row := db.Raw(sql, doctorName).Row()
	row.Scan(&result)
	return
}

func Select_patientNo_by_patientName(patientName string) (result string) {
	sql := "SELECT patientNo FROM patient WHERE patientName = ?"
	row := db.Raw(sql, patientName).Row()
	row.Scan(&result)
	return
}

func Select_patientPhone_by_patientName(patientName string) (result string) {
	sql := "SELECT patientPhone FROM patient WHERE patientName = ?"
	row := db.Raw(sql, patientName).Row()
	row.Scan(&result)
	return
}
