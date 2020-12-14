package service

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Page struct {
	PageSize  int   `json:"pageSize"`
	PageNo    int   `json:"pageNo"`
	ItemCount int64 `json:"itemCount"`
	PageCount int64 `json:"pageCount"`
}

type PageResult struct {
	Page Page        `json:"pageInfo"`
	Data interface{} `json:"dataList"`
}

func NewPage(pageNo int, pageSize int, itemCount int64) Page {
	return Page{PageSize: pageSize, PageNo: pageNo, ItemCount: itemCount, PageCount: itemCount / int64(pageSize)}
}

func NewPageResult(pageNo int, pageSize int, itemCount int64, data interface{}) PageResult {
	page := Page{PageSize: pageSize, PageNo: pageNo, ItemCount: itemCount, PageCount: itemCount / int64(pageSize)}
	d := data
	return PageResult{Page: page, Data: d}
}

func NewFailResponse() Response {
	return Response{"999", "fail", nil}
}

func NewSuccessResponse() Response {
	return Response{"200", "success", nil}
}

func NewSuccessResponseWithData(data interface{}) Response {
	return Response{"200", "success", data}
}
