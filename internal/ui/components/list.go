package components

import (
	"fmt"

	"github.com/whookdev/cli/internal/projects"
	"github.com/whookdev/cli/internal/ui"
)

type ProjectList struct {
	projects []projects.Project
	cursor   int
}

func NewProjectList(projects []projects.Project) ProjectList {
	return ProjectList{
		projects: projects,
	}
}

func (l ProjectList) Render() string {
	var s string
	s += ui.TitleStyle.Render("Project Selector") + "\n"

	for i, project := range l.projects {
		cursor := " "
		if l.cursor == i {
			cursor = ">"
			s += ui.SelectedItemStyle.Render(fmt.Sprintf("%s %s", cursor, project.Name))
		} else {
			s += ui.NormalItemStyle.Render(fmt.Sprintf("%s %s", cursor, project.Name))
		}
		if i < len(l.projects)-1 {
			s += "\n"
		}
	}

	return s
}

func (l *ProjectList) MoveCursor(up bool) {
	if up {
		if l.cursor > 0 {
			l.cursor--
		}
	} else {
		if l.cursor < len(l.projects)-1 {
			l.cursor++
		}
	}
}

func (l ProjectList) SelectedProject() projects.Project {
	return l.projects[l.cursor]
}
