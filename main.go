package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joaquinicolas/NewsReader/rs-232-port"
	"github.com/joaquinicolas/NewsReader/server"
	"github.com/joaquinicolas/NewsReader/sqlite"
)

//Read 3 argument.
//

func main() {

	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			fmt.Println("Post Alive status")
			server.PostAlive()
		}
	}()

	var port string
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		panic("Missing program argument")
		os.Exit(1)
	}

	rs_232_port.SetupPort(port, 115200)
	go rs_232_port.Read(func(msg string, n int) {

		go sqlite.Store(strings.Trim(msg, "\u0000"))
		fmt.Println(strings.Trim(msg, "\u0000"))
		go server.PostNews(strings.Trim(msg, "\u0000"))
	})
	http.HandleFunc("/novedades", server.GetOnly(server.HandleNews))

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080.")

}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
