package checksum

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileChecksums(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.dat")
	if err := os.WriteFile(filePath, []byte("gleedos-checksum-payload"), 0644); err != nil {
		t.Fatal(err)
	}

	sha, err := FileSHA256(filePath)
	if err != nil || len(sha) != 64 {
		t.Fatalf("unexpected SHA256: %s, %v", sha, err)
	}

	md5Sum, err := FileMD5(filePath)
	if err != nil || len(md5Sum) != 32 {
		t.Fatalf("unexpected MD5: %s, %v", md5Sum, err)
	}
}
