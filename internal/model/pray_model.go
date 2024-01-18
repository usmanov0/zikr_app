package model

type Pray struct {
	Id         int
	Language   string
	Definition string
}

type PrayRepository interface {
	Save(pray *Pray) error
	Read(prayId int) (*Pray, error)
	Put(prayId int, language, definition string) error
	Delete(prayId int) error
}
