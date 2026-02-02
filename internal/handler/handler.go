package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/baync180705/raddish/internal/resp"
	"github.com/baync180705/raddish/internal/store"
)

type Handler struct {
	store *store.Raddish
}

func New(s *store.Raddish) *Handler {
	return &Handler{store: s}
}

func (h *Handler) HandleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Fprint(conn, ">> ")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		tokens := resp.Tokenize(cmd)
		ParsedCmd, err := resp.Parse(tokens)

		if err != nil {
			fmt.Fprintln(conn, err)
			fmt.Fprint(conn, ">> ")
			continue
		}

		h.executeCommand(conn, ParsedCmd)

		fmt.Fprint(conn, ">> ")
	}
}

func (h *Handler) executeCommand(conn net.Conn, cmd *resp.ParsedCmd) {
	switch strings.ToUpper(cmd.Op) {
	case "PING":
		fmt.Fprintln(conn, "PONG")
	case "CREATE":
		code := h.store.CREATE(cmd.Db)
		fmt.Fprintln(conn, code)
	case "SET":
		code := h.store.SET(cmd.Db, cmd.K, cmd.V)
		fmt.Fprintln(conn, code)
	case "GET":
		resp, code := h.store.GET(cmd.Db, cmd.K)
		fmt.Fprintln(conn, resp)
		fmt.Fprintln(conn, code)
	case "DEL":
		code := h.store.DEL(cmd.Db, cmd.K)
		fmt.Fprintln(conn, code)
	case "LISTDB":
		resp, code := h.store.LISTDB()
		fmt.Fprintln(conn, resp)
		fmt.Fprintln(conn, code)
	case "LISTKEYS":
		resp, code := h.store.LISTKEYS(cmd.Db)
		fmt.Fprintln(conn, resp)
		fmt.Fprintln(conn, code)
	case "EXIT":
		fmt.Fprintln(conn, "connection terminated, see ya !")
		conn.Close()
		return
	}
}
