package twitch

import (
	"fmt"
	"github.com/zhuharev/streams"
	"time"
)

type TwitchResponse struct {
	Streams []TwitchStream `json:"streams"`
	Total   int            `json:"_total"`
	Links   struct {
		Self     string `json:"self"`
		Next     string `json:"next"`
		Featured string `json:"featured"`
		Summary  string `json:"summary"`
		Followed string `json:"followed"`
	} `json:"_links"`
}

type TwitchStream struct {
	ID          int64     `json:"_id"`
	Game        string    `json:"game"`
	Viewers     int       `json:"viewers"`
	CreatedAt   time.Time `json:"created_at"`
	VideoHeight int       `json:"video_height"`
	AverageFps  float64   `json:"average_fps"`
	Delay       int       `json:"delay"`
	IsPlaylist  bool      `json:"is_playlist"`
	Links       struct {
		Self string `json:"self"`
	} `json:"_links"`
	Preview struct {
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Large    string `json:"large"`
		Template string `json:"template"`
	} `json:"preview"`
	Channel TwitchChannel `json:"channel"`
}

func (ts TwitchStream) ToStreamsStream() streams.Stream {
	str := streams.Stream{}
	str.ServiceId = fmt.Sprint(ts.ID)
	str.CreatedAt = ts.CreatedAt
	str.Game = ts.Game
	str.Preview = ts.Preview.Large
	str.Viewers = ts.Viewers
	return str
}

type TwitchChannel struct {
	Mature                       bool        `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Delay                        interface{} `json:"delay"`
	Logo                         string      `json:"logo"`
	Banner                       interface{} `json:"banner"`
	VideoBanner                  string      `json:"video_banner"`
	Background                   interface{} `json:"background"`
	ProfileBanner                interface{} `json:"profile_banner"`
	ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
	Links                        struct {
		Self          string `json:"self"`
		Follows       string `json:"follows"`
		Commercial    string `json:"commercial"`
		StreamKey     string `json:"stream_key"`
		Chat          string `json:"chat"`
		Features      string `json:"features"`
		Subscriptions string `json:"subscriptions"`
		Editors       string `json:"editors"`
		Teams         string `json:"teams"`
		Videos        string `json:"videos"`
	} `json:"_links"`
}

func (tc TwitchChannel) ToStreamsChannel() *streams.Channel {
	ch := new(streams.Channel)
	ch.DisplayName = tc.DisplayName
	ch.Followers = tc.Followers
	ch.Game = tc.Game
	ch.Language = tc.BroadcasterLanguage
	ch.Logo = tc.Logo
	ch.Mature = tc.Mature
	ch.ServiceId = tc.Name
	ch.ServiceProvider = ProviderName
	ch.Status = tc.Status
	ch.Updated = tc.UpdatedAt
	ch.Views = tc.Views
	return ch
}
