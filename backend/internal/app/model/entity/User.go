package entity

type User struct {
	Id            *int
	Email         *string
	Name          *string
	Password      *string
	IsDeactivated *bool
}
