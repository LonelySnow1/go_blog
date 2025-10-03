package service

/*
 为什么要这样设计：
	1. 服务集中管理（Service Group 模式）：将相关业务集成到服务组中，方便管理
	2. 单例复用，避免资源浪费：ServiceGroupApp是一个全局单例，内部的示例会被复用
	3. 代码分层与职责边界：通过入口调用服务，而不是直接依赖具体的服务实现。 符合「依赖倒置原则」,方便后续对具体实现进行修改
*/

type ServiceGroup struct {
	EsService
	BaseService
}

var ServiceGroupApp = new(ServiceGroup)
