package kafka



type SimEventStatus int

const (
	// when client makes a simulation event request to command server
	SimEventStatusRequested SimEventStatus = iota
	// when event is reached to model server, yet waiting to be allocated to worker
	SimEventStatusPending
	// when event is allocated to worker
	SimEventStatusAllocated
	// when the work failed while worker is processing
	SimEventStatusFailed
	// when the work finished successfully                        
	SimEventStatusFinished
)

var SimEventTopic = map[SimEventStatus]string{
	SimEventStatusRequested: "sim.requested",
	SimEventStatusPending:   "sim.pending",
	SimEventStatusAllocated: "sim.allocated",
	SimEventStatusFailed:    "sim.failed",
	SimEventStatusFinished:  "sim.finished",
}

// A corespondance to SimEventTopic
var SimEventRollbackTopic = map[SimEventStatus]string{
	SimEventStatusRequested: "sim.requested_rollback",
	SimEventStatusPending:   "sim.pending_rollback",
	SimEventStatusAllocated: "sim.allocated_rollback",
	SimEventStatusFailed:    "sim.failed_rollback",
	SimEventStatusFinished:  "sim.finished_rollback",
}



type ValEventStatus int

const (
	// when client makes a validation event request to command server
	ValEventStatusRequested ValEventStatus = iota
	// when event is reached to model server, yet waiting to be allocated to worker
	ValEventStatusPending
	// when event is allocated to worker
	ValEventStatusAllocated
	// when the result is validated
	ValEventStatusValidated
	// when the result is invalidated
	ValEventStatusInvalidated
	// when the work failed while worker is processing
	ValEventStatusFailed
	// when the work finished successfully
	ValEventStatusFinished
)

var ValEventTopic = map[ValEventStatus]string{
	ValEventStatusRequested:   "val.requested",
	ValEventStatusPending:     "val.pending",
	ValEventStatusAllocated:   "val.allocated",
	ValEventStatusFailed:      "val.failed",
	ValEventStatusFinished:    "val.finished",
}

var ValEventRollbackTopic = map[ValEventStatus]string{
	ValEventStatusRequested:   "val.requested_rollback",
	ValEventStatusPending:     "val.pending_rollback",
	ValEventStatusAllocated:   "val.allocated_rollback",
	ValEventStatusFailed:      "val.failed_rollback",
	ValEventStatusFinished:    "val.finished_rollback",
}



type RealEventStatus int

const (
	// when client makes a realtime event request to command server
	RealEventStatusRequested RealEventStatus = iota
	// when event is reached to model server, yet waiting to be allocated to worker
	RealEventStatusPending
	// when event is allocated to worker
	RealEventStatusAllocated
	// when unexpected error occured, so unable to proceed
	RealEventStatusFailed
	// when client makes a cease request to command server
	RealEventStatusCeased
)

var RealEventTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real.requested",
	RealEventStatusPending:   "real.pending",
	RealEventStatusAllocated: "real.allocated",
	RealEventStatusFailed:	  "real.failed",
	RealEventStatusCeased:    "real.ceased",
}

var RealEventRollbackTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real.requested_rollback",
	RealEventStatusPending:   "real.pending_rollback",
	RealEventStatusAllocated: "real.allocated_rollback",
	RealEventStatusFailed:	  "real.failed_rollback",
	RealEventStatusCeased:    "real.ceased_rollback",
}