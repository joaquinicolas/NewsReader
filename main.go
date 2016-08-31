package main

import (
	"net/http"
	"github.com/joaquinicolas/newsReader/server"
	"github.com/joaquinicolas/newsReader/rs-232-port"
	//"os"
	"github.com/joaquinicolas/newsReader/sqlite"
	"strings"
)





//Read 3 argument.
//
//name (COM PORT), baud, time_out
func main() {

	rs_232_port.SetupPort("/dev/pts/4",115200)
	go rs_232_port.Read(func(msg string,n int) {
		//sqlite.Store(bytes.Trim(msg,"\u0000"))
		sqlite.Store(strings.Trim(msg,"\u0000"))
	})
	http.HandleFunc("/novedades",server.GetOnly(server.HandleNews))

	http.ListenAndServe(":8080",nil)
}

