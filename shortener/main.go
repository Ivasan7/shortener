/*
Copyright © 2020 Ivasan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
//	"./cmd"
	_ "github.com/mattn/go-sqlite3"
	//"fmt"
	//"strconv"
	//"errors"
	"os"
)

var (
	baseconv, _    = NewBaseConvertor(62)
)

func main() {
	dir, _ := os.Getwd()
	DB, _ := NewDBManager("test1")
	fmt.Println(DB.getLastID())


	// database, _ := sql.Open("sqlite3", "./urlList.db")
    // statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urlList (id INTEGER PRIMARY KEY, longUrl TEXT, shortUrl TEXT)")
    // statement.Exec()
    // statement, _ = database.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
    // statement.Exec("fake.it/"+baseconv.Encode(938641), "www.google.com")
    // rows, _ := database.Query("SELECT id, longUrl, shortUrl FROM urlList")
    // var id int
    // var longName string
    // var shortName string
    // for rows.Next() {
    //     rows.Scan(&id, &longName, &shortName)
    //     fmt.Println(strconv.Itoa(id) + ": " + longName + " " + shortName)
    // }	

	// cmd.Execute()
}
