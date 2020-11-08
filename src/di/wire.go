//+build wireinject

package di

import (
	"clean-architecture-go/application/usecase"
	fs "clean-architecture-go/infrastructure/persistent/firestore"
	"clean-architecture-go/infrastructure/waf/echo/handler"
	"clean-architecture-go/interface/controller"
	jsonp "clean-architecture-go/interface/presenter/json"
	"github.com/google/wire"
)

func InitializeTaskHandler() handler.ITaskHandler {
	wire.Build(
		handler.NewTaskHandler,
		controller.NewTaskController,
		usecase.NewTaskUseCase,
		jsonp.NewTaskJsonPresenter,
		fs.NewTaskFirestoreRepository,
	)
	return &handler.TaskHandler{}
}
