package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"snippetbox.timpham.net/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	envValue := os.Getenv("DB_PASS") // Grabs DB .env file
	defaultDNS := "web:" + envValue + "@/snippetbox?parseTime=true"

	// Flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dns := flag.String("dns", defaultDNS, "MySQL data source name") // MySQL DSN string
	flag.Parse()

	// INFO and ERROR logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creates DB connection
	db, err := openDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close() // close connection right before main functions exits

	// initialize new Application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct to use the errorLog logger instead of standard logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Wraps sql.Open() and returns an sql.DB connection pool for given dsn
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
