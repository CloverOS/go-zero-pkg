package captcha

import (
	"bytes"
	"github.com/dchest/captcha"
)

func CreateImage() string {
	captchaId := captcha.NewLen(captcha.DefaultLen)
	return captchaId
}

func Reload(captchaId string) bool {
	return captcha.Reload(captchaId)
}

func Verify(captchaId, val string) bool {
	return captcha.VerifyString(captchaId, val)
}

func GetImageByte(captchaId string) ([]byte, error) {
	var content bytes.Buffer
	err := captcha.WriteImage(&content, captchaId, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		return nil, err
	}
	return content.Bytes(), nil
}
