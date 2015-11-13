package main
import "time"
type Student struct {
	Name, Address, Class string
	Age                  int
	Gender               bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}