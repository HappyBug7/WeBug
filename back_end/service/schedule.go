package service

import "back_end/model"

type Schedule struct {
}

func (s *Schedule) Class_add(class model.Class) error {
	if err := model.DB.Model(&model.Class{}).Create(&class).Error; err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Class_broadcast(day int) (error, []model.Class) {
	classes := []model.Class{}
	if err := model.DB.Model(&model.Class{}).Where("day = ?", day).Find(&classes).Error; err != nil {
		return err, nil
	}
	return nil, classes
}

func (s *Schedule) Class_get(day int, class int) model.Class {
	empty_class := model.Class{}
	empty_class.Class = class
	empty_class.Day = day
	class_ := model.Class{}
	if err := model.DB.Model(&model.Class{}).Where("day = ? AND class = ?", day, class).Find(&class_).Error; err != nil {
		return empty_class
	}
	return class_
}
