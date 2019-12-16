package arms

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

// MetricQuery invokes the arms.MetricQuery API synchronously
// api document: https://help.aliyun.com/api/arms/metricquery.html
func (client *Client) MetricQuery(request *MetricQueryRequest) (response *MetricQueryResponse, err error) {
	response = CreateMetricQueryResponse()
	err = client.DoAction(request, response)
	return
}

// MetricQueryWithChan invokes the arms.MetricQuery API asynchronously
// api document: https://help.aliyun.com/api/arms/metricquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetricQueryWithChan(request *MetricQueryRequest) (<-chan *MetricQueryResponse, <-chan error) {
	responseChan := make(chan *MetricQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetricQuery(request)
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

// MetricQueryWithCallback invokes the arms.MetricQuery API asynchronously
// api document: https://help.aliyun.com/api/arms/metricquery.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetricQueryWithCallback(request *MetricQueryRequest, callback func(response *MetricQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetricQueryResponse
		var err error
		defer close(result)
		response, err = client.MetricQuery(request)
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

// MetricQueryRequest is the request struct for api MetricQuery
type MetricQueryRequest struct {
	*requests.RpcRequest
	EndTime        requests.Integer      `position:"Query" name:"EndTime"`
	OrderBy        string                `position:"Query" name:"OrderBy"`
	Filters        *[]MetricQueryFilters `position:"Query" name:"Filters"  type:"Repeated"`
	StartTime      requests.Integer      `position:"Query" name:"StartTime"`
	IintervalInSec requests.Integer      `position:"Query" name:"IintervalInSec"`
	Measures       *[]string             `position:"Query" name:"Measures"  type:"Repeated"`
	Metric         string                `position:"Query" name:"Metric"`
	SecurityToken  string                `position:"Query" name:"SecurityToken"`
	Limit          requests.Integer      `position:"Query" name:"Limit"`
	Dimensions     *[]string             `position:"Query" name:"Dimensions"  type:"Repeated"`
	Order          string                `position:"Query" name:"Order"`
}

// MetricQueryFilters is a repeated param struct in MetricQueryRequest
type MetricQueryFilters struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// MetricQueryResponse is the response struct for api MetricQuery
type MetricQueryResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMetricQueryRequest creates a request to invoke MetricQuery API
func CreateMetricQueryRequest() (request *MetricQueryRequest) {
	request = &MetricQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-02-19", "MetricQuery", "", "")
	return
}

// CreateMetricQueryResponse creates a response to parse from MetricQuery response
func CreateMetricQueryResponse() (response *MetricQueryResponse) {
	response = &MetricQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
