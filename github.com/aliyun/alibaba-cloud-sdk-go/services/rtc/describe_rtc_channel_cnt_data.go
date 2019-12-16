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

// DescribeRtcChannelCntData invokes the rtc.DescribeRtcChannelCntData API synchronously
// api document: https://help.aliyun.com/api/rtc/describertcchannelcntdata.html
func (client *Client) DescribeRtcChannelCntData(request *DescribeRtcChannelCntDataRequest) (response *DescribeRtcChannelCntDataResponse, err error) {
	response = CreateDescribeRtcChannelCntDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcChannelCntDataWithChan invokes the rtc.DescribeRtcChannelCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcchannelcntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcChannelCntDataWithChan(request *DescribeRtcChannelCntDataRequest) (<-chan *DescribeRtcChannelCntDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcChannelCntDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcChannelCntData(request)
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

// DescribeRtcChannelCntDataWithCallback invokes the rtc.DescribeRtcChannelCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcchannelcntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcChannelCntDataWithCallback(request *DescribeRtcChannelCntDataRequest, callback func(response *DescribeRtcChannelCntDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcChannelCntDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcChannelCntData(request)
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

// DescribeRtcChannelCntDataRequest is the request struct for api DescribeRtcChannelCntData
type DescribeRtcChannelCntDataRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	Interval    string           `position:"Query" name:"Interval"`
}

// DescribeRtcChannelCntDataResponse is the response struct for api DescribeRtcChannelCntData
type DescribeRtcChannelCntDataResponse struct {
	*responses.BaseResponse
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	ChannelCntDataPerInterval ChannelCntDataPerInterval `json:"ChannelCntDataPerInterval" xml:"ChannelCntDataPerInterval"`
}

// CreateDescribeRtcChannelCntDataRequest creates a request to invoke DescribeRtcChannelCntData API
func CreateDescribeRtcChannelCntDataRequest() (request *DescribeRtcChannelCntDataRequest) {
	request = &DescribeRtcChannelCntDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcChannelCntData", "rtc", "openAPI")
	return
}

// CreateDescribeRtcChannelCntDataResponse creates a response to parse from DescribeRtcChannelCntData response
func CreateDescribeRtcChannelCntDataResponse() (response *DescribeRtcChannelCntDataResponse) {
	response = &DescribeRtcChannelCntDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
