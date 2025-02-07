package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// Task Model
type Task struct {
	Id          int       `orm:"auto;pk"`                     // Primary Key
	Title       string    `orm:"size(100);not null"`          // Task Title
	Description string    `orm:"type(text);null"`             // Task Description
	Category    *Category `orm:"rel(fk);on_delete(cascade)"`  // Foreign Key to Category
	Status      string    `orm:"size(20);default(TODO)"`      // Enum-like status
	Priority    string    `orm:"size(10);default(MEDIUM)"`    // Enum-like priority
	DueDate     time.Time `orm:"null"`                        // Due Date
	CompletedAt time.Time `orm:"null"`                        // Completion Timestamp
	CreatedBy   *User     `orm:"rel(fk)"`                     // Foreign Key to User
	AssignedTo  *User     `orm:"rel(fk);null"`                // Task Assigned To
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"` // Creation Timestamp
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`     // Update Timestamp
}

// Register Model
func init() {
	orm.RegisterModel(new(Task))
}
