package integration

func (s *Suite) TestLogin() {
	userCreated, err := s.client.CreateUser(ctx, "John_Create", "new_john@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("John_Create", userCreated.Name)
	s.Equal("new_john@gmail.com", userCreated.Email)
	s.NotNil(userCreated.Token)

	res, err := s.client.Login(ctx, "new_john@gmail.com", "12345678")

	s.NoError(err)
	s.Equal("John_Create", res.Name)
	s.Equal("new_john@gmail.com", res.Email)
	s.NotNil(res.Token)
}
