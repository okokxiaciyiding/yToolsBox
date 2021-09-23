package dao

import (
	"main/dto"
	"main/model"
	"main/service"

	"main/util"

	"gorm.io/gorm"
)

func SelectAllUser(db *gorm.DB, query string, pagenum int, pagesize int, param string) ([]map[string]interface{}, int) {
	struct_userList := []dto.UserDTO{}
	map_userList := []map[string]interface{}{}
	// 不带Query，返回全部
	// 否则返回like搜索后的结果
	if query == "" {
		db.Order("id").Model(&model.User{}).Find(&struct_userList)
	} else {
		db.Order("id").Where("username LIKE ?", "%"+query+"%").Model(&model.User{}).Find(&struct_userList)
	}

	DefaultLength := len(struct_userList)

	// 把一个自定义结构体的array 转换成map的array
	// 这里用了json的方法 虽然效率低 但是解决了返回给前端大小写的问题
	for i := 0; i < len(struct_userList); i++ {
		map_item := util.Struct2MapViaJson(struct_userList[i])
		map_userList = append(map_userList, map_item)
	}

	// 计算一下需要如何切割数组
	ArrayStart, ArrayEnd := service.CalculateReturnMapLength(pagenum, pagesize, map_userList)
	// 返回切片后的结果
	return map_userList[ArrayStart:ArrayEnd], DefaultLength
}

func SelectSpecifiedUser(db *gorm.DB, userID int) dto.UserDTO {
	struct_userList := dto.UserDTO{}
	db.Model(&model.User{}).Where("id = ?", userID).Find(&struct_userList)
	return struct_userList
}

func DeleteSpecifiedUser(db *gorm.DB, userID int) model.User {
	struct_user := model.User{}
	db.Delete(&struct_user, userID)
	return struct_user
}

func UpdateUserState(db *gorm.DB, mgstate string, userID int) {
	db.Model(&model.User{}).Where("id = ?", userID).Update("mgstate", mgstate)
}

func UpdateSpecifiedUser(db *gorm.DB, userID int, email string, mobile string) {
	db.Model(&model.User{}).Where("id = ?", userID).Updates(map[string]interface{}{"email": email, "mobile": mobile})
}

func InsertNewUser(db *gorm.DB, username string, password string, mobile string, email string, worknum string) (model.User, *gorm.DB) {
	newUser := model.User{UserName: username, PassWord: password, Mobile: mobile, Email: email, WorkNum: worknum}
	result := db.Create(&newUser)

	return newUser, result
}
