package api_management


type(
	Field struct {
		Name string
		Description string
		Type string
		Null bool
		Validation interface{} //TODO://handled for later user..
		ReadOnly bool
		Default interface{}
		Unique bool
		Hidden bool
		Resolve interface{}
		Args interface{}

	}

)
