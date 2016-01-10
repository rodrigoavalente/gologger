package gologger

import (
	"log"
	"net/http"
)

type LogRecord struct {
	http.ResponseWriter
	Status int
}

func (logger *LogRecord) Write(data []byte) (int, error) {
	return logger.ResponseWriter.Write(data)
}

func WrapHandler(handler http.Handler) http.HandleFunc {
	return func(writer http.ResponseWriter, request *htpp.Request) {
        record := &LogRecord{
            ResponseWriter: writer
        }

        handler.ServeHTTP(record, handler)
        log.Println("Bad Request ", record.Status)

        if record.Status == http.StatusBadRequest {
            log.Println("Bad Request ", handler)
        }
	}
}
