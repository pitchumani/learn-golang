package customerror

import (
	"fmt"
	"time"
)

type ErrTableNotFound struct {
	Table      string
	OccurredAt time.Time
}

func (e ErrTableNotFound) Error() string {
	return fmt.Sprintf("[%s] table not found %s", e.OccurredAt, e.Table)
}
