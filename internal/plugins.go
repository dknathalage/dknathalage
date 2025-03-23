package internal

type ReadmeSection interface {
	Name() string
	Generate() (string, error)
	Order() int
}

var registeredSections []ReadmeSection

func RegisterSection(section ReadmeSection) {
	registeredSections = append(registeredSections, section)
}

func GetRegisteredSections() []ReadmeSection {
	return registeredSections
}
