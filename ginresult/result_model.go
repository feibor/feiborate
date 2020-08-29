package ginresult

import (
  "feibor.com/feibor"
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
)

const (
  // HTTPStatusSuccess -- HTTP编码
  HTTPStatusSuccess = 200
)

// ResultPage 分页查询结果集的对象
type ResultPage struct {
  *Result
  *Page
}

// InvokeResult Callback函数方式处理
func (r *ResultPage) InvokeResult(ctx *gin.Context, fc func(result *ResultPage) error) (err error) {
  result := r
  err = fc(result)
  if err != nil {
    result.ExecWithNoDataPage(ctx, err)
    return
  }
  return
}

// InvokeResultResp Callback函数方式处理
func (r *ResultPage) InvokeResultResp(ctx *gin.Context, fc func(result *ResultPage) error) (resp interface{}, err error) {
  err = fc(r)
  if err != nil {
    r.ExecWithNoDataPage(ctx, err)
    return
  }
  r.ExecPage(ctx, nil, resp)
  return
}

// Result 结果集
type Result struct {
  Code string `json:"code"`
  Msg  string `json:"msg"`
  // Dada  数据接口
  Data           interface{} `json:"data"`
  HTTPStatusCode int         `json:"-"`
}

// Page 通用的分页对象
type Page struct {
  PageSize int `json:"pageSize"` // 每页数据大小
  PageNum  int `json:"pageNum"`  // 当前页码
  Total    int `json:"total"`    // 数据总量
  Offset   int `json:"-"`        // 数据查询偏移量，即 pageSize*pageNum - pageSize
}

// InvokeResult 执行result call back
func (r *Result) InvokeResult(ctx *gin.Context, fc func(result *Result) error) (err error) {
  err = fc(r)
  if err != nil {
    r.ExecWithNoData(ctx, err)
    return
  }
  return
}

// InvokeResultWithNodata 执行result call back
func (r *Result) InvokeResultWithNodata(ctx *gin.Context, fc func(result *Result) error) (err error) {
  err = r.InvokeResult(ctx, fc)
  if err == nil {
    r.ExecWithNoData(ctx, nil)
    return
  }
  return
}

// NewPageResult 创建pageResult
func NewPageResult() *ResultPage {
  return &ResultPage{
    Result: &Result{},
  }
}

// NewResult 创建pageResult
func NewResult() *Result {
  return &Result{
    HTTPStatusCode: HTTPStatusSuccess,
  }
}

// NewPage 创建Page对象
func NewPage(pageSize int, pageNum int) (p *Page) {
  p = &Page{
    PageSize: pageSize,
    PageNum:  pageNum,
  }
  p.PageInit()
  return
}

// DefaultPageSize 设置默认分页大小
func (p *Page) DefaultPageSize() {
  if p.PageSize == 0 {
    // 默认显示15条数据
    p.PageSize = 15
  } else if p.PageSize > 100 {
    p.PageSize = 100
  }

  // 默认查第一页数据
  if p.PageNum <= 0 {
    p.PageNum = 1
  }
}

// PageInit 返回sql查询时的limit的值
func (p *Page) PageInit() (offset int) {
  p.DefaultPageSize()
  p.Offset = p.PageSize*p.PageNum - p.PageSize
  return p.Offset
}

// PageInfoFeed Feed流形式的分页信息
type PageInfoFeed struct {
  PageSize int    `json:"pageSize"` // 每页数据大小
  PageNum  int    `json:"pageNum"`  // 当前数据数量
  HasMore  string `json:"hasMore"`  // 0:没有更多了,1:还有更多数据
  LastID   string `json:"lastID"`   // 最后一条数的ID
}

// DefaultPageSize 设置默认的pageSize的值
func (p *PageInfoFeed) DefaultPageSize() {
  if p.PageSize == 0 {
    // 默认显示15条数据
    p.PageSize = 15
  }
}

// Success 默认成功返回
func (r *Result) Success() {
  r.Msg = "Success!"
  r.Code = feibor.Success
  r.HTTPStatusCode = HTTPStatusSuccess
}

// Fail 返回失败
func (r *Result) Fail() {
  // Fail 方法主要提供返回错误的json数据
  r.Msg = "Something get error."
  r.Code = feibor.Fail
}

// FailMessage 方法主要提供返回错误的json数据
func (r *Result) FailMessage(msg string) {
  r.Msg = msg
  r.Code = feibor.Fail
}

// FailErr 携带error信息,如果是respError，则
// 必然存在errorCode和msg，因此进行赋值。否则不赋值
func (r *Result) FailErr(err error) {
  switch vtype := err.(type) {
  case *feibor.RespError:
    r.Msg = vtype.ErrorMsg
    r.Code = vtype.ErrorCode
    r.HTTPStatusCode = vtype.HTTPStatusCode
  default:
    r.Code = feibor.Fail
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
func (r *Result) ExecWithNoData(ctx *gin.Context, err error) {
  r.Exec(ctx, err, nil)
}

// Exec 直接执行c.context的语法，不需要再单独执行
func (r *Result) Exec(ctx *gin.Context, err error, resObj interface{}) {
  if err != nil {
    logrus.Error(err)
    r.FailErr(err)
    ctx.JSON(r.HTTPStatusCode, r)
    return
  }
  r.SuccessData(resObj)
  ctx.JSON(HTTPStatusSuccess, r)
  if err != nil {
    logrus.Error(err)
  }
}

// ExecWithNoDataPage 无需数据传输时的语法
func (r *ResultPage) ExecWithNoDataPage(c *gin.Context, err error) {
  r.ExecPage(c, err, nil)
}

// ExecPage 直接执行c.context的语法，不需要再单独执行
func (r *ResultPage) ExecPage(c *gin.Context, err error, resObj interface{}) {
  if err != nil {
    logrus.Error(err)
    r.FailErr(err)
    c.JSON(r.HTTPStatusCode, r)
    return
  }
  r.SuccessData(resObj)
  c.JSON(HTTPStatusSuccess, r)
}
