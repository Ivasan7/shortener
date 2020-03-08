package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	"log"
	"os"
	//"strconv"
)
type DBManager struct {
	 db     *sql.DB
	 itemNr int
}

var (
	errInvalidLink = errors.New("short link too large")
	errDBOpen = errors.New("Database can not be opened")
	errDBExists = errors.New("Database already exists: ")
	errDBTableCreate = errors.New("Table can not be generated")
	errDBInsert = errors.New("Insert into DB failed")
	errDBQuery = errors.New("Query reqest failed")

)

// NewBaseConvertor instantiates a new BaseConvertor object
func NewDBManager(dbIn string) (*DBManager, error) {
	dir, err := os.Getwd()
	dbFile := dir + "/" +dbIn + ".db" 
	if _, err := os.Stat(dbFile); err == nil {
		log.Fatal(errDBExists)
	}

	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(errDBOpen)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS urlList (id INTEGER PRIMARY KEY, longUrl TEXT, shortUrl TEXT)")
	if err != nil {
		log.Fatal(errDBTableCreate)
	}
	statement.Exec()
	// statement, err = database.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
	// if err != nil {
	// 	fmt.Println(errDBInsert)
	// }
	// statement.Exec("fake.it/"+baseconv.Encode(938641), "www.google.com")
	// rows, err := database.Query(fmt.Sprintf("SELECT id, longUrl, shortUrl FROM %s",dbIn))
	// if err != nil {
	// 	fmt.Println(errDBQuery)
	// }	
	
	// var id int
    // var longName string
    // var shortName string
	// for rows.Next() {
    //     rows.Scan(&id, &longName, &shortName)
    //     fmt.Println(strconv.Itoa(id) + ": " + longName + " " + shortName)
    // }	

	 return &DBManager {
		 db:        database,
		 itemNr:    0,
	 }, nil
}

func (e *DBManager) getLastID() int {
	return e.itemNr
}

// func (e *NewDBManager) getShortLink(longUrl string) string {
// 	var shortUrl string
// 	sqlStmt := `SELECT shortUrl FROM urlList WHERE longUrl = ?`
// 	row := e.QueryRow(sqlStmt, longUrl)
// 	switch err := row.Scan(&shortUrl); err {
// 	case sqlErrNoRows:
// 		fmt.Println("Element "+ longUrl +" not present in DB" )	
// 	case nil:
// 		return shortUrl;
// 	}	
// }

// func (e * NewDBManager) getLongLink(shortUrl string) string {
// 	var longUrl string
// 	sqlStmt := `SELECT longUrl FROM urlList WHERE shortUrl = ?`
// 	row := e.QueryRow(sqlStmt, shortUrl)
// 	switch err := row.Scan(&longUrl); err {
// 	case sqlErrNoRows:
// 		fmt.Println("Element "+ shortUrl +" not present in DB" )
// 	case nil:
// 		return shortUrl;
// 	}	
// }

//func (e* NewDBManager) insert2DB(longUrl string, shortUrl string) {
// 	newID := e.getLastID() + 1
// 	statement, _ = e.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
// 	statement.Exec(shortUrl, longUrl)
//}

// func (e *NewDBManager)ShortLinkExists(link string) otherLink string {
//     sqlStmt := `SELECT shortUrl FROM urlList WHERE shortUrls = ?`
//     err := e.QueryRow(sqlStmt, link).Scan(&shortUrl)
//     if err != nil {
//         if err != sql.ErrNoRows {
//             log.Print(err)
//         }

//         return false
//     }

//     return true
// }