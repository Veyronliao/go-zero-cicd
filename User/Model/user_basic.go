package models

// // 创建用户
// func CreateUser(user *UserBasic) (int, error) {
// 	result := DB.Create(user)
// 	return int(result.RowsAffected), result.Error
// }

// // 删除用户
// func DeleteUser(user *UserBasic) (int, error) {
// 	fmt.Println(user.ID)
// 	result := DB.Delete(&user, user.ID)
// 	fmt.Println(result.RowsAffected)
// 	return int(result.RowsAffected), result.Error
// }

// // 更新user信息，更新所有字段
// func UpdateUser(user *UserBasic) (int, error) {
// 	result := DB.Save(user)
// 	return int(result.RowsAffected), result.Error
// }

// // 修改user密码
// func UpdateUserPassWord(user *UserBasic) (int, error) {
// 	result := DB.Model(&user).Update("PassWord", user.PassWord)
// 	return int(result.RowsAffected), result.Error
// }

// // 获取单个用户信息
// func GetUserInfoById(id uint) *UserBasic {
// 	user := &UserBasic{}
// 	result := DB.First(user, id)
// 	if result.Error != nil {
// 		return nil
// 	}
// 	return user
// }

// //分页查询

// func GetUserListPage(user UserBasic) *gorm.DB {
// 	tx := DB.Debug().Model(new(UserBasic)).Select("Id,Name,PassWord,Gender,Email,Telephone,CreatedAt,UpdatedAt,DeletedAt")
// 	if user.UserName != "" {
// 		tx.Where("UserName like ?", "%"+user.UserName+"%")
// 	}
// 	if user.Email != "" {
// 		tx.Where("Email like ?", "%"+user.Email+"%")
// 	}
// 	if user.Telephone != "" {
// 		tx.Where("Telephone like ?", "%"+user.Telephone+"%")
// 	}
// 	return tx
// }
