package sqlite

import (
	"testing"
)



func TestStore(t *testing.T)  {

	if Store("New Message") <= 0{
		t.Error("Insert Fail")
	}

	rows := Read()
	if len(rows) <= 0 {
		t.Error("Select Fail. No return data")
	}

}
