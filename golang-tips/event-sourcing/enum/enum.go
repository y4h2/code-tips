package enum

var EventType = struct {
	Transfer string
	Open     string
	Store    string
	Withdraw string
}{
	Transfer: "transfer",
	Open:     "open",
	Store:    "store",
	Withdraw: "withdraw",
}

var EventEntity = struct {
	Account string
}{
	Account: "account",
}
