package models

type FileInfo struct {
	Owner     string   `bson:"owner"` // 파일 소유자
	Name      string   `bson:"name"`
	Uri       string   `bson:"uri"`
	Type      string   `bson:"type"`
	Tags      []string `bson:"tag"`
	Size      int64    `bson:"size"`
	CreatedAt int64    `bson:"createdAt"`
}
