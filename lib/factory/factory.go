// added by swh@admpub.com
package factory

import (
	"errors"
	"strings"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/sqlbuilder"
)

const (
	R = iota
	W
)

var (
	ErrNotFoundKey = errors.New(`not found the key`)
)

func New() *Factory {
	return &Factory{
		databases: make([]*Cluster, 0),
	}
}

type Factory struct {
	databases []*Cluster
	cacher    Cacher
}

func (f *Factory) Debug() bool {
	return db.Debug
}

func (f *Factory) SetDebug(on bool) *Factory {
	db.Debug = on
	return f
}

func (f *Factory) SetCacher(cc Cacher) *Factory {
	f.cacher = cc
	return f
}

func (f *Factory) AddDB(databases ...sqlbuilder.Database) *Factory {
	if len(f.databases) > 0 {
		f.databases[0].AddW(databases...)
	} else {
		c := NewCluster()
		c.AddW(databases...)
		f.databases = append(f.databases, c)
	}
	return f
}

func (f *Factory) AddSlaveDB(databases ...sqlbuilder.Database) *Factory {
	if len(f.databases) > 0 {
		f.databases[0].AddR(databases...)
	} else {
		c := NewCluster()
		c.AddR(databases...)
		f.databases = append(f.databases, c)
	}
	return f
}

func (f *Factory) SetCluster(index int, cluster *Cluster) *Factory {
	if len(f.databases) > index {
		f.databases[index] = cluster
	}
	return f
}

func (f *Factory) AddCluster(clusters ...*Cluster) *Factory {
	f.databases = append(f.databases, clusters...)
	return f
}

func (f *Factory) Cluster(index int) *Cluster {
	if len(f.databases) > index {
		return f.databases[index]
	}
	if index == 0 {
		panic(`Not connected to any database`)
	}
	return f.Cluster(0)
}

func (f *Factory) GetCluster(index int) *Cluster {
	return f.Cluster(index)
}

func (f *Factory) Collection(collection string, args ...int) db.Collection {
	var index int
	switch len(args) {
	case 2:
		index = args[0]
		if args[1] == R {
			c := f.Cluster(index)
			collection = c.Table(collection)
			return c.R().Collection(collection)
		}
	case 1:
		index = args[0]
	}
	c := f.Cluster(index)
	collection = c.Table(collection)
	return c.W().Collection(collection)
}

func (f *Factory) Find(collection string, args ...interface{}) db.Result {
	return f.Collection(collection).Find(args...)
}

func (f *Factory) FindR(collection string, args ...interface{}) db.Result {
	return f.Collection(collection, 0, R).Find(args...)
}

func (f *Factory) FindDB(index int, collection string, args ...interface{}) db.Result {
	return f.Collection(collection, index).Find(args...)
}

func (f *Factory) FindDBR(index int, collection string, args ...interface{}) db.Result {
	return f.Collection(collection, index, R).Find(args...)
}

func (f *Factory) CloseAll() {
	for _, cluster := range f.databases {
		cluster.CloseAll()
	}
}

func (f *Factory) Result(param *Param) db.Result {
	return f.Collection(param.Collection, param.Index, param.ReadOrWrite).Find(param.Args...)
}

// ================================
// API
// ================================

// Read ==========================

func (f *Factory) SelectAll(param *Param) error {
	selector := f.Select(param)
	if param.Size > 0 {
		selector = selector.Limit(param.Size).Offset(param.Offset())
	}
	if param.SelectorMiddleware != nil {
		selector = param.SelectorMiddleware(selector)
	}
	return selector.All(param.ResultData)
}

func (f *Factory) SelectOne(param *Param) error {
	selector := f.Select(param).Limit(1)
	if param.SelectorMiddleware != nil {
		selector = param.SelectorMiddleware(selector)
	}
	return selector.One(param.ResultData)
}

func (f *Factory) Select(param *Param) sqlbuilder.Selector {
	c := f.Cluster(param.Index)
	collection := c.Table(param.Collection)
	selector := c.R().Select(param.Cols...).From(collection)
	if param.Joins == nil {
		return selector
	}
	for _, join := range param.Joins {
		coll := c.Table(join.Collection)
		if len(join.Alias) > 0 {
			coll += ` AS ` + join.Alias
		}
		switch strings.ToUpper(join.Type) {
		case "LEFT":
			selector = selector.LeftJoin(coll)
		case "RIGHT":
			selector = selector.RightJoin(coll)
		case "CROSS":
			selector = selector.CrossJoin(coll)
		case "INNER":
			selector = selector.FullJoin(coll)
		default:
			selector = selector.FullJoin(coll)
		}
		if len(join.Condition) > 0 {
			selector = selector.On(join.Condition)
		}
	}
	return selector
}

func (f *Factory) All(param *Param) error {
	if param.Lifetime > 0 && f.cacher != nil {
		data, err := f.cacher.Get(param.CachedKey())
		if err == nil && data != nil {
			if v, ok := data.(*Param); ok {
				param = v
				param.factory = f
				return nil
			}
		}
		defer f.cacher.Put(param.CachedKey(), param, param.Lifetime)
	}
	var res db.Result
	if param.Middleware == nil {
		res = f.FindDBR(param.Index, param.Collection, param.Args...)
	} else {
		res = param.Middleware(f.FindDBR(param.Index, param.Collection, param.Args...))
	}
	return res.All(param.ResultData)
}

func (f *Factory) List(param *Param) (func() int64, error) {

	if param.Lifetime > 0 && f.cacher != nil {
		data, err := f.cacher.Get(param.CachedKey())
		if err == nil && data != nil {
			if v, ok := data.(*Param); ok {
				param = v
				param.factory = f
				return func() int64 {
					return param.Total
				}, nil
			}
		}
		defer f.cacher.Put(param.CachedKey(), param, param.Lifetime)
	}

	var res db.Result
	if param.Middleware == nil {
		param.CountFunc = func() int64 {
			if param.Total <= 0 {
				res := f.FindDBR(param.Index, param.Collection, param.Args...)
				count, _ := res.Count()
				param.Total = int64(count)
			}
			return param.Total
		}
		res = f.FindDBR(param.Index, param.Collection, param.Args...).Limit(param.Size).Offset(param.Offset())
	} else {
		param.CountFunc = func() int64 {
			if param.Total <= 0 {
				res := param.Middleware(f.FindDBR(param.Index, param.Collection, param.Args...))
				count, _ := res.Count()
				param.Total = int64(count)
			}
			return param.Total
		}
		res = param.Middleware(f.FindDBR(param.Index, param.Collection, param.Args...).Limit(param.Size).Offset(param.Offset()))
	}
	return param.CountFunc, res.All(param.ResultData)
}

func (f *Factory) One(param *Param) error {

	if param.Lifetime > 0 && f.cacher != nil {
		data, err := f.cacher.Get(param.CachedKey())
		if err == nil && data != nil {
			if v, ok := data.(*Param); ok {
				param = v
				param.factory = f
				return nil
			}
		}
		defer f.cacher.Put(param.CachedKey(), param, param.Lifetime)
	}

	var res db.Result
	if param.Middleware == nil {
		res = f.FindDBR(param.Index, param.Collection, param.Args...)
	} else {
		res = param.Middleware(f.FindDBR(param.Index, param.Collection, param.Args...))
	}
	return res.One(param.ResultData)
}

func (f *Factory) Count(param *Param) (int64, error) {

	if param.Lifetime > 0 && f.cacher != nil {
		data, err := f.cacher.Get(param.CachedKey())
		if err == nil && data != nil {
			if v, ok := data.(*Param); ok {
				param = v
				param.factory = f
				return param.Total, nil
			}
		}
		defer f.cacher.Put(param.CachedKey(), param, param.Lifetime)
	}

	var cnt uint64
	var err error
	var res db.Result
	if param.Middleware == nil {
		res = f.FindDBR(param.Index, param.Collection, param.Args...)
	} else {
		res = param.Middleware(f.FindDBR(param.Index, param.Collection, param.Args...))
	}
	cnt, err = res.Count()
	param.Total = int64(cnt)
	return param.Total, err
}

// Write ==========================

func (f *Factory) Insert(param *Param) (interface{}, error) {
	return f.Collection(param.Collection, param.Index, W).Insert(param.SaveData)
}

func (f *Factory) Update(param *Param) error {
	var res db.Result
	if param.Middleware == nil {
		res = f.FindDB(param.Index, param.Collection, param.Args...)
	} else {
		res = param.Middleware(f.FindDB(param.Index, param.Collection, param.Args...))
	}
	return res.Update(param.SaveData)
}

func (f *Factory) Delete(param *Param) error {
	var res db.Result
	if param.Middleware == nil {
		res = f.FindDB(param.Index, param.Collection, param.Args...)
	} else {
		res = param.Middleware(f.FindDB(param.Index, param.Collection, param.Args...))
	}
	return res.Delete()
}
