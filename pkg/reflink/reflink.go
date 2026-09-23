package reflink

func IsReflinkSupported(fsType string) bool {
	return fsType == "apfs" || fsType == "btrfs" || fsType == "zfs"
}
