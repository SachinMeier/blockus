package bitcoind

func reverseBytes(data []byte) []byte {
	dlen := len(data)
	res := make([]byte, dlen)
	for i, b := range data {
		res[dlen-i-1] = b
	}
	return res
}
