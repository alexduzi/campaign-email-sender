package campaign

import (
	internalerrors "campaignemailsender/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
	Started  string = "Started"
	Done     string = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50;primaryKey"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50;not null"` // Make sure this matches Campaign.ID type
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;primaryKey"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive" gorm:"foreignKey:CampaignId;references:ID"`
	Status    string    `gorm:"size:20"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	campaignId := xid.New().String()

	contacts := make([]Contact, len(emails))

	for index, value := range emails {
		contacts[index].Email = value
		contacts[index].ID = xid.New().String()
		contacts[index].CampaignId = campaignId
	}

	campaign := &Campaign{
		ID:        campaignId,
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
