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
	 dbFile string
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
		 dbFile: dbFile,
	 }, nil
}

func (e *DBManager) getLastID() int {
	return e.itemNr
}

func (e *DBManager) incrementID() {
	e.itemNr = e.itemNr + 1 
}

func (e *DBManager) close() {
	e.db.Close()
}

func (e *DBManager) getUrlByID(ID int) (string,string) {
	sqlStatement := `SELECT shortUrl, longUrl FROM urlList WHERE id=$1;`
	var shortUrl string
	var longUrl string

	row := e.db.QueryRow(sqlStatement, ID)
	switch err := row.Scan(&shortUrl, &longUrl); err {
	case sql.ErrNoRows:
	  log.Println("No rows were returned!")
	  return "",""
	case nil:
	  log.Printf("The following links has been inserted. ID: %d, shortUrl: %s, longUrl: %s ",ID,shortUrl, longUrl)
	  return shortUrl,longUrl
	default:
	  panic(err)
}

}

func (e *DBManager) getShortUrl(longUrl string) (int,string) {
	var shortUrl string
	var ID int
	sqlStmt := `SELECT id,shortUrl FROM urlList WHERE longUrl = ?`
	row := e.db.QueryRow(sqlStmt, longUrl)
	switch err := row.Scan(&ID,&shortUrl); err {
	case sql.ErrNoRows:
		log.Println("Element "+ longUrl +" not present in DB" )
		return -1,""
	case nil:
		return ID,shortUrl
	default:
		panic(err)
	}
}

func (e * DBManager) getLongLink(shortUrl string) (int,string) {
	var longUrl string
	var ID int
	sqlStmt := `SELECT id,longUrl FROM urlList WHERE shortUrl = ?`
	row := e.db.QueryRow(sqlStmt, shortUrl)
	switch err := row.Scan(&ID,&longUrl); err {
	case sql.ErrNoRows:
		log.Println("Element "+ shortUrl +" not present in DB" )
		return -1,""
	case nil:
		return ID,longUrl
	default:
		panic(err)
	}	
}

func (e* DBManager) insert2DB(longUrl string, shortUrl string)  int {
	// TODO: check if long URL alredy present, if yes, just return ID and SHORTURL
	statement, err := e.db.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
	if err != nil {
		log.Fatal(errDBInsert)
	}
	statement.Exec(shortUrl, longUrl)
	e.incrementID()
	return e.getLastID()
}



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