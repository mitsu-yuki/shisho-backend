package size

import "time"

type Size struct {
	id           string
	name         string
	createAt     time.Time
	lastUpdateAt time.Time
	deletedAt    *time.Time
}

func (s *Size) ID() string {
	return s.id
}

func (s *Size) Name() string {
	return s.name
}

func (s *Size) CreateAt() time.Time {
	return s.createAt
}

func (s *Size) LastUpdateAt() time.Time {
	return s.lastUpdateAt
}

func (s *Size) DeletedAt() *time.Time {
	return s.deletedAt
}
