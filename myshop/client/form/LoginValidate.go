package form

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
