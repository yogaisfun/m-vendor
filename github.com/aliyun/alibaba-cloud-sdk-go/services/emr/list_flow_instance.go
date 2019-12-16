package emr

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

// ListFlowInstance invokes the emr.ListFlowInstance API synchronously
// api document: https://help.aliyun.com/api/emr/listflowinstance.html
func (client *Client) ListFlowInstance(request *ListFlowInstanceRequest) (response *ListFlowInstanceResponse, err error) {
	response = CreateListFlowInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowInstanceWithChan invokes the emr.ListFlowInstance API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowInstanceWithChan(request *ListFlowInstanceRequest) (<-chan *ListFlowInstanceResponse, <-chan error) {
	responseChan := make(chan *ListFlowInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowInstance(request)
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

// ListFlowInstanceWithCallback invokes the emr.ListFlowInstance API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowInstanceWithCallback(request *ListFlowInstanceRequest, callback func(response *ListFlowInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowInstanceResponse
		var err error
		defer close(result)
		response, err = client.ListFlowInstance(request)
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

// ListFlowInstanceRequest is the request struct for api ListFlowInstance
type ListFlowInstanceRequest struct {
	*requests.RpcRequest
	Owner      string           `position:"Query" name:"Owner"`
	TimeRange  string           `position:"Query" name:"TimeRange"`
	StatusList *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	FlowName   string           `position:"Query" name:"FlowName"`
	Id         string           `position:"Query" name:"Id"`
	FlowId     string           `position:"Query" name:"FlowId"`
	ProjectId  string           `position:"Query" name:"ProjectId"`
	OrderType  string           `position:"Query" name:"OrderType"`
}

// ListFlowInstanceResponse is the response struct for api ListFlowInstance
type ListFlowInstanceResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	Total         int           `json:"Total" xml:"Total"`
	FlowInstances FlowInstances `json:"FlowInstances" xml:"FlowInstances"`
}

// CreateListFlowInstanceRequest creates a request to invoke ListFlowInstance API
func CreateListFlowInstanceRequest() (request *ListFlowInstanceRequest) {
	request = &ListFlowInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowInstance", "emr", "openAPI")
	return
}

// CreateListFlowInstanceResponse creates a response to parse from ListFlowInstance response
func CreateListFlowInstanceResponse() (response *ListFlowInstanceResponse) {
	response = &ListFlowInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
