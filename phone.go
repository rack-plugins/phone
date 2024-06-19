package phone

import (
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/xluohome/phonedata"
)

func PhoneInfo(c *gin.Context) {
	phoneNumber := c.Param("phoneNumber")

	pr, err := phonedata.Find(phoneNumber)
	if err != nil {
		ezap.Errorf("查询手机号[%s]信息失败: %s", phoneNumber, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"result": "500", "message": "查询手机号信息失败, 请检查输入是否正确"})
		return
	}
	ezap.Infof("查询手机号信息: %s, %-v", phoneNumber, pr)
	c.JSON(http.StatusOK, pr)
}
