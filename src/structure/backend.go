package structure

import "util"

type Backend struct {
	Host 	string `json:"host"`
	Port 	uint16 `json:"port"`
}

func (backend *Backend)Url() string {
	return util.HostAndPortToAddress(backend.Host, backend.Port)
}