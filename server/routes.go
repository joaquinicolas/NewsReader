package server

import (
	"net/http"
	"github.com/joaquinicolas/newsReader/sqlite"
	"encoding/json"
	"net"
	"bytes"
	"github.com/joaquinicolas/newsReader/config"
)

func HandleNews(w http.ResponseWriter,r *http.Request)  {
	data := sqlite.Read()
	result,_ := json.Marshal(data)
	w.Header().Set("Content-Type","application/json")
	w.Write(result)
}


//Post to server that raspberry is online
func PostAlive()  {
	mac := getMAC()
	if mac == nil {
		return
	}
	json := []byte(`{"MAC":"` + mac.String() + `"}`)
	req, err := http.NewRequest("POST",config.ReadConfig().RemoteServer,bytes.NewBuffer(json))
	req.Header.Set("Content-Type","application/json")

	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}

func getMAC() net.HardwareAddr{
	interfaces,_ := net.Interfaces()
	for _,i := range interfaces{
		if i.HardwareAddr != nil {

			return i.HardwareAddr
		}
	}

	return nil
}
