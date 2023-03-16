package model

type Perm struct {
	UniqueKey   string
	Description string
}

func (p Perm) Key() string {
	return p.UniqueKey
}
