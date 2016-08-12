// added by swh@admpub.com
package factory

import (
	"github.com/webx-top/db"
)

type Param struct {
	factory    *Factory
	Index      int
	Collection string
	Middleware func(db.Result) db.Result
	Result     interface{}
	Args       []interface{}
	SaveData   interface{}
	Page       int
	Size       int
	Total      int64
}

func NewParam(factory *Factory) *Param {
	p := &Param{
		factory: factory,
		Args:    make([]interface{}, 0),
	}
	return p
}

func (p *Param) SetIndex(index int) *Param {
	p.Index = index
	return p
}

func (p *Param) SetCollection(collection string) *Param {
	p.Collection = collection
	return p
}

func (p *Param) SetMiddleware(middleware func(db.Result) db.Result) *Param {
	p.Middleware = middleware
	return p
}

func (p *Param) SetResult(result interface{}) *Param {
	p.Result = result
	return p
}

func (p *Param) SetArgs(args ...interface{}) *Param {
	p.Args = args
	return p
}

func (p *Param) SetSave(save interface{}) *Param {
	p.SaveData = save
	return p
}

func (p *Param) SetPage(n int) *Param {
	p.Page = n
	return p
}

func (p *Param) SetSize(size int) *Param {
	p.Size = size
	return p
}

func (p *Param) SetTotal(total int64) *Param {
	p.Total = total
	return p
}

func (p *Param) Offset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.Size
}

// Read ==========================

func (p *Param) All() error {
	return p.factory.All(p)
}

func (p *Param) PageList() (func() int64, error) {
	return p.factory.PageList(p)
}

func (p *Param) One() error {
	return p.factory.One(p)
}

func (p *Param) Count() (int64, error) {
	return p.factory.Count(p)
}

// Write ==========================

func (p *Param) Insert() (interface{}, error) {
	return p.factory.Insert(p)
}

func (p *Param) Update() error {
	return p.factory.Update(p)
}

func (p *Param) Delete() error {
	return p.factory.Delete(p)
}
