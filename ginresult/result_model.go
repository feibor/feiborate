package ginresult

import (
	"github.com/feibor/feiborate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// HTTPStatusSuccess -- HTTP编码
	HTTPStatusSuccess = 200
)

// NewResult 创建pageResult
func NewResult(ctx *gin.Context) *Result {
	return &Result{
		HTTPStatusCode: HTTPStatusSuccess,
		ctx:            ctx,
	}
}

// Result 结果集
type Result struct {
	Code           string       `json:"code"`
	Msg            string       `json:"msg"`
	Data           interface{}  `json:"data"` // Dada  数据接口
	HTTPStatusCode int          `json:"-"`    // HTTP status编码
	ctx            *gin.Context // 上下文
}

// Validate 校验
func (r *Result) Validate() (err error) {
	if r.ctx == nil {
		err = feiborate.NewCommonRespError("While init result ,ctx cannot be nil ")
		return
	}
	return
}

// InvokeResult 执行result call back
func (r *Result) InvokeResult(fc func(result *Result) error) (err error) {
	if err = r.Validate(); err != nil {
		r.ExecWithNoData(err)
		return
	}
	if err = fc(r); err != nil {
		r.ExecWithNoData(err)
		return
	}
	return
}

// InvokeResultResp 执行result call back，且可以直接使用resp
func (r *Result) InvokeResultResp(fc func(result *Result) (resp interface{}, err error)) (err error) {
	if err = r.Validate(); err != nil {
		r.ExecWithNoData(err)
		return
	}
	resp, err := fc(r)
	if err != nil {
		r.ExecWithNoData(err)
		return
	}
	r.Exec(nil, resp)
	return
}

// InvokeResultWithNodata 执行result call back
func (r *Result) InvokeResultWithNodata(fc func(result *Result) error) (err error) {
	err = r.InvokeResult(fc)
	if err == nil {
		r.ExecWithNoData(nil)
		return
	}
	return
}

// Success 默认成功返回
func (r *Result) Success() {
	r.Msg = "Success!"
	r.Code = feiborate.Success
	r.HTTPStatusCode = HTTPStatusSuccess
}

// Fail 返回失败
func (r *Result) Fail() {
	// Fail 方法主要提供返回错误的json数据
	r.Msg = "Something get error."
	r.Code = feiborate.Fail
}

// FailMessage 方法主要提供返回错误的json数据
func (r *Result) FailMessage(msg string) {
	r.Msg = msg
	r.Code = feiborate.Fail
}

// FailErr 携带error信息,如果是respError，则
// 必然存在errorCode和msg，因此进行赋值。否则不赋值
func (r *Result) FailErr(err error) {
	switch vtype := err.(type) {
	case *feiborate.RespError:
		r.Msg = vtype.ErrorMsg
		r.Code = vtype.ErrorCode
		r.HTTPStatusCode = vtype.HTTPStatusCode
	default:
		r.Code = feiborate.Fail
		r.Msg = err.Error()
		r.HTTPStatusCode = HTTPStatusSuccess
		logrus.Error(err)
	}
}

// SuccessData 即 JSON中的data数据需要加入相关构造函数
func (r *Result) SuccessData(obj interface{}) {
	r.Success()
	r.Data = obj
}

// FailData 带有Data的错误信息
func (r *Result) FailData(obj interface{}) {
	r.Fail()
	r.Data = obj
	r.HTTPStatusCode = HTTPStatusSuccess
}

// ExecWithNoData 无需数据传输时的语法
func (r *Result) ExecWithNoData(err error) {
	r.Exec(err, nil)
}

// Exec 直接执行c.context的语法，不需要再单独执行
func (r *Result) Exec(err error, resObj interface{}) {
	if err != nil {
		logrus.Error(err)
		r.FailErr(err)
		r.ctx.JSON(r.HTTPStatusCode, r)
		return
	}
	r.SuccessData(resObj)
	r.ctx.JSON(HTTPStatusSuccess, r)
	if err != nil {
		logrus.Error(err)
	}
}

// AbortErr 终止ctx的方法
func (r *Result) AbortErr(err error) {
	logrus.Error(err)
	r.FailErr(err)
	r.ctx.AbortWithStatusJSON(r.HTTPStatusCode, &r)
}
