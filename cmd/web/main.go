package main

import (
    "database/sql"
    "flag"
    "log"
    "net/http"
    "os"

    "github.com/Krokin/PinBox/pkg/models/mysql"
     
    _"github.com/go-sql-driver/mysql"
)

type application struct {
    errorLog      *log.Logger
    infoLog       *log.Logger
    pins          *mysql.PinModel
}

func main() {
    addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
    dsn := flag.String("dsn", "", "(user:password@/pinbox?parseTime=true) Название MySQL источника данных")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    sqlbase, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer sqlbase.Close()

    app := &application{
        errorLog:      errorLog,
        infoLog:       infoLog,
        pins:      &mysql.PinModel{DB: sqlbase},
    }

    srv := app.routes()
    app.infoLog.Printf("Запуск сервера на http://127.0.0.1%s", *addr)
    app.errorLog.Fatal(http.ListenAndServe(*addr, srv))
}

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