package jira

import (
	"context"
	"fmt"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/azisuazusa/todo-cli/internal/domain/entity"
)

type RepoImpl struct{}

func New() *RepoImpl {
	return &RepoImpl{}
}

func (ri *RepoImpl) initJIRAClient(ctx context.Context, integrationEntity entity.Integration) (*jira.Client, error) {
	jiraAuth := jira.BasicAuthTransport{
		Username: integrationEntity.Details["username"],
		Password: integrationEntity.Details["password"],
	}

	client, err := jira.NewClient(jiraAuth.Client(), integrationEntity.Details["url"])
	if err != nil {
		return nil, fmt.Errorf("error while creating jira client: %w", err)
	}

	return client, nil
}

func (ri *RepoImpl) GetTasks(ctx context.Context, projectID string, integrationEntity entity.Integration) ([]entity.Task, error) {
	client, err := ri.initJIRAClient(ctx, integrationEntity)
	if err != nil {
		return []entity.Task{}, fmt.Errorf("error while initializing jira client: %w", err)
	}

	last := 0
	var issues []jira.Issue
	for {
		opt := &jira.SearchOptions{
			StartAt:    last,
			MaxResults: 1000,
		}

		chunk, resp, errSearch := client.Issue.Search(integrationEntity.Details["jql"], opt)
		if errSearch != nil {
			err = errSearch
			return []entity.Task{}, fmt.Errorf("error while searching issues: %w", err)
		}

		total := resp.Total
		if issues == nil {
			issues = make([]jira.Issue, 0, total)
		}

		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)
		if last >= resp.Total {
			break
		}
	}

	var tasks []entity.Task
	for _, issue := range issues {
		task := entity.Task{
			ProjectID:   projectID,
			Name:        issue.Fields.Summary,
			Description: issue.Fields.Description,
			CompletedAt: time.Time{},
			Integration: entity.TaskIntegration{
				ID:   issue.ID,
				Type: entity.IntegrationTypeJIRA,
			},
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (ri *RepoImpl) AddWorklog(ctx context.Context, issueID, taskName string, timeSpent time.Duration, integrationEntity entity.Integration) error {
	client, err := ri.initJIRAClient(ctx, integrationEntity)
	if err != nil {
		return fmt.Errorf("error while initializing jira client: %w", err)
	}

	worklog := jira.WorklogRecord{
		Comment:   taskName,
		TimeSpent: timeSpent.String(),
	}

	_, _, err = client.Issue.AddWorklogRecordWithContext(ctx, issueID, &worklog)
	if err != nil {
		return fmt.Errorf("error while adding worklog: %w", err)
	}

	return nil

}
