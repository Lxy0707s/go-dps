package template_dp

var (
	defaultMsg = "\n尊敬的%s：\n" +
		"您好，我公司近期即将推出新的产品/服务，诚挚邀请你通过官网:\n" +
		"www.xxx.com,在本月底之前报名加入内测，有惊喜礼包等着你哦。\n" +
		"                                 2023年2月28日\n" +
		"                                   xxx公司"
	welcomeMsg = "\n尊敬的%s: \n" +
		"    欢迎使用本公司提供的服务，诚挚邀请您于2023年3月1日参加我司举办的\n" +
		"线下交流会，届时将对本相关产品退出更新服务和优化，敬请光临\n" +
		"                                        2023年2月28日\n" +
		"                                          xxx公司"
	postMsg = "\n尊敬的%s: \n" +
		"    感谢购买本公司的产品，您在收货后如有任何不满意或者产品使用、改进意见，\n " +
		"请及时致电官方客户热线：000851-8888，或者关注公众号：你好呀，咨询客户，以\n " +
		"便保障您的合法权益\n" +
		"                                          2023年2月28日\n" +
		"                                          %s"
	emailTmp = map[string]string{
		"default":     defaultMsg,
		"welcome-tmp": welcomeMsg,
		"post-tmp":    postMsg,
	}
)
