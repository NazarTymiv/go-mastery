package tests

import (
	"testing"
)

func TestTransfer(t *testing.T) {
	tests := []struct {
		name        string
		from        string
		to          string
		amount      float64
		expectErr   bool
		expectedMsg string
	}{
		{
			name:   "successful transfer from user1 to user2",
			from:   User1,
			to:     User2,
			amount: 30,
		},
		{
			name:        "transfer with insufficient balance",
			from:        User1,
			to:          User2,
			amount:      200,
			expectErr:   true,
			expectedMsg: "Sender does not have enough money to send",
		},
		{
			name:        "transfer to non-existent user",
			from:        User1,
			to:          "ghost",
			amount:      10,
			expectErr:   true,
			expectedMsg: "Account not found",
		},
		{
			name:        "transfer from non-existent user",
			from:        "ghost",
			to:          User2,
			amount:      10,
			expectErr:   true,
			expectedMsg: "Account not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := setup()

			err := bank.Transfer(tt.from, tt.to, tt.amount)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if tt.expectedMsg != "" && err.Error() != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			fromBal := bank.GetBalance(tt.from)
			toBal := bank.GetBalance(tt.to)

			expectedFrom := 100.0
			expectedTo := 50.0
			expectedFrom -= tt.amount
			expectedTo += tt.amount

			if fromBal != expectedFrom {
				t.Errorf("expected sender balance %.2f, got %.2f", expectedFrom, fromBal)
			}
			if toBal != expectedTo {
				t.Errorf("expected receiver balance %.2f, got %.2f", expectedTo, toBal)
			}

			if len(bank.Transactions) != 1 {
				t.Errorf("expected 1 transaction, got %d", len(bank.Transactions))
			}
		})
	}
}
