package gotest

import (
	"encoding/json"
	"testing"
	"timewheel/Lib"
)

func Test_json_time(t *testing.T) {

	testTimeVal := `{"time":"2020-04-05 13:32:11"}`

	type TestTime struct {
		Time Lib.Time `json:"time"`
	}

	testTime := &TestTime{}

	if err := json.Unmarshal([]byte(testTimeVal), testTime); err != nil {
		t.Log(err.Error())
		return
	}

	t.Log(testTime.Time)

}
