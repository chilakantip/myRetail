package pg_persist

import (
	"github.com/pkg/errors"
)

var (
	ErrNoRecords      = errors.New("no_records")
	ErrNoRowsAffected = errors.New("no_rows_affected")
)

const ()
