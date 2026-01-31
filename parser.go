package main

import (
	"errors"
	"strings"
)

type parsedCmd struct {
	op string
	db string
	k  string
	v  string
}

func Tokenize(cmd string) []string {
	return strings.Fields(cmd)
}

func Parse(commands []string) (*parsedCmd, error) {
	if len(commands) == 0 {
		return &parsedCmd{}, errors.New("No command found")
	}

	op := commands[0]
	args := commands[1:]

	switch strings.ToUpper(op) {
	case "PING":
		return &parsedCmd{op: op}, nil
	case "CREATE":
		if len(args) != 1 {
			return &parsedCmd{}, errors.New("usage: CREATE <dbname>")
		}
		db := args[0]
		return &parsedCmd{op: op, db: db}, nil
	case "SET":
		if len(args) != 3 {
			return &parsedCmd{}, errors.New("usage: SET <dbname> <key> <value>")
		}
		db := args[0]
		k := args[1]
		v := args[2]
		return &parsedCmd{op: op, db: db, k: k, v: v}, nil
	case "GET":
		if len(args) != 2 {
			return &parsedCmd{}, errors.New("usage: GET <dbname> <key>")
		}
		db := args[0]
		k := args[1]
		return &parsedCmd{op: op, db: db, k: k}, nil
	case "DEL":
		if len(args) != 2 {
			return &parsedCmd{}, errors.New("usage: DEL <dbname> <key>")
		}
		db := args[0]
		k := args[1]
		return &parsedCmd{op: op, db: db, k: k}, nil
	case "LISTDB":
		return &parsedCmd{op: op}, nil
	case "LISTKEYS":
		if len(args) != 1 {
			return &parsedCmd{}, errors.New("usage: LISTKEYS <dbname>")
		}
		db := args[0]
		return &parsedCmd{op: op, db: db}, nil
	case "EXIT":
		return &parsedCmd{op: op}, nil
	default:
		return &parsedCmd{}, errors.New("unknown command")
	}
}
