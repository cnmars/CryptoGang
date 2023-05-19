package service

import (
	"errors"
	"fmt"
	request "golddigger/app/system/models/request/common"
	system "golddigger/app/system/models"
	"golddigger/global"

	"gorm.io/gorm"
)

type FeedbackService struct{}
func (fbService *FeedbackService) CreateFeedback(fb system.FeedBack) (err error) {
	// if !errors.Is(global.GD_DB.Where("address = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("存在相同api")
	// }
	return global.GD_DB.Create(&fb).Error
}

func (fbService *FeedbackService) GetAllFeedback() (fbs []system.FeedBack, err error) {
	err = global.GD_DB.Find(&fbs).Error
	return
}
func  (fbService *FeedbackService) SetReaded(id int) (err error) {
	err = global.GD_DB.Where("id = ?", id).First(&system.FeedBack{}).Update("status", 1).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                      // api记录不存在
		return err
	}
	return nil
}
func  (fbService *FeedbackService) DeleteById(id int) (err error) {
	var entity system.FeedBack
	err = global.GD_DB.Where("id = ?", id).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                      // api记录不存在
		return err
	}
	err = global.GD_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}
func (fbService *FeedbackService) GetFeedbackList(fb system.FeedBack, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GD_DB.Model(&system.FeedBack{})
	var fbList []system.FeedBack

	if fb.Topic != "" {
		db = db.Where("topic LIKE ?", "%"+fb.Topic+"%")
	}

	if fb.Type != 0 {
		db = db.Where("type = ?", fb.Type)
	}
	if fb.Status != 0 {
		db = db.Where("status = ?", fb.Status)
	}


	err = db.Count(&total).Error

	if err != nil {
		return fbList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["created_at"] = true
			orderMap["topic"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't matched any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return fbList, total, err
			}

			err = db.Order(OrderStr).Find(&fbList).Error
		} else {
			err = db.Order("created_at desc").Find(&fbList).Error
		}
	}
	return fbList, total, err
}
