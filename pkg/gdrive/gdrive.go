package gdrive

type DriveUploadRequest struct {
	FileName string
	ParentID string
}

func NewUploadRequest(name, parent string) DriveUploadRequest {
	return DriveUploadRequest{FileName: name, ParentID: parent}
}
