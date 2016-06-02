package router

var Handlers *HandlerManager

func init() {
	Handlers = &HandlerManager{}
}

type HandlerManager struct {
	Handlers []LogHandler
}

func (h *HandlerManager) Add(lh LogHandler) {
	h.Handlers = append(h.Handlers, lh)
}
