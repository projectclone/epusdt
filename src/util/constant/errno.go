package constant

var Errno = map[int]string{
	400:   "System Error",
	401:   "Signature authentication error",
	10001: "The wallet address already exists, please do not add it again",
	10002: "The payment transaction already exists, please do not create it again",
	10003: "No available wallet address, unable to initiate payment",
	10004: "The payment amount is incorrect and cannot meet the minimum payment unit",
	10005: "No available amount channel",
	10006: "Exchange rate calculation error",
	10007: "Order block processed",
	10008: "Order does not exist",
	10009: "Unable to parse request parameters",
}

var (
	SystemErr                  = Err(400)
	SignatureErr               = Err(401)
	WalletAddressAlreadyExists = Err(10001)
	OrderAlreadyExists         = Err(10002)
	NotAvailableWalletAddress  = Err(10003)
	PayAmountErr               = Err(10004)
	NotAvailableAmountErr      = Err(10005)
	RateAmountErr              = Err(10006)
	OrderBlockAlreadyProcess   = Err(10007)
	OrderNotExists             = Err(10008)
	ParamsMarshalErr           = Err(10009)
)

type RspError struct {
	Code int
	Msg  string
}

func (re *RspError) Error() string {
	return re.Msg
}

func Err(code int) (err error) {
	err = &RspError{
		Code: code,
		Msg:  Errno[code],
	}
	return err
}

func (re *RspError) Render() (code int, msg string) {
	return re.Code, re.Msg
}
