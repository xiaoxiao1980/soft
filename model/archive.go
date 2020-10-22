package model

type Field struct {
	ID            uint
	FileFieldCode string
	FileFieldName string
}

type File struct {
	ID          uint
	TableID     uint
	FileID      string
	FileName    string
	ArchiveID   string
	ArchiveName string
	Status      string
}

type Version struct {
	ID        uint
	ArchID    uint
	FileName  string
	Extension string
	Filemd5   string
	FileSize  uint
	Time      string
}

type Record struct {
	//gorm.Model
	ID           uint
	ArchID       uint
	ArchCode     string
	Title        string
	SaveFileName string
	Extension    string
	FilePath     string
}

type Result struct {
	data map[string]interface{}
}

func (s Field) TableName() string {
	return "tt_file_arch_field"
}

func (s File) TableName() string {
	return "tt_file_archive"
}

func (s Version) TableName() string {
	return "tt_file_e_version"
}

func (s Record) TableName() string {
	return "tt_archive_e_record"
}
