package sessionManagement

import "github.com/edipermadi/softhsm/pkg/transport/pb"

type session struct {
	sessionId uint64
	slotId    uint64
	flags     uint64
}

func makeSession(id uint64, req *pb.OpenSessionRequest) session {
	return session{sessionId: id, slotId: req.SlotId, flags: req.Flags}
}
