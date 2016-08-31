package rs_232_port

import (
	"github.com/tarm/serial"
	"log"
	"fmt"
)


var config *serial.Config;


//Create Config for port.
//
//Return *Config type.
//
//If name is null, log Fatal Exception.
func SetupPort(name string,baud int)  {
	if name == "" || baud == 0{
		log.Fatal("Port Name cannot be null")
		return
	}
	config = &serial.Config{Name:name,Baud:baud}

}

//Read input data from Serial port
//
//Execute callback function in a goroutine when new data enter by Serial port
func Read(f func(data string, n int))  {

	if config == nil{
		log.Fatal("Configuration cannot be null")
		return
	}

	s,err := serial.OpenPort(config)

	checkError(err)
	for  {
		fmt.Println("Waiting for a message")
		buffer := make([]byte,256)
		n,_ := s.Read(buffer)

		go f(string(buffer),n)
	}

}

func checkError(e error)  {
	if e != nil{
		log.Fatal(e)
	}
}