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

// DescribeRtcPeakUserCntData invokes the rtc.DescribeRtcPeakUserCntData API synchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakusercntdata.html
func (client *Client) DescribeRtcPeakUserCntData(request *DescribeRtcPeakUserCntDataRequest) (response *DescribeRtcPeakUserCntDataResponse, err error) {
	response = CreateDescribeRtcPeakUserCntDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcPeakUserCntDataWithChan invokes the rtc.DescribeRtcPeakUserCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakusercntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcPeakUserCntDataWithChan(request *DescribeRtcPeakUserCntDataRequest) (<-chan *DescribeRtcPeakUserCntDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcPeakUserCntDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcPeakUserCntData(request)
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

// DescribeRtcPeakUserCntDataWithCallback invokes the rtc.DescribeRtcPeakUserCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakusercntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcPeakUserCntDataWithCallback(request *DescribeRtcPeakUserCntDataRequest, callback func(response *DescribeRtcPeakUserCntDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcPeakUserCntDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcPeakUserCntData(request)
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

// DescribeRtcPeakUserCntDataRequest is the request struct for api DescribeRtcPeakUserCntData
type DescribeRtcPeakUserCntDataRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	Interval    string           `position:"Query" name:"Interval"`
}

// DescribeRtcPeakUserCntDataResponse is the response struct for api DescribeRtcPeakUserCntData
type DescribeRtcPeakUserCntDataResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	PeakUserCntDataPerInterval PeakUserCntDataPerInterval `json:"PeakUserCntDataPerInterval" xml:"PeakUserCntDataPerInterval"`
}

// CreateDescribeRtcPeakUserCntDataRequest creates a request to invoke DescribeRtcPeakUserCntData API
func CreateDescribeRtcPeakUserCntDataRequest() (request *DescribeRtcPeakUserCntDataRequest) {
	request = &DescribeRtcPeakUserCntDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcPeakUserCntData", "rtc", "openAPI")
	return
}

// CreateDescribeRtcPeakUserCntDataResponse creates a response to parse from DescribeRtcPeakUserCntData response
func CreateDescribeRtcPeakUserCntDataResponse() (response *DescribeRtcPeakUserCntDataResponse) {
	response = &DescribeRtcPeakUserCntDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
