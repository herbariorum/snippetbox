package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/herbariorum/snippetbox/internal/models"
	_ "github.com/lib/pq"
)

type Usuarios struct {
	Id    int
	Nome  string
	Cargo string
	Email string
	Cpf   string
	Roles string
}

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	snippets       *models.SnippedModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	addr := flag.String("addr", ":7000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := dbConn()

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		snippets:       &models.SnippedModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Iniciando servidor na porta %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func dbConn() (db *sql.DB) {
	dbHost := "localhost"
	dbPort := 5432
	dbDriver := "postgres"
	dbUser := "postgres"
	dbPass := "bot901"
	dbName := "snippetbox"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open(dbDriver, psqlInfo)
	if err != nil {
		log.Println("Não abre o banco de dados ", err.Error())
	}

	return db
}
