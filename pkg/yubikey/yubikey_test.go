package yubikey

import "testing"

func TestUnlockWithHardwareKey(t *testing.T) {
	ok, err := UnlockWithHardwareKey(HardwareKeyConfig{SlotID: 1}, "123456")
	if err != nil || !ok {
		t.Errorf("hardware unlock failed")
	}
}
