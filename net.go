package vmctl

import (
	"errors"
	"fmt"
)

// Net is all net related stuff in VM
type Net struct {
	Redirects []Redirect `yaml:"redir"`
}

func (n Net) ToQemu() (args []string, err error) {
	for _, r := range n.Redirects {
		res, rerr := r.ToQemu()
		if rerr != nil {
			err = rerr
			return
		}
		args = append(args, res...)
	}
	return
}

// Redirect is a redirection
type Redirect struct {
	Protocol string `yaml:"protocol"`
	Host     uint16 `yaml:"host"`
	Guest    uint16 `yaml:"guest"`
}

// TODO: check protocol

func (r Redirect) ToQemu() ([]string, error) {
	if r.Protocol == "" {
		r.Protocol = "tcp"
	}
	if r.Host == 0 {
		return nil, errors.New("missing host")
	}
	if r.Guest == 0 {
		return nil, errors.New("missing guest")
	}
	return []string{"-redir", fmt.Sprintf("%s:%d::%d", r.Protocol, r.Host, r.Guest)}, nil
}
