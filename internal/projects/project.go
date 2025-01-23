package projects

type Project struct {
	Name string
}

type Manager struct {
	projects []Project
}

func NewManager() *Manager {
	return &Manager{
		projects: []Project{
			{Name: "Project 1"},
			{Name: "Project 2"},
			{Name: "Project 3"},
		},
	}
}

func (m *Manager) AddProject(name string) error {
	if name == "" {
		return ErrEmptyProjectName
	}

	m.projects = append(m.projects, Project{Name: name})
	return nil
}

func (m *Manager) Projects() []Project {
	return m.projects
}
