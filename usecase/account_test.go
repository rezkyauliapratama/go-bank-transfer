package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/gsabadini/go-bank-transfer/domain"
	"github.com/gsabadini/go-bank-transfer/infrastructure/database"
	"github.com/gsabadini/go-bank-transfer/repository"
)

func TestCreate(t *testing.T) {
	type args struct {
		repository repository.Account
		account    domain.Account
	}

	tests := []struct {
		name     string
		args     args
		expected interface{}
	}{
		{
			name: "Create account successful",
			args: args{
				repository: repository.NewAccount(database.MongoHandlerSuccessMock{}),
				account: domain.Account{
					Name:      "Test",
					Cpf:       "44451598087",
					Ballance:  0,
					CreatedAt: time.Now(),
				},
			},
			expected: nil,
		},
		{
			name: "Create account error",
			args: args{
				repository: repository.NewAccount(database.MongoHandlerErrorMock{}),
				account: domain.Account{
					Name:      "Test",
					Cpf:       "44451598087",
					Ballance:  0,
					CreatedAt: time.Now(),
				},
			},
			expected: "Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := Create(tt.args.repository, tt.args.account); (err != nil) && (err.Error() != tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, err, tt.expected)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	var timeNow = time.Now()

	type args struct {
		repository repository.Account
		account    []domain.Account
	}

	tests := []struct {
		name          string
		args          args
		expected      []domain.Account
		expectedError interface{}
	}{
		{
			name: "Success return list accounts",
			args: args{
				repository: repository.NewAccount(database.MongoHandlerSuccessMock{}),
				account: []domain.Account{
					{
						Id:        "0",
						Name:      "Test-0",
						Cpf:       "",
						Ballance:  0,
						CreatedAt: timeNow,
					},
					{
						Id:        "1",
						Name:      "Test-1",
						Cpf:       "",
						Ballance:  120,
						CreatedAt: timeNow,
					},
				},
			},
			expected: []domain.Account{
				{
					Id:        "0",
					Name:      "Test-0",
					Cpf:       "",
					Ballance:  0,
					CreatedAt: timeNow,
				},
				{
					Id:        "1",
					Name:      "Test-1",
					Cpf:       "",
					Ballance:  120,
					CreatedAt: timeNow,
				},
			},
		},
		{
			name: "Empty return list accounts",
			args: args{
				repository: repository.NewAccount(database.MongoHandlerSuccessMock{}),
				account:    []domain.Account{},
			},
			expected: []domain.Account{},
		},
		{
			name: "Error return list accounts",
			args: args{
				repository: repository.NewAccount(database.MongoHandlerErrorMock{}),
				account:    []domain.Account{},
			},
			expectedError: "Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := FindAll(tt.args.repository, tt.args.account)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, err, tt.expected)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, got, tt.expected)
			}
		})
	}
}