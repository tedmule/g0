// Copyright 2015 flannel authors
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

package netswatch

import (
	"fmt"
	"net/http"
	"time"
)

const (
	consulServicesAPI   = "/v1/agent/services"
	consulRegisterAPI   = "/v1/agent/service/register"
	consulDeregisterAPI = "/v1/agent/service/deregister"
)

type NWService struct {
	ID      string
	Address string
	Tags    []string
	Name    string
}

type DNSRegistry struct {
	Endpoint string
	Token    string
}

func (dnsr *DNSRegistry) ListService() {
	url := dnsr.Endpoint + consulServicesAPI
	fmt.Println(url)

	cli := &http.Client{
		Timeout: time.Second * 5,
	}

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Status)

}

func (dnsr *DNSRegistry) RegisterSvc() {
}
