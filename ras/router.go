package ras

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func RASRouter(mail_channel chan mail.Mail, r *gin.Engine) {
	api := r.Group("/api/ras")
	{
		api.GET("", HelloWorldController)
		api.GET("/programs", getPrograms)
		api.GET("/departments", getDepartments)
		api.GET("/program-departments", getProgramsDepartments)
		api.GET("/testmail", MailController(mail_channel))
	}
}
