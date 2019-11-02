package functional

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
	LastLogin time.Time `dorm:"last_login" gorm:"column:last_login;default:CURRENT_TIMESTAMP;not null;" json:"last_login"`

	NickName     string `dorm:"nick_name" gorm:"column:nick_name;unique;not_null"`
	Name         string `dorm:"name" gorm:"column:name;unique;not_null"`
	Password     string `dorm:"password" gorm:"column:password;not_null"`
	Phone        string `dorm:"phone" gorm:"column:phone;unique;not_null"`
	//Rank         string `dorm:"rank" gorm:"column:rank;unique;not_null"`
	RegisterCity string `dorm:"register_city" gorm:"column:register_city;not_null"`
}

func BenchmarkBaseTraits_objectFactory(b *testing.B) {
	var info = NewBaseTraits(User{})
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		_ = info.ObjectFactory()
	}
}

func cccccc() interface{} { return new(User) }

func BenchmarkBaseTraits_objectFactoryPure(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		_ = cccccc()
	}
}

func BenchmarkBaseTraits_objectFactoryPureN(b *testing.B) {
	var New = cccccc
	b.ResetTimer()

	for i := 0; i < b.N; i ++ {
		_ = New()
	}
}

func TestTraits_objectFactory(t *testing.T) {
	var info = NewBaseTraits(User{})
	if a, ok := info.ObjectFactory().(*User); ok {
		fmt.Printf("%T %v\n", a, a)
	} else {
		t.Errorf("bad creation %T", a)
	}
}

func TestTraits_sliceFactory(t *testing.T) {
	var info = NewBaseTraits(User{})
	if a, ok := info.SliceFactory().([]User); ok {
		fmt.Printf("%T %v\n", a, a)
	} else {
		t.Errorf("bad creation %T", a)
	}
}