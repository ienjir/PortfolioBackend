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
	if err != nil || {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid starttime"})
	}

	endTime, err := strconv.Atoi(strEndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endtime'"})
		return
	}

	breakTime, err := strconv.Atoi(strBreak)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid break time"})
	}

	sum := a + b

	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
