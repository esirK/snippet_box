package main

import (
	"log"
	"net"
	"os"
)

var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)

var conn, _ = getConn()
var NetLogger = log.New(conn, "NET: ", log.Ldate|log.Ltime|log.Lshortfile)

func getConn() (net.Conn, error){
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		return nil, err
	}
	return conn, err
}
