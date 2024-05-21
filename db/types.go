package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type StringArray []string

func (t StringArray) GormDataType() string {
	return "text[]"
}

func (t *StringArray) Scan(src interface{}) error {
	v := pq.StringArray{}
	res := v.Scan(src)
	*t = StringArray(v)
	return res
}

func (t StringArray) Value() (driver.Value, error) {
	return pq.StringArray(t).Value()
}

type Int64Array []int64

func (t Int64Array) GormDataType() string {
	return "integer[]"
}

func (t *Int64Array) Scan(src interface{}) error {
	v := pq.Int64Array{}
	res := v.Scan(src)
	*t = Int64Array(v)
	return res
}

func (t Int64Array) Value() (driver.Value, error) {
	return pq.Int64Array(t).Value()
}

func (t Int64Array) ToUintArray() []uint {
	res := make([]uint, len(t))
	for i, v := range t {
		res[i] = uint(v)
	}
	return res
}

func (t *Int64Array) FromUintArray(l []uint) {
	*t = make(Int64Array, len(l))
	for i, v := range l {
		(*t)[i] = int64(v)
	}
}

// type UintArray []uint

// func (t UintArray) GormDataType() string {
// 	return "integer[]"
// }

// func (t *UintArray) Scan(src interface{}) error {
// 	v := pq.Int64Array{}
// 	res := v.Scan(src)
// 	*t = make([]uint, len(v))
// 	for i, val := range v {
// 		(*t)[i] = uint(val)
// 	}
// 	return res
// }

// func (t UintArray) Value() (driver.Value, error) {
// 	pqarr := make(pq.Int64Array, len(t))
// 	for i, val := range t {
// 		pqarr[i] = int64(val)
// 	}
// 	return pqarr.Value()
// }

type Float64Array []float64

func (t Float64Array) GormDataType() string {
	return "real[]"
}

func (t *Float64Array) Scan(src interface{}) error {
	v := pq.Float64Array{}
	res := v.Scan(src)
	*t = Float64Array(v)
	return res
}

func (t Float64Array) Value() (driver.Value, error) {
	return pq.Float64Array(t).Value()
}

// // util.JsonPsql use for db.Raw and db.Exec type json
type JsonStruct[T any] struct {
	Val T
}

func NewJsonStruct[T any](v T) JsonStruct[T] {
	return JsonStruct[T]{Val: v}
}

func (x JsonStruct[T]) GormDataType() string {
	return "json"
}

// // Value to write to database
func (x JsonStruct[T]) Value() (driver.Value, error) {
	return json.Marshal(x.Val)
}

// // Scan to read from database
func (x *JsonStruct[T]) Scan(src interface{}) error {
	if src == nil {
		var v T
		x.Val = v
		return nil
	}
	var source []byte
	switch v := src.(type) {
	case string:
		source = []byte(v)
	case []byte:
		source = v
	default:
		return errors.New("need type string or []byte for json")
	}
	return json.Unmarshal(source, &x.Val)
}

func (x JsonStruct[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Val)
}

func (x *JsonStruct[T]) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &x.Val)
}

func (x *JsonStruct[T]) Copy(obj JsonStruct[T]) error {
	bytes, err := obj.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &x.Val)
}

// type JsonPsql JsonStruct[interface{}]

func NewJsonPsql(src interface{}) *string {
	if src == nil {
		return nil
	}
	res, err := json.Marshal(src)
	if err != nil {
		panic(fmt.Sprintf("util.NewJsonPsql marshal error: %v", err))
	}
	x := string(res)
	return &x
}

// func PrettyStruct(data interface{}) string {
// 	val, err := json.MarshalIndent(data, "", "    ")
// 	if err != nil {
// 		fmt.Println("PrettyStruct error", err)
// 		return ""
// 	}
// 	return string(val)
// }

func IsValidUUID(s string) bool {
	_, err := uuid.FromString(s)
	return err == nil
}

type CustomUUID string

func (t CustomUUID) GormDataType() string {
	return "uuid"
}

func (t CustomUUID) IsNull() bool {
	return len(t) == 0
}

func (t *CustomUUID) Scan(src interface{}) error {
	if src == nil {
		*t = ""
		return nil
	}
	var source string
	switch v := src.(type) {
	case string:
		source = v
	case []byte:
		source = string(v)
	default:
		return errors.New("need type string or []byte for uuid")
	}
	if len(source) == 0 {
		*t = ""
		return nil
	}
	if IsValidUUID(source) {
		*t = CustomUUID(source)
		return nil
	} else {
		return errors.New("invalid uuid")
	}
}

func (t CustomUUID) Value() (driver.Value, error) {
	if t.IsNull() {
		return nil, nil
	}
	return string(t), nil
}

func (t CustomUUID) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	}
	res, err := json.Marshal(string(t))
	return res, err
}

func (t *CustomUUID) UnmarshalJSON(bytes []byte) error {
	v := ""
	*t = ""
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}
	if !IsValidUUID(string(v)) {
		return nil
	}
	*t = CustomUUID(v)
	return nil
}

type StringChannelReader struct {
	channel   <-chan string
	currChunk []byte
}

func NewStringChannelReader(channel <-chan string) *StringChannelReader {
	return &StringChannelReader{
		channel: channel,
	}
}

func (scr *StringChannelReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, io.ErrShortBuffer
	}

	if len(scr.currChunk) == 0 {
		if !scr.readNextChunk() {
			return 0, io.EOF
		}
	}
	// fmt.Println(260, len(p), string(scr.currChunk))

	n := copy(p, scr.currChunk)
	// fmt.Println(263, n)
	scr.currChunk = scr.currChunk[n:]

	return n, nil
}

func (scr *StringChannelReader) readNextChunk() bool {
	// fmt.Println(268, "readnextchunk")
	str, ok := <-scr.channel
	// fmt.Println(270, str, ok)
	if !ok {
		return false
	}

	scr.currChunk = []byte(str)
	return true
}

// func main() {
// 	// Create a channel of strings
// 	stringChannel := make(chan string)

// 	// Start a goroutine to produce strings and send them to the channel
// 	go func() {
// 		defer close(stringChannel)
// 		stringChannel <- "Hello"
// 		stringChannel <- ", "
// 		stringChannel <- "World!"
// 	}()

// 	// Create a reader from the string channel
// 	reader := NewStringChannelReader(stringChannel)

// 	// Read from the reader and print the contents
// 	buf := make([]byte, 128)
// 	for {
// 		n, err := reader.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println("Error:", err)
// 			return
// 		}

// 		fmt.Print(string(buf[:n]))
// 	}
// }

type TkInfo struct {
	Token            string
	Sub              int
	Name             string
	Email            string
	OrganizationCode string `json:"organization_code"`
}

func (tk TkInfo) CanUpdateConfig() bool {
	return tk.Email == "thao.nguyen@olli-ai.com" || tk.Email == "son@olli-ai.com" || tk.Email == "quang.nguyen@olli-ai.com" || tk.Email == "ducanh@olli-ai.com" || tk.Email == "quang.num@gmail.com"
}

func MustMarshal(v any) []byte {
	d, _ := json.Marshal(v)
	return d
}
