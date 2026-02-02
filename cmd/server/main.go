package main

import (
    "fmt"
    "log"
    "net"

    "github.com/baync180705/raddish/internal/handler"
    "github.com/baync180705/raddish/internal/store"
)

func main() {
    l, err := net.Listen("tcp", ":1112")
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()
    fmt.Println("Raddish is running on :1112")

    raddishDB := store.INIT()
    h := handler.New(raddishDB)

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Connection error:", err)
            continue
        }
        go h.HandleConnection(conn)
    }
}