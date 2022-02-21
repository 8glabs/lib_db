package lib_db

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type OrmUtils struct {
	pageNum      int
	pageSize     int
	order        string
	selects      string
	model        interface{}
	listTemplate map[string]interface{}
	mainTable    string
	preloads     string
	db           *gorm.DB
}

func NewOrmUtils(db *gorm.DB, model interface{}) *OrmUtils {
	t := &OrmUtils{}
	t.pageNum = 0
	t.pageSize = 10
	t.db = db
	t.model = model

	t.listTemplate = make(map[string]interface{})
	t.listTemplate["data"] = nil
	t.listTemplate["page"] = 0
	t.listTemplate["pageSize"] = 10
	t.listTemplate["pageCount"] = 0
	return t
}

func (p *OrmUtils) SetPage(n int) *OrmUtils {
	if n >= 0 {
		p.pageNum = int(n)
	}
	return p
}

func (p *OrmUtils) SetPageSize(size int) *OrmUtils {
	if size > 0 {
		p.pageSize = int(size)
	}
	return p
}

func (p *OrmUtils) SetOrder(order string) *OrmUtils {
	p.order = order
	return p
}

func (p *OrmUtils) Table(main string) *OrmUtils {
	p.mainTable = main
	p.db = p.db.Table(main)
	return p
}

func (p *OrmUtils) Select(query string) *OrmUtils {
	p.selects = query
	return p
}

func (p *OrmUtils) GroupBy(cond string) *OrmUtils {
	p.db = p.db.Group(cond)
	return p
}

func (p *OrmUtils) GroupByProject() *OrmUtils {
	p.db = p.db.Group("projects.name")
	return p
}

func (p *OrmUtils) Preload(preloads string) *OrmUtils {
	p.db = p.db.Preload(preloads)
	return p
}

func (p *OrmUtils) Or(query string, args ...interface{}) *OrmUtils {
	p.db = p.db.Or(query, args...)
	return p
}

func (p *OrmUtils) Join(join string) *OrmUtils {
	p.db = p.db.Joins(join)
	return p
}

func (p *OrmUtils) Search(like string, fields ...string) *OrmUtils {
	if len(like) == 0 {
		return p
	}
	fnum := len(fields)
	where := make([]string, fnum)
	for i := 0; i < fnum; i++ {
		l := fmt.Sprintf("%%%s%%", like)
		where[i] = fmt.Sprintf("(%s LIKE '%s')", fields[i], l)
	}
	search := strings.Join(where, " OR ")
	p.db = p.db.Where(search)
	return p
}

func (p *OrmUtils) Where(query string, args ...interface{}) *OrmUtils {
	p.db = p.db.Where(query, args...)
	return p
}

func (p *OrmUtils) List() (map[string]interface{}, error) {
	p.preList()
	var count int64
	if err := p.db.Model(p.model).Count(&count).Error; err != nil {
		return nil, err
	}
	if p.order != "" {
		p.db = p.db.Order(p.order)
	}
	p.db = p.db.Find(p.model)
	if errors.Is(p.db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if p.db.Error != nil {
		return nil, errors.New(p.db.Error.Error())
	}

	p.listTemplate["data"] = p.model
	p.listTemplate["page"] = 0
	p.listTemplate["pageSize"] = 0
	p.listTemplate["pageCount"] = 1
	return p.listTemplate, nil
}

func (p *OrmUtils) PageList() (map[string]interface{}, error) {
	p.preList()

	var count int64
	if err := p.db.Model(p.model).Count(&count).Error; err != nil {
		return nil, err
	}

	var totalPage int
	totalPage = (int(count) + p.pageSize - 1) / p.pageSize

	if p.pageNum < 0 {
		p.pageNum = 0
	}
	if totalPage > 0 && p.pageNum >= totalPage {
		p.pageNum = totalPage - 1
	}
	if p.order != "" {
		p.db = p.db.Order(p.order)
	}

	p.db = p.db.Offset(p.pageNum * p.pageSize).Limit(p.pageSize).Find(p.model)
	if errors.Is(p.db.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if p.db.Error != nil {
		return nil, errors.New(p.db.Error.Error())
	}
	p.listTemplate["data"] = p.model
	p.listTemplate["page"] = p.pageNum
	p.listTemplate["pageSize"] = p.pageSize
	p.listTemplate["pageCount"] = totalPage
	// p.listTemplate["count"] = count
	return p.listTemplate, nil
}

func (p *OrmUtils) preList() {

	if len(p.selects) > 0 {
		p.db = p.db.Select(p.selects)
	}

}
