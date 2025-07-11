package tests

import (
	"testing"

	"github.com/nazartymiv/go-mastery/external-tasks/safe-bank-transaction/logic"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		balance     float64
		wantName    string
		wantBalance float64
		expectErr   bool
	}{
		{
			name:        "normal account creation",
			userName:    User1,
			balance:     100,
			wantName:    User1,
			wantBalance: 100,
			expectErr:   false,
		},
		{
			name:      "account creation without user name",
			userName:  "",
			balance:   100,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := logic.NewBank()

			err := bank.CreateAccount(tt.userName, tt.balance)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			account := bank.Accounts[tt.userName]
			if account == nil {
				t.Fatalf("account was not created")
			}

			if account.Name != tt.userName {
				t.Errorf("expected name %s, got %s", tt.wantName, account.Name)
			}

			if account.Balance != tt.wantBalance {
				t.Errorf("expected balance %.2f, got %.2f", tt.wantBalance, account.Balance)
			}
		})
	}
}
