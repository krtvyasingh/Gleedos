package pqc

func VerifyCiphertextLength(ct []byte) bool {
	return len(ct) == 1568
}
