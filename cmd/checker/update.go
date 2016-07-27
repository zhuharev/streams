package main

import (
	"fmt"
	"github.com/urfave/cli"

	"github.com/zhuharev/streams"
	_ "github.com/zhuharev/streams/providers/twitch"
)

func update(c *cli.Context) error {

	for _, provider := range streams.RegisteredProviders() {
		var (
			chs    []*streams.Channel
			newChs []*streams.Channel
		)
		e := x.Where("service_provider = ?", provider).Find(&chs)
		if e != nil {
			return e
		}
		var ids []string
		if len(chs) > 1 {
			for _, ch := range chs {
				ids = append(ids, ch.ServiceId)
			}
			newChs, e = streams.GetStreams(provider, ids[0], ids[1:]...)
			if e != nil {
				return e
			}
		} else if len(chs) == 1 {
			newChs, e = streams.GetStreams(provider, chs[0].ServiceId)
			if e != nil {
				return e
			}
		}

		for _, oldCh := range chs {
			var newCh *streams.Channel
			for _, v := range newChs {
				if oldCh.ServiceProvider == v.ServiceProvider && oldCh.ServiceId == v.ServiceId {
					newCh = v
				}
			}

			if newCh == nil {
				continue
			}

			_, e := x.Id(oldCh.Id).Cols("is_online").Update(newCh)
			if e != nil {
				return e
			}
			fmt.Printf("update channel %d set online is %v \n", oldCh.Id, newCh.IsOnline)

			fmt.Println(newCh.IsOnline, newCh.Stream)

			if oldCh.IsOnline == newCh.IsOnline {
				if newCh.Stream != nil {
					_, e := x.Where("service_id = ?", newCh.Stream.ServiceId).Cols("viewers").Update(newCh.Stream)
					if e != nil {
						return e
					}
				}
				continue
			}
			oldCh.IsOnline = newCh.IsOnline
			if newCh.Stream != nil {
				newCh.Stream.ChannelId = oldCh.Id
				_, e := x.Insert(newCh.Stream)
				if e != nil {
					return e
				}
			}

		}
	}

	fmt.Println(Config.Db.Driver)
	return nil
}

func get(service, user string) (*streams.Channel, error) {

	var (
		channel = new(streams.Channel)

		strms []*streams.Stream
	)

	_, e := x.Where("service_provider = ? and service_id = ?", service, user).Get(channel)
	if e != nil {
		return nil, e
	}

	e = x.Where("channel_id = ? and service_id = ?", channel.Id, user).OrderBy(Config.Db.IdName + " desc").Limit(1).Find(&strms)
	if e != nil {
		return nil, e
	}

	if strms != nil {
		channel.Stream = strms[0]
	}

	return channel, e
}

func add(c *cli.Context) error {
	var (
		// todo
		//u       = c.String("url")

		user    = c.String("username")
		service = c.String("service")
	)
	fmt.Println(streams.RegisteredProviders)
	ch, e := streams.GetChannel(service, user)
	if e != nil {
		return e
	}

	channel := new(streams.Channel)

	has, e := x.Where("service_provider = ? and service_id = ?", service, user).Get(channel)
	if e != nil {
		return e
	}
	if has {
		return fmt.Errorf("User %s exist!", user)
	}

	_, e = x.Insert(ch)
	if e != nil {
		return e
	}

	return nil
}

func status(c *cli.Context) error {
	fmt.Println("run status")
	var (
		// todo
		//u       = c.String("url")

		user    = c.String("username")
		service = c.String("service")
	)

	channel, e := get(service, user)
	if e != nil {
		return e
	}

	fmt.Println(channel.DisplayName, channel.IsOnline)
	return nil
}
