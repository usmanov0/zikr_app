package app

import "zikar-app/internal/model"

type PrayService interface {
	SavePray(pray *model.Pray) error
	Update(count int) error
	Refresh() error
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

func (p *prayService) Update(_ int) error {
	err := p.prayRepo.UpdaterCount()
	if err != nil {
		return err
	}
	return nil
}

func (p *prayService) Refresh() error {
	var count model.Pray
	count.CountPray = 0

	return nil
}
