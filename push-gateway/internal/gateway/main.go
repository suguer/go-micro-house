package gateway

type SmsGateway interface {
}

func Factory() *SmsGateway {
	return nil
}
