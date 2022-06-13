package mesh

import (
	//"fmt"
	//proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
)

type Mesh struct {
	str string
}

func (m *Mesh) String() string {
	return m.str
}

func FromString(s string) *Mesh {
	return &Mesh{
		str: s,
	}
}
