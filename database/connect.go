package database

// 데이터베이스 연결
import (
	//"github.com/go-sql-driver/mysql"

	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// db정보와 db_dns
	sqlDB, err := sql.Open("mysql", "glassKT:1234@tcp(13.125.236.0:3306)/glasskt?parseTime=true")
	if err != nil {
		panic(err)
	}

	// 기존 db와 연결(gorm)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 연결된 db를 리턴
	DB = db
	log.Printf("[DB 연결 성공]")
}
