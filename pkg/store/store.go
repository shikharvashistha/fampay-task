package store

import "github.com/shikharvashistha/fampay/pkg/store/relational"

type store struct {
	rl relational.RL
}

func (s *store) RL() relational.RL {
	return s.rl
}
