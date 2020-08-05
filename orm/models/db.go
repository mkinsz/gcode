package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Db global db object
var Db *sqlx.DB

func init() {
	fmt.Println("Model Init...")
	db, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	Db = db
	fmt.Println("DB connected...")
}

type Member struct {
	Username string         `db:"username"`
	Money    float64        `db:"money"`
	Birthday sql.NullString `db:"birthday"`
}

func Select() {
	//查一条
	var info Member
	err := Db.Get(&info, "select username,money,birthday from member where id=?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)

	//查多条
	var list []Member
	err = Db.Select(&list, "select username,money,birthday from member")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}

func Insert() {
	result, err := Db.Exec("insert into member(username,money,created_at)values (?,?,?)", "test", 20, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

// func Update() {
// 	result, err := Db.Exec("update member set money=money+3 where id=?", 2)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(rows)
// }

func Delete() {
	result, err := Db.Exec("delete from member where id=?", 7)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows)
}

func Transaction() {
	db, err := Db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := db.Exec("update member set money=money+3 where id=?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		db.Rollback()
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	result, err = Db.Exec("update member set money=money-3 where id=?", 2)
	if err != nil {
		db.Rollback()
		fmt.Println(err)
		return
	}
	num, err = result.RowsAffected()
	if err != nil {
		db.Rollback()
		fmt.Println(err)
		return
	}
	err = db.Commit()
	if err != nil {
		defer clearTransaction(db) //出现异常，用来收回
	}
	fmt.Println("Commit...")
}

func clearTransaction(tx *sql.Tx) {
	fmt.Println("rollback...")
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		fmt.Println(err)
	}
}

// Close db close
func Close() {
	Db.Close()
}
