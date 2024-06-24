package user

import (
	"github.com/notiku/gofm/utils"
)

// Retrieves information about a user.
func (c *UserService) GetInfo(username string) (*User, error) {
	params := map[string]string{
		"method": "user.getinfo",
		"user":   username,
	}

	var result struct {
		User User `json:"user"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.User, nil
}
