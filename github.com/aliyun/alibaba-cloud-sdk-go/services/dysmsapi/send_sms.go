package dysmsapi

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

// SendSms invokes the dysmsapi.SendSms API synchronously
// api document: https://help.aliyun.com/api/dysmsapi/sendsms.html
func (client *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
	response = CreateSendSmsResponse()
	err = client.DoAction(request, response)
	return
}

// SendSmsWithChan invokes the dysmsapi.SendSms API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/sendsms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendSmsWithChan(request *SendSmsRequest) (<-chan *SendSmsResponse, <-chan error) {
	responseChan := make(chan *SendSmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendSms(request)
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

// SendSmsWithCallback invokes the dysmsapi.SendSms API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/sendsms.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendSmsWithCallback(request *SendSmsRequest, callback func(response *SendSmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendSmsResponse
		var err error
		defer close(result)
		response, err = client.SendSms(request)
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

// SendSmsRequest is the request struct for api SendSms
type SendSmsRequest struct {
	*requests.RpcRequest
	SmsUpExtendCode      string           `position:"Query" name:"SmsUpExtendCode"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SignName             string           `position:"Query" name:"SignName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PhoneNumbers         string           `position:"Query" name:"PhoneNumbers"`
	OutId                string           `position:"Query" name:"OutId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateCode         string           `position:"Query" name:"TemplateCode"`
	TemplateParam        string           `position:"Query" name:"TemplateParam"`
}

// SendSmsResponse is the response struct for api SendSms
type SendSmsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BizId     string `json:"BizId" xml:"BizId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSendSmsRequest creates a request to invoke SendSms API
func CreateSendSmsRequest() (request *SendSmsRequest) {
	request = &SendSmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "SendSms", "dysmsapi", "openAPI")
	return
}

// CreateSendSmsResponse creates a response to parse from SendSms response
func CreateSendSmsResponse() (response *SendSmsResponse) {
	response = &SendSmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
