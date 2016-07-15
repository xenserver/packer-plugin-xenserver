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
	"fmt"
	"net/url"
	"strings"
)

const (
	key = iota
	value
)

type csparams map[string]string

// Return a checked url.Values value
func processParams(cs *CloudStackClient, cmd *command, rawParams interface{}) (url.Values, error) {
	params, err := convertType(rawParams)
	if err != nil {
		return nil, err
	}
	if err := verifyRequiredParams(cs, cmd, &params); err != nil {
		return nil, err
	}
	if err := verifyAllUsedParams(cs, cmd, &params); err != nil {
		return nil, err
	}
	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Set(key, value)
	}
	return urlValues, nil
}

func convertType(rawParams interface{}) (params csparams, err error) {
	switch rawParams.(type) {
	case string:
		params, err = parseParams(rawParams.(string))
	case map[string]string:
		params = csparams(rawParams.(map[string]string))
	default:
		return nil, fmt.Errorf("Wrong type: '%T'; only 'string' and 'map[string]string' are valid rawParams types", rawParams)
	}
	return
}

func parseParams(rawParams string) (csparams, error) {
	params := csparams{}
	if len(rawParams) == 0 {
		return params, nil
	}
	paramSlice := strings.Split(rawParams, ",")
	for _, paramSet := range paramSlice {
		keyValue := strings.SplitN(paramSet, ":", 2)
		if len(keyValue) != 2 || len(keyValue[value]) == 0 {
			return nil, fmt.Errorf("Error processing parameters around: ... %s ...", paramSet)
		}
		params[strings.ToLower(strings.Trim(strings.TrimSpace(keyValue[key]), `"`))] = strings.Trim(strings.TrimSpace(keyValue[value]), `"`)
	}
	return params, nil
}

// Test to see of all the required parameters are set
func verifyRequiredParams(cs *CloudStackClient, cmd *command, rawParams *csparams) error {
	params := *rawParams
	for key, _ := range cmd.RequiredParams {
		if _, found := params[key]; !found {
			if _, found := params[strings.TrimSuffix(key, "id")]; found {
				if err := convertToId(cs, strings.TrimSuffix(key, "id"), &params); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("Missing required parameter: %s", key)
			}
		}
	}
	return nil
}

// Test to see is all used parameters are valid parameters for the requested command
func verifyAllUsedParams(cs *CloudStackClient, cmd *command, rawParams *csparams) error {
	params := *rawParams
	// Make a slice of keys to walk over, as we could be changing the params value
	// during the loop, if we find a key that needs to be resolved to an id
	keys := []string{}
	for key, _ := range params {
		keys = append(keys, key)
	}
	for _, key := range keys {
		if _, found := cmd.RequiredParams[key]; !found {
			if _, found := cmd.OptionalParams[key]; !found {
				if _, found := cmd.OptionalParams[fmt.Sprintf("%sid", key)]; found {
					if err := convertToId(cs, key, &params); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("Parameter '%s' is not a valid parameter", key)
				}
			}
		}
	}
	return nil
}

func convertToId(cs *CloudStackClient, key string, rawParams *csparams) error {
	params := *rawParams
	rawJSON, err := cs.RawRequest(fmt.Sprintf("list%ss", key), fmt.Sprintf("name:%s", params[key]))
	if err != nil {
		return fmt.Errorf("Error trying to resolve the id for %s %s: %s", key, params[key], err)
	}
	id, err := unmarshalListId(key, rawJSON)
	if err != nil {
		return fmt.Errorf("Error trying to resolve the id for %s %s: %s", key, params[key], err)
	}
	delete(params, key)
	params[fmt.Sprintf("%sid", key)] = id
	return nil
}
