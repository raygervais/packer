package ram

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

// UpdateRole invokes the ram.UpdateRole API synchronously
// api document: https://help.aliyun.com/api/ram/updaterole.html
func (client *Client) UpdateRole(request *UpdateRoleRequest) (response *UpdateRoleResponse, err error) {
	response = CreateUpdateRoleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRoleWithChan invokes the ram.UpdateRole API asynchronously
// api document: https://help.aliyun.com/api/ram/updaterole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRoleWithChan(request *UpdateRoleRequest) (<-chan *UpdateRoleResponse, <-chan error) {
	responseChan := make(chan *UpdateRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRole(request)
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

// UpdateRoleWithCallback invokes the ram.UpdateRole API asynchronously
// api document: https://help.aliyun.com/api/ram/updaterole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRoleWithCallback(request *UpdateRoleRequest, callback func(response *UpdateRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRoleResponse
		var err error
		defer close(result)
		response, err = client.UpdateRole(request)
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

// UpdateRoleRequest is the request struct for api UpdateRole
type UpdateRoleRequest struct {
	*requests.RpcRequest
	NewAssumeRolePolicyDocument string `position:"Query" name:"NewAssumeRolePolicyDocument"`
	RoleName                    string `position:"Query" name:"RoleName"`
}

// UpdateRoleResponse is the response struct for api UpdateRole
type UpdateRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Role      Role   `json:"Role" xml:"Role"`
}

// CreateUpdateRoleRequest creates a request to invoke UpdateRole API
func CreateUpdateRoleRequest() (request *UpdateRoleRequest) {
	request = &UpdateRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "UpdateRole", "ram", "openAPI")
	return
}

// CreateUpdateRoleResponse creates a response to parse from UpdateRole response
func CreateUpdateRoleResponse() (response *UpdateRoleResponse) {
	response = &UpdateRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
