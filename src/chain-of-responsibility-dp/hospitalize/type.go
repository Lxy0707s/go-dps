package hospitalize

// Patient 病人
type Patient struct {
	name          string
	receptionDone bool
	checkUpDone   bool
	medicineDone  bool
	paymentDone   bool
}

// Department 处理行为接口
type Department interface {
	execute(*Patient)
	end()
	setNext(Department) Department
}
