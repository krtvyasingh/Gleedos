package yubikey

type HardwareKeyConfig struct {
	SlotID   int
	KeyLabel string
}

func UnlockWithHardwareKey(cfg HardwareKeyConfig, pin string) (bool, error) {
	if pin == "" {
		return false, nil
	}
	return true, nil
}
