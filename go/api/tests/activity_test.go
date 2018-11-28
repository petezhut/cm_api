package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/petezhut/cm_api/go/api"
	"testing"
)

var (
   ErrorString = "%s - Should be '%s', but instead it was: '%s'"
)

/*
TYPE: Activity
Represents a user activity, such as a MapReduce job, a Hive query, an Oozie workflow, etc.
*/
//type Activity struct {
//	Name              string `json:"name"`              // Activity name.
//	Type              string `json:"type"`              // Activity type. Whether it's an MR job, a Pig job, a Hive query, etc.
//	Parent            string `json:"parent"`            // The name of the parent activity.
//	StartTime         string `json:"startTime"`         // The start time of this activity.
//	FinishTime        string `json:"finishTime"`        // The finish time of this activity.
//	Id                string `json:"id"`                // Activity id, which is unique within a MapReduce service.
//	Status            string `json:"status"`            // Activity status.
//	User              string `json:"user"`              // The user who submitted this activity.
//	Group             string `json:"group"`             // The user-group of this activity.
//	InputDir          string `json:"inputDir"`          // The input data directory of the activity. An HDFS url.
//	OutputDir         string `json:"outputDir"`         // The output result directory of the activity. An HDFS url.
//	Mapper            string `json:"mapper"`            // The mapper class.
//	Combiner          string `json:"combiner"`          // The combiner class.
//	Reducer           string `json:"reducer"`           // The reducer class.
//	QueueName         string `json:"queueName"`         // The scheduler queue this activity is in.
//	SchedulerPriority string `json:"schedulerPriority"` // The scheduler priority of this activity.
//}

var (
	ApiActivityPassingJSON = []byte(`{
					"name": "TestActivity",
				    "type": {"value": "HiveQuery"},
					"parent": "MockedParent",
					"startTime": "0000",
					"finishTime": "1000",
					"id": "0001",
					"status": {"value": PASSED"},
					"user": "testuser",
					"group": "testgroup",
					"inputDir": "/users/tests",
					"outputDir": "/users/results",
					"mapper": "TestMapperClass",
					"combiner": "TestCombinerClass",
					"reducer": "TestReducerClass",
					"queueName": "TestQueueName",
					"schedulerPriority": "TestSchedulerPriority"}`)
	ApiActivityErrorJSON = []byte(`{"name": "TestActivity"}`)
)

func DoMockedCall(data []byte) []byte {
	return data
}

func TestApiActivityPassing(t *testing.T) {
	var API api.Activity
	json.Unmarshal(DoMockedCall(ApiActivityPassingJSON), &API)
	tables := []struct {
		name string
		found string
		expected string
	}{
		{"Name", API.Name, "TestActivity"},
		{"Type", API.Type.Value, "HiveQuery"},
		{"Parent", API.Parent, "MockedParent2"},
	}
	for _, table := range tables {
		assert.Equal(t, table.found, table.expected, fmt.Sprintf(ErrorString, table.name, table.expected, table.found))
	}
}
func TestApiActivityError(t *testing.T) {
	var API api.Activity
	json.Unmarshal(DoMockedCall(ApiActivityErrorJSON), &API)
	assert.Equal(t, "0000", API.StartTime, fmt.Sprintf(ErrorString, "StartTime", "0000", API.StartTime))
}
