package exception

type InvalidState struct {

}

func (e *InvalidState) Error() string{
	return "invalid state"
}