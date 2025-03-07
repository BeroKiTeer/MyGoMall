package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
	"user/biz/dal/redis"
)

type User struct {
	Base
	Username       string
	PasswordHashed string
	PhoneNumber    string `gorm:"unique"`
	Email          string `gorm:"unique"`
	AddressId      int64
	Status         int8
	Role           int8
}

func (u User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

// CreateUser 创建一条新的用户记录
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// DeleteUser 删除一条用户记录
func DeleteUser(db *gorm.DB, user *User) error {
	return db.Delete(user).Error
}

func GetUserById(db *gorm.DB, ctx context.Context, id int32) (user *User, err error) {
	err = db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

// UserExistsByID 检查用户是否存在
func UserExistsByID(db *gorm.DB, userID int32) (bool, error) {
	var exists bool
	err := db.Raw("SELECT EXISTS(SELECT 1 FROM user WHERE id = ?)", userID).
		Scan(&exists).
		Error

	if err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, ctx context.Context, userID int32, updates map[string]interface{}) error {
	// 1️⃣ 检查用户是否存在
	var user User
	err := db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}

	// 2️⃣ 执行更新操作
	err = db.WithContext(ctx).Model(&user).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func GetCachedUserById(ctx context.Context, db *gorm.DB, ID int32) (data string, err error) {
	cachedUser, err := redis.RedisClusterClient.Get(ctx, fmt.Sprintf("UserId:%d", ID)).Result()
	if errors.Is(err, redis.Nil) {
		// 如果缓存没有命中，从数据库获取
		row, err := GetUserById(db, ctx, ID)
		if err != nil {
			klog.Error(err)
			return "", err
		}

		//将查询到的数据存入缓存
		cacheData := row
		//序列化后存储到redis里
		UserJson, _ := json.Marshal(cacheData)
		if err := redis.RedisClusterClient.Set(
			ctx,
			fmt.Sprintf("UserId:%d", row.ID),
			UserJson,
			1*time.Hour).Err(); err != nil {
			klog.Error("用户缓存失败:", err)
		}
		cachedUser = string(UserJson)
	} else if err != nil {
		klog.Error(err)
	}
	return cachedUser, nil
}

func GetCachedUserByEmail(ctx context.Context, db *gorm.DB, email string) (data string, err error) {
	cachedUser, err := redis.RedisClusterClient.Get(ctx, fmt.Sprintf("email:%s", email)).Result()
	if errors.Is(err, redis.Nil) {
		// 如果缓存没有命中，从数据库获取
		row, err := GetByEmail(db, ctx, email)
		if err != nil {
			klog.Error(err)
			return "", err
		}

		//将查询到的数据存入缓存
		cacheData := row
		//序列化后存储到redis里
		UserJson, _ := json.Marshal(cacheData)
		if err := redis.RedisClusterClient.Set(
			ctx,
			fmt.Sprintf("email:%s", email),
			UserJson,
			1*time.Hour).Err(); err != nil {
			klog.Error("用户缓存失败:", err)
		}
		cachedUser = string(UserJson)
	} else if err != nil {
		klog.Error(err)
	}
	return cachedUser, nil
}

func GetCachedAddressById(ctx context.Context, db *gorm.DB, AddressId int64) (data string, err error) {
	addressStr, err := redis.RedisClusterClient.Get(ctx, fmt.Sprintf("AddressId:%d", AddressId)).Result()
	//若redis中不存在，则从mysql中获取
	if errors.Is(err, redis.Nil) {
		//从数据库中查询
		address, err := GetAddressById(db, ctx, AddressId)
		if err != nil {
			klog.Error("查询失败:", err)
		}
		if address == nil {
			address = &Address{
				Address: "",
			}
		}
		// 序列化为JSON字符串存储(存储结构体时的策略)
		addressJSON, _ := json.Marshal(address)
		addressStr = string(addressJSON)
		if err := redis.RedisClusterClient.Set(
			ctx,
			fmt.Sprintf("AddressId:%d", AddressId),
			addressJSON, // 存储序列化后的JSON
			24*time.Hour).Err(); err != nil {
			klog.Error("地址缓存失败:", err)
		}
	} else if err != nil {
		klog.Error(err)
		return "", err
	}

	return addressStr, nil
}
