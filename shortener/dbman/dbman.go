package dbman

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"errors"
	"log"
	"os"
)
type DBManager struct {
	 db     *sql.DB
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

func NewDBManager(dbIn string) (*DBManager, error) {
	dir, err := os.Getwd()
	dbFile := dir + "/" +dbIn + ".db" 

	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(errDBOpen)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS urlList (id INTEGER PRIMARY KEY, longUrl TEXT, shortUrl TEXT)")
	if err != nil {
		log.Fatal(errDBTableCreate)
	}
	statement.Exec()

	 return &DBManager {
		 db:        database,
		 dbFile: dbFile,
	 }, nil
}

func (e *DBManager) GetLastID() int {
	//return e.itemNr
	sqlStatement := `SELECT MAX(id) from urlList`
	var maxId int

	row := e.db.QueryRow(sqlStatement)
	row.Scan(&maxId)
	return maxId
}

func (e *DBManager) Close() {
	e.db.Close()
}

func (e *DBManager) GetUrlByID(ID int) (string,string) {
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

func (e *DBManager) GetShortUrl(longUrl string) (int,string) {
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

func (e * DBManager) GetLongUrl(shortUrl string) (int,string) {
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

func (e* DBManager) Insert2DB(longUrl string, shortUrl string)  int {
	ID,name := e.GetShortUrl("longUrl")
	if ID != -1 {
		log.Printf("The %s long URL is already present in the DB with ID: %d, shortURL: %s",longUrl,ID, name)
	}
	statement, err := e.db.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
	if err != nil {
		log.Fatal(errDBInsert)
	}
	statement.Exec(shortUrl, longUrl)
	return e.GetLastID()
}