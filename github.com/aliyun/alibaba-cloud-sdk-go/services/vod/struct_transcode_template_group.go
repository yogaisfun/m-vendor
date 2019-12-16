package vod

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

// TranscodeTemplateGroup is a nested struct in vod response
type TranscodeTemplateGroup struct {
	Name                     string              `json:"Name" xml:"Name"`
	AppId                    string              `json:"AppId" xml:"AppId"`
	ModifyTime               string              `json:"ModifyTime" xml:"ModifyTime"`
	TranscodeMode            string              `json:"TranscodeMode" xml:"TranscodeMode"`
	CreationTime             string              `json:"CreationTime" xml:"CreationTime"`
	Locked                   string              `json:"Locked" xml:"Locked"`
	IsDefault                string              `json:"IsDefault" xml:"IsDefault"`
	TranscodeTemplateGroupId string              `json:"TranscodeTemplateGroupId" xml:"TranscodeTemplateGroupId"`
	TranscodeTemplateList    []TranscodeTemplate `json:"TranscodeTemplateList" xml:"TranscodeTemplateList"`
}
