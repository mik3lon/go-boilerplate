package kernel

import "go-boilerplate/internal/pkg/domain/bus"

type CQRSModuleInterface interface {
	AddCommand(cb bus.CommandBus, c bus.Command, handler bus.CommandHandler)
	AddQuery(qb bus.QueryBus, q bus.Query, handler bus.QueryHandler)
}

type CQRSModule struct {
}

func (cqrsm *CQRSModule) AddCommand(cb bus.CommandBus, c bus.Command, handler bus.CommandHandler) {
	err := cb.RegisterCommand(c, handler)
	if err != nil {
		panic(err)
	}
}

func (cqrsm *CQRSModule) AddQuery(qb bus.QueryBus, q bus.Query, handler bus.QueryHandler) {
	err := qb.RegisterQuery(q, handler)
	if err != nil {
		panic(err)
	}
}
