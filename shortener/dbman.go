package main

var (
	ErrInvalidDB = errors.New("Invalid DB")
)


type DBManager struct {
	 db   DB
}

// NewBaseConvertor instantiates a new BaseConvertor object
func NewDBManager(db string) (*DBManager, error) {
	info, err := os.Stat(db)
	database, _ := sql.Open("sqlite3", "./urlList.db")
	if os.IsNotExist(err) == 0 {
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS urlList (id INTEGER PRIMARY KEY, longUrl TEXT, shortUrl TEXT)")
		statement.Exec()
	} 

	return &DBManager{
		db:        database,
	}, nil
}

func (e *NewDBManager) getLastID() int {
	statement,_ := e.Prepare("INSERT table SET unique_id=? ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)")
	res,_ := statement.Exec(unique_id)
	lid, _ := res.LastInsertId()

	return lid
}

func (e *NewDBManager) getShortLink(longUrl string) string {
	var shortUrl string
	sqlStmt := `SELECT shortUrl FROM urlList WHERE longUrl = ?`
	row := e.QueryRow(sqlStmt, longUrl)
	switch err := row.Scan(&shortUrl); err {
	case sqlErrNoRows:
		fmt.Println("Element "+ longUrl +" not present in DB" )	
	case nil:
		return shortUrl;
	}	
}

func (e * NewDBManager) getLongLink(shortUrl string) string {
	var longUrl string
	sqlStmt := `SELECT longUrl FROM urlList WHERE shortUrl = ?`
	row := e.QueryRow(sqlStmt, shortUrl)
	switch err := row.Scan(&longUrl); err {
	case sqlErrNoRows:
		fmt.Println("Element "+ shortUrl +" not present in DB" )
	case nil:
		return shortUrl;
	}	
}

func (e* NewDBManager) insert2DB(longUrl string, shortUrl string){
	newID := e.getLastID() + 1
	statement, _ = e.Prepare("INSERT INTO urlList (longUrl, shortUrl) VALUES (?, ?)")
	statement.Exec(shortUrl, longUrl)
}

func (e *NewDBManager)ShortLinkExists(link string) otherLink string{
    sqlStmt := `SELECT shortUrl FROM urlList WHERE shortUrls = ?`
    err := e.QueryRow(sqlStmt, link).Scan(&shortUrl)
    if err != nil {
        if err != sql.ErrNoRows {
            log.Print(err)
        }

        return false
    }

    return true
}