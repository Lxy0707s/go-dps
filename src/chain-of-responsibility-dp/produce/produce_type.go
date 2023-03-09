package produce

const (
	Strict = "Strict"
	Normal = "Normal"
)

// FactoryProcess 生产流程接口
type FactoryProcess interface {
	Execute(*Material)                                                // 开始处理
	qcQa(*Material) bool                                              // 质检
	produce(*Material)                                                // 生产
	withModel(model string)                                           // 生产模式
	NextProcess(fac FactoryProcess, strictType string) FactoryProcess // 交付
	End()                                                             // 结束流程
}

// Material 产品原型/原材料
type Material struct {
	Manufacturer    string  // 厂家
	Models          string  // 型号
	Price           float64 // 原材料价格
	TCO             float64 // 生产成本，每一个流程附加一部分
	Name            string  // 名称
	Description     string  // 描述
	QuarantineLabel bool    // 质检标识，判断是否废品
}

/**
  生产流程： 采购检验-加工生产-包装-交付发货
  标准：严格流程，流程不能跳过，某个原材料，某一流程质量不达标则放入废品仓库
*/

// SetQuarantineLabel 设置质检标识
func (m *Material) SetQuarantineLabel(label bool) {
	m.QuarantineLabel = label
}

// SetDescription 设置生产成本
func (m *Material) SetDescription(description string) {
	m.Description = description
}

// SetTCO 设置描述
func (m *Material) SetTCO(tco float64) {
	m.TCO = m.TCO + tco
}
