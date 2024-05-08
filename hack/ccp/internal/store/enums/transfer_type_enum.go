// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

const (
	// TransferTypeTransfer is a TransferType of type Transfer.
	TransferTypeTransfer TransferType = "Transfer"
	// TransferTypeSIP is a TransferType of type SIP.
	TransferTypeSIP TransferType = "SIP"
	// TransferTypeDIP is a TransferType of type DIP.
	TransferTypeDIP TransferType = "DIP"
)

var ErrInvalidTransferType = fmt.Errorf("not a valid TransferType, try [%s]", strings.Join(_TransferTypeNames, ", "))

var _TransferTypeNames = []string{
	string(TransferTypeTransfer),
	string(TransferTypeSIP),
	string(TransferTypeDIP),
}

// TransferTypeNames returns a list of possible string values of TransferType.
func TransferTypeNames() []string {
	tmp := make([]string, len(_TransferTypeNames))
	copy(tmp, _TransferTypeNames)
	return tmp
}

// String implements the Stringer interface.
func (x TransferType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x TransferType) IsValid() bool {
	_, err := ParseTransferType(string(x))
	return err == nil
}

var _TransferTypeValue = map[string]TransferType{
	"Transfer": TransferTypeTransfer,
	"transfer": TransferTypeTransfer,
	"SIP":      TransferTypeSIP,
	"sip":      TransferTypeSIP,
	"DIP":      TransferTypeDIP,
	"dip":      TransferTypeDIP,
}

// ParseTransferType attempts to convert a string to a TransferType.
func ParseTransferType(name string) (TransferType, error) {
	if x, ok := _TransferTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _TransferTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return TransferType(""), fmt.Errorf("%s is %w", name, ErrInvalidTransferType)
}

func (x TransferType) Ptr() *TransferType {
	return &x
}

// MarshalText implements the text marshaller method.
func (x TransferType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *TransferType) UnmarshalText(text []byte) error {
	tmp, err := ParseTransferType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errTransferTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *TransferType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = TransferType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseTransferType(v)
	case []byte:
		*x, err = ParseTransferType(string(v))
	case TransferType:
		*x = v
	case *TransferType:
		if v == nil {
			return errTransferTypeNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errTransferTypeNilPtr
		}
		*x, err = ParseTransferType(*v)
	default:
		return errors.New("invalid type for TransferType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x TransferType) Value() (driver.Value, error) {
	return x.String(), nil
}

// Set implements the Golang flag.Value interface func.
func (x *TransferType) Set(val string) error {
	v, err := ParseTransferType(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *TransferType) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *TransferType) Type() string {
	return "TransferType"
}
