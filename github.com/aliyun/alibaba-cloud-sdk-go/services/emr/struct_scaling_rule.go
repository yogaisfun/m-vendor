package emr

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

// ScalingRule is a nested struct in emr response
type ScalingRule struct {
	Id                   string            `json:"Id" xml:"Id"`
	RuleCategory         string            `json:"RuleCategory" xml:"RuleCategory"`
	RuleName             string            `json:"RuleName" xml:"RuleName"`
	AdjustmentType       string            `json:"AdjustmentType" xml:"AdjustmentType"`
	AdjustmentValue      int               `json:"AdjustmentValue" xml:"AdjustmentValue"`
	Cooldown             int               `json:"Cooldown" xml:"Cooldown"`
	Status               string            `json:"Status" xml:"Status"`
	LaunchTime           string            `json:"LaunchTime" xml:"LaunchTime"`
	LaunchExpirationTime int               `json:"LaunchExpirationTime" xml:"LaunchExpirationTime"`
	RecurrenceType       string            `json:"RecurrenceType" xml:"RecurrenceType"`
	RecurrenceValue      string            `json:"RecurrenceValue" xml:"RecurrenceValue"`
	RecurrenceEndTime    string            `json:"RecurrenceEndTime" xml:"RecurrenceEndTime"`
	SchedulerTrigger     SchedulerTrigger  `json:"SchedulerTrigger" xml:"SchedulerTrigger"`
	CloudWatchTrigger    CloudWatchTrigger `json:"CloudWatchTrigger" xml:"CloudWatchTrigger"`
}
