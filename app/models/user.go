package models

import (
	"github.com/jinzhu/gorm"
    "regexp"
    "github.com/revel/revel"
)

type User struct{
    gorm.Model
    Name string `gorm:"size:100"`
    Username string `gorm:"size:100;unique"`
    Password []byte
}

var userRegex = regexp.MustCompile("^\\w*$")


func (user *User) Validate(v *revel.Validation){
    v.Check(
        user.Name,
        revel.Required{},
        revel.MinSize{2},
        revel.MaxSize{100},
    )

    v.Check(
        user.Username,
        revel.Required{},
        revel.MinSize{4},
        revel.MaxSize{100},
        revel.Match{userRegex},
    )
}