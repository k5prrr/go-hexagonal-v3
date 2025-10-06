package app

type diContainer struct {
	knowApi
	knowService
	knowRepository
	issClient
}

func (d *diContainer) KnowAPI() api.KnowAPI {
	if d.knowApi == nill {
		d.knowApi = KnowAPI.NewAPI(d.knowService())
	}
}
