package worker

import (
	"github.com/luqman-v1/absent/repo"
)

type Payload struct {
	Status string
}

func (p *Payload) Handle() error {
	repoPresent := repo.NewRepo()
	repoPresent.Login()
	repoPresent.Present(p.Status)
	return nil
}
