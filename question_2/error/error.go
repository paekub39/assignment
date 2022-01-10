package error

type CashierError struct {
	Msg string
}

func (ce CashierError) Error() string {
	return ce.Msg
}
