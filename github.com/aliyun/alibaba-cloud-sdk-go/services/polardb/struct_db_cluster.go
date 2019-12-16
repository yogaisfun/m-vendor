package polardb

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

// DBCluster is a nested struct in polardb response
type DBCluster struct {
	DBClusterId          string                      `json:"DBClusterId" xml:"DBClusterId"`
	DBClusterDescription string                      `json:"DBClusterDescription" xml:"DBClusterDescription"`
	PayType              string                      `json:"PayType" xml:"PayType"`
	DBClusterNetworkType string                      `json:"DBClusterNetworkType" xml:"DBClusterNetworkType"`
	RegionId             string                      `json:"RegionId" xml:"RegionId"`
	ExpireTime           string                      `json:"ExpireTime" xml:"ExpireTime"`
	Expired              string                      `json:"Expired" xml:"Expired"`
	DBClusterStatus      string                      `json:"DBClusterStatus" xml:"DBClusterStatus"`
	Engine               string                      `json:"Engine" xml:"Engine"`
	DBType               string                      `json:"DBType" xml:"DBType"`
	DBVersion            string                      `json:"DBVersion" xml:"DBVersion"`
	LockMode             string                      `json:"LockMode" xml:"LockMode"`
	CreateTime           string                      `json:"CreateTime" xml:"CreateTime"`
	VpcId                string                      `json:"VpcId" xml:"VpcId"`
	DBNodeNumber         int                         `json:"DBNodeNumber" xml:"DBNodeNumber"`
	DBNodeClass          string                      `json:"DBNodeClass" xml:"DBNodeClass"`
	StorageUsed          int                         `json:"StorageUsed" xml:"StorageUsed"`
	DBNodes              DBNodesInDescribeDBClusters `json:"DBNodes" xml:"DBNodes"`
	Tags                 TagsInDescribeDBClusters    `json:"Tags" xml:"Tags"`
}
