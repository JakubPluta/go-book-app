package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Books []*Book

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageUrl      string
	Description   string
}

type BookOuts []*BookOut

type BookOut struct {
    ID            uint   `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedDate string `json:"published_date"`
    ImageUrl      string `json:"image_url"`
    Description   string `json:"description"`


}

type BookForm struct {
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedDate string `json:"published_date"`
    ImageUrl      string `json:"image_url"`
    Description   string `json:"description"`
}

func (f *BookForm) ToModel() (*Book, error) {
    pubDate, err := time.Parse("2022-01-02", f.PublishedDate)
    if err != nil {
        return nil, err
    }
    return &Book{
        Title:         f.Title,
        Author:        f.Author,
        PublishedDate: pubDate,
        ImageUrl:      f.ImageUrl,
        Description:   f.Description,
    }, nil
}


func (b Book) ToDto() *BookOut{
	return &BookOut{
		ID:            b.ID,
        Title:         b.Title,
        Author:        b.Author,
        PublishedDate: b.PublishedDate.Format("2006-01-02"),
        ImageUrl:      b.ImageUrl,
        Description:   b.Description,
	}
}


func (bs Books) ToDto() BookOuts {
    dtos := make([]*BookOut, len(bs))
    for i, b := range bs {
        dtos[i] = b.ToDto()
    }
    return dtos
}