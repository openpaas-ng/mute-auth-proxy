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

package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config represents the structure containing the information from the config file
type Config struct {
	Port             int
	ConiksServerAddr string      `toml:"coniksserver_addr"`
	OauthPrefs       OauthConfig `toml:"oauth"`
}

type OauthConfig struct {
	RedirectURL string        `toml:"redirect_url"`
	GooglePrefs ProviderPrefs `toml:"google"`
	GithubPrefs ProviderPrefs `toml:"github"`
}

type ProviderPrefs struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

// LoadConfig loads and parses the information from the config file and fill the Config struct
func LoadConfig(file string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return nil, fmt.Errorf("Failed to load config: %v", err)
	}

	return &conf, nil
}
