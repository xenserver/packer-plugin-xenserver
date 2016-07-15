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
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type commands map[string]*command

type CloudStackClient struct {
	client      *http.Client // The http client for communicating
	apiCommands commands     // All available API commands
	baseURL     string       // The base URL of the API
	apiKey      string       // Api key
	secret      string       // Secret key
	timeout     int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 60 seconds
}

// Creates a new client for communicating with CloudStack
func newClient(apiurl string, apikey string, secret string, verifyssl bool) *CloudStackClient {
	cs := CloudStackClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}, // If verifyssl is true, skipping the verify should be false and vice versa
			},
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		timeout: 60,
	}
	return &cs
}

// Creates a new non-caching CloudStack client that doesn't use any local cache for the API command info. This means that every request
// made using this client will result in (at least) 2 calls to the API. The first one will retrieve and process all available commands
// (including their details like required and optional parameters) and the second one will be the actual request. In between the
// retrieved API command info will be used to validate the request before actually executing it.
//
// When using HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings.
func NewClient(apiurl string, apikey string, secret string, verifyssl bool) (*CloudStackClient, error) {
	cs := newClient(apiurl, apikey, secret, verifyssl)
	err := cacheAPICommands(cs, 0)
	return cs, err
}

// Creates a new caching CloudStack client which uses a local cache file for the API command info which in turn is used to validate a
// request before executing it. The 'cacheExpirationInDays' parameter is used to determine when the cache needs to be updated. It's
// recommended to always use this client when working with a (more or less) stable CloudStack environment as using this client will speed
// up requests considerably and will prevent unnecessary load on the API server. When you are developing agains multiple versions of
// CloudStack or if you are using this package as part of a tool/program to migrate or update CloudStack, it's probably a good idea to
// use the non-caching client instead.
//
// When using HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings.
func NewCachingClient(apiurl string, apikey string, secret string, cacheExpirationInDays int, verifyssl bool) (*CloudStackClient, error) {
	cs := newClient(apiurl, apikey, secret, verifyssl)
	err := cacheAPICommands(cs, cacheExpirationInDays)
	return cs, err
}

// When using a synced request the function will wait for the async call to finish before returning. The default is to poll for 60
// seconds, to check if the async job is finished. If the job is not finished within this time, the async job details will be returned
// instead of the actual result of the async command requested. Using this method you can tweak the timeout to fit your needs.
func (cs *CloudStackClient) AsyncTimeout(timeoutInSeconds int64) {
	cs.timeout = timeoutInSeconds
}

// Executes a raw request after checking and verifing the command and parameters first. The parameters can be passed as a string or
// as a map[string]string. In the latter case the keys of the map are treated as the parameter names and the values as, well, the values :)
// When sending the parameters as a single string, it needs to have the following format: "param1:value1, param2:value with space2,
// param3:value3". The return value of a successful raw request is always the raw JSON that is received as response from the API and nil.
// The return value of a failed raw request is nil and a detailed error describing the problem.
func (cs *CloudStackClient) RawRequest(command string, params interface{}) (json.RawMessage, error) {
	cmd, found := cs.apiCommands[strings.ToLower(command)]
	if !found {
		return nil, fmt.Errorf("'%s' is not a valid CloudStack command", command)
	}
	parms, err := processParams(cs, cmd, params)
	if err != nil {
		return nil, err
	}
	return newRequest(cs, cmd, parms)
}

// Executes a request after checking and verifing the command and parameters first. The parameters can be passed as a string or
// as a map[string]string. In the latter case the keys of the map are treated as the parameter names and the values as, well,
// the values :) When sending the parameters as a single string, it needs to have the following format: "param1:value1,
// param2:value with space2, param3:value3". The return value of a successful request is either a single string representing the
// returned ID or an empty string when the requested command doesn't return an ID and nil. The return value of a failed request
// is an empty string and a detailed error describing the problem. So you would definitely want to check the returned error
// value here to determine if the call was successful or not.
func (cs *CloudStackClient) Request(command string, params interface{}) (string, error) {
	rawJSON, err := cs.RawRequest(command, params)
	if err != nil {
		return "", err
	}
	cmd := cs.apiCommands[strings.ToLower(command)]
	if strings.HasPrefix(strings.ToLower(cmd.Name), "list") {
		return unmarshalListId(strings.TrimSuffix(strings.TrimPrefix(strings.ToLower(cmd.Name), "list"), "s"), rawJSON)
	}
	if cmd.ReturnsAnId {
		return unmarshalId(strings.ToLower(cmd.Name), rawJSON)
	}
	return "", nil
}

// For sync API commands this function behaves exactly the same as a standard RawRequest() call, but for async API commands
// this function will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return the raw JSON response received from the API and nil, but when the timout is
// reached it will return nil and an error containing the async job ID for the running job.
func (cs *CloudStackClient) RawSyncedRequest(command string, params interface{}) (json.RawMessage, error) {
	rawJSON, err := cs.RawRequest(command, params)
	if err != nil {
		return nil, err
	}
	cmd := cs.apiCommands[strings.ToLower(command)]
	if cmd.Isasync {
		asyncJobResult, err := getAsyncJobResult(cs, rawJSON)
		if err != nil {
			return nil, err
		}
		if asyncJobResult.Jobstatus == 2 {
			if asyncJobResult.Jobresulttype == "text" {
				return nil, fmt.Errorf(string(asyncJobResult.Jobresult))
			} else {
				return nil, fmt.Errorf("Undefined error: %s", string(asyncJobResult.Jobresult))
			}
		}
		return asyncJobResult.Jobresult, nil
	}
	return rawJSON, nil
}

// For sync API commands this function behaves exactly the same as a standard Request() call, but for async API commands
// this function will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return the ID or an empty string when the requested command doesn't return an ID and
// nil. When the timout is reached it will return an empty string and an error containing the async job ID for the running job.
func (cs *CloudStackClient) SyncedRequest(command string, params interface{}) (string, error) {
	rawJSON, err := cs.RawSyncedRequest(command, params)
	if err != nil {
		return "", err
	}
	cmd := cs.apiCommands[strings.ToLower(command)]
	if strings.HasPrefix(strings.ToLower(cmd.Name), "list") {
		return unmarshalListId(strings.TrimSuffix(strings.TrimPrefix(strings.ToLower(cmd.Name), "list"), "s"), rawJSON)
	}
	if cmd.ReturnsAnId {
		return unmarshalId(strings.ToLower(cmd.Name), rawJSON)
	}
	return "", nil
}

func getAsyncJobResult(cs *CloudStackClient, rawJSON json.RawMessage) (*asyncJobResult, error) {
	result, err := unmarshalAsyncResponse(rawJSON)
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().Unix()
	for {
		rawJSON, err := cs.RawRequest("queryAsyncJobResult", fmt.Sprintf("jobid:%s", result.Jobid))
		if err != nil {
			return nil, err
		}
		jobResult, err := unmarshalAsyncJobResponse(rawJSON)
		if err != nil {
			return nil, err
		}
		if jobResult.Jobstatus > 0 {
			return jobResult, nil
		}
		if time.Now().Unix()-currentTime > cs.timeout {
			return nil, fmt.Errorf("Timeout while waiting for async job to finish (jobid:%s)", result.Jobid)
		}
		time.Sleep(3 * time.Second)
	}
}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occured. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func newRequest(cs *CloudStackClient, cmd *command, params url.Values) (json.RawMessage, error) {
	client := cs.client

	params.Set("apikey", cs.apiKey)
	params.Set("command", cmd.Name)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters and sort them by key, done by Encode
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := params.Encode()
	s2 := strings.ToLower(s)
	s3 := strings.Replace(s2, "+", "%20", -1)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signature = url.QueryEscape(signature)

	// Create the final URL before we issue the request
	url := cs.baseURL + "?" + s + "&signature=" + signature

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	rawJSON, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errorInfo, err := unmarshalCsError(rawJSON)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("CloudStack API error %d (CSExceptionErrorCode: %d): %s", errorInfo.Errorcode, errorInfo.CsErrorcode, errorInfo.Errortext)
	}
	return rawJSON, nil
}
