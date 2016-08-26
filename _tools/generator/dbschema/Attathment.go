//Generated by webx-top/db
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
)

type Attathment struct {
	trans	*factory.Transaction
	
	Id             	int     	`db:"id,omitempty,pk" comment:"ID"`
	Name           	string  	`db:"name" comment:"文件名"`
	Path           	string  	`db:"path" comment:"保存路径"`
	Extension      	string  	`db:"extension" comment:"扩展名"`
	Type           	string  	`db:"type" comment:"文件类型"`
	Size           	int64   	`db:"size" comment:"文件尺寸"`
	Uid            	int     	`db:"uid" comment:"UID"`
	Deleted        	int     	`db:"deleted" comment:"被删除时间"`
	Created        	int     	`db:"created" comment:"创建时间"`
	Audited        	int     	`db:"audited" comment:"审核时间"`
	RcId           	int     	`db:"rc_id" comment:"关联id"`
	RcType         	string  	`db:"rc_type" comment:"关联类型"`
	Tags           	string  	`db:"tags" comment:"标签"`
}

func (this *Attathment) SetTrans(trans *factory.Transaction) *Attathment {
	this.trans = trans
	return this
}

func (this *Attathment) Param() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("attathment")
}

func (this *Attathment) Get(mw func(db.Result) db.Result) error {
	return this.Param().SetRecv(this).SetMiddleware(mw).One()
}

func (this *Attathment) List(mw func(db.Result) db.Result, page, size int) ([]*Attathment, func() int64, error) {
	r := []*Attathment{}
	counter, err := this.Param().SetPage(page).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Attathment) ListByOffset(mw func(db.Result) db.Result, offset, size int) ([]*Attathment, func() int64, error) {
	r := []*Attathment{}
	counter, err := this.Param().SetOffset(offset).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Attathment) Add(args ...*Attathment) (interface{}, error) {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).Insert()
}

func (this *Attathment) Edit(mw func(db.Result) db.Result, args ...*Attathment) error {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).SetMiddleware(mw).Update()
}

func (this *Attathment) Delete(mw func(db.Result) db.Result) error {
	return this.Param().SetMiddleware(mw).Delete()
}

