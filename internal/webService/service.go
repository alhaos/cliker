package webService

func (s *service) Run() {
	err := s.Router.Run(s.conf.Address)
	if err != nil {
		panic(err)
	}
}
