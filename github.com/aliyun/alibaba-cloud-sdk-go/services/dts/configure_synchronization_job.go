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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ConfigureSynchronizationJob invokes the dts.ConfigureSynchronizationJob API synchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjob.html
func (client *Client) ConfigureSynchronizationJob(request *ConfigureSynchronizationJobRequest) (response *ConfigureSynchronizationJobResponse, err error) {
	response = CreateConfigureSynchronizationJobResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSynchronizationJobWithChan invokes the dts.ConfigureSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobWithChan(request *ConfigureSynchronizationJobRequest) (<-chan *ConfigureSynchronizationJobResponse, <-chan error) {
	responseChan := make(chan *ConfigureSynchronizationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSynchronizationJob(request)
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

// ConfigureSynchronizationJobWithCallback invokes the dts.ConfigureSynchronizationJob API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobWithCallback(request *ConfigureSynchronizationJobRequest, callback func(response *ConfigureSynchronizationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSynchronizationJobResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSynchronizationJob(request)
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

// ConfigureSynchronizationJobRequest is the request struct for api ConfigureSynchronizationJob
type ConfigureSynchronizationJobRequest struct {
	*requests.RpcRequest
	SynchronizationJobName   string                                         `position:"Query" name:"SynchronizationJobName"`
	SynchronizationJobId     string                                         `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string                                         `position:"Query" name:"SynchronizationDirection"`
	StructureInitialization  requests.Boolean                               `position:"Query" name:"StructureInitialization"`
	DataInitialization       requests.Boolean                               `position:"Query" name:"DataInitialization"`
	SynchronizationObjects   string                                         `position:"Query" name:"SynchronizationObjects"`
	MigrationReserved        string                                         `position:"Query" name:"MigrationReserved"`
	Checkpoint               string                                         `position:"Query" name:"Checkpoint"`
	OwnerId                  string                                         `position:"Query" name:"OwnerId"`
	SourceEndpoint           ConfigureSynchronizationJobSourceEndpoint      `position:"Query" name:"SourceEndpoint" type:"Struct"`
	DestinationEndpoint      ConfigureSynchronizationJobDestinationEndpoint `position:"Query" name:"DestinationEndpoint" type:"Struct"`
	PartitionKey             ConfigureSynchronizationJobPartitionKey        `position:"Query" name:"PartitionKey" type:"Struct"`
}

type ConfigureSynchronizationJobSourceEndpoint struct {
	InstanceId   string `name:"InstanceId"`
	InstanceType string `name:"InstanceType"`
	IP           string `name:"IP"`
	Port         string `name:"Port"`
	UserName     string `name:"UserName"`
	Password     string `name:"Password"`
	OwnerID      string `name:"OwnerID"`
	Role         string `name:"Role"`
}

type ConfigureSynchronizationJobDestinationEndpoint struct {
	InstanceId   string `name:"InstanceId"`
	InstanceType string `name:"InstanceType"`
	IP           string `name:"IP"`
	Port         string `name:"Port"`
	UserName     string `name:"UserName"`
	Password     string `name:"Password"`
}

type ConfigureSynchronizationJobPartitionKey struct {
	ModifyTime_Year   requests.Boolean `name:"ModifyTime_Year"`
	ModifyTime_Month  requests.Boolean `name:"ModifyTime_Month"`
	ModifyTime_Day    requests.Boolean `name:"ModifyTime_Day"`
	ModifyTime_Hour   requests.Boolean `name:"ModifyTime_Hour"`
	ModifyTime_Minute requests.Boolean `name:"ModifyTime_Minute"`
}

// ConfigureSynchronizationJobResponse is the response struct for api ConfigureSynchronizationJob
type ConfigureSynchronizationJobResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigureSynchronizationJobRequest creates a request to invoke ConfigureSynchronizationJob API
func CreateConfigureSynchronizationJobRequest() (request *ConfigureSynchronizationJobRequest) {
	request = &ConfigureSynchronizationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ConfigureSynchronizationJob", "dts", "openAPI")
	return
}

// CreateConfigureSynchronizationJobResponse creates a response to parse from ConfigureSynchronizationJob response
func CreateConfigureSynchronizationJobResponse() (response *ConfigureSynchronizationJobResponse) {
	response = &ConfigureSynchronizationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}