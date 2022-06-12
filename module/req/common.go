package req

type IdReq struct {
	Id uint `uri:"id" binding:"required"`
}
