package entity

type Result struct {
	Msg    string `json:"msg,omitempty"`
	Object any    `json:"object,omitempty"`
}

func (Result) Ok() Result {
	return Result{
		Msg:    "success",
		Object: nil,
	}
}

func (Result) OkWithObj(object any) Result {
	return Result{
		Msg:    "success",
		Object: object,
	}
}

func (Result) OkWithMsgAndObj(msg string, object any) Result {
	return Result{
		Msg:    msg,
		Object: object,
	}
}
