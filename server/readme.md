# CLI-TypeRacer Server

## Communication

Messages are defined in communication package, each message contains:

- playerId - unique uuid user identifier, generated on connection
- command - enum used to determine handler function
- content - message payload, can be used to deliver data

## Messages

This section contains description of all supported message commands.
Commands can be found in `server/communication/message.go `

### ERROR

Returned in case handler returns error or provided message was not validated successfully

### UNRECOGNIZED

Returned in case message command handler was not found or its command is not supported

### WELCOME

Base request used to register in a server as a player, without it you can't access any other functions. Server creates a new player instance and returns its id.
Request:

```
{
	"playerId":  "", // Should be empty
	"command":  "WELCOME",
	"content":  "" // Should be empty
}
```

Response:

```
{
	"playerId":  "814a2b70-4418-4aec-a65c-e05730e764da", // random id
	"command":  "WELCOME",
	"content":  "Welcome to the server!"
}
```

This is a successful response, from this point user can host/join games etc.
