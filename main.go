package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joaquinicolas/newsReader/rs-232-port"
	"github.com/joaquinicolas/newsReader/server"
	"github.com/joaquinicolas/newsReader/sqlite"
)

//Read 3 argument.
//
//name (COM PORT), baud, time_out
func main() {

	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			fmt.Println("Post Alive status")
			server.PostAlive()
		}
	}()

	port := os.Args[1]

	rs_232_port.SetupPort(port, 115200)
	go rs_232_port.Read(func(msg string, n int) {
		//sqlite.Store(bytes.Trim(msg,"\u0000"))
		sqlite.Store(strings.Trim(msg, "\u0000"))
	})
	http.HandleFunc("/novedades", server.GetOnly(server.HandleNews))

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080.")

}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
