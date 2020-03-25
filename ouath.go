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

func (c *Client) AuthUser(name, token string) error {
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
	}

	userIdString, err := c.storage.ServiceID("vk", name)
	if err != nil {
		return errors.Wrap(err, "getting service user id")
	}
	uid, err := strconv.Atoi(userIdString)
	if err != nil {
		return errors.New("service user id is not value")
	}

	_, err = c.AccountGetAppPermissions(uid)
	if err != nil {
		return errors.Wrap(err, "checking token validity")
	}

	return c.auth(name, token)
}

func (c *Client) AuthGroup(group, token string) error {
	tmpClient := &Client{
		parentClient: c,
		accessKey:    token,
	}

	_, err := tmpClient.GroupsGetTokenPermissions()
	if err != nil {
		return errors.Wrap(err, "checking token validity")
	}

	return c.auth(group, token)
}

func (c *Client) auth(id, token string) error {
	for _, r := range token {
		if (r < '0' || r > '9') && (r < 'a' || r > 'f') {
			return errors.New("token is invalid")
		}
	}

	err := c.storage.AuthService("vk", id, token)
	if err != nil {
		return errors.Wrap(err, "saving token")

	}

	return nil
}
