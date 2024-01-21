package task

import (
	"fmt"

	"github.com/azisuazusa/todo-cli/internal/domain/entity"
	"github.com/manifoldco/promptui"
)

var taskSelectTemplate = &promptui.SelectTemplates{
	Active:   `{{ if .ParentTaskID }}  ▸ {{ .Name | cyan }}{{ else }}▸ {{ .Name | cyan }}{{ end }}`,
	Inactive: `{{ if .ParentTaskID }}    {{ .Name }}{{ else }}  {{ .Name }}{{ end }}`,
	Selected: `{{ "✔" | green }} {{ "Selected" | bold }}: {{ .Name | cyan }}`,
	Details:  `{{ "Description:" }} {{ .Description }}`,
}

type TaskView struct {
	Name        string
	Description string
}

func CreateTaskView(taskNumber int, t entity.Task) TaskView {
	taskName := fmt.Sprintf("%d. %s", taskNumber, t.Name)
	if t.IsStarted {
		taskName = fmt.Sprintf("%s (Started)", taskName)
	}

	if t.Integration.ID != "" {
		taskName = fmt.Sprintf("%s (%s)", taskName, t.Integration.ID)
	}

	return TaskView{
		Name:        taskName,
		Description: t.Description,
	}
}
