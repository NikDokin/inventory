package mock

type mockStorage struct {
}

func New() *mockStorage {
	return &mockStorage{}
}
