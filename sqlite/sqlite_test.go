package sqlite

import (
	"testing"
	"fmt"
)



func TestStore(t *testing.T)  {

	result := Store("New Message")
	if result <= 0{
		t.Error("Insert Fail")
	}

	fmt.Println("Rows afectadas ",result )
}

func TestRead(t *testing.T) {
	result := Read()
	if len(result) == 0 {
		t.Error("Read function return array with no rows")
	}
	fmt.Println("Read rows length ",len(result))
}
