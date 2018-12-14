// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.1.0.0

package protoex

import "fmt"
import "strconv"
import "strings"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// Balance key
type BalanceKey struct {
    proto.BalanceKey
}

// Convert Balance flags key to string
func (k *BalanceKey) String() string {
    var sb strings.Builder
    sb.WriteString("BalanceKey(")
    sb.WriteString(k.BalanceKey.String())
    sb.WriteString(")")
    return sb.String()
}

// Balance struct
type Balance struct {
    *proto.Balance
    Locked float64 `json:"locked"`
}

// Create a new Balance struct
func NewBalance() *Balance {
    return &Balance{
        Balance: proto.NewBalance(),
        Locked: float64(0.0),
    }
}

// Create a new Balance struct from JSON
func NewBalanceFromJSON(buffer []byte) (*Balance, error) {
    result := *NewBalance()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *Balance) Copy() *Balance {
    var result = *s
    return &result
}

// Struct deep clone
func (s *Balance) Clone() *Balance {
    // Serialize the struct to the FBE stream
    writer := NewBalanceModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewBalanceModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *Balance) Key() BalanceKey {
    return BalanceKey{
        BalanceKey: s.Balance.Key(),
    }
}

// Convert struct to optional
func (s *Balance) Optional() *Balance {
    return s
}

// Convert struct to string
func (s *Balance) String() string {
    var sb strings.Builder
    sb.WriteString("Balance(")
    sb.WriteString(s.Balance.String())
    sb.WriteString(",locked=")
    sb.WriteString(strconv.FormatFloat(float64(s.Locked), 'g', -1, 64))
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *Balance) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}
