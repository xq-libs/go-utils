package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	twepoch        = int64(1525705533000)
	workeridBits   = uint(10)
	sequenceBits   = uint(12)
	workeridMax    = int64(-1 ^ -1<<workeridBits)
	sequenceMask   = int64(-1 ^ -1<<sequenceBits)
	workeridShift  = sequenceBits
	timestampShift = workeridBits + sequenceBits
)

type SnowFlake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

func NewSnowFlake(workerid int64) (*SnowFlake, error) {
	if workerid < 0 || workerid > workeridMax {
		return nil, errors.New("error happen")
	}
	return &SnowFlake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

func (s *SnowFlake) Generate() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1000000
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	r := (now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence)
	return r
}
