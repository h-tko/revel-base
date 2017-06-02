package models

import (
	"github.com/jinzhu/gorm"
)

// 一応外部キーをタグに設定してるけど、この内容ならなくても動く
type User struct {
	gorm.Model
	UserParam   UserParam `gorm:"ForeignKey:UserID"`
	Name        string
	MailAddress string
	Password    string
}

func (User) TableName() string {
	return "users"
}

func NewUser() *User {
	return new(User)
}

// 1レコードだけ取る場合のサンプル.
//
// func (user *User)にして自身のプロパティに設定もできるけど、わかりづらくなるので一律戻り値として返すようにする
func (*User) One(mailAddress, password string) (result *User) {
	db.Select("id, name").
		Where("mail_address = ?", mailAddress).
		Where("password = ?", password).
		First(&result)

	return
}

// Joinを使った場合のサンプル.
//
// これだとstructの中を平で持たないとマッピングしてくれないので辛い
//
// 使わないやつ
func (*User) AllJoinSample() (users []*User) {
	db.Select("users.id, name, user_params.stamina").Joins("inner join user_params on users.id = user_params.user_id").Find(&users)

	return
}

// gorm.Preloadを使った場合のサンプル
//
// 結合先のテーブルがJoinで必要だったりInner結合の場合はJoinsも指定が必要（無駄に検索されちゃうので）
//
// こっちを使いましょう
func (*User) AllPreloadSample() (users []*User) {
	db.Select("users.id, name").Joins("inner join user_params on users.id = user_params.user_id").Preload("UserParam").Find(&users)

	return
}
