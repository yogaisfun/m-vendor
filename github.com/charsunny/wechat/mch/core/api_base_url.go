package core

func APIBaseURL() string {
	// TODO(chanxuehong): 后期做容灾功能
	return "https://api.mch.weixin.qq.com"
}

func APIAppURL() string  {
	return "https://payapp.weixin.qq.com"
}