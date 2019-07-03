package controllers

import (
	"github.com/Murray-LIANG/microservices-example/users/common"
	"gopkg.in/mgo.v2"
)

type Context struct {
	MongoSession *mgo.Session
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

func NewContext() *Context {
	session := common.GetDBSession().Copy()
	return &Context{MongoSession: session}
}

func (c *Context) DBCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}
