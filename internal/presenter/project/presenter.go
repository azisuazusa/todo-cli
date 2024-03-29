package project

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/azisuazusa/todo-cli/internal/domain/entity"
	"github.com/azisuazusa/todo-cli/internal/domain/project"
	"github.com/azisuazusa/todo-cli/internal/domain/syncintegration"
	"github.com/azisuazusa/todo-cli/internal/domain/task"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
)

type Presenter struct {
	projectUseCase         project.UseCase
	syncintegrationUseCase syncintegration.UseCase
	taskUseCase            task.UseCase
}

func New(projectUseCase project.UseCase, settingUseCase syncintegration.UseCase, taskUseCase task.UseCase) *Presenter {
	return &Presenter{
		projectUseCase:         projectUseCase,
		syncintegrationUseCase: settingUseCase,
		taskUseCase:            taskUseCase,
	}
}

func (p *Presenter) GetProjects(ctx context.Context) error {
	projects, err := p.projectUseCase.GetAll(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Projects:")
	t.AppendHeader(table.Row{"#", "Name", "Description"})

	for i, project := range projects {
		projectView := CreateProjectView(project)
		t.AppendRow(table.Row{i + 1, projectView.Name, projectView.Description})
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	return nil
}

func (p *Presenter) Add(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Project name",
	}

	name, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	prompt = promptui.Prompt{
		Label: "Description",
	}

	description, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	project := entity.Project{
		Name:        name,
		Description: description,
	}

	if err = p.projectUseCase.Insert(ctx, project); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Project added successfully!")

	return nil
}

func (p *Presenter) Remove(ctx context.Context) error {
	projects, err := p.projectUseCase.GetAll(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	projectViews := []ProjectView{}
	for _, project := range projects {
		projectViews = append(projectViews, CreateProjectView(project))
	}

	prompt := promptui.Select{
		Label:     "Select project to remove",
		Items:     projectViews,
		Templates: projectSelectTemplate,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	err = p.projectUseCase.Delete(ctx, projects[i].ID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Project removed successfully!")

	return nil
}

func (p *Presenter) Select(ctx context.Context) error {
	projects, err := p.projectUseCase.GetAll(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	projectViews := []ProjectView{}
	for _, project := range projects {
		projectViews = append(projectViews, CreateProjectView(project))
	}

	prompt := promptui.Select{
		Label:     "Select project",
		Items:     projectViews,
		Templates: projectSelectTemplate,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	// Makesure there is no task running before change project
	if err = p.taskUseCase.Stop(ctx); err != nil && errors.Unwrap(err) != task.ErrTaskNotFound {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.projectUseCase.Select(ctx, projects[i].ID); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Project selected successfully!")

	return nil
}

func (p *Presenter) SyncTasks(ctx context.Context) error {
	if err := p.projectUseCase.SyncTasks(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Tasks synced successfully!")

	return nil
}

func (p *Presenter) AddIntegration(ctx context.Context) error {
	prompt := promptui.Select{
		Label: "Select integration",
		Items: []string{string(entity.IntegrationTypeJIRA)},
	}

	_, name, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	integration := entity.Integration{
		IsEnabled: true,
		Type:      entity.IntegrationType(name),
	}

	if integration.Type == entity.IntegrationTypeJIRA {
		details, errPrompt := p.jiraPrompt(ctx)
		if errPrompt != nil {
			err = errPrompt
			fmt.Printf("Error: %v\n", err)
			return err
		}

		integration.Details = details
		if err = p.projectUseCase.AddIntegration(ctx, integration); err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}
	}

	if integration.Type == entity.IntegrationTypeGitHub {
		details, errPrompt := p.githubPrompt(ctx)
		if errPrompt != nil {
			err = errPrompt
			fmt.Printf("Error: %v\n", err)
			return err
		}

		integration.Details = details
		if err = p.projectUseCase.AddIntegration(ctx, integration); err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Integration added successfully!")

	return nil
}

func (p *Presenter) jiraPrompt(ctx context.Context) (map[string]string, error) {
	prompt := promptui.Prompt{
		Label: "JIRA URL",
	}

	url, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "JIRA Username",
	}

	username, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "JIRA Token",
	}

	token, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	prompt = promptui.Prompt{
		Label:   "JIRA JQL",
		Default: "assignee = currentUser() AND resolution = Unresolved",
	}

	jql, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return map[string]string{
		"url":      url,
		"username": username,
		"token":    token,
		"jql":      jql,
	}, nil
}

func (p *Presenter) githubPrompt(ctx context.Context) (map[string]string, error) {
	// TODO: implement github integration
	return nil, nil
}

func (p *Presenter) RemoveIntegration(ctx context.Context) (err error) {
	project, err := p.projectUseCase.GetSelected(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	prompt := promptui.Select{
		Label: "Select integration to remove",
		Items: project.Integrations,
		Templates: &promptui.SelectTemplates{
			Active:   `▸ {{ .Type | cyan }}`,
			Inactive: `  {{ .Type }}`,
			Selected: `{{ "✔" | green }} {{ .Type | cyan }}`,
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	err = p.projectUseCase.RemoveIntegration(ctx, project.Integrations[i].Type)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Integration removed successfully!")

	return nil
}

func (p *Presenter) EnableIntegration(ctx context.Context) error {
	project, err := p.projectUseCase.GetSelected(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	prompt := promptui.Select{
		Label: "Select integration to enable",
		Items: project.Integrations,
		Templates: &promptui.SelectTemplates{
			Active:   `▸ {{ .Type | cyan }}`,
			Inactive: `  {{ .Type }}`,
			Selected: `{{ "✔" | green }} {{ .Type | cyan }}`,
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	err = p.projectUseCase.EnableIntegration(ctx, project.Integrations[i].Type)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Integration enabled successfully!")

	return nil
}

func (p *Presenter) DisableIntegration(ctx context.Context) error {
	project, err := p.projectUseCase.GetSelected(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	prompt := promptui.Select{
		Label: "Select integration to disable",
		Items: project.Integrations,
		Templates: &promptui.SelectTemplates{
			Active:   `▸ {{ .Type | cyan }}`,
			Inactive: `  {{ .Type }}`,
			Selected: `{{ "✔" | green }} {{ .Type | cyan }}`,
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	err = p.projectUseCase.DisableIntegration(ctx, project.Integrations[i].Type)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err = p.syncintegrationUseCase.Upload(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Println("Integration disabled successfully!")

	return nil
}
