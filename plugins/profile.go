package plugins

import (
	"github.com/dknathalage/dknathalage/internal"
)

type ProfileSection struct{}

func (s *ProfileSection) Name() string {
	return "Profile 👔"
}

func (s *ProfileSection) Generate() (string, error) {
	return "Software engineer passionate about distributed systems and Go programming.", nil
}

func (s *ProfileSection) Order() int {
	return 1
}

func init() {
	internal.RegisterSection(&ProfileSection{})
}
