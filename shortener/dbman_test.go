package main

import (
	"testing"
	"os"
)

var (
	testdbName  string = "plan"
)

func TestDB(t *testing.T){
	//Create DB
	db,err := NewDBManager("plan")
	if err != nil {
		t.Error("DB creation failed")
	}
	//Last ID
	id_num := db.getLastID()
	if id_num != 0{
		t.Errorf("Number of idea is not 0, got: %d",id_num)
	}
	// Insertion check
	idNum := db.insert2DB("short","long")
	s,l := db.getUrlByID(idNum)
	if s != "short" && l != "long"{
		t.Errorf("Insertion test faliure")
	}



	defer db.close()
	defer os.Remove(db.dbFile); if err != nil {
		t.Errorf("DB File can not be removed")
	}
}