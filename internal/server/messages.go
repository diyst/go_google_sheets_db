package server

const (
	AuthenticationOk     = 'R'
	ParameterStatus      = 'S'
	BackendKeyData       = 'K'
	ReadyForQuery        = 'Z'
	Idle                 = 'I'
	QueryMessage         = 'Q'
	StartupMessage       = 0
	RowDescription       = 'T'
	DataRow              = 'D'
	CommandComplete      = 'C'
	Transaction          = 'T'
	Failed               = 'E'
	ParseMessage         = 'P'
	ParseCompleteMessage = 'B'
	Terminate            = 'X'
)
