package model

type User struct {
	Firstname string     `gorm:"firstname"  validate:"required,max=50,alpha" `
	Lastname  string     `gorm:"lastname" validate:"required,max=50,alpha" `
	Email     string     `gorm:"email"    validate:"required,email"`
	Phone     string     `gorm:"phone" validate:"required,max=20,numeric"`
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Password  string     `gorm:"password" validate:"required,min=6"`
	Status    UserStatus `gorm:"type:int" json:"status"`
	Role      UserRole   `gorm:"type:int" json:"role"`
}
