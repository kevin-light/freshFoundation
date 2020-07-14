package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)


// 默认值 gorm.Model 定义
//type Model struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time
//}

// UserInfo --> 数据结构
type GpUser struct {
	Name         string
	Age 		 int
	NameSv 	 	 sql.NullString `gorm:"default:'小王子'"` // sql.NullString 实现了Scanner/Valuer接口	Age          int64
	Birthday     *time.Time
	//Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	//MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
	gorm.Model

}
//type Userdemo1 struct {
//	ID   string // 名为`ID`的字段会默认作为表的主键
//	Name string
//}

//type Animal struct {
//	AnimalID int64 `gorm:"primary_key"`		  // 使用`AnimalID`作为主键     +`gorm:"column:beast_id"`  // set column name to `beast_id`
//	Name     string
//	Age      int64
//}
//func (Animal) TableName() string {			// 将 Animal 的表名设置为 `Animal_change`
//	return "Animal_change"
//}
func main() {
	db,err:=gorm.Open("mysql","root:123123@(127.0.0.1:3306)/dtdba?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		panic(err)
	}
	defer db.Close()
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "gp_" + defaultTableName;		// GORM还支持更改默认表名称规则：
	//
	//}

	db.AutoMigrate(&GpUser{})	// 创建表 自动迁移（把数据结构体和数据表对应）
	//u1 := UserInfo{1,"小明","男","游泳"}	// 创建数据行
	//db.SingularTable(true)  	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.Table("Animal_users").CreateTable(&Animal{})  	// 使用User结构体创建名为`deleted_users`的表
	//var deleted_users []User
	//db.Table("deleted_users").Find(&deleted_users)			//// SELECT * FROM deleted_users;
	//db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()  	//// DELETE FROM deleted_users WHERE name = 'jinzhu';

	//var user = GpUser{Name: "",Age: 18}
	//db.Create(&user)		// 实际执行的SQL语句是INSERT INTO users("age") values('99');，排除了零值字段Name;,如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口，如下：
	//var nameSv = GpUser{NameSv: sql.NullString{"",true},Age: 17}
	//db.Create(&nameSv)

	var user []GpUser
	//db.First(&user)
	//fmt.Println("根据主键查询第一条记录:",user)
	//db.Last(&user)
	//fmt.Println("根据主键查询的最后一条记录",user)
	//db.Find(&user)
	//fmt.Println("查询所有记录",user)
	//db.First(&user,10)		// 查询指定的某条记录(仅当主键为整型时可用)//// SELECT * FROM users WHERE id = 10;
	//db.Where("name=?","大").First(&user)
	//db.Where("name=?","大").Find(&user)
	//db.Where("name<>?","大").Find(&user)
	//db.Where("name IN (?)",[]string{"大"}).Find(&user)
	//db.Where("name LIKE ?","%中%").Find(&user)
	//db.Where("name=? AND age>=?","大","18").Find(&user)
	//db.Where("create_at between ? and ?","lastWeek","today").Find(&user)
	//
	//db.Not("name", "xxx").First(&user)	//// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;
	//fmt.Println("not first:",user)
	//db.Not("name = ?", "xxx").First(&user)
	//fmt.Println("plain sql ",user)
	//db.Not("name",[]string{"中","xxx"}).Find(&user)

	//db.Where("name=?","大").Or("name=?","中").Find(&user)
	//// 根据主键获取记录 (只适用于整形主键)
	//db.First(&user, 23)
	////// SELECT * FROM users WHERE id = 23 LIMIT 1;
	//// 根据主键获取记录, 如果它是一个非整形主键
	//db.First(&user, "id = ?", "string_primary_key")
	////// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;
	//// Plain SQL
	//db.Find(&user, "name = ?", "jinzhu")
	////// SELECT * FROM users WHERE name = "jinzhu";
	//db.Find(&user, "name <> ? AND age > ?", "jinzhu", 20)
	//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	db.First(&user)
	db.Model(&user).Update("name", "hello") 	// 更新单个属性，如果它有变化	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Where("active = ?", true).Update("name", "hello")	// 根据给定的条件更新单个属性	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})	// 使用 map 更新多个属性，只会更新其中有变化的属性	//// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Updates(User{Name: "hello", Age: 18})	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段	//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	db.Model(&user).Updates(User{Name: "", Age: 0, Active: false})
	//使用SQL表达式更新  //先查询表中的第一条数据保存至user变量。
	db.First(&user)
	db.Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 100))
	//// UPDATE `users` SET `age` = age * 2 + 100, `updated_at` = '2020-02-16 13:10:20'  WHERE `users`.`id` = 1;
	db.Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 100)})
	//// UPDATE "users" SET "age" = age * '2' + '100', "updated_at" = '2020-02-16 13:05:51' WHERE `users`.`id` = 1;
	db.Model(&user).UpdateColumn("age", gorm.Expr("age - ?", 1))
	//// UPDATE "users" SET "age" = age - 1 WHERE "id" = '1';
	db.Model(&user).Where("age > 10").UpdateColumn("age", gorm.Expr("age - ?", 1))
	//// UPDATE "users" SET "age" = age - 1 WHERE "id" = '1' AND quantity > 10;





	//fmt.Println("db end......")

	//db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)	// 	例如PostgreSQL数据库中可以使用下面的方式实现合并插入, 有则更新, 无则插入。 // INSERT INTO products (name, code) VALUES ("name", "code") ON CONFLICT;

}
