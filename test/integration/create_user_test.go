package integration

func (s *Suite) TestCreateUser() {
	res, err := s.client.CreateUser(ctx, "John_Create", "john@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("John_Create", res.Name)
	s.Equal("john@gmail.com", res.Email)
	s.NotNil(res.Token)
}
