package buffers

import (
	"buffers_test/DAL/core"
)

type recordStatus int8

const (
	new      recordStatus = iota
	updated               = iota
	original              = iota
	deleted               = iota
)

type cachedVal struct {
	current  core.CoreModel
	verified core.CoreModel
	status   recordStatus
}
