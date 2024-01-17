package model

type Pray struct {
	Id         int
	Language   string
	Definition string
	CountPray  int
}

type PrayRepository interface {
	Save(pray *Pray) error
	UpdaterCount() error
}
