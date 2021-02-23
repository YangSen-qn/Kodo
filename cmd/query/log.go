package query

import (
	"context"
	"github.com/YangSen-qn/cmd"
)

type queryLogCMDHandler struct {
	isTime bool
	isUser bool
}

func configLogCMD(superCMD *cmd.Command) {

	var handler = &queryLogCMDHandler{}

	queryCMDBuilder := &cmd.CommandBuilder{
		Use:                    "uplog",
		Short:                  "query uplog",
		Long:                   "",
		Example:                "",
		UserData:               handler,
		ExecuteFunction:        handler.execute,
	}
	queryCMDBuilder.FlagsBoolVar(&handler.isTime, "time", "t", false, "")
	queryCMDBuilder.FlagsBoolVar(&handler.isUser, "user", "u", false, "")

	superCMD.AddCMD(queryCMDBuilder.Build())
}

func (handler *queryLogCMDHandler) execute(context context.Context) error {


	return nil
}