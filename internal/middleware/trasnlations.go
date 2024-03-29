package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_HK"
	"github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_HK.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_trans.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_trans.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_trans.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
