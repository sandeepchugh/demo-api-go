package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Delhi", "110075", "2000-01-01", "1"},
		{"1002", "Rohit", "New Delhi", "110075", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
