package form

type Good struct {
	Id int `form:"id" binding:"required"`
}
