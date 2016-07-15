//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gocs

import (
	"encoding/gob"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mitchellh/osext"
)

var cacheFile string

func setCacheFile() error {
	if cacheFile == "" {
		exe, err := osext.Executable()
		if err != nil {
			return err
		}
		strings.TrimSuffix(exe, path.Ext(exe))
		cacheFile = exe + ".cache"
	}
	return nil
}

// Updates (or creates) a cache containing the API's structure
func cacheAPICommands(cs *CloudStackClient, cacheExpirationInDays int) (err error) {
	if cacheExpirationInDays > 0 {
		if err := setCacheFile(); err != nil {
			return err
		}
		beforeDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-cacheExpirationInDays, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
		if fileInfo, err := os.Stat(cacheFile); err != nil || fileInfo.ModTime().Before(beforeDate) {
			if err := updateAPICommandFile(cs); err != nil {
				return err
			}
		}
		cs.apiCommands, err = readAPICommandFile(cs)
	} else {
		cs.apiCommands, err = requestAPICommands(cs)
	}
	return
}

func readAPICommandFile(cs *CloudStackClient) (commands, error) {
	reader, err := os.Open(cacheFile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decoder := gob.NewDecoder(reader)
	var apiCommands commands
	if err := decoder.Decode(&apiCommands); err != nil {
		return nil, err
	}
	return apiCommands, nil
}

// Calls newRequest() directly as there is no cache yet, so it's not
// possible to check and verify this call which is done by the exported
// request functions
func requestAPICommands(cs *CloudStackClient) (commands, error) {
	rawJSON, err := newRequest(cs, &command{Name: "listApis"}, url.Values{})
	if err != nil {
		return nil, err
	}
	return unmarshalApiCommands(rawJSON)
}

func updateAPICommandFile(cs *CloudStackClient) error {
	result, err := requestAPICommands(cs)
	if err != nil {
		return err
	}

	writer, err := os.Create(cacheFile)
	if err != nil {
		return err
	}
	defer writer.Close()

	encoder := gob.NewEncoder(writer)
	return encoder.Encode(result)
}
