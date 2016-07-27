package twitch

import (
	"fmt"
	"github.com/Unknwon/com"
	"net/http"
	"strings"

	"github.com/zhuharev/streams"
)

const (
	ProviderName = "twitch"
)

var (
	baseUrl = "https://api.twitch.tv/kraken"

	channelsEndpoint = "channels"
	streamsEndpoint  = "streams"
)

func getChannelEndpoint(channel string) string {
	return fmt.Sprintf("%s/%s/%s", baseUrl, channelsEndpoint, channel)
}

func getStreamsEndpoint(channel string, channels ...string) string {
	channels = append(channels, channel)
	return fmt.Sprintf("%s/%s?limit=100&channel=%s", baseUrl, streamsEndpoint, strings.Join(channels, ","))
}

type TwitchServiceProvider struct {
}

func (tsp *TwitchServiceProvider) Name() string {
	return ProviderName
}

func (tsp *TwitchServiceProvider) TestUrl(u string) bool {
	var (
		prefixes = []string{
			"twitch.tv",
			"www.twitch.tv",
			"http://twitch.tv",
			"https://twitch.tv",
			"http://www.twitch.tv",
			"https://www.twitch.tv",
		}
	)
	for _, prefix := range prefixes {
		if strings.HasPrefix(u, prefix) {
			return true
		}
	}
	return false
}

func (tsp *TwitchServiceProvider) GetChannel(channel string) (*streams.Channel, error) {

	var (
		u  = getChannelEndpoint(channel)
		tc = new(TwitchChannel)
	)

	e := com.HttpGetJSON(http.DefaultClient, u, &tc)
	if e != nil {
		return nil, e
	}

	return tc.ToStreamsChannel(), nil
}

func (tsp *TwitchServiceProvider) GetStream(channel string) (*streams.Channel, error) {
	response, e := tsp.GetStreams(channel)
	if response != nil {
		return response[0], e
	} else {
		return nil, e
	}
}

func (tsp *TwitchServiceProvider) GetStreams(channel string, channels ...string) ([]*streams.Channel, error) {

	var (
		u  = getStreamsEndpoint(channel, channels...)
		tr = new(TwitchResponse)

		res = make([]*streams.Channel, len(channels)+1)
	)

	channels = append(channels, channel)

	e := com.HttpGetJSON(http.DefaultClient, u, &tr)
	if e != nil {
		return nil, e
	}

	for index, channelName := range channels {
		var (
			has = false
			id  = 0
			ch  *streams.Channel
		)

		for i, stream := range tr.Streams {
			if stream.Channel.Name == channelName {
				has = true
				id = i
			}
		}
		if has {

			ch = tr.Streams[id].Channel.ToStreamsChannel()
			ch.IsOnline = true

			channelStream := tr.Streams[id].ToStreamsStream()

			ch.Stream = &channelStream
			res[index] = ch
		} else {
			ch = &streams.Channel{}
			ch.ServiceId = channelName
			ch.ServiceProvider = ProviderName
			res[index] = ch
		}

	}

	return res, nil
}
