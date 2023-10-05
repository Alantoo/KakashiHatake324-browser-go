package crigo

import "errors"

var (
	errConnectionClosed    = errors.New("connection is closed")
	errServerConnection    = errors.New("unknown error connecting to server")
	errContextCancelled    = errors.New("context canceled")
	errSessionDoesntExists = errors.New("session doesn't exists")
	errSessionExists       = errors.New("session already exists")
	errCannotFindListner   = errors.New("could not find the listener channel")
)
