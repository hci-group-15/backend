package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserProfileResponse struct {
	Username string `json:"username" example:"My Nice Username"`
}

// Request handler to get a user's profile
//
// @Summary		Show user profile
// @Description	Get user's profile. Includes username
// @Tags		user
// @Produce		json
// @Success		200	{object}		UserProfileResponse
// @Failure		401	{object}		nil
// @Failure		500	{object}		httputil.HTTPError
// @Router		/user/profile [get]
func (*Controller) GetUserProfile(c *gin.Context) {
	ses := sessions.Default(c)
	username, ok := ses.Get("username").(string)
	if !ok {
		c.AbortWithStatus(401)
		return
	}
	c.JSON(200, UserProfileResponse{Username: username})
}
