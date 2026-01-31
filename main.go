package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":1112")

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	raddish := INIT()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Could not connect to Raddish, try again !")
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				cmd := scanner.Text()
				commands := Tokenize(cmd)
				parsedCmd, err := Parse(commands)

				if err != nil {
					fmt.Println(err)
					continue
				}

				switch parsedCmd.op {
				case "PING":
					fmt.Fprintln(conn, "PONG")
				case "CREATE":
					code := raddish.CREATE(parsedCmd.db)
					fmt.Fprintln(conn, code)
				case "SET":
					code := raddish.SET(parsedCmd.db, parsedCmd.k, parsedCmd.v)
					fmt.Fprintln(conn, code)
				case "GET":
					resp, code := raddish.GET(parsedCmd.db, parsedCmd.k)
					fmt.Fprintln(conn, resp)
					fmt.Fprintln(conn, code)
				case "DEL":
					code := raddish.DEL(parsedCmd.db, parsedCmd.k)
					fmt.Fprintln(conn, code)
				}
			}
		}(conn)
	}
}
