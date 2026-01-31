# Raddish

Raddish is a lightweight, in‑memory key–value store that speaks a simple text protocol over TCP.  
It supports multiple logical databases, basic CRUD operations on string keys, and listing of DBs/keys.

> Status: Experimental / learning project.

---

## Features

- TCP server (default `:1112`)
- Multiple logical databases
- Commands:
  - `PING`
  - `CREATE <dbname>`
  - `SET <dbname> <key> <value>`
  - `GET <dbname> <key>`
  - `DEL <dbname> <key>`
  - `LISTDB`
  - `LISTKEYS <dbname>`
- Concurrent safe access using mutexes

---

## Getting Started

### Prerequisites

- Go 1.20+ (or compatible version)

### Build

```bash
git clone https://github.com/baync180705/raddish.git
cd raddish

go build -o raddish
```

### Run

```bash
./raddish
# Server listens on :1112
```

You should see:

```text
Raddish initialized successfully
```

---

## Connecting

You can use `nc` (`netcat`), `telnet`, or any TCP client.

### Using netcat

```bash
nc localhost 1112
```

Now you can type commands (one per line).

---

## Protocol & Commands

The protocol is line‑based, space‑separated commands.  
The server usually returns either a message and/or a numeric status code:

- `1` means success
- `0` means error

### PING

Health check.

```text
> PING
PONG
```

### CREATE

Create a new logical database.

```text
> CREATE mydb
1            # success

> CREATE mydb
Cannot create an existing key
0            # failure
```

### SET

Set a key to a value in a given DB.

```text
> SET mydb foo bar
1            # success
```

Usage:

```text
SET <dbname> <key> <value>
```

### GET

Get the value of a key in a given DB.

```text
> GET mydb foo
bar
1            # success

> GET mydb unknown
Given key is unavailable
<INVALID>
0            # failure
```

Usage:

```text
GET <dbname> <key>
```

### DEL

Delete a key from a given DB.

```text
> DEL mydb foo
1            # success

> DEL mydb foo
Given key not fonund
0            # failure
```

Usage:

```text
DEL <dbname> <key>
```

### LISTDB

List all existing databases (one per line).

```text
> LISTDB
mydb
anotherdb
1            # success
```

If none exist:

```text
No DBs available, use CREATE <dbname> to create a DB
0
```

### LISTKEYS

List all keys in a given DB (one per line).

```text
> LISTKEYS mydb
foo
bar
1            # success
```

If the DB does not exist:

```text
> LISTKEYS unknown
given DB does not exist
0
```

If the DB exists but has no keys:

```text
No keys exist in the mentioned DB, use SET <dbname> <key> <value> to set a key
0
```

Usage:

```text
LISTKEYS <dbname>
```

---

## Implementation Notes

- `main.go` starts the TCP server on `:1112` and handles client connections.
- `parser.go` tokenizes and parses incoming text commands into a structured form.
- `store.go` implements:
  - `Raddish` – top‑level store that holds databases
  - `registry` – per‑DB key–value store
  - Methods: `CREATE`, `SET`, `GET`, `DEL`, `LISTDB`, `LISTKEYS`
- Concurrency:
  - A mutex guards the DB map and DB list.
  - Each DB has its own RW mutex for key operations.
