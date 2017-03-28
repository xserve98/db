//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

func init(){
	factory.Fields=map[string]map[string]*factory.FieldInfo{"comment":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"bigint", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"主键", GoType:"uint64", GoName:"Id"}, "quote":&factory.FieldInfo{Name:"quote", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"引用内容", GoType:"string", GoName:"Quote"}, "related_times":&factory.FieldInfo{Name:"related_times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"本身被回复次数", GoType:"uint", GoName:"RelatedTimes"}, "root_times":&factory.FieldInfo{Name:"root_times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"根节点下的所有回复次数", GoType:"uint", GoName:"RootTimes"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联内容类型", GoType:"string", GoName:"RcType"}, "uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"发布者用户名", GoType:"string", GoName:"Uname"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "status":&factory.FieldInfo{Name:"status", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"0", Comment:"状态", GoType:"int", GoName:"Status"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"关联内容ID", GoType:"uint64", GoName:"RcId"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"html", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}, "root_id":&factory.FieldInfo{Name:"root_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"", GoType:"uint64", GoName:"RootId"}, "r_type":&factory.FieldInfo{Name:"r_type", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"reply", "append"}, DefaultValue:"", Comment:"关联类型（reply-回复，append-追加）", GoType:"string", GoName:"RType"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"发布者id", GoType:"uint64", GoName:"Uid"}, "down":&factory.FieldInfo{Name:"down", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被踩次数", GoType:"uint64", GoName:"Down"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "for_uid":&factory.FieldInfo{Name:"for_uid", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被评人id", GoType:"uint64", GoName:"ForUid"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"内容", GoType:"string", GoName:"Content"}, "r_id":&factory.FieldInfo{Name:"r_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"关联本表的id", GoType:"uint64", GoName:"RId"}, "up":&factory.FieldInfo{Name:"up", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"0", Comment:"被顶次数", GoType:"uint64", GoName:"Up"}, "for_uname":&factory.FieldInfo{Name:"for_uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"被评人用户名", GoType:"string", GoName:"ForUname"}}, "post":map[string]*factory.FieldInfo{"comments":&factory.FieldInfo{Name:"comments", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被评论次数", GoType:"uint", GoName:"Comments"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "title":&factory.FieldInfo{Name:"title", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:180, Options:[]string{}, DefaultValue:"", Comment:"标题", GoType:"string", GoName:"Title"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"简介", GoType:"string", GoName:"Description"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "views":&factory.FieldInfo{Name:"views", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被浏览次数", GoType:"uint", GoName:"Views"}, "month":&factory.FieldInfo{Name:"month", DataType:"tinyint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"", Comment:"归档月份", GoType:"uint", GoName:"Month"}, "allow_comment":&factory.FieldInfo{Name:"allow_comment", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"是否允许评论", GoType:"string", GoName:"AllowComment"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类ID", GoType:"uint", GoName:"Catid"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"内容", GoType:"string", GoName:"Content"}, "uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"作者名", GoType:"string", GoName:"Uname"}, "passwd":&factory.FieldInfo{Name:"passwd", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"访问密码", GoType:"string", GoName:"Passwd"}, "likes":&factory.FieldInfo{Name:"likes", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被喜欢次数", GoType:"uint", GoName:"Likes"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被删除时间", GoType:"uint", GoName:"Deleted"}, "year":&factory.FieldInfo{Name:"year", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:5, Options:[]string{}, DefaultValue:"", Comment:"归档年份", GoType:"uint", GoName:"Year"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"html", "markdown"}, DefaultValue:"html", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "display":&factory.FieldInfo{Name:"display", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"ALL", "SELF", "FRIEND", "PWD"}, DefaultValue:"ALL", Comment:"显示", GoType:"string", GoName:"Display"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"UID", GoType:"uint", GoName:"Uid"}}, "user":map[string]*factory.FieldInfo{"mobile":&factory.FieldInfo{Name:"mobile", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:15, Options:[]string{}, DefaultValue:"", Comment:"手机号", GoType:"string", GoName:"Mobile"}, "login_ip":&factory.FieldInfo{Name:"login_ip", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:40, Options:[]string{}, DefaultValue:"", Comment:"上次登录IP", GoType:"string", GoName:"LoginIp"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "active":&factory.FieldInfo{Name:"active", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"激活", GoType:"string", GoName:"Active"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"UID", GoType:"uint", GoName:"Id"}, "uname":&factory.FieldInfo{Name:"uname", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"用户名", GoType:"string", GoName:"Uname"}, "salt":&factory.FieldInfo{Name:"salt", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"盐值", GoType:"string", GoName:"Salt"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "avatar":&factory.FieldInfo{Name:"avatar", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"头像", GoType:"string", GoName:"Avatar"}, "passwd":&factory.FieldInfo{Name:"passwd", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"密码", GoType:"string", GoName:"Passwd"}, "email":&factory.FieldInfo{Name:"email", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"邮箱", GoType:"string", GoName:"Email"}, "login_time":&factory.FieldInfo{Name:"login_time", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"上次登录时间", GoType:"uint", GoName:"LoginTime"}}, "album":map[string]*factory.FieldInfo{"title":&factory.FieldInfo{Name:"title", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:180, Options:[]string{}, DefaultValue:"", Comment:"标题", GoType:"string", GoName:"Title"}, "comments":&factory.FieldInfo{Name:"comments", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"评论次数", GoType:"uint", GoName:"Comments"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"简介", GoType:"string", GoName:"Description"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"正文", GoType:"string", GoName:"Content"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"编辑时间", GoType:"uint", GoName:"Updated"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"删除时间", GoType:"uint", GoName:"Deleted"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类ID", GoType:"uint", GoName:"Catid"}, "display":&factory.FieldInfo{Name:"display", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"ALL", "SELF", "FRIEND", "PWD"}, DefaultValue:"ALL", Comment:"显示", GoType:"string", GoName:"Display"}, "allow_comment":&factory.FieldInfo{Name:"allow_comment", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"Y", Comment:"是否允许评论", GoType:"string", GoName:"AllowComment"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "views":&factory.FieldInfo{Name:"views", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"浏览次数", GoType:"uint", GoName:"Views"}, "likes":&factory.FieldInfo{Name:"likes", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"喜欢次数", GoType:"uint", GoName:"Likes"}}, "attathment":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"UID", GoType:"uint", GoName:"Uid"}, "deleted":&factory.FieldInfo{Name:"deleted", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"被删除时间", GoType:"uint", GoName:"Deleted"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "size":&factory.FieldInfo{Name:"size", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"文件尺寸", GoType:"uint64", GoName:"Size"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "audited":&factory.FieldInfo{Name:"audited", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"审核时间", GoType:"uint", GoName:"Audited"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"关联id", GoType:"uint", GoName:"RcId"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"文件名", GoType:"string", GoName:"Name"}, "path":&factory.FieldInfo{Name:"path", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"保存路径", GoType:"string", GoName:"Path"}, "extension":&factory.FieldInfo{Name:"extension", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:5, Options:[]string{}, DefaultValue:"", Comment:"扩展名", GoType:"string", GoName:"Extension"}, "type":&factory.FieldInfo{Name:"type", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"image", "media", "other"}, DefaultValue:"image", Comment:"文件类型", GoType:"string", GoName:"Type"}, "tags":&factory.FieldInfo{Name:"tags", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"标签", GoType:"string", GoName:"Tags"}}, "category":map[string]*factory.FieldInfo{"pid":&factory.FieldInfo{Name:"pid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"上级分类", GoType:"uint", GoName:"Pid"}, "haschild":&factory.FieldInfo{Name:"haschild", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否有子分类", GoType:"string", GoName:"Haschild"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"分类名称", GoType:"string", GoName:"Name"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"说明", GoType:"string", GoName:"Description"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "sort":&factory.FieldInfo{Name:"sort", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"排序", GoType:"int", GoName:"Sort"}, "tmpl":&factory.FieldInfo{Name:"tmpl", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"模板", GoType:"string", GoName:"Tmpl"}}, "config":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"主键ID", GoType:"uint", GoName:"Id"}, "key":&factory.FieldInfo{Name:"key", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:60, Options:[]string{}, DefaultValue:"", Comment:"配置项", GoType:"string", GoName:"Key"}, "val":&factory.FieldInfo{Name:"val", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"配置值", GoType:"string", GoName:"Val"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}}, "link":map[string]*factory.FieldInfo{"updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "sort":&factory.FieldInfo{Name:"sort", DataType:"int", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"排序", GoType:"int", GoName:"Sort"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"主键ID", GoType:"uint", GoName:"Id"}, "url":&factory.FieldInfo{Name:"url", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"网址", GoType:"string", GoName:"Url"}, "logo":&factory.FieldInfo{Name:"logo", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"LOGO", GoType:"string", GoName:"Logo"}, "show":&factory.FieldInfo{Name:"show", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否显示", GoType:"string", GoName:"Show"}, "verified":&factory.FieldInfo{Name:"verified", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"验证时间", GoType:"uint", GoName:"Verified"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "catid":&factory.FieldInfo{Name:"catid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"分类", GoType:"uint", GoName:"Catid"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"名称", GoType:"string", GoName:"Name"}}, "ocontent":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"bigint", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint64", GoName:"Id"}, "rc_id":&factory.FieldInfo{Name:"rc_id", DataType:"bigint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"关联ID", GoType:"uint64", GoName:"RcId"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "content":&factory.FieldInfo{Name:"content", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"博客原始内容", GoType:"string", GoName:"Content"}, "etype":&factory.FieldInfo{Name:"etype", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"markdown"}, DefaultValue:"markdown", Comment:"编辑器类型", GoType:"string", GoName:"Etype"}}, "tag":map[string]*factory.FieldInfo{"times":&factory.FieldInfo{Name:"times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"使用次数", GoType:"uint", GoName:"Times"}, "rc_type":&factory.FieldInfo{Name:"rc_type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"post", Comment:"关联类型", GoType:"string", GoName:"RcType"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"标签名", GoType:"string", GoName:"Name"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建者", GoType:"uint", GoName:"Uid"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}}}

}

