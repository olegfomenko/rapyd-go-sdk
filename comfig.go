package rapyd

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"net/url"
)

type Rapyder interface {
	Rapyd() Client
}

type rapyder struct {
	getter kv.Getter
	once   comfig.Once
}

func NewRapyder(getter kv.Getter) Rapyder {
	return &rapyder{
		getter: getter,
	}
}

func (r *rapyder) Rapyd() Client {
	return r.once.Do(func() interface{} {
		var config struct {
			AccessKey string   `fig:"access_key"`
			SecretKey string   `fig:"secret_key"`
			Endpoint  *url.URL `fig:"endpoint"`
			EWallet   string   `fig:"wallet"`
		}

		if err := figure.Out(&config).From(kv.MustGetStringMap(r.getter, "rapyd")).Please(); err != nil {
			panic(err)
		}

		return NewClient([]byte(config.AccessKey), []byte(config.SecretKey), config.Endpoint, config.EWallet)
	}).(Client)
}
