package model
import "time"

type User struct {
	Name, Address, Email string
	Age                  int
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}