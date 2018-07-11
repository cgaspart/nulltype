package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

/* SQL and JSon null.Float64 */

type Float64 sql.NullFloat64

func NewFloat64(f float64) Float64 {
	nf := Float64{}
	nf.Valid = true
	nf.Float64 = f
	return nf
}

func (nf *Float64) UnmarshalJSON(b []byte) error {
	var i float64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Compare(b, []byte("null")) == 0 {
		nf.Valid = false
		return nil
	}
	nf.Float64 = i
	nf.Valid = true
	return nil
}

func (nf Float64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	}
	return json.Marshal(nil)
}

func (nf *Float64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = Float64{f.Float64, false}
	} else {
		*nf = Float64{f.Float64, true}
	}

	return nil
}

func (nf Float64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}
	return nf.Float64, nil
}
