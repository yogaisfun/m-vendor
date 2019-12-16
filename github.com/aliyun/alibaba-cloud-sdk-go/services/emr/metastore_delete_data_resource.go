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

// MetastoreDeleteDataResource invokes the emr.MetastoreDeleteDataResource API synchronously
// api document: https://help.aliyun.com/api/emr/metastoredeletedataresource.html
func (client *Client) MetastoreDeleteDataResource(request *MetastoreDeleteDataResourceRequest) (response *MetastoreDeleteDataResourceResponse, err error) {
	response = CreateMetastoreDeleteDataResourceResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreDeleteDataResourceWithChan invokes the emr.MetastoreDeleteDataResource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredeletedataresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDeleteDataResourceWithChan(request *MetastoreDeleteDataResourceRequest) (<-chan *MetastoreDeleteDataResourceResponse, <-chan error) {
	responseChan := make(chan *MetastoreDeleteDataResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreDeleteDataResource(request)
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

// MetastoreDeleteDataResourceWithCallback invokes the emr.MetastoreDeleteDataResource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredeletedataresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDeleteDataResourceWithCallback(request *MetastoreDeleteDataResourceRequest, callback func(response *MetastoreDeleteDataResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreDeleteDataResourceResponse
		var err error
		defer close(result)
		response, err = client.MetastoreDeleteDataResource(request)
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

// MetastoreDeleteDataResourceRequest is the request struct for api MetastoreDeleteDataResource
type MetastoreDeleteDataResourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Id              string           `position:"Query" name:"Id"`
}

// MetastoreDeleteDataResourceResponse is the response struct for api MetastoreDeleteDataResource
type MetastoreDeleteDataResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMetastoreDeleteDataResourceRequest creates a request to invoke MetastoreDeleteDataResource API
func CreateMetastoreDeleteDataResourceRequest() (request *MetastoreDeleteDataResourceRequest) {
	request = &MetastoreDeleteDataResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDeleteDataResource", "emr", "openAPI")
	return
}

// CreateMetastoreDeleteDataResourceResponse creates a response to parse from MetastoreDeleteDataResource response
func CreateMetastoreDeleteDataResourceResponse() (response *MetastoreDeleteDataResourceResponse) {
	response = &MetastoreDeleteDataResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
