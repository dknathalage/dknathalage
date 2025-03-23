package plugins

import (
	"github.com/dknathalage/dknathalage/internal"
)

type ProjectsSection struct{}

func (s *ProjectsSection) Name() string {
	return "Projects"
}

func (s *ProjectsSection) Generate() (string, error) {
	return "- **This Repo**: Self-updating GitHub profile\n" +
		"- **Other Projects**: Coming soon!", nil
}

func (s *ProjectsSection) Order() int {
	return 2
}

func init() {
	internal.RegisterSection(&ProjectsSection{})
}
