package plugins

import "github.com/dknathalage/dknathalage/internal"

type SkillSection struct{}

func (s *SkillSection) Name() string {
	return "Skills"
}

func (s *SkillSection) Generate() (string, error) {
	template := `
### Programming Languages
- Go
- Python
- JavaScript
- Swift

### DevOps Tools
- Docker
- Kubernetes
- Terraform
- GitHub Actions

### Cloud Platforms
- AWS
- GCP
`

	return template, nil
}

func (s *SkillSection) Order() int {
	return 3
}

func init() {
	internal.RegisterSection(&SkillSection{})
}
