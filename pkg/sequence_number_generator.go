package pkg

import (
	"time"
)

type SequenceNumberGenerator interface {
	Generate() int
}

type sequenceNumberGenerator struct {
	next int
}

func NewSequenceNumberGenerator() *sequenceNumberGenerator {
	return &sequenceNumberGenerator{next: int(time.Now().Unix())}
}

func (sng *sequenceNumberGenerator) Generate() int {
	seqnum := sng.next
	sng.next = sng.next + 1
	return seqnum
}