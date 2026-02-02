package resp

import (
	"errors"
	"strings"
)

type ParsedCmd struct {
	Op string
	Db string
	K  string
	V  string
}

func Tokenize(cmd string) []string {
	return strings.Fields(cmd)
}

func Parse(commands []string) (*ParsedCmd, error) {
	if len(commands) == 0 {
		return &ParsedCmd{}, errors.New("No command found")
	}

	op := commands[0]
	args := commands[1:]

	switch strings.ToUpper(op) {
	case "PING":
		return &ParsedCmd{Op: op}, nil
	case "CREATE":
		if len(args) != 1 {
			return &ParsedCmd{}, errors.New("usage: CREATE <dbname>")
		}
		db := args[0]
		return &ParsedCmd{Op: op, Db: db}, nil
	case "SET":
		if len(args) != 3 {
			return &ParsedCmd{}, errors.New("usage: SET <dbname> <key> <value>")
		}
		db := args[0]
		k := args[1]
		v := args[2]
		return &ParsedCmd{Op: op, Db: db, K: k, V: v}, nil
	case "GET":
		if len(args) != 2 {
			return &ParsedCmd{}, errors.New("usage: GET <dbname> <key>")
		}
		db := args[0]
		k := args[1]
		return &ParsedCmd{Op: op, Db: db, K: k}, nil
	case "DEL":
		if len(args) != 2 {
			return &ParsedCmd{}, errors.New("usage: DEL <dbname> <key>")
		}
		db := args[0]
		k := args[1]
		return &ParsedCmd{Op: op, Db: db, K: k}, nil
	case "LISTDB":
		return &ParsedCmd{Op: op}, nil
	case "LISTKEYS":
		if len(args) != 1 {
			return &ParsedCmd{}, errors.New("usage: LISTKEYS <dbname>")
		}
		db := args[0]
		return &ParsedCmd{Op: op, Db: db}, nil
	case "EXIT":
		return &ParsedCmd{Op: op}, nil
	default:
		return &ParsedCmd{}, errors.New("unknown command")
	}
}
