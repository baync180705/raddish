package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/baync180705/raddish/internal/msg"
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

	fmt.Fprint(conn, msg.InfoPrompt)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		tokens := resp.Tokenize(cmd)
		ParsedCmd, err := resp.Parse(tokens)

		if err != nil {
			fmt.Fprintln(conn, err)
			fmt.Fprint(conn, msg.InfoPrompt)
			continue
		}

		h.executeCommand(conn, ParsedCmd)

		fmt.Fprint(conn, msg.InfoPrompt)
	}
}

func (h *Handler) executeCommand(conn net.Conn, cmd *resp.ParsedCmd) {
	switch strings.ToUpper(cmd.Op) {
	case "PING":
		fmt.Fprintln(conn, msg.InfoPong)
	case "CREATE":
		code := h.store.CREATE(cmd.Db)
		fmt.Fprintln(conn, code)
	case "SET":
		code := h.store.SET(cmd.Db, cmd.K, cmd.V)
		fmt.Fprintln(conn, code)
	case "GET":
		respStr, code := h.store.GET(cmd.Db, cmd.K)
		fmt.Fprintln(conn, respStr)
		fmt.Fprintln(conn, code)
	case "DEL":
		code := h.store.DEL(cmd.Db, cmd.K)
		fmt.Fprintln(conn, code)
	case "LISTDB":
		respStr, code := h.store.LISTDB()
		fmt.Fprintln(conn, respStr)
		fmt.Fprintln(conn, code)
	case "LISTKEYS":
		respStr, code := h.store.LISTKEYS(cmd.Db)
		fmt.Fprintln(conn, respStr)
		fmt.Fprintln(conn, code)
	case "EXIT":
		fmt.Fprintln(conn, msg.InfoExit)
		conn.Close()
		return
	default:
		fmt.Fprintf(conn, msg.ErrorUnknownCommandFmt+"\n", cmd.Op)
	}
}
