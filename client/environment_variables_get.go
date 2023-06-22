package client

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c *Client) getEnvironmentVariables(ctx context.Context, projectID, teamID string) ([]EnvironmentVariable, error) {
	url := fmt.Sprintf("%s/v8/projects/%s/env?decrypt=true", c.baseURL, projectID)
	if c.teamID(teamID) != "" {
		url = fmt.Sprintf("%s&teamId=%s", url, c.teamID(teamID))
	}

	envResponse := struct {
		Env []EnvironmentVariable `json:"envs"`
	}{}
	tflog.Trace(ctx, "getting environment variables", map[string]interface{}{
		"url": url,
	})
	err := c.doRequest(clientRequest{
		ctx:    ctx,
		method: "GET",
		url:    url,
		body:   "",
	}, &envResponse)
	for _, env := range envResponse.Env {
		env.TeamID = c.teamID(teamID)
	}
	return envResponse.Env, err
}

// GetEnvironmentVariable gets a singluar environment variable from Vercel based on its ID.
func (c *Client) GetEnvironmentVariable(ctx context.Context, projectID, teamID, envID string) (e EnvironmentVariable, err error) {
	url := fmt.Sprintf("%s/v1/projects/%s/env/%s", c.baseURL, projectID, envID)
	if c.teamID(teamID) != "" {
		url = fmt.Sprintf("%s?teamId=%s", url, c.teamID(teamID))
	}

	tflog.Trace(ctx, "getting environment variable", map[string]interface{}{
		"url": url,
	})
	err = c.doRequest(clientRequest{
		ctx:    ctx,
		method: "GET",
		url:    url,
		body:   "",
	}, &e)
	e.TeamID = c.teamID(teamID)
	return e, err
}
