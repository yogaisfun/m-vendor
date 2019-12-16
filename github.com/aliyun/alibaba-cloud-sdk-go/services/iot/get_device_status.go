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

// GetDeviceStatus invokes the iot.GetDeviceStatus API synchronously
// api document: https://help.aliyun.com/api/iot/getdevicestatus.html
func (client *Client) GetDeviceStatus(request *GetDeviceStatusRequest) (response *GetDeviceStatusResponse, err error) {
	response = CreateGetDeviceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceStatusWithChan invokes the iot.GetDeviceStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/getdevicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceStatusWithChan(request *GetDeviceStatusRequest) (<-chan *GetDeviceStatusResponse, <-chan error) {
	responseChan := make(chan *GetDeviceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceStatus(request)
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

// GetDeviceStatusWithCallback invokes the iot.GetDeviceStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/getdevicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceStatusWithCallback(request *GetDeviceStatusRequest, callback func(response *GetDeviceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceStatusResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceStatus(request)
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

// GetDeviceStatusRequest is the request struct for api GetDeviceStatus
type GetDeviceStatusRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
}

// GetDeviceStatusResponse is the response struct for api GetDeviceStatus
type GetDeviceStatusResponse struct {
	*responses.BaseResponse
	RequestId    string                `json:"RequestId" xml:"RequestId"`
	Success      bool                  `json:"Success" xml:"Success"`
	Code         string                `json:"Code" xml:"Code"`
	ErrorMessage string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetDeviceStatus `json:"Data" xml:"Data"`
}

// CreateGetDeviceStatusRequest creates a request to invoke GetDeviceStatus API
func CreateGetDeviceStatusRequest() (request *GetDeviceStatusRequest) {
	request = &GetDeviceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetDeviceStatus", "iot", "openAPI")
	return
}

// CreateGetDeviceStatusResponse creates a response to parse from GetDeviceStatus response
func CreateGetDeviceStatusResponse() (response *GetDeviceStatusResponse) {
	response = &GetDeviceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
