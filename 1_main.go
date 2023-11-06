package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

var mainAdmin string
var mainAdminId int

func main() {
	go start_bot()
	addr := flag.String("addr", ("127.0.4.14:56441"), "Сетевой адрес веб-сервера")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	mainAdmin, mainAdminId = ShowMainAdmin()
	infoLog.Printf("Админ: %s", mainAdmin)
	infoLog.Printf("ID: %d", mainAdminId)

	dataWorkes = ShowWorkersIDs()
	fmt.Println("Сотрудники: ", dataWorkes)
	fmt.Println("-------------------------------------------------------------")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
