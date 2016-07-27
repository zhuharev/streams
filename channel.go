package streams

import (
	"time"
)

type Channel struct {
	// for sql db
	Id int64

	IsOnline bool
	Stream   *Stream `xorm:"-"`

	Followers int
	Views     int

	// image http link
	Logo string

	// ie twitch test_channel
	ServiceId string

	// service provider
	ServiceProvider string

	// twitch
	Game string

	// title
	DisplayName string

	// lang
	Language string

	// can be html
	Status string

	// 18+
	Mature bool

	// for db
	Updated time.Time `xorm:"updated"`
	Created time.Time `xorm:"created"`
}

func (ch Channel) SetStream(stream *Stream) {
	if stream != nil {
		ch.Stream = stream
		ch.IsOnline = true
	} else {
		ch.IsOnline = false
	}
}
