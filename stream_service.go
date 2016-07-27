package streams

type StreamService interface {
	Name() string
	TestUrl(string) bool

	GetChannel(channel string) (*Channel, error)
	GetStream(channel string) (*Channel, error)

	GetStreams(channel string, channels ...string) ([]*Channel, error)
}

var providers = map[string]StreamService{}

func Register(name string, ss StreamService) {
	providers[name] = ss
}

func RegisteredProviders() (names []string) {
	for key := range providers {
		names = append(names, key)
	}
	return
}

func GetChannel(provider, channel string) (*Channel, error) {
	return providers[provider].GetChannel(channel)
}

func GetStreams(provider, channel string, channels ...string) ([]*Channel, error) {
	return providers[provider].GetStreams(channel, channels...)
}
