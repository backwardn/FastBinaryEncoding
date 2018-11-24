// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.1.0.0

package proto

import "encoding/json"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// OrderSide enum key
type OrderSideKey byte

// Convert OrderSide enum key to string
func (k OrderSideKey) String() string {
    return OrderSide(k).String()
}

// OrderSide enum
type OrderSide byte

// OrderSide enum values
//noinspection GoSnakeCaseUsage
const (
    OrderSide_buy = OrderSide(0 + 0)
    OrderSide_sell = OrderSide(0 + 1)
)

// Create a new OrderSide enum
func NewOrderSide() *OrderSide {
    return new(OrderSide)
}

// Get the enum key
func (e OrderSide) Key() OrderSideKey {
    return OrderSideKey(e)
}

// Convert enum to string
func (e OrderSide) String() string {
    if e == OrderSide_buy {
        return "buy"
    }
    if e == OrderSide_sell {
        return "sell"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e OrderSide) MarshalJSON() ([]byte, error) {
    return json.Marshal(byte(e))
}

// Convert JSON to enum
func (e *OrderSide) UnmarshalJSON(buffer []byte) error {
    var result byte
    err := json.Unmarshal(buffer, &result)
    if err != nil {
        return err
    }
    *e = OrderSide(result)
    return nil
}