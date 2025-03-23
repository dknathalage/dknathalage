package plugins

import "github.com/dknathalage/dknathalage/internal"

type SkillSection struct{}

func (s *SkillSection) Name() string {
	return "Skills"
}

func (s *SkillSection) Generate() (string, error) {
	template := `
<table>
	<tr>
		<td>
			<h3>Programming Languages</h3>
			<ul>
				<li>Go</li>
				<li>Python</li>
				<li>JavaScript</li>
				<li>Swift</li>
			</ul>
		</td>
		<td>
			<h3>DevOps Tools</h3>
			<ul>
				<li>Docker</li>
				<li>Kubernetes</li>
				<li>Terraform</li>
				<li>GitHub Actions</li>
			</ul>
		</td>
		<td>
			<h3>Cloud Platforms</h3>
			<ul>
				<li>AWS</li>
				<li>GCP</li>
			</ul>
		</td>
	</tr>
</table>
`

	return template, nil
}

func (s *SkillSection) Order() int {
	return 3
}

func init() {
	internal.RegisterSection(&SkillSection{})
}
