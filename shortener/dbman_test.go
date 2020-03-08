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
	// Get short URL
	idNum1,s1 := db.getShortUrl("long")
	if s1 != "short"  {
		t.Errorf("Retrived long URL is not correct")
	}
	if idNum1 != 1 {
		t.Errorf("Retrived  URL ID is not correct")	
	}
	// Get long URL
	idNum2,s2 := db.getLongUrl("short")
	if s2 != "long" {
		t.Errorf("Retrived long URL is not correct")
	}
	if idNum2 != 1 {
		t.Errorf("Retrived  URL ID is not correct")	
	}
	defer db.close()
	defer os.Remove(db.dbFile); if err != nil {
		t.Errorf("DB File can not be removed")
	}
}