package vk

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xelaj/vk/types"
)

const (
	OauthHost = "oauth.vk.com"
)

func (c *Client) OauthUrl(permissions types.UserPermissions) string {
	v := make(url.Values)
	v.Set("client_id", fmt.Sprint(c.id))
	v.Set("scope", fmt.Sprint(permissions))
	v.Set("response_type", "token")
	u := url.URL{
		Scheme:   "https",
		Host:     OauthHost,
		Path:     "authorize",
		RawQuery: v.Encode(),
	}
	return u.String()
}

func (c *Client) AuthUser(token string) error {
	if strings.HasPrefix(token, "https://oauth.vk.com") {
		u, err := url.Parse(token)
		if err != nil {
			return errors.Wrap(err, "token is not valid url as excpected. cannot parse")
		}
		v, err := url.ParseQuery(u.Fragment)
		if err != nil {
			return errors.Wrap(err, "server side error: url is invalid")
		}

		token := v.Get("access_token")

		if token == "" {
			return errors.New("token not found in url")
		}

		user := v.Get("user_id")
		if user != "" {
			i, _ := strconv.Atoi(user)
			c.clientTokens[i] = token
		}
		return nil
	}

	for _, r := range token {
		if (r < '0' || r > '9') && (r < 'a' || r > 'f') {
			return errors.New("token is invalid")
		}
	}
	if c.NeedToSaveTokens {
		c.SaveTo("")
	}
	return nil
}

func (c *Client) AuthGroup(groupID int, token string) error {
	c.groupTokens[groupID] = token

	// test
	_, err := c.By(-groupID).GroupsGetTokenPermissions()
	if err != nil {
		delete(c.groupTokens, groupID)
		return errors.Wrap(err, "checking group token")
	}

	return nil
}
