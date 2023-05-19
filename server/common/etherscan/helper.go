
 package etherscan

 import (
	 "math/big"
	 "reflect"
	 "strconv"
	 "time"
 )
 
 // compose adds input to param, whose key is tag
 // if input is nil or nil of some type, compose is a no-op.
 func compose(param map[string]interface{}, tag string, input interface{}) {
	 // simple situation
	 if input == nil {
		 return
	 }
 
	 // needs dig further
	 v := reflect.ValueOf(input)
	 switch v.Kind() {
	 case reflect.Ptr, reflect.Slice, reflect.Interface:
		 if v.IsNil() {
			 return
		 }
	 }
 
	 param[tag] = input
 }
 
 // M is a type shorthand for param input
 
 //包装原始big.Int实现解析json的接口
 type BigInt big.Int
 
 // UnmarshalText implements the encoding.TextUnmarshaler interface.
 func (b *BigInt) UnmarshalText(text []byte) (err error) {
	 var bigInt = new(big.Int)
	 err = bigInt.UnmarshalText(text)
	 if err != nil {
		 return
	 }
 
	 *b = BigInt(*bigInt)
	 return nil
 }
 
 //encoding.TextMarshaler
 func (b *BigInt) MarshalText() (text []byte, err error) {
	 return []byte(b.Int().String()), nil
 }
 
 func (b *BigInt) Int() *big.Int {
	 return (*big.Int)(b)
 }
 
 // 包装原始Time实现json解析的接口
 type Time time.Time
 
 func (t *Time) UnmarshalText(text []byte) (err error) {
	 input, err := strconv.ParseInt(string(text), 10, 64)
	 if err != nil {
		 err = wrapErr(err, "strconv.ParseInt")
		 return
	 }
 
	 var timestamp = time.Unix(input, 0)
	 *t = Time(timestamp)
 
	 return nil
 }
 
 func (t Time) Time() time.Time {
	 return time.Time(t)
 }
 
 func (t Time) MarshalText() (text []byte, err error) {
	 return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
 }