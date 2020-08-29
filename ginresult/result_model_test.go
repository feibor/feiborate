package ginresult

import (
	"testing"
)

func TestNewResult(t *testing.T) {
	result := NewResult()
	err := result.InvokeResult(func(result *Result) (err error) {
		resp := `{"status":"success"}`
		result.Exec(nil, resp)
		return
	})
	if err != nil {
		panic(err)
	}
}

func TestNewPageResult(t *testing.T) {
	result := NewPageResult()
	err := result.InvokeResultWithNodata(func(result *Result) (err error) {
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		panic(err)
	}
}
