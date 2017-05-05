package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sqlite3"
	"github.com/unknwon/com"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//分类
type Category struct {
	Id              int64
	Title           string    //名称
	Created         time.Time `orm:"index"` //创建时间，建立索引
	views           int64     `orm:"index"` //浏览次数，建立索引
	TopicTime       time.Time `orm:"index"` //最后操作时间
	TopicCount      int64     //浏览次数
	TopicLastUserId int64     //最后操作人
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(5000)"`
	Created time.Time `orm:"index"`
}

//文章
type Topic struct {
	Id              int64
	Uid             int64  //用户ID
	Title           string //标题
	Category        string
	Content         string    `orm:"size(5000)"` //文章内容，string默认长度为255 使用“`”重新设定长度
	Attachment      string    //附件
	Created         time.Time `orm:"index"` //创建时间，建立索引
	Updated         time.Time `orm:"index"` //最后更新时间
	Views           int64     `orm:"index"` //浏览次数，建立索引
	Author          string    //作者
	ReplyTime       time.Time `orm:"index"` //评论时间
	ReplyCount      int64     //评论条数
	ReplyLastUserId int64     //最后评论ID
}

func AddReply(id, name, content string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Comment{
		Tid:     cid,
		Name:    name,
		Content: content,
		Created: time.Now(),
	}
	_, err = o.Insert(reply)

	return err
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Comment), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	nowtime := time.Now()
	cate := &Category{
		Title:     name,
		Created:   nowtime,
		TopicTime: nowtime,
	}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	var err error
	if !isDesc {
		_, err = qs.All(&topics)
	} else {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	}
	return topics, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetTopic(id string) (*Topic, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", cid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(id, category, title, content string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: cid}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		_, err = o.Update(topic)
	}
	return err
}

func UpdTopic(id string, flag bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: cid}
	if o.Read(topic) == nil {
		if flag {
			topic.ReplyCount++
		} else {
			topic.ReplyCount--
			if topic.ReplyCount < 0 {
				topic.ReplyCount = 0
			}
		}

		topic.ReplyTime = time.Now()
		_, err = o.Update(topic)
	}
	return err
}

func DeleteTopic(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: cid}
	_, err = o.Delete(topic)
	return err
}

func GetAllReplies(id string) ([]*Comment, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	replies := make([]*Comment, 0)

	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", cid).All(&replies)
	return replies, err
}

func DelReply(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Comment{Id: cid}
	_, err = o.Delete(reply)
	return err
}
