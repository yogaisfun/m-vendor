package rtc

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

// DescribeRtcUserCntData invokes the rtc.DescribeRtcUserCntData API synchronously
// api document: https://help.aliyun.com/api/rtc/describertcusercntdata.html
func (client *Client) DescribeRtcUserCntData(request *DescribeRtcUserCntDataRequest) (response *DescribeRtcUserCntDataResponse, err error) {
	response = CreateDescribeRtcUserCntDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcUserCntDataWithChan invokes the rtc.DescribeRtcUserCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcusercntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcUserCntDataWithChan(request *DescribeRtcUserCntDataRequest) (<-chan *DescribeRtcUserCntDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcUserCntDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcUserCntData(request)
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

// DescribeRtcUserCntDataWithCallback invokes the rtc.DescribeRtcUserCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcusercntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcUserCntDataWithCallback(request *DescribeRtcUserCntDataRequest, callback func(response *DescribeRtcUserCntDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcUserCntDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcUserCntData(request)
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

// DescribeRtcUserCntDataRequest is the request struct for api DescribeRtcUserCntData
type DescribeRtcUserCntDataRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	Interval    string           `position:"Query" name:"Interval"`
}

// DescribeRtcUserCntDataResponse is the response struct for api DescribeRtcUserCntData
type DescribeRtcUserCntDataResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	UserCntDataPerInterval UserCntDataPerInterval `json:"UserCntDataPerInterval" xml:"UserCntDataPerInterval"`
}

// CreateDescribeRtcUserCntDataRequest creates a request to invoke DescribeRtcUserCntData API
func CreateDescribeRtcUserCntDataRequest() (request *DescribeRtcUserCntDataRequest) {
	request = &DescribeRtcUserCntDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcUserCntData", "rtc", "openAPI")
	return
}

// CreateDescribeRtcUserCntDataResponse creates a response to parse from DescribeRtcUserCntData response
func CreateDescribeRtcUserCntDataResponse() (response *DescribeRtcUserCntDataResponse) {
	response = &DescribeRtcUserCntDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
