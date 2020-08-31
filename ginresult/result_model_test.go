package ginresult

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestInvokeResult(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := NewResult(ctx)
	err := result.InvokeResult(func(result *Result) (err error) {
		resp := `{"status":"success"}`
		result.Exec(nil, resp)
		return
	})
	if err != nil {
		panic(err)
	}
}

func TestResult_InvokeResultResp(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	result := NewResult(ctx)
	err := result.InvokeResultResp(func(result *Result) (resp interface{}, err error) {
		resp = `{"status":"success"}`
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
