package common

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
	POST   Action = 1
	PUT    Action = 2
	PATCH  Action = 3
	DELETE Action = 4
	GET    Action = 5
	Reply  Action = 6
	Notify Action = 7
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

type IMessage interface {
	Source() string
	Vnet() string
	Destination() string
	ServiceArea() uint16
	ServiceName() string
	Sequence() uint32
	Priority() Priority
	Action() Action
	Timeout() uint16
	Request() bool
	Reply() bool
	FailMessage() string
	Data() string
	SetData(string)
	Tr() ITransaction
	Serialize() []byte
}

type ITransaction interface {
	Id() string
	State() TransactionState
	ErrorMessage() string
	StartTime() int64
}
