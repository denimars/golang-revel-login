package models

import (
    "github.com/jinzhu/gorm"
    "github.com/revel/revel"
)

type Book struct{
    gorm.Model
    Title string `gorm:"size:100"`
    Description string `gorm:"size:255"`
}

func (book *Book) Validate(v *revel.Validation){
    v.Check(
        book.Title,
        revel.MinSize{1},
        revel.MaxSize{100},
    )

    v.Check(
        book.Description,
        revel.MinSize{4},
        revel.MaxSize{255},
    )
}