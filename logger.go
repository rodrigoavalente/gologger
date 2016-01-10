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

func (logger *LogRecord) WriteHeader(status int) {
	logger.Status = status
	logger.ResponseWriter.WriteHeader(status)
}

func WrapHandler(handler http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		record := &LogRecord{
			ResponseWriter: writer,
		}

		handler.ServeHTTP(record, request)
		log.Println("Bad Request ", record.Status)

		if record.Status == http.StatusBadRequest {
			log.Println("Bad Request ", request)
		}
	}
}
