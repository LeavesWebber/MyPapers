package service

func (s *MPSService) Mint(toAddresses []string, amount uint64) error { return nil }
func (s *MPSService) Transfer(recipient string, amount uint64) error { return nil }
func (s *MPSService) GetBalanceOf(address string) (uint64, error)    { return 0, nil }
func (s *MPSService) StoreHash(hash string) error                    { return nil }
func (s *MPSService) GetRecipientByHash(hash string) (string, error) { return "", nil }
func (s *MPSService) StoreReview(content string) error               { return nil }
func (s *MPSService) GetReviewByHash(content string) (string, error) { return "", nil }
func (s *MPSService) RegisterUser(address string) error              { return nil }
