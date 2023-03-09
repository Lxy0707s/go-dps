package hospitalize

// Patient 病人
type Patient struct {
	name          string
	receptionDone bool
	checkUpDone   bool
	medicineDone  bool
	paymentDone   bool
}

// Process 处理流程 行为接口
type Process interface {
	execute(*Patient)
	end()
	setNext(Process) Process
}
