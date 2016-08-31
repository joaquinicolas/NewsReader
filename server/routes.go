package server

import (
	"net/http"
	"github.com/joaquinicolas/newsReader/sqlite"
	"encoding/json"
)

func HandleNews(w http.ResponseWriter,r *http.Request)  {
	data := sqlite.Read()
	result,_ := json.Marshal(data)
	w.Header().Set("Content-Type","application/json")
	w.Write(result)
}