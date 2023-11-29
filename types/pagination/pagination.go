package pagination

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

const DefaultPagingSize = 100

const MaxPagingSize = 2000

const FirstPage = 1

type Page struct {
	Size int64 `json:"size"` // Page Size, default 100
	Numb int64 `json:"numb"` // Page Numb, From One
}

func PageALL() *Page {
	return &Page{
		Size: MaxPagingSize,
		Numb: FirstPage,
	}
}

type Paging struct {
	Size  int64 `json:"size"`  // Page Size, default 100
	Numb  int64 `json:"numb"`  // Page Numb, From One
	Total int64 `json:"total"` // Page Total, The Page Total
	Count int64 `json:"count"` // Item Count, The Item Count
}

func (paging *Paging) ToString() string {
	return fmt.Sprintf("size: %d|numb: %d|total: %d|count: %d",
		paging.Size, paging.Numb, paging.Total, paging.Count)
}

func (paging *Paging) WithCount(count int64) {
	if paging.Size == 0 {
		paging.Size = DefaultPagingSize
	}
	if paging.Numb < FirstPage {
		paging.Numb = FirstPage
	}

	if count == 0 {
		paging.Total = 0
		paging.Numb = FirstPage
		return
	}
	paging.Count = count
	paging.Total = paging.Count / paging.Size
	if paging.Total%paging.Size != 0 {
		paging.Total += 1
	}

	if paging.Numb > paging.Total {
		paging.Numb = paging.Total
	}

	if paging.Numb < 1 {
		paging.Numb = 1
	}
}

func (paging *Paging) Skip() int64 {
	return (paging.Numb - 1) * paging.Size
}

func (paging *Paging) Limit() int64 {
	return paging.Size
}

func PagingOfPage(page *Page) Paging {
	if page == nil {
		return PagingOf(DefaultPagingSize, FirstPage)
	}
	return PagingOf(page.Size, page.Numb)
}

func PagingOf(size int64, current int64) Paging {
	if size <= 0 {
		size = DefaultPagingSize
	}
	if size > MaxPagingSize {
		size = MaxPagingSize
	}
	if current < FirstPage {
		current = FirstPage
	}
	return Paging{
		Size:  size,
		Numb:  current,
		Total: 0,
		Count: 0,
	}
}

func PagingALL() Paging {
	return PagingOf(MaxPagingSize, FirstPage)
}

type Pagination struct {
	Paging *Paging       `bson:"paging" json:"paging"`
	Items  []interface{} `bson:"items" json:"items"`
}

func PagingWrap[T interface{}](paging *Paging, arr []*T, wrap func(m *T) interface{}) Pagination {
	if len(arr) == 0 {
		return Pagination{
			Paging: paging,
			Items:  []interface{}{},
		}
	}
	items := make([]any, len(arr))
	for i, m := range arr {
		items[i] = wrap(m)
	}
	return Pagination{
		Paging: paging,
		Items:  items,
	}
}

type Limit int64

func (limit Limit) To() int64 {
	return int64(limit)
}

func (limit Limit) Iter(exec func(i int64) (bool, *errors.Error), skipErr bool) *errors.Error {
	max := limit.To()
	var n int64
	for n = 0; n < max; n++ {
		end, err := exec(n)
		if err != nil {
			if skipErr {
				logger.Logger.Error("call exec err", zap.Error(err))
				continue
			}
			return err
		}
		if end {
			break
		}
	}
	return nil
}

func LimitOf(l int64) Limit {
	if l < 0 {
		return Limit(DefaultPagingSize)
	}
	if l > MaxPagingSize {
		return Limit(MaxPagingSize)
	}
	return Limit(l)
}
