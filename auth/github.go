// Copyright 2017 Jean-Philippe Eisenbarth
//
// This file is part of Mute Authentication Proxy.
//
// Foobar is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Mute Authentication Proxy is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Foobar. See the file COPYING.  If not, see <http://www.gnu.org/licenses/>.

package auth

import (
	"net/http"

	"github.com/coast-team/mute-auth-proxy/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// MakeGithubLoginHandler returns the handler for the Github login route
func MakeGithubLoginHandler(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		githubOauthConfig := oauth2.Config{
			RedirectURL:  conf.OauthPrefs.RedirectURL,
			ClientID:     conf.OauthPrefs.GithubPrefs.ClientID,
			ClientSecret: conf.OauthPrefs.GithubPrefs.ClientSecret,
			Scopes:       []string{""},
			Endpoint:     github.Endpoint,
		}
		HandleProviderLogin(w, r, "github", githubOauthConfig)
	}
}
