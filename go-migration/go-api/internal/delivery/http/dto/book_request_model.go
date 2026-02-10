package dto

type BookModel struct {
	Id     int    `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,gte=1000,lte=2026"`
	ISBN   string `json:"isbn" binding:"required,len=10|len=13"`
}
