package kms

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

// ListSecretVersionIds invokes the kms.ListSecretVersionIds API synchronously
// api document: https://help.aliyun.com/api/kms/listsecretversionids.html
func (client *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (response *ListSecretVersionIdsResponse, err error) {
	response = CreateListSecretVersionIdsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSecretVersionIdsWithChan invokes the kms.ListSecretVersionIds API asynchronously
// api document: https://help.aliyun.com/api/kms/listsecretversionids.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSecretVersionIdsWithChan(request *ListSecretVersionIdsRequest) (<-chan *ListSecretVersionIdsResponse, <-chan error) {
	responseChan := make(chan *ListSecretVersionIdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSecretVersionIds(request)
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

// ListSecretVersionIdsWithCallback invokes the kms.ListSecretVersionIds API asynchronously
// api document: https://help.aliyun.com/api/kms/listsecretversionids.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSecretVersionIdsWithCallback(request *ListSecretVersionIdsRequest, callback func(response *ListSecretVersionIdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSecretVersionIdsResponse
		var err error
		defer close(result)
		response, err = client.ListSecretVersionIds(request)
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

// ListSecretVersionIdsRequest is the request struct for api ListSecretVersionIds
type ListSecretVersionIdsRequest struct {
	*requests.RpcRequest
	IncludeDeprecated string           `position:"Query" name:"IncludeDeprecated"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	SecretName        string           `position:"Query" name:"SecretName"`
	PageNumber        requests.Integer `position:"Query" name:"PageNumber"`
}

// ListSecretVersionIdsResponse is the response struct for api ListSecretVersionIds
type ListSecretVersionIdsResponse struct {
	*responses.BaseResponse
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	SecretName string     `json:"SecretName" xml:"SecretName"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	VersionIds VersionIds `json:"VersionIds" xml:"VersionIds"`
}

// CreateListSecretVersionIdsRequest creates a request to invoke ListSecretVersionIds API
func CreateListSecretVersionIdsRequest() (request *ListSecretVersionIdsRequest) {
	request = &ListSecretVersionIdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListSecretVersionIds", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSecretVersionIdsResponse creates a response to parse from ListSecretVersionIds response
func CreateListSecretVersionIdsResponse() (response *ListSecretVersionIdsResponse) {
	response = &ListSecretVersionIdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
