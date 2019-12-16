package iot

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

// DeleteDeviceFile invokes the iot.DeleteDeviceFile API synchronously
// api document: https://help.aliyun.com/api/iot/deletedevicefile.html
func (client *Client) DeleteDeviceFile(request *DeleteDeviceFileRequest) (response *DeleteDeviceFileResponse, err error) {
	response = CreateDeleteDeviceFileResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeviceFileWithChan invokes the iot.DeleteDeviceFile API asynchronously
// api document: https://help.aliyun.com/api/iot/deletedevicefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceFileWithChan(request *DeleteDeviceFileRequest) (<-chan *DeleteDeviceFileResponse, <-chan error) {
	responseChan := make(chan *DeleteDeviceFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDeviceFile(request)
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

// DeleteDeviceFileWithCallback invokes the iot.DeleteDeviceFile API asynchronously
// api document: https://help.aliyun.com/api/iot/deletedevicefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDeviceFileWithCallback(request *DeleteDeviceFileRequest, callback func(response *DeleteDeviceFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeviceFileResponse
		var err error
		defer close(result)
		response, err = client.DeleteDeviceFile(request)
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

// DeleteDeviceFileRequest is the request struct for api DeleteDeviceFile
type DeleteDeviceFileRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	FileId        string `position:"Query" name:"FileId"`
}

// DeleteDeviceFileResponse is the response struct for api DeleteDeviceFile
type DeleteDeviceFileResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteDeviceFileRequest creates a request to invoke DeleteDeviceFile API
func CreateDeleteDeviceFileRequest() (request *DeleteDeviceFileRequest) {
	request = &DeleteDeviceFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteDeviceFile", "iot", "openAPI")
	return
}

// CreateDeleteDeviceFileResponse creates a response to parse from DeleteDeviceFile response
func CreateDeleteDeviceFileResponse() (response *DeleteDeviceFileResponse) {
	response = &DeleteDeviceFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
