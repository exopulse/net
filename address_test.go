package net

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAddress_ParseAddress(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Address
		wantErr bool
	}{
		{"address-and-port", "mail:25", Address{"mail", 25}, false},
		{"address", "mail", Address{"mail", 0}, false},
		{"no-address", ":25", Address{"", 0}, true},
		{"empty", "", Address{"", 0}, true},
		{"invalid-port", "loclahost:port1", Address{"", 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddress(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Host() != tt.want.Host() || got.Port() != tt.want.Port() {
				t.Errorf("ParseAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_WithDefaultPort(t *testing.T) {
	type fields struct {
		host string
		port int
	}

	tests := []struct {
		name   string
		fields fields
		port   int
		want   Address
	}{
		{"assigned", fields{"mail", 25}, 589, Address{"mail", 25}},
		{"unassigned", fields{"mail", 0}, 589, Address{"mail", 589}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				host: tt.fields.host,
				port: tt.fields.port,
			}

			if got := a.WithDefaultPort(tt.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.WithDefaultPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_String(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"address-and-port", "mail:25"},
		{"address", "mail"},
		{"address", "127.0.0.1"},
		{"address", "127.0.0.1:3000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddress(tt.value)

			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			if got.String() != tt.value {
				t.Errorf("ParseAddress() = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestAddress_MustParseAddress_Success(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()

	s := "mail:25"
	address := MustParseAddress(s)

	if address.String() != s {
		t.Errorf("MustParseAddress = %v, want %v", address.String(), s)
	}
}

func TestAddress_MustParseAddress_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expected")
		}
	}()

	s := "mail:invalid"
	address := MustParseAddress(s)

	if address.String() != s {
		t.Errorf("MustParseAddress = %v, want %v", address.String(), s)
	}
}

func TestAddress_JSON(t *testing.T) {
	v1 := struct {
		ID     string
		Server Address
	}{"123", MustParseAddress("127.0.0.1:4000")}

	bytes, err := json.Marshal(&v1)

	if err != nil {
		t.Fatal(err)
	}

	s := string(bytes)
	expected := `{"ID":"123","Server":"127.0.0.1:4000"}`

	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}

	v2 := struct {
		ID     string
		Server Address
	}{}

	if err := json.Unmarshal(bytes, &v2); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("expected %v, got %v", v1, v2)
	}
}
