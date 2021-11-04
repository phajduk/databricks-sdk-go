package clusters

import (
	"encoding/json"

	"github.com/phajduk/databricks-sdk-go/client"
	"github.com/phajduk/databricks-sdk-go/models"
)

type Endpoint struct {
	Client *client.Client
}

func (c *Endpoint) List() (*models.JobsListResponse, error) {
	bytes, err := c.Client.Query("GET", "jobs/list", nil)
	if err != nil {
		return nil, err
	}

	resp := models.JobsListResponse{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
