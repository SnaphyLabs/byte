package api_management


type(
	Field struct {
		Name string
		Description string
		Null bool
		Validation interface{} //TODO://handled for later user..
		ReadOnly bool
		Default string
		Unique bool
		Hidden bool
		Resolve interface{}

	}

)
