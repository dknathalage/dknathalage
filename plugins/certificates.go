package plugins

import "github.com/dknathalage/dknathalage/internal"

type CertificateSection struct{}

func (s *CertificateSection) Name() string {
	return "Certificates 📋"
}

func (s *CertificateSection) Generate() (string, error) {
	template := `
- AWS Certified Solutions Architect - Associate
- AWS Certified Developer - Associate
- AWS Certified DevOps - Professional
- GCP Associate Cloud Engineer
	`

	return template, nil
}

func (s *CertificateSection) Order() int {
	return 4
}

func init() {
	internal.RegisterSection(&CertificateSection{})
}
