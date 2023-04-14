package app

type BirdData struct {
	Idx 	 string `json:"index" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}


type BirdSearchKey struct {
	Key 	 *string `form:"key"`
}