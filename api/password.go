package api

import (
	"github.com/gin-gonic/gin"
)

type PasswordApi struct {
}

func (p *PasswordApi) Save(c *gin.Context) {
	//var saveInfoReq req.SavePassword
	//if err := c.ShouldBindJSON(&saveInfoReq); err != nil {
	//	c.JSON(http.StatusBadRequest, callback.BackFail("参数错误"))
	//	return
	//}
	//userId := User.GetUserId(c)
	//passwd := &entity.Password{}
	//if isSuccessful := service.Password.Save(passwd, userId); isSuccessful {
	//	c.JSON(http.StatusOK, callback.Success())
	//	return
	//}
	//c.JSON(http.StatusBadRequest, callback.BackFail("保存失败"))
}

func (p *PasswordApi) QueryById(c *gin.Context) {
	//id := c.Param("id")
	//userId := User.GetUserId(c)
	//password := service.QueryPasswordById(id, userId)
	//c.JSON(http.StatusOK, callback.Success(password))
}

func (p *PasswordApi) QueryList(c *gin.Context) {
	//userId := User.GetUserId(c)
	//var reqParam req.QueryPasswd
	//if err := c.ShouldBindQuery(&reqParam); err != nil {
	//	log.Printf("参数绑定错误,原因=%v", err)
	//	c.JSON(http.StatusOK, callback.BackFail("参数错误"))
	//	return
	//}
	//c.JSON(http.StatusOK, callback.Success(service.QueryList(reqParam, userId)))
}

func (p *PasswordApi) DeleteById(c *gin.Context) {
	//userId := User.GetUserId(c)
	//id := c.Param("id")
	//service.DeleteById(id, userId)
	//c.JSON(http.StatusOK, callback.Success(service.DeleteById(id, userId)))
}

func (p *PasswordApi) SearchPasswdList(c *gin.Context) {
	//userId := User.GetUserId(c)
	//var query req.QueryPasswd
	//err := c.ShouldBindQuery(&query)
	//if err != nil {
	//	log.Printf("参数绑定,原因=%v", err.Error())
	//	c.JSON(http.StatusOK, callback.BackFail("参数错误"))
	//	return
	//}
	//c.JSON(http.StatusOK, callback.Success(service.SearchList(query, userId)))
}
