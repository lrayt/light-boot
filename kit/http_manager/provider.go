package http_manager

// RouterHandler 路由处理器
type RouterHandler interface{}

type RouterGroup interface {
	Get(route string, handler RouterHandler)
	Post(router string, handler RouterHandler)
}

type HttpProvider interface {
	Group(name string) RouterGroup
	Run()
}
