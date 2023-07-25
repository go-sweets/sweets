package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/mix-plus/go-mixplus/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StringTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StringTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return 15
	}
	if pageSize > 15 {
		return 15
	}

	return pageSize
}

func GetPageOffset(c *gin.Context) (offset, limit int) {
	page := convert.StringTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}

	limit = convert.StringTo(c.Query("page_size")).MustInt()
	if limit <= 0 {
		limit = 15
	} else if limit > 15 {
		limit = 15
	}
	offset = (page - 1) * limit
	return
}
