package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"google.golang.org/appengine"

	gmysql "github.com/go-sql-driver/mysql"
)

var (
	secrets *config
	db      *sql.DB
	dbName  = "playground"
)

func connect() error {
	var err error
	cfg := &gmysql.Config{
		Addr:   secrets.SQL.Address,
		User:   secrets.SQL.Username,
		Passwd: secrets.SQL.Password,
		DBName: secrets.SQL.DBName,
	}
	if db, err = mysql.DialCfg(cfg); err != nil {
		return fmt.Errorf("failed to connect to db: %s", err)
	}
	return db.Ping()
}

func createTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS gotest (
          		timestamp  BIGINT,
              userip     VARCHAR(255)
          )`
	_, err := db.Exec(stmt)
	return err
}

func main() {
	var err error
	if secrets, err = getConfig("config.yaml"); err != nil {
		log.Fatalf("failed to get secrets: %s", err)
	}
	if err = connect(); err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	log.Println("db connected successfully.")
	if err = createTable(); err != nil {
		log.Fatalf("failed to create gotest table: %s", err)
	}
	log.Println("gotest table created successfully")
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get a list of the most recent visits.
	visits, err := queryVisits(10)
	if err != nil {
		msg := fmt.Sprintf("Could not get recent visits: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Record this visit.
	if err := recordVisit(time.Now().UnixNano(), r.RemoteAddr); err != nil {
		msg := fmt.Sprintf("Could not save visit: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Previous visits:")
	for _, v := range visits {
		fmt.Fprintf(w, "[%s] %s\n", time.Unix(0, v.timestamp), v.userIP)
	}
	fmt.Fprintln(w, "\nSuccessfully stored an entry of the current request.")
}

type visit struct {
	timestamp int64
	userIP    string
}

func recordVisit(timestamp int64, userIP string) error {
	stmt := "INSERT INTO gotest (timestamp, userip) VALUES (?, ?)"
	_, err := db.Exec(stmt, timestamp, userIP)
	return err
}

func queryVisits(limit int64) ([]visit, error) {
	rows, err := db.Query("SELECT timestamp, userip FROM gotest ORDER BY timestamp DESC LIMIT ?", limit)
	if err != nil {
		return nil, fmt.Errorf("Could not get recent visits: %v", err)
	}
	defer rows.Close()

	var visits []visit

	for rows.Next() {
		var v visit
		if err := rows.Scan(&v.timestamp, &v.userIP); err != nil {
			return nil, fmt.Errorf("Could not get timestamp/user IP out of row: %v", err)
		}
		visits = append(visits, v)
	}

	return visits, rows.Err()
}
