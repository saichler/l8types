package ifs

import "reflect"

type Priority byte

const (
	P8 Priority = 0
	P7 Priority = 1
	P6 Priority = 2
	P5 Priority = 3
	P4 Priority = 4
	P3 Priority = 5
	P2 Priority = 6
	P1 Priority = 7
)

type Action byte

const (
	POST      Action = 1
	PUT       Action = 2
	PATCH     Action = 3
	DELETE    Action = 4
	GET       Action = 5
	Reply     Action = 6
	Notify    Action = 7
	Sync      Action = 8
	EndPoints Action = 9
)

type TransactionState uint8

const (
	Create     TransactionState = 1
	Created    TransactionState = 2
	Start      TransactionState = 3
	Lock       TransactionState = 4
	Locked     TransactionState = 5
	LockFailed TransactionState = 6
	Commit     TransactionState = 7
	Commited   TransactionState = 8
	Rollback   TransactionState = 9
	Rollbacked TransactionState = 10
	Finish     TransactionState = 11
	Finished   TransactionState = 12
	Errored    TransactionState = 13
)

func (t TransactionState) String() string {
	switch t {
	case Create:
		return "Create"
	case Created:
		return "Created"
	case Start:
		return "Start"
	case Lock:
		return "Lock"
	case Locked:
		return "Locked"
	case LockFailed:
		return "LockFailed"
	case Commit:
		return "Commit"
	case Commited:
		return "Commited"
	case Rollback:
		return "Rollback"
	case Rollbacked:
		return "Rollbacked"
	case Finish:
		return "Finish"
	case Finished:
		return "Finished"
	case Errored:
		return "Errored"
	}
	return "Unknown"
}

type IMessage interface {
	Source() string
	Vnet() string
	Destination() string
	ServiceArea() uint16
	ServiceName() string
	Sequence() uint32
	Priority() Priority
	Action() Action
	SetAction(Action)
	Timeout() uint16
	Request() bool
	Reply() bool
	FailMessage() string
	Data() string
	SetData(string)
	Tr() ITransaction
	SetTr(transaction ITransaction)
}

type ITransaction interface {
	Id() string
	State() TransactionState
	SetState(TransactionState)
	ErrorMessage() string
	SetErrorMessage(string)
	StartTime() int64
}

func IsNil(any interface{}) bool {
	if any == nil {
		return true
	}
	v := reflect.ValueOf(any)
	isNil := v.IsNil()
	if !isNil {
		if v.Kind() == reflect.Func {
			panic("Trying to check nil on a function!")
		}
	}
	return isNil
}
