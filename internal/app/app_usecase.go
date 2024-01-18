package app

import "zikar-app/internal/model"

type PrayService interface {
	SavePray(pray *model.Pray) error
	ReadPray(prayId int) (*model.Pray, error)
	PutPray(id int, language, definition string) error
	Delete(id int) error
}

type prayService struct {
	prayRepo model.PrayRepository
}

func NewBasketService(pray model.PrayRepository) PrayService {
	return &prayService{prayRepo: pray}
}

func (p *prayService) SavePray(pray *model.Pray) error {
	err := p.prayRepo.Save(pray)
	if err != nil {
		return err
	}
	return nil
}

func (p *prayService) ReadPray(prayId int) (*model.Pray, error) {
	pray, err := p.prayRepo.Read(prayId)
	if err != nil {
		return nil, err
	}
	return pray, nil
}

func (p *prayService) PutPray(id int, language, definition string) error {
	err := p.prayRepo.Put(id, language, definition)
	if err != nil {
		return err
	}
	return nil
}

func (p *prayService) Delete(id int) error {
	err := p.prayRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
