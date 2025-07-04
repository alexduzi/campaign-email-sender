package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Campaign X"
	createdBy = "test@test.com.br"
	content   = "Body Hi!"
	contacts  = []string{"email1@e.com", "email2@e.com"}
	fake      = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
	assert.Equal(createdBy, campaign.CreatedBy)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts, createdBy)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts, createdBy)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"}, createdBy)

	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(Pending, campaign.Status)
}

func Test_NewCampaign_CreatedByIsEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, contacts, "")

	assert.Equal("createdby is required with min 5", err.Error())
}

func Test_Done_ChangeStatus(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Done()

	assert.Equal(Done, campaign.Status)
}

func Test_Started_ChangeStatus(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Started()

	assert.Equal(Started, campaign.Status)
}

func Test_Cancel_ChangeStatus(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	campaign.Cancel()

	assert.Equal(Canceled, campaign.Status)
}
