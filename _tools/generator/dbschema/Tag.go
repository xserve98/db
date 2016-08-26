//Generated by webx-top/db
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
)

type Tag struct {
	trans	*factory.Transaction
	
	Id             	int     	`db:"id,omitempty,pk" comment:"ID"`
	Name           	string  	`db:"name" comment:"标签名"`
	Uid            	int     	`db:"uid" comment:"创建者"`
	Created        	int     	`db:"created" comment:"创建时间"`
	Times          	int     	`db:"times" comment:"使用次数"`
	RcType         	string  	`db:"rc_type" comment:"关联类型"`
}

func (this *Tag) SetTrans(trans *factory.Transaction) *Tag {
	this.trans = trans
	return this
}

func (this *Tag) Param() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("tag")
}

func (this *Tag) Get(mw func(db.Result) db.Result) error {
	return this.Param().SetRecv(this).SetMiddleware(mw).One()
}

func (this *Tag) List(mw func(db.Result) db.Result, page, size int) ([]*Tag, func() int64, error) {
	r := []*Tag{}
	counter, err := this.Param().SetPage(page).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Tag) ListByOffset(mw func(db.Result) db.Result, offset, size int) ([]*Tag, func() int64, error) {
	r := []*Tag{}
	counter, err := this.Param().SetOffset(offset).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Tag) Add(args ...*Tag) (interface{}, error) {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).Insert()
}

func (this *Tag) Edit(mw func(db.Result) db.Result, args ...*Tag) error {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).SetMiddleware(mw).Update()
}

func (this *Tag) Delete(mw func(db.Result) db.Result) error {
	return this.Param().SetMiddleware(mw).Delete()
}

