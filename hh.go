package hhgo

import "errors"

type HeadHunter interface {
}

type headHunter struct {
	options *Options
}

func New(options *Options) (HeadHunter, error) {
	if options.PublicId == "" {
		return nil, errors.New("empty public id")
	}

	if options.PrivateKey == "" {
		return nil, errors.New("empty private key")
	}

	if options.Endpoint == "" {
		options.Endpoint = ApiEndpoint
	}

	return headHunter{options: options}, nil
}
