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
	EndPoints Action = 9
	Handle    Action = 8

	// Leader Election Actions
	ElectionRequest    Action = 10 // Initiate election for a service
	ElectionResponse   Action = 11 // Response from higher-priority node (still alive)
	LeaderAnnouncement Action = 12 // Announce new leader for service
	LeaderHeartbeat    Action = 13 // Periodic heartbeat from current leader
	LeaderQuery        Action = 14 // Query who is the current leader
	LeaderResign       Action = 15 // Graceful leader resignation
	LeaderChallenge    Action = 16 // Challenge current leader validity

	// Participant Registry Actions
	ServiceRegister   Action = 17 // Announce: "I'm hosting this service"
	ServiceUnregister Action = 18 // Announce: "I'm no longer hosting this service"
	ServiceQuery      Action = 19 // Query: "Who is hosting this service?"

	MapR_POST   Action = 21
	MapR_PUT    Action = 22
	MapR_PATCH  Action = 23
	MapR_DELETE Action = 24
	MapR_GET    Action = 25
)

type TransactionState uint8

const (
	NotATransaction TransactionState = 0
	Created         TransactionState = 1
	Queued          TransactionState = 2
	Running         TransactionState = 3
	Committed       TransactionState = 4
	Rollback        TransactionState = 5
	Failed          TransactionState = 6
	Cleanup         TransactionState = 7
)

func (t TransactionState) String() string {
	switch t {
	case NotATransaction:
		return "NotATransaction"
	case Created:
		return "Created"
	case Queued:
		return "Queued"
	case Running:
		return "Running"
	case Committed:
		return "Committed"
	case Rollback:
		return "Rollback"
	case Failed:
		return "Failed"
	case Cleanup:
		return "Cleanup"
	}
	return "Unknown"
}

const (
	DESTINATION_Single = "signleXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	DESTINATION_Leader = "leaderXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
)
