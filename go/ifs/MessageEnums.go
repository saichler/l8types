package ifs

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

type MulticastMode byte

const (
	M_All        MulticastMode = 0
	M_RoundRobin MulticastMode = 1
	M_Proximity  MulticastMode = 2
	M_Local      MulticastMode = 3
	M_Leader     MulticastMode = 4
	M_Unicast    MulticastMode = 128
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

	// Leader Election Actions
	ElectionRequest    Action = 10 // Initiate election for a service
	ElectionResponse   Action = 11 // Response from higher-priority node (still alive)
	LeaderAnnouncement Action = 12 // Announce new leader for service
	LeaderHeartbeat    Action = 13 // Periodic heartbeat from current leader
	LeaderQuery        Action = 14 // Query who is the current leader
	LeaderResign       Action = 15 // Graceful leader resignation
	LeaderChallenge    Action = 16 // Challenge current leader validity
)

type TransactionState uint8

const (
	Empty      TransactionState = 0
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

const (
	DESTINATION_Single = "signleXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	DESTINATION_Leader = "leaderXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
)
