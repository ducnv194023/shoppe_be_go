package services

type IMailService interface {
	SendVerificationEmail()
}

type MailService struct {
}

// SendVerificationEmail implements IMailService.
func (ms *MailService) SendVerificationEmail() {
	panic("unimplemented")
}

func NewMailService() IMailService {
	return &MailService{}
}
