package models

type FileTypes string

const FILE_TYPE_GENERIC FileTypes = "GENERIC"

type S3FileModel struct {
	Key          string    `json:"key" bson:"key,omitempty"  `
	OriginalName string    `json:"originalName" bson:"originalName,omitempty"  `
	Url          string    `json:"url" bson:"url,omitempty"  `
	Mimetype     string    `json:"mimeType" bson:"mimeType,omitempty"  `
	FileType     FileTypes `json:"fileType" bson:"fileType,omitempty"  `
	Size         int64     `json:"size" bson:"size,omitempty"  `
}

func NewS3FileModel(
	key string,
	originalName string,
	url string,
	mimetype string,
	fileType FileTypes,
	size int64) *S3FileModel {
	v := &S3FileModel{OriginalName: originalName, Key: key, Url: url, Mimetype: mimetype, FileType: fileType, Size: size}
	return v
}
