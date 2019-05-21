package models

import (
	"github.com/cocotyty/gentool/daogen/itfc"

	"bytes"
	"database/sql"
	"time"
)

type User struct {
	Id               int       `db:"id" json:"id"`                               // id
	Uid              int64     `db:"uid" json:"uid"`                             // 用户唯一id
	Name             string    `db:"name" json:"name"`                           // 真实姓名
	IdCard           string    `db:"id_card" json:"id_card"`                     // 身份证
	Status           int       `db:"status" json:"status"`                       // 状态：0未认证，1认证成功，2认证失败
	Mobile           int64     `db:"mobile" json:"mobile"`                       // 手机号
	Password         string    `db:"password" json:"password"`                   // 支付密码
	TotalBalance     int64     `db:"total_balance" json:"total_balance"`         // 总余额
	FrozenBalance    int64     `db:"frozen_balance" json:"frozen_balance"`       // 冻结余额
	AvailableBalance int64     `db:"available_balance" json:"available_balance"` // 可用余额
	Fee              int       `db:"fee" json:"fee"`                             // 个人手续费率
	CreatedTime      time.Time `db:"created_time" json:"created_time"`           // 创建时间
	UpdatedTime      time.Time `db:"updated_time" json:"updated_time"`           // 更新时间
	Baned            int       `db:"baned" json:"baned"`                         // baned
}
type UserDao struct {
	ColId               string
	ColUid              string
	ColName             string
	ColIdCard           string
	ColStatus           string
	ColMobile           string
	ColPassword         string
	ColTotalBalance     string
	ColFrozenBalance    string
	ColAvailableBalance string
	ColFee              string
	ColCreatedTime      string
	ColUpdatedTime      string
	ColBaned            string
	Columns             []string
	Table               string

	DB itfc.SqlxHandle
}

func NewUserDao(db itfc.SqlxHandle) *UserDao {
	dao := &UserDao{DB: db}
	dao.Init()
	return dao
}

func (dao *UserDao) Init() {
	dao.Columns = []string{
		"id",
		"uid",
		"name",
		"id_card",
		"status",
		"mobile",
		"password",
		"total_balance",
		"frozen_balance",
		"available_balance",
		"fee",
		"created_time",
		"updated_time",
		"baned",
	}

	dao.ColId = "id"
	dao.ColUid = "uid"
	dao.ColName = "name"
	dao.ColIdCard = "id_card"
	dao.ColStatus = "status"
	dao.ColMobile = "mobile"
	dao.ColPassword = "password"
	dao.ColTotalBalance = "total_balance"
	dao.ColFrozenBalance = "frozen_balance"
	dao.ColAvailableBalance = "available_balance"
	dao.ColFee = "fee"
	dao.ColCreatedTime = "created_time"
	dao.ColUpdatedTime = "updated_time"
	dao.ColBaned = "baned"
	dao.Table = "user"
}
func (dao *UserDao) Fields() string {
	return "`id` ,`uid` ,`name` ,`id_card` ,`status` ,`mobile` ,`password` ,`total_balance` ,`frozen_balance` ,`available_balance` ,`fee` ,`created_time` ,`updated_time` ,`baned` "
}

func (dao *UserDao) FindByID(ID int) (one *User, err error) {
	one = &User{}
	err = dao.DB.Get(one, "select "+dao.Fields()+" from `user` where `id`=? limit 1", ID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (dao *UserDao) Page(page, pageSize int) (list []*User, err error) {
	list = []*User{}
	err = dao.DB.Select(&list, "select "+dao.Fields()+" from `user` limit ?,?", pageSize*(page-1), pageSize)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}
func (dao *UserDao) WherePage(whereSql string, page, pageSize int, args ...interface{}) (list []*User, err error) {
	list = []*User{}
	args = append(args, pageSize*(page-1), pageSize)
	err = dao.DB.Select(&list, "select "+dao.Fields()+" from `user` where "+whereSql+" limit ?,?", args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}

func (dao *UserDao) RawSelect(str string, args []interface{}) (list []*User, err error) {
	list = []*User{}
	err = dao.DB.Select(&list, str, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}

func (dao *UserDao) All() (list []*User, err error) {
	list = []*User{}
	err = dao.DB.Select(&list, "select "+dao.Fields()+" from `user` ")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}
func (dao *UserDao) Where(whereSql string, args ...interface{}) (list []*User, err error) {
	list = []*User{}
	err = dao.DB.Select(&list, "select "+dao.Fields()+" from `user` where "+whereSql, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}
func (dao *UserDao) Find(whereSql string, args ...interface{}) (one *User, err error) {
	one = &User{}
	err = dao.DB.Get(one, "select "+dao.Fields()+" from `user` where "+whereSql, args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (dao *UserDao) Insert(one *User) (res sql.Result, err error) {
	return dao.DB.Exec("insert into `user` ( `uid` ,`name` ,`id_card` ,`status` ,`mobile` ,`password` ,`total_balance` ,`frozen_balance` ,`available_balance` ,`fee` ,`created_time` ,`updated_time` ,`baned` )values(   ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", one.Uid, one.Name, one.IdCard, one.Status, one.Mobile, one.Password, one.TotalBalance, one.FrozenBalance, one.AvailableBalance, one.Fee, one.CreatedTime, one.UpdatedTime, one.Baned)
}
func (dao *UserDao) UpdateAll(one *User) (res sql.Result, err error) {
	return dao.DB.Exec("update `user` set `uid`=? ,`name`=? ,`id_card`=? ,`status`=? ,`mobile`=? ,`password`=? ,`total_balance`=? ,`frozen_balance`=? ,`available_balance`=? ,`fee`=? ,`created_time`=? ,`updated_time`=? ,`baned`=?  where id = ? ", one.Uid, one.Name, one.IdCard, one.Status, one.Mobile, one.Password, one.TotalBalance, one.FrozenBalance, one.AvailableBalance, one.Fee, one.CreatedTime, one.UpdatedTime, one.Baned, one.Id)
}
func (dao *UserDao) InsertX(one *User) (err error) {
	res, err := dao.Insert(one)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	one.Id = int(id)
	return nil
}
func (dao *UserDao) UpdateWhere(sqlStr string, kv map[string]interface{}, args ...interface{}) (res sql.Result, err error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("update `user` set ")
	i := 0
	length := len(kv)
	var vs []interface{}
	for k, v := range kv {
		i++
		buf.WriteString("`")
		buf.WriteString(k)
		buf.WriteString("`")
		buf.WriteString("=?")

		if i < length {
			buf.WriteString(",")
		}
		vs = append(vs, v)
	}
	vs = append(vs, args...)
	buf.WriteString(" where " + sqlStr)
	return dao.DB.Exec(buf.String(), vs...)
}
func (dao *UserDao) Update(ID int, kv map[string]interface{}) (res sql.Result, err error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("update `user` set ")
	i := 0
	length := len(kv)
	var vs []interface{}
	for k, v := range kv {
		i++
		buf.WriteString("`")
		buf.WriteString(k)
		buf.WriteString("`")
		buf.WriteString("=?")

		if i < length {
			buf.WriteString(",")
		}
		vs = append(vs, v)
	}
	vs = append(vs, ID)
	buf.WriteString(" where id = ?")
	return dao.DB.Exec(buf.String(), vs...)
}
