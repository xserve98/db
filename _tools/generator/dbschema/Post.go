//Generated by webx-top/db
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
)

type Post struct {
	trans	*factory.Transaction
	
	Id              	int     	`db:"id,omitempty,pk" comment:"ID"`
	Title           	string  	`db:"title" comment:"标题"`
	Description     	string  	`db:"description" comment:"简介"`
	Content         	string  	`db:"content" comment:"内容"`
	Etype           	string  	`db:"etype" comment:"编辑器类型"`
	Created         	int     	`db:"created" comment:"创建时间"`
	Updated         	int     	`db:"updated" comment:"修改时间"`
	Display         	string  	`db:"display" comment:"显示"`
	Uid             	int     	`db:"uid" comment:"UID"`
	Uname           	string  	`db:"uname" comment:"作者名"`
	Passwd          	string  	`db:"passwd" comment:"访问密码"`
	Views           	int     	`db:"views" comment:"被浏览次数"`
	Comments        	int     	`db:"comments" comment:"被评论次数"`
	Likes           	int     	`db:"likes" comment:"被喜欢次数"`
	Deleted         	int     	`db:"deleted" comment:"被删除时间"`
	Year            	int     	`db:"year" comment:"归档年份"`
	Month           	string  	`db:"month" comment:"归档月份"`
	AllowComment    	string  	`db:"allow_comment" comment:"是否允许评论"`
	Tags            	string  	`db:"tags" comment:"标签"`
	Catid           	int     	`db:"catid" comment:"分类ID"`
}

func (this *Post) SetTrans(trans *factory.Transaction) *Post {
	this.trans = trans
	return this
}

func (this *Post) Param() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("post")
}

func (this *Post) Get(mw func(db.Result) db.Result) error {
	return this.Param().SetRecv(this).SetMiddleware(mw).One()
}

func (this *Post) List(mw func(db.Result) db.Result, page, size int) ([]*Post, func() int64, error) {
	r := []*Post{}
	counter, err := this.Param().SetPage(page).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Post) ListByOffset(mw func(db.Result) db.Result, offset, size int) ([]*Post, func() int64, error) {
	r := []*Post{}
	counter, err := this.Param().SetOffset(offset).SetSize(size).SetRecv(&r).SetMiddleware(mw).List()
	return r, counter, err
}

func (this *Post) Add(args ...*Post) (interface{}, error) {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).Insert()
}

func (this *Post) Edit(mw func(db.Result) db.Result, args ...*Post) error {
	var data = this
	if len(args)>0 {
		data = args[0]
	}
	return this.Param().SetSend(data).SetMiddleware(mw).Update()
}

func (this *Post) Delete(mw func(db.Result) db.Result) error {
	return this.Param().SetMiddleware(mw).Delete()
}

