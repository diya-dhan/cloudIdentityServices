package cli

import (
	"context"
	"encoding/json"

	"terraform-provider-ias/internal/cli/apiObjects/applications"
)

type ApplicationsCli struct {
	cliClient *Client
}

func NewApplicationCli(cliClient *Client) ApplicationsCli {
	return ApplicationsCli{cliClient: cliClient}
}

func (a *ApplicationsCli) getUrl() string {
	return "Applications/v1/"
}

func (a *ApplicationsCli) Get(ctx context.Context) (applications.ApplicationsResponse, error) {
	var app applications.ApplicationsResponse

	res, err, _ := a.cliClient.Execute(ctx, "GET", a.getUrl(), nil, ApplicationHeader, nil)

	if err != nil {
		return app, err
	}

	if err = json.Unmarshal(res, &app); err != nil {
		return app, err
	}
	return app, nil
}

func (a *ApplicationsCli) GetByAppId(ctx context.Context, appId string) (applications.ApplicationResponse, error) {
	var app applications.ApplicationResponse

	res, err, _ := a.cliClient.Execute(ctx, "GET", a.getUrl()+appId[1:len(appId)-1], nil, ApplicationHeader, nil)

	if err != nil {
		return app, err
	}

	if err = json.Unmarshal(res, &app); err != nil {
		return app, err
	}
	return app, nil
}

type ApplicationCreateInput struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *ApplicationsCli) Create(ctx context.Context, args *ApplicationCreateInput) (string, error) {

	var appId string

	_, err, res := a.cliClient.Execute(ctx, "POST", a.getUrl(), args, ApplicationHeader, []string{
		"location",
	})

	if err != nil {
		return appId, err
	}

	appId = res["location"]

	return appId, nil
}
