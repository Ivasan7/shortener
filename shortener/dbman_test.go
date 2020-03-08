package main

import (
	"testing"
)

var (
	testdbName  string = "plan"
)

func TestDB(t *testing.T){
	//Create DB
	_,err := NewDBManager("plan")
	if err != nil {
		t.Error("DB creation failed")
	}
	//db.getLastID()

}