package streams

import (
	"time"
)

type Stream struct {
	// for sql db
	Id        int64
	ChannelId int64
	ServiceId string

	Game string

	Viewers int

	CreatedAt time.Time
	EndedAt   time.Time

	Preview string

	// for db
	Updated time.Time `xorm:"updated"`
	Created time.Time `xorm:"created"`
}
