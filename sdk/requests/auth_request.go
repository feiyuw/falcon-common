// Copyright 2017 Xiaomi, Inc.
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

package requests

import (
	"encoding/json"
	"errors"
	"github.com/toolkits/net/httplib"
	"time"
)

// CurlPlus is used to send curl request to falcon API
func CurlPlus(uri, method, tokenName, tokenSig string, headers, params map[string]string) (req *httplib.BeegoHttpRequest, err error) {
	switch method {
	case "GET":
		req = httplib.Get(uri)
	case "POST":
		req = httplib.Post(uri)
	case "PUT":
		req = httplib.Put(uri)
	case "DELETE":
		req = httplib.Delete(uri)
	case "HEAD":
		req = httplib.Head(uri)
	default:
		err = errors.New("invalid http method")
		return
	}

	req = req.SetTimeout(1*time.Second, 5*time.Second)

	token, _ := json.Marshal(map[string]string{
		"name": tokenName,
		"sig":  tokenSig,
	})
	req.Header("Apitoken", string(token))

	for hk, hv := range headers {
		req.Header(hk, hv)
	}

	for pk, pv := range params {
		req.Param(pk, pv)
	}

	return
}
