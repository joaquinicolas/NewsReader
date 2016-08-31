package sqlite
/*
import (
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
	err error
	Configuartion *Config
)

type Config struct {
	driver string
	conString string
}

type ContentValues struct {
	valuePair []*keyValuePair

}

func (v *ContentValues) Add(key,value string)  {
	append(v.valuePair,  keyValuePair{Key:key, Value:value})
}

func (v *ContentValues) Get(position int)  *keyValuePair {
	return v.valuePair[position]
}

func (v *ContentValues) iterator() int{
	return len(v.valuePair)
}

type keyValuePair struct {
	Key string
	Value string
}

func checkErr(e error)  {
	if e != nil{
		panic(e)
	}
}

func Open() {
	db, err = sql.Open(&Configuartion.driver,Configuartion.conString)
}

func Insert(values ContentValues,table string) int {
	query := fmt.Sprintf("INSERT INTO %s(",table)
	v := ""
	vargs := make([]string,values.iterator())
	for i := 0;i <= values.iterator() ; i++ {
		if i > 0 {
			query += ","
			v += ","

		}
		query += values.Get(i).Key
		v += "?"
		append(vargs,values.Get(i).Value)

	}

	query += fmt.Sprintf(") VALUES(%s)",v)
	stmt,err := db.Prepare(query)
	checkErr(err)
	res, err := stmt.Exec(vargs)
	checkErr(err)
	id,err := res.LastInsertId()
	checkErr(err)
	return id
}

*/