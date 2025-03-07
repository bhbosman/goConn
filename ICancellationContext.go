package goConn

import "context"

type ICancellationContext interface {
	Add(connectionId string, f func()) (bool, error)
	Remove(connectionId string) error
	Cancel(string)
	CancelWithError(string, error)
	Err() error
	CancelContext() context.Context
	CancelFunc() context.CancelFunc
}
