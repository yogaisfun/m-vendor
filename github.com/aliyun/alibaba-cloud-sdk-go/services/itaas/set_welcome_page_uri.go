package itaas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetWelcomePageURI invokes the itaas.SetWelcomePageURI API synchronously
// api document: https://help.aliyun.com/api/itaas/setwelcomepageuri.html
func (client *Client) SetWelcomePageURI(request *SetWelcomePageURIRequest) (response *SetWelcomePageURIResponse, err error) {
	response = CreateSetWelcomePageURIResponse()
	err = client.DoAction(request, response)
	return
}

// SetWelcomePageURIWithChan invokes the itaas.SetWelcomePageURI API asynchronously
// api document: https://help.aliyun.com/api/itaas/setwelcomepageuri.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetWelcomePageURIWithChan(request *SetWelcomePageURIRequest) (<-chan *SetWelcomePageURIResponse, <-chan error) {
	responseChan := make(chan *SetWelcomePageURIResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetWelcomePageURI(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetWelcomePageURIWithCallback invokes the itaas.SetWelcomePageURI API asynchronously
// api document: https://help.aliyun.com/api/itaas/setwelcomepageuri.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetWelcomePageURIWithCallback(request *SetWelcomePageURIRequest, callback func(response *SetWelcomePageURIResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetWelcomePageURIResponse
		var err error
		defer close(result)
		response, err = client.SetWelcomePageURI(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetWelcomePageURIRequest is the request struct for api SetWelcomePageURI
type SetWelcomePageURIRequest struct {
	*requests.RpcRequest
	Clientappid string `position:"Query" name:"Clientappid"`
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Welcomeuri  string `position:"Query" name:"Welcomeuri"`
}

// SetWelcomePageURIResponse is the response struct for api SetWelcomePageURI
type SetWelcomePageURIResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	ErrorCode int                          `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string                       `json:"ErrorMsg" xml:"ErrorMsg"`
	Success   bool                         `json:"Success" xml:"Success"`
	ErrorList ErrorListInSetWelcomePageURI `json:"ErrorList" xml:"ErrorList"`
}

// CreateSetWelcomePageURIRequest creates a request to invoke SetWelcomePageURI API
func CreateSetWelcomePageURIRequest() (request *SetWelcomePageURIRequest) {
	request = &SetWelcomePageURIRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-05", "SetWelcomePageURI", "itaas", "openAPI")
	return
}

// CreateSetWelcomePageURIResponse creates a response to parse from SetWelcomePageURI response
func CreateSetWelcomePageURIResponse() (response *SetWelcomePageURIResponse) {
	response = &SetWelcomePageURIResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
