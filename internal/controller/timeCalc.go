package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CalcTime(c *gin.Context) {
	strStartTime := c.Query("start")
	strEndTime := c.Query("end")
	strBreak := c.Query("break")

	startTime, err := strconv.Atoi(strStartTime)
	if err != nil || startTime >= 24 || startTime <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid starttime"})
		return
	}

	endTime, err := strconv.Atoi(strEndTime)
	if err != nil || endTime >= 24 || endTime <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endtime'"})
		return
	}

	breakTime, err := strconv.Atoi(strBreak)
	if err != nil || breakTime >= 24 || breakTime >= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid break time"})
	}

	worktime := 24 - startTime - (-24 + endTime) - breakTime

	c.JSON(http.StatusOK, gin.H{"worktime": worktime})
}
