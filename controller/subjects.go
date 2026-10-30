package controller

import "github.com/gin-gonic/gin"

// Request handler to get a user's subjects
//
// @Summary 	Get a user's subjects
// @Description	Returns a list of user subjects, without any topics. Those need to be requested separately
// @Tags		subjects
// @Produce		json
// @Success		200	{object}		[]model.Subject
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/user/subjects/get [get]
func (m *Controller) GetUserSubjects(c *gin.Context) {
	c.JSON(200, gin.H{})
}
