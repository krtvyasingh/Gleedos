package ramshred

func MultiPassShred(data []byte) {
	// Pass 1: Zeroes
	for i := range data {
		data[i] = 0x00
	}
	// Pass 2: Ones
	for i := range data {
		data[i] = 0xFF
	}
	// Pass 3: Final Zeroes
	for i := range data {
		data[i] = 0x00
	}
}
