package template_dp

type Message interface {
	genMSG(string)
	setMSG(info *MsgInfo)
	getMSG(string) string
	sendMSG() error
}

type (
	MsgTmp struct {
		Message
	}
	MsgInfo struct {
		Name    string
		Address string
	}
)

func (m *MsgTmp) genAndSendOTP(tmpName string, info *MsgInfo) error {
	m.genMSG(tmpName)
	m.setMSG(info)
	err := m.sendMSG()
	if err != nil {
		return err
	}
	return nil
}
