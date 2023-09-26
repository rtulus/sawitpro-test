package handler

import (
	"reflect"
	"testing"
)

func Test_validateRegistrationInput(t *testing.T) {
	type args struct {
		ph   string
		name string
		pass string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "valid input",
			args: args{
				ph:   "+62788155269",
				name: "Tommy",
				pass: "Hihihi123!@#",
			},
			want: []string{},
		},
		{
			name: "invalid phone number",
			args: args{
				ph:   "62788155269999",
				name: "Tommy",
				pass: "Hihihi123!@#",
			},
			want: []string{
				"Phone number must start with +62",
				"Phone number must be at minimum 10 characters and maximum 13 characters",
			},
		},
		{
			name: "invalid name",
			args: args{
				ph:   "+62788155269",
				name: "as",
				pass: "Hihihi123!@#",
			},
			want: []string{
				"Full name must be at minimum 3 characters and maximum 60 characters",
			},
		},
		{
			name: "invalid password",
			args: args{
				ph:   "+62788155269",
				name: "Tommy",
				pass: "123",
			},
			want: []string{
				"Password must be at minimum 6 characters and maximum 64 characters",
				"Password must contain at least 1 capital character AND 1 number AND 1 special (non-alphanumeric) character",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateRegistrationInput(tt.args.ph, tt.args.name, tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateRegistrationInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateSalt(t *testing.T) {
	tests := []struct {
		name       string
		bytelength int
		wantErr    bool
	}{
		{
			name:       "generate random salt with 16 byte length",
			bytelength: 16,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateSalt()
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.bytelength {
				t.Errorf("generateSalt() = byte length is %v, want %v", len(got), tt.bytelength)
			}
		})
	}
}
