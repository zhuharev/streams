package twitch

import (
	"github.com/zhuharev/streams"
)

func init() {
	streams.Register(ProviderName, &TwitchServiceProvider{})
}
