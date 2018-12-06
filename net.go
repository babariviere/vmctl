package vmctl

import (
	"errors"
	"fmt"
)

// Net is all net related stuff in VM
type Net struct {
	Enable    bool       `yaml:"enable"`
	Redirects []Redirect `yaml:"redir"`
}

func (n Net) ToQemu() (args []string, err error) {
	if n.Enable {
		args = append(args, "-net", "nic")
	}
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
	HostIP    string `yaml:"host_ip"`
	HostPort  uint16 `yaml:"host_port"`
	GuestIP   string `yaml:"guest_ip"`
	GuestPort uint16 `yaml:"guest_port"`
}

// TODO: check protocol

func (r Redirect) ToQemu() ([]string, error) {
	if r.HostPort == 0 {
		return nil, errors.New("missing host port")
	}
	if r.GuestPort == 0 {
		return nil, errors.New("missing guest port")
	}
	return []string{"-net", fmt.Sprintf("user,hostfwd=tcp:%s:%d-%s:%d", r.HostIP, r.HostPort, r.GuestIP, r.GuestPort)}, nil
}
