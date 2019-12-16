package dbs

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

// ModifyBackupSourceEndpoint invokes the dbs.ModifyBackupSourceEndpoint API synchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsourceendpoint.html
func (client *Client) ModifyBackupSourceEndpoint(request *ModifyBackupSourceEndpointRequest) (response *ModifyBackupSourceEndpointResponse, err error) {
	response = CreateModifyBackupSourceEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupSourceEndpointWithChan invokes the dbs.ModifyBackupSourceEndpoint API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsourceendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupSourceEndpointWithChan(request *ModifyBackupSourceEndpointRequest) (<-chan *ModifyBackupSourceEndpointResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupSourceEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupSourceEndpoint(request)
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

// ModifyBackupSourceEndpointWithCallback invokes the dbs.ModifyBackupSourceEndpoint API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsourceendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupSourceEndpointWithCallback(request *ModifyBackupSourceEndpointRequest, callback func(response *ModifyBackupSourceEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupSourceEndpointResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupSourceEndpoint(request)
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

// ModifyBackupSourceEndpointRequest is the request struct for api ModifyBackupSourceEndpoint
type ModifyBackupSourceEndpointRequest struct {
	*requests.RpcRequest
	SourceEndpointRegion       string           `position:"Query" name:"SourceEndpointRegion"`
	BackupGatewayId            requests.Integer `position:"Query" name:"BackupGatewayId"`
	SourceEndpointInstanceID   string           `position:"Query" name:"SourceEndpointInstanceID"`
	SourceEndpointUserName     string           `position:"Query" name:"SourceEndpointUserName"`
	ClientToken                string           `position:"Query" name:"ClientToken"`
	SourceEndpointPassword     string           `position:"Query" name:"SourceEndpointPassword"`
	BackupPlanId               string           `position:"Query" name:"BackupPlanId"`
	BackupObjects              string           `position:"Query" name:"BackupObjects"`
	OwnerId                    string           `position:"Query" name:"OwnerId"`
	SourceEndpointPort         requests.Integer `position:"Query" name:"SourceEndpointPort"`
	SourceEndpointDatabaseName string           `position:"Query" name:"SourceEndpointDatabaseName"`
	SourceEndpointInstanceType string           `position:"Query" name:"SourceEndpointInstanceType"`
	SourceEndpointIP           string           `position:"Query" name:"SourceEndpointIP"`
	SourceEndpointOracleSID    string           `position:"Query" name:"SourceEndpointOracleSID"`
}

// ModifyBackupSourceEndpointResponse is the response struct for api ModifyBackupSourceEndpoint
type ModifyBackupSourceEndpointResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	NeedPrecheck   bool   `json:"NeedPrecheck" xml:"NeedPrecheck"`
}

// CreateModifyBackupSourceEndpointRequest creates a request to invoke ModifyBackupSourceEndpoint API
func CreateModifyBackupSourceEndpointRequest() (request *ModifyBackupSourceEndpointRequest) {
	request = &ModifyBackupSourceEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyBackupSourceEndpoint", "cbs", "openAPI")
	return
}

// CreateModifyBackupSourceEndpointResponse creates a response to parse from ModifyBackupSourceEndpoint response
func CreateModifyBackupSourceEndpointResponse() (response *ModifyBackupSourceEndpointResponse) {
	response = &ModifyBackupSourceEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
