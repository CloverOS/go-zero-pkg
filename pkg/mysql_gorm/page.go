package mysql_gorm

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Page interface {
	GetCount() int
	GetPage() int
}

type DefaultPage struct {
	Count int `json:"size"` //大小
	Page  int `json:"page"` //页码
}

type PageData[T any] struct {
	Total int64       `json:"total"` //总数
	Page  DefaultPage `json:"page"`  //页数信息
	Data  T           `json:"data"`  //数据
}

func (p PageData[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func GetPageData[T any, M any](db *gorm.DB, page Page, isScan bool) (pagedata PageData[T], err error) {
	pagedata.Page = DefaultPage{
		Count: page.GetCount(),
		Page:  page.GetPage(),
	}
	var model M
	temp := db.Model(&model)
	err = temp.Count(&pagedata.Total).Error
	if err != nil {
		return pagedata, err
	}
	offset := (page.GetPage() - 1) * page.GetCount()
	temp.Offset(offset).Limit(page.GetCount())
	if isScan {
		err = temp.Scan(&pagedata.Data).Error
	} else {
		err = temp.Find(&pagedata.Data).Error
	}
	if err != nil {
		return pagedata, err
	}
	return pagedata, err
}

func GetPageDataNilModel[T any](db *gorm.DB, page Page, isScan bool) (pagedata PageData[T], err error) {
	pagedata.Page = DefaultPage{
		Count: page.GetCount(),
		Page:  page.GetPage(),
	}
	temp := db
	err = temp.Count(&pagedata.Total).Error
	if err != nil {
		return pagedata, err
	}
	offset := (page.GetPage() - 1) * page.GetCount()
	temp.Offset(offset).Limit(page.GetCount())
	if isScan {
		err = temp.Scan(&pagedata.Data).Error
	} else {
		err = temp.Find(&pagedata.Data).Error
	}
	if err != nil {
		return pagedata, err
	}
	return pagedata, err
}

func FindPageData[T any, M any](db *gorm.DB, page Page) (pagedata PageData[T], err error) {
	return GetPageData[T, M](db, page, false)
}

func ScanPageData[T any, M any](db *gorm.DB, page Page) (pagedata PageData[T], err error) {
	return GetPageData[T, M](db, page, true)
}

func FindPageDataNilModel[T any](db *gorm.DB, page Page) (pagedata PageData[T], err error) {
	return GetPageDataNilModel[T](db, page, false)
}

func ScanPageDataNilModel[T any](db *gorm.DB, page Page) (pagedata PageData[T], err error) {
	return GetPageDataNilModel[T](db, page, true)
}
