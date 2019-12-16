package edas

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

// DeleteK8sApplication invokes the edas.DeleteK8sApplication API synchronously
// api document: https://help.aliyun.com/api/edas/deletek8sapplication.html
func (client *Client) DeleteK8sApplication(request *DeleteK8sApplicationRequest) (response *DeleteK8sApplicationResponse, err error) {
	response = CreateDeleteK8sApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteK8sApplicationWithChan invokes the edas.DeleteK8sApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/deletek8sapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteK8sApplicationWithChan(request *DeleteK8sApplicationRequest) (<-chan *DeleteK8sApplicationResponse, <-chan error) {
	responseChan := make(chan *DeleteK8sApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteK8sApplication(request)
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

// DeleteK8sApplicationWithCallback invokes the edas.DeleteK8sApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/deletek8sapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteK8sApplicationWithCallback(request *DeleteK8sApplicationRequest, callback func(response *DeleteK8sApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteK8sApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeleteK8sApplication(request)
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

// DeleteK8sApplicationRequest is the request struct for api DeleteK8sApplication
type DeleteK8sApplicationRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// DeleteK8sApplicationResponse is the response struct for api DeleteK8sApplication
type DeleteK8sApplicationResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
}

// CreateDeleteK8sApplicationRequest creates a request to invoke DeleteK8sApplication API
func CreateDeleteK8sApplicationRequest() (request *DeleteK8sApplicationRequest) {
	request = &DeleteK8sApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeleteK8sApplication", "/pop/v5/k8s/acs/k8s_apps", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteK8sApplicationResponse creates a response to parse from DeleteK8sApplication response
func CreateDeleteK8sApplicationResponse() (response *DeleteK8sApplicationResponse) {
	response = &DeleteK8sApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
