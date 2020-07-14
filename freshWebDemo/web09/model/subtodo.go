package model

import "webGinDemo/web09/dao"

type Subtodo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`

}

// Subtodo - model的增删改查
func CreateSubtodo(subtodo *Subtodo) (err error) {
	err = dao.DB.Create(&subtodo).Error
	return
}

func GetSubtodoAll() (subtodoAll []*Subtodo, err error)  {
	if err = dao.DB.Find(&subtodoAll).Error; err != nil{
		return nil, err
	}
	return
}

func GetSubtodo(id string) (subtodo *Subtodo,err error) {
	subtodo = new(Subtodo)
	if err = dao.DB.Debug().Where("id=?",id).First(subtodo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateSubtodo(subtodo *Subtodo) (err error) {
	err = dao.DB.Save(subtodo).Error
	return
}

func DeleteSubtodo(id string) (err error) {
	err = dao.DB.Where("id=?",id).Delete(&Subtodo{}).Error
	return
}