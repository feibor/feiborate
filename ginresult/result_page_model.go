package ginresult

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Page 通用的分页对象
type Page struct {
	PageSize int `json:"pageSize"` // 每页数据大小
	PageNum  int `json:"pageNum"`  // 当前页码
	Total    int `json:"total"`    // 数据总量
	Offset   int `json:"-"`        // 数据查询偏移量，即 pageSize*pageNum - pageSize
}

// NewPageResult 创建pageResult
func NewPageResult() *ResultPage {
	return &ResultPage{
		Result: &Result{},
	}
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

// NewPage 创建Page对象
func NewPage(ctx *gin.Context, pageSize int, pageNum int) (p *Page) {
	p = &Page{
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	p.PageInit()
	return
}

// ResultPage 分页查询结果集的对象
type ResultPage struct {
	*Result
	*Page
}

// InvokeResult Callback函数方式处理
func (r *ResultPage) InvokeResult(fc func(result *ResultPage) error) (err error) {
	result := r
	err = fc(result)
	if err != nil {
		result.ExecWithNoDataPage(err)
		return
	}
	return
}

// InvokeResultResp Callback函数方式处理
func (r *ResultPage) InvokeResultResp(fc func(result *ResultPage) error) (resp interface{}, err error) {
	err = fc(r)
	if err != nil {
		r.ExecWithNoDataPage(err)
		return
	}
	r.ExecPage(nil, resp)
	return
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

// ExecWithNoDataPage 无需数据传输时的语法
func (r *ResultPage) ExecWithNoDataPage(err error) {
	r.ExecPage(err, nil)
}

// ExecPage 直接执行c.context的语法，不需要再单独执行
func (r *ResultPage) ExecPage(err error, resObj interface{}) {
	if err != nil {
		logrus.Error(err)
		r.FailErr(err)
		r.ctx.JSON(r.HTTPStatusCode, r)
		return
	}
	r.SuccessData(resObj)
	r.ctx.JSON(HTTPStatusSuccess, r)
}
