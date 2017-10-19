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
	"golang.org/x/oauth2/google"
)

// MakeGoogleLoginHandler returns the handler for the Google login route
func MakeGoogleLoginHandler(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		googleOauthConfig := oauth2.Config{
			RedirectURL:  conf.OauthPrefs.RedirectURL,
			ClientID:     conf.OauthPrefs.GooglePrefs.ClientID,
			ClientSecret: conf.OauthPrefs.GooglePrefs.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
		HandleProviderLogin(w, r, "google", googleOauthConfig)
	}
}
