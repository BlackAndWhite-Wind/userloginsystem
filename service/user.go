package service

import (
	"UserLoginSystem/config"
	"UserLoginSystem/model"
	"UserLoginSystem/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var otpStore = make(map[string]string) // 简单的内存存储方式

func SendOTPByEmail(email string) error {
	otp := utils.GenerateOTP(6)
	otpStore[email] = otp
	subject := "Your OTP Code"
	body := fmt.Sprintf("Your OTP code is %s and is valid for 5 minutes.", otp)
	err := utils.SendEmail(email, subject, body)
	if err != nil {
		return err
	}
	go func() {
		time.Sleep(5 * time.Minute)
		delete(otpStore, email)
	}()
	return nil
}

func VerifyOTPByEmail(email, otp string) error {
	storedOTP, exists := otpStore[email]
	if !exists || storedOTP != otp {
		return errors.New("invalid or expired OTP")
	}
	delete(otpStore, email)
	return nil
}

func SendOTPByPhone(phone string) error {
	otp := utils.GenerateOTP(6)
	otpStore[phone] = otp
	message := fmt.Sprintf("Your OTP code is %s and is valid for 5 minutes.", otp)
	err := utils.SendSms(phone, message)
	if err != nil {
		return err
	}
	go func() {
		time.Sleep(5 * time.Minute)
		delete(otpStore, phone)
	}()
	return nil
}

func VerifyOTPByPhone(identifier, otp string) error {
	storedOTP, exists := otpStore[identifier]
	if !exists || storedOTP != otp {
		return errors.New("invalid or expired OTP")
	}
	delete(otpStore, identifier)
	return nil
}

func RegisterUser(username, email, phone, password string) error {
	if !utils.IsValidEmail(email) {
		return errors.New("invalid email format")
	}
	if !utils.IsValidPhone(phone) {
		return errors.New("invalid phone format")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		UserName:    username,
		Email:       email,
		PhoneNumber: phone,
		Password:    string(hashedPassword),
	}
	return config.DB.Create(&user).Error
}

func AuthenticateUserByUsername(username, password string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func ChangePassword(userID uint, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return config.DB.
		Model(&model.User{}).
		Where("id = ?", userID).
		Update("password", string(hashedPassword)).
		Error
}
