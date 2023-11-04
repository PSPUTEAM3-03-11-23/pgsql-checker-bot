package response

import "time"

type Incident struct {
	Id    int
	Db    *Database
	Error *string
	Date  *time.Time
}
