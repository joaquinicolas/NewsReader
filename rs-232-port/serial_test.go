package rs_232_port

import (
	"testing"
	"fmt"
)



func TestOpenPort(t *testing.T)  {
	SetupPort("/dev/pts/4",115200)
	if config == nil {
		t.Error("Configuration error")
	}

	fmt.Println(config.Name)
}

func TestRead(t *testing.T)  {
	Read(func(data string,n int) {
		fmt.Print(data)

	})
}


