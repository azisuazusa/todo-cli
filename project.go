package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

type JIRAConfig struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Token    string `json:"token"`
	JQL      string `json:"jql"`
}

type DropboxConfig struct {
	AccessToken string `json:"access_token"`
}

type Project struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	IsSelected  bool          `json:"is_selected"`
	Tasks       Tasks         `json:"tasks"`
	JIRA        JIRAConfig    `json:"jira"`
	Dropbox     DropboxConfig `json:"dropbox"`
}

type Projects []Project

func (p *Project) Add() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	id := uuid.New().String()
	p.ID = id
	projects = append(projects, *p)

	if err := projects.save(); err != nil {
		return err
	}

	fmt.Println("Successfully added project:", p.Name)

	return nil
}

func (p *Projects) load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	file, err := os.ReadFile(homeDir + "/.todo-projects.json")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading file:", err)
		return err
	}

	if os.IsNotExist(err) {
		fmt.Println("File does not exist, creating...")
		if err := os.WriteFile(homeDir+"/.todo-projects.json", []byte(""), 0643); err != nil {
			fmt.Println("Error creating file:", err)
			return err
		}
	}

	if len(file) == 0 {
		fmt.Println("File is empty, returning...")
		return nil
	}

	if err := json.Unmarshal(file, p); err != nil {
		fmt.Println("Error unmarshalling file:", err)
		return err
	}

	return nil
}

func (p *Projects) save() error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling file:", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}

	if err := os.WriteFile(homeDir+"/.todo-projects.json", data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

func ProjectList() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	fmt.Println("Projects:")
	for i, project := range projects {
		fmt.Printf("%d. %s", i+1, project.Name)
		if project.Description != "" {
			fmt.Printf(" | %s", project.Description)
		}
		fmt.Println()
	}

	return nil
}

func ProjectSelect() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	prompt := promptui.Select{
		Label:     "Select project",
		Items:     projects,
		Templates: projectSelectTemplate(),
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Error selecting project:", err)
		return err
	}

	for j := range projects {
		projects[j].IsSelected = false
	}

	projects[i].IsSelected = true
	if err := projects.save(); err != nil {
		return err
	}

	fmt.Println("Successfully selected project:", projects[i].Name)

	return nil
}

func (p *Projects) remove(index int) {
	*p = append((*p)[:index], (*p)[index+1:]...)
}

func ProjectRemove() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	prompt := promptui.Select{
		Label:     "Select project",
		Items:     projects,
		Templates: projectSelectTemplate(),
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Error selecting project:", err)
		return err
	}

	projects.remove(i)
	if err := projects.save(); err != nil {
		fmt.Println("Error removing projects:", err)
		return err
	}

	fmt.Println("Successfully removed project:", projects[i].Name)

	return nil
}

func (p *Projects) getSelectedIndex() int {
	for i, project := range *p {
		if project.IsSelected {
			return i
		}
	}

	return -1
}

func projectSelectTemplate() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Active:   `▸ {{ .Name | cyan }}`,
		Inactive: `  {{ .Name }}`,
		Selected: `{{ "✔" | green }} {{ "Selected" | bold }}: {{ .Name | cyan }}`,
		Details:  `{{ "Description:" }} {{ .Description }}`,
	}
}

func JIRASetting() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	selectedIndex := projects.getSelectedIndex()
	if selectedIndex == -1 {
		fmt.Println("No project selected")
		return nil
	}

	prompt := promptui.Prompt{
		Label: "JIRA URL",
	}

	jiraURL, err := prompt.Run()
	if err != nil {
		fmt.Println("Error getting jira url:", err)
		return err
	}

	prompt = promptui.Prompt{
		Label: "JIRA Username",
	}

	jiraUsername, err := prompt.Run()
	if err != nil {
		fmt.Println("Error getting jira username:", err)
		return err
	}

	prompt = promptui.Prompt{
		Label: "JIRA Token",
		Mask:  '*',
	}

	jiraToken, err := prompt.Run()
	if err != nil {
		fmt.Println("Error getting jira token:", err)
		return err
	}

	projects[selectedIndex].JIRA.URL = jiraURL
	projects[selectedIndex].JIRA.Username = jiraUsername
	projects[selectedIndex].JIRA.Token = jiraToken
	projects[selectedIndex].JIRA.JQL = "assignee = currentUser() AND resolution = Unresolved"
	if err := projects.save(); err != nil {
		return err
	}

	fmt.Println("Successfully updated project:", projects[selectedIndex].Name)

	return nil
}

func DropboxSetting() error {
	projects := Projects{}
	if err := projects.load(); err != nil {
		return err
	}

	selectedIndex := projects.getSelectedIndex()
	if selectedIndex == -1 {
		fmt.Println("No project selected")
		return nil
	}

	prompt := promptui.Prompt{
		Label: "Dropbox Token",
	}

	dropboxToken, err := prompt.Run()
	if err != nil {
		fmt.Println("Error getting dropbox token:", err)
		return err
	}

	projects[selectedIndex].Dropbox.AccessToken = dropboxToken
	if err := projects.save(); err != nil {
		return err
	}

	fmt.Println("Successfully updated project:", projects[selectedIndex].Name)

	return nil
}