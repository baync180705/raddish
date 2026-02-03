package msg

// Info / log messages.
const (
	InfoRaddishInitialized = "(info) Raddish initialized successfully"
	InfoPrompt             = ">> "
	InfoPong               = "PONG"
	InfoExit               = "(info) connection terminated, see ya !"
	InfoCouldNotConnect    = "(info) could not connect to Raddish, try again !"
)

// Store / engine messages.
const (
	ErrorDBAlreadyExists = "(error) cannot create an existing key"
	ErrorDBNotFound      = "(error) given DB does not exist"
	ErrorKeyNotFound     = "(error) given key is unavailable"
	ErrorKeyNotFoundDel  = "(error) given key not found"
	ErrorNoDBsAvailable  = "(error) no DBs available, use CREATE <dbname> to create a DB"
	ErrorNoKeysInDB      = "(error) no keys exist in the mentioned DB, use SET <dbname> <key> <value> to set a key"
)

// Parser / protocol messages.
const (
	ErrorNoCommandFound = "(error) no command found"
	ErrorUsageCreate    = "(error) usage: CREATE <dbname>"
	ErrorUsageSet       = "(error) usage: SET <dbname> <key> <value>"
	ErrorUsageGet       = "(error) usage: GET <dbname> <key>"
	ErrorUsageDel       = "(error) usage: DEL <dbname> <key>"
	ErrorUsageListKeys  = "(error) usage: LISTKEYS <dbname>"
	ErrorUnknownCommand = "(error) unknown command"
)

// Handler / connection messages.
const (
	ErrorUnknownCommandFmt = "(error) unknown command - %s"
)
