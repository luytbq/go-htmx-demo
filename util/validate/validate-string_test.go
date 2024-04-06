package validate

import "testing"

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantOk  bool
		wantMsg string
	}{
		{
			name:    "TestValidateEmail #1",
			email:   "",
			wantOk:  false,
			wantMsg: emailEmptyMsg,
		},
		{
			name:    "TestValidateEmail #3",
			email:   "abc@gmail.com",
			wantOk:  true,
			wantMsg: emptyString,
		},
		{
			name:    "TestValidateEmail #4",
			email:   "abc@gmail.a.com",
			wantOk:  true,
			wantMsg: emptyString,
		},
		{
			name:    "TestValidateEmail #4",
			email:   "@gmail.a.com",
			wantOk:  false,
			wantMsg: emailInvalidMsg,
		},
		{
			name:    "TestValidateEmail #5",
			email:   "@gmail.a.com",
			wantOk:  false,
			wantMsg: emailInvalidMsg,
		},
		{
			name:    "TestValidateEmail #6",
			email:   "23432$@gmail.com",
			wantOk:  false,
			wantMsg: emailInvalidMsg,
		},
		{
			name:    "TestValidateEmail #7",
			email:   "abc-def@gmail.com",
			wantOk:  true,
			wantMsg: emptyString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, msg := ValidateEmail(tt.email)
			if ok != tt.wantOk {
				t.Errorf("ValidateEmail() ok = %v, want %v", ok, tt.wantOk)
			}
			if msg != tt.wantMsg {
				t.Errorf("ValidateEmail() msg = %v, want %v", msg, tt.wantMsg)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantOk   bool
		wantMsg  string
	}{
		{
			name:     "TestValidatePassword #1",
			password: "",
			wantOk:   false,
			wantMsg:  passwordEmptyMsg,
		},
		{
			name:     "TestValidatePassword #2",
			password: "1",
			wantOk:   false,
			wantMsg:  passwordTooShortMsg,
		},
		{
			name:     "TestValidatePassword #3",
			password: "1234567",
			wantOk:   false,
			wantMsg:  passwordTooShortMsg,
		},
		{
			name:     "TestValidatePassword #4",
			password: "12345678",
			wantOk:   true,
			wantMsg:  emptyString,
		},
		{
			name:     "TestValidatePassword #5",
			password: "123456789",
			wantOk:   true,
			wantMsg:  emptyString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, msg := ValidatePassword(tt.password)
			if ok != tt.wantOk {
				t.Errorf("ValidatePassword() ok = %v, want %v", ok, tt.wantOk)
			}
			if msg != tt.wantMsg {
				t.Errorf("ValidatePassword() msg = %v, want %v", msg, tt.wantMsg)
			}
		})
	}
}
