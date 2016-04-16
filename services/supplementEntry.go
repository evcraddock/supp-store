package services

import (
	"time"

	"github.com/evcraddock/supp-store/models"
	"gopkg.in/mgo.v2"
)

type (
	SupplementEntryService struct {
		session *mgo.Session
	}
)

func InitSupplmentEntryService(s *mgo.Session) *SupplementEntryService {

	return &SupplementEntryService{s}
}

func (this *SupplementEntryService) Create(entry models.SupplementEntry) error {

	entry.DateTaken = time.Now()

	err := this.session.DB("supplements").C("entries").Insert(entry)

	return err
}

func (this *SupplementEntryService) GetAll() ([]models.SupplementEntry, error) {
	var entries []models.SupplementEntry

	if err := this.session.DB("supplements").C("entries").Find(nil).All(&entries); err != nil {
		return nil, err
	}

	return entries, nil
}
