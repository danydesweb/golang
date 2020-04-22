package main
import (
 "fmt"
 "log"
 "time"
"github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/postgres"
 "github.com/satori/go.uuid"
)
// Base contains common columns for all tables.
type Base struct {
 ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt *time.Time `sql:"index"`
}
// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
 uuid, err := uuid.NewV4()
 if err != nil {
  return err
 }
 return scope.SetColumn("ID", uuid)
}
// User is the model for the user table.
type User struct {
 Base
 SomeFlag bool    `gorm:"column:some_flag;not null;default:true"`
 Profile  Profile
}
// Profile is the model for the profile table.
type Profile struct {
 Base
 Name   string    `gorm:"column:name;size:128;not null;"`
 UserID uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
}
func main() {
 db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=	Sty59HjeuNLpFjjhRA5HNok1gHc58lVs")
 if err != nil {
  panic(err)
 }
db.LogMode(true)
 db.AutoMigrate(&User{}, &Profile{})
user := &User{SomeFlag: false}
 if db.Create(&user).Error != nil {
  log.Panic("Unable to create user.")
 }
profile := &Profile{Name: "New User", UserID: user.Base.ID}
 if db.Create(&profile).Error != nil {
  log.Panic("Unable to create profile.")
 }
fetchedUser := &User{}
 if db.Where("id = ?", profile.UserID).Preload("Profile").First(&fetchedUser).RecordNotFound() {
  log.Panic("Unable to find created user.")
 }
fmt.Printf("User: %+v\n", fetchedUser)
}