package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/joaquinicolas/NewsReader/sqlite"
	"log"
)

// HandleNews asdasdasdasdas
func HandleNews(w http.ResponseWriter, r *http.Request) {
	data := sqlite.Read()
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

//PostNews send news to server
func PostNews(msg string) {
	if msg == "" {
		return
	}
	values := map[string]string{"Mac":getMAC().String(),"Data":msg}
	json,_ := json.Marshal(values)
	req, _ := http.NewRequest("POST", "http://131.255.5.183:9090/News", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Post Request /News finished!")
	defer resp.Body.Close()
}

//PostAlive notify to server that
func PostAlive() {

	mac := getMAC()
	fmt.Println(mac)
	if mac == nil {
		return
	}

	//131.255.5.183
	values := map[string]string{"Mac": mac.String()}
	json,_ := json.Marshal(values)
	fmt.Println(json)
	req, _ := http.NewRequest("POST", "http://131.255.5.183:9090/Alive", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Post request /Alive finished!")
	defer resp.Body.Close()

}

func getMAC() net.HardwareAddr {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		if i.HardwareAddr != nil {

			return i.HardwareAddr
		}
	}

	return nil
}
