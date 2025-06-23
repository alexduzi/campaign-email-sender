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
	Fail     string = "Fail"
	Done     string = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50;primaryKey;not null"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50;not null"` // Make sure this matches Campaign.ID type
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;primaryKey;not null"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive" gorm:"foreignKey:CampaignId;references:ID"`
	Status    string    `gorm:"size:20;not null"`
	CreatedBy string    `validate:"min=5,max=100,email" gorm:"size:100;not null"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = Deleted
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Fail() {
	c.Status = Fail
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Started() {
	c.Status = Started
	c.UpdatedOn = time.Now()
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {
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
		CreatedBy: createdBy,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
