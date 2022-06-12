package phonebook

// Source: https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/

import "testing"

//-------- Dummy
// O dummy nao influencia na execução. Ele pode até ser executado, mas nao altera o teste.

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindReturnsError(t *testing.T) {
	//p := Phonebook{}
	//phonebook := &p
	phonebook := &Phonebook{}

	expected := ErrMissingArgs
	_, got := phonebook.Find(DummySearcher{}, "", "")

	if got != expected {
		t.Errorf("Want '%s', got '%s'", expected, got)
	}
}

//-------- Stub - devolve valores geralmente fixos. Quando é necessário testar uma resposta do método influenciada pelos valores do stub.
type StubSearcher struct {
	phone string
}

var expectedPhone = "+31 65 222 333"

func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: "Foo",
		LastName:  "Bar",
		Phone:     expectedPhone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	dummyFirstName := "Jane"
	dummLastName := "Doe"
	phonebook := &Phonebook{}
	stubSearcher := StubSearcher{}

	phone, _ := phonebook.Find(stubSearcher, dummyFirstName, dummLastName)

	if phone != expectedPhone {
		t.Errorf("Want '%s', got '%s'", expectedPhone, phone)
	}
}

//----- Spy - verifica se as implentações esperadas foram chamadas (e quantas vezes)

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	searchCallCount int
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	ss.searchCallCount++
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}

//----- Mocks - geralmente possui funcionalidades do spy e stub, mais validações do teste em si.

type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}
	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")
	mock.ExpectToCall("Save")

	phone, _ := phonebook.Find(mock, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}

//------Fake - implementa certa lógica, geralmente simplificada, para testes.

type FakeSearcher struct{}

func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}
}
