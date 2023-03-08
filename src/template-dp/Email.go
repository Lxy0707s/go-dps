package template_dp

import (
	"fmt"
	"strings"
)

type Email struct {
	name       string // 消息模板名称
	msgContext string // 消息内容
	MsgTmp            // 接口

	tmpCache map[string]string // 消息模板缓存
}

const EmailTemplateName = "EMAIL"

func (s *Email) genMSG(tmpName string) {
	s.msgContext = ">>模板初始内容<<"
	if _, ok := emailTmp[s.name]; ok {
		s.msgContext = emailTmp[s.name]
	}
	fmt.Printf("%s: msg template generating templated \n", EmailTemplateName)
	return
}

func (s *Email) setMSG(info *MsgInfo) {
	strs := strings.Split(s.msgContext, "%s")
	switch len(strs) {
	case 2:
		s.msgContext = fmt.Sprintf(s.msgContext, info.Name)
	case 3:
		s.msgContext = fmt.Sprintf(s.msgContext, info.Name, info.Address)
	default:
		s.msgContext = ""
	}
	return
}

func (s *Email) getMSG(tp string) string {
	return s.msgContext
}

func (s *Email) sendMSG() error {
	fmt.Printf("%s: sending msg:\n", EmailTemplateName)
	fmt.Print("--------------------start-------------------------")
	fmt.Printf("%s\n", s.msgContext)
	fmt.Println("---------------------end--------------------------")
	return nil
}
