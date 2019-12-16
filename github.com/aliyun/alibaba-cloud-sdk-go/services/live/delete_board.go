package live

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

// DeleteBoard invokes the live.DeleteBoard API synchronously
// api document: https://help.aliyun.com/api/live/deleteboard.html
func (client *Client) DeleteBoard(request *DeleteBoardRequest) (response *DeleteBoardResponse, err error) {
	response = CreateDeleteBoardResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBoardWithChan invokes the live.DeleteBoard API asynchronously
// api document: https://help.aliyun.com/api/live/deleteboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBoardWithChan(request *DeleteBoardRequest) (<-chan *DeleteBoardResponse, <-chan error) {
	responseChan := make(chan *DeleteBoardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBoard(request)
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

// DeleteBoardWithCallback invokes the live.DeleteBoard API asynchronously
// api document: https://help.aliyun.com/api/live/deleteboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBoardWithCallback(request *DeleteBoardRequest, callback func(response *DeleteBoardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBoardResponse
		var err error
		defer close(result)
		response, err = client.DeleteBoard(request)
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

// DeleteBoardRequest is the request struct for api DeleteBoard
type DeleteBoardRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	BoardId string           `position:"Query" name:"BoardId"`
}

// DeleteBoardResponse is the response struct for api DeleteBoard
type DeleteBoardResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBoardRequest creates a request to invoke DeleteBoard API
func CreateDeleteBoardRequest() (request *DeleteBoardRequest) {
	request = &DeleteBoardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteBoard", "live", "openAPI")
	return
}

// CreateDeleteBoardResponse creates a response to parse from DeleteBoard response
func CreateDeleteBoardResponse() (response *DeleteBoardResponse) {
	response = &DeleteBoardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
