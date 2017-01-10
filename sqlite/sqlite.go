package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"database/sql/driver"
)

const CREATE_TABLE  = "CREATE TABLE IF NOT EXISTS news(" +
	"uid INTEGER PRIMARY KEY AUTOINCREMENT," +
	"data TEXT," +
	"checkin datetime DEFAULT CURRENT_TIMESTAMP," +
	"checkout datetime NULL);"

const (
	DATA_FIELD  = "data"
	CHECKIN = "checkin"
	CHECKOUT = "checkout")


type New struct {
	Data     string `json:"data"`
	Checkin  time.Time `json:"checkin"`
	Checkout NullTime	`json:"checkout"`
}

type NullTime struct {
	Time time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error  {
	nt.Time,nt.Valid = value.(time.Time)
	return nil
}

func (nt NullTime) Value() (driver.Value,error)  {
	if !nt.Valid {
		return nil,nil
	}
	return nt.Time,nil
}

func checkError(e error)  {
	if e != nil{
		panic(e)
	}
}

func openCon()  *sql.DB{
	db, err := sql.Open("sqlite3","./rs.db")
	checkError(err)
	stmt,err := db.Prepare(CREATE_TABLE)
	checkError(err)

	stmt.Exec()
	fmt.Println("Base de datos creada")
	return db
}

func Store(values string) int64 {
	stmt,err := openCon().Prepare(
		fmt.Sprintf("INSERT INTO news(%s)" +
			"VALUES(?)",DATA_FIELD,))

	checkError(err)
	defer  stmt.Close()
	fmt.Println("Insertando data")
	res,err := stmt.Exec(&values)
	checkError(err)

	rowsAffected,err := res.RowsAffected()
	checkError(err)

	return rowsAffected
}

// Read News from database
// return New array or empty New array if the query no has rows
func Read()  []*New {
	sql := `
	SELECT data,checkin,checkout FROM news;
	`

	rows, err := openCon().Query(sql)
	checkError(err)

	var result []*New
	for rows.Next(){
		item := New{}
		err := rows.Scan(&item.Data,&item.Checkin,&item.Checkout)
		checkError(err)
		result = append(result,&item)
	}
	return result
}