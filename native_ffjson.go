// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: native.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Native) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Native) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "request":`)
	fflib.WriteJsonString(buf, string(mj.Request))
	buf.WriteByte(',')
	if len(mj.Ver) != 0 {
		buf.WriteString(`"ver":`)
		fflib.WriteJsonString(buf, string(mj.Ver))
		buf.WriteByte(',')
	}
	if len(mj.API) != 0 {
		buf.WriteString(`"api":`)
		if mj.API != nil {
			buf.WriteString(`[`)
			for i, v := range mj.API {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.BAttr) != 0 {
		buf.WriteString(`"battr":`)
		if mj.BAttr != nil {
			buf.WriteString(`[`)
			for i, v := range mj.BAttr {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Ext != nil {
		if true {
			buf.WriteString(`"ext":`)

			{

				obj, err = mj.Ext.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Nativebase = iota
	ffj_t_Nativeno_such_key

	ffj_t_Native_Request

	ffj_t_Native_Ver

	ffj_t_Native_API

	ffj_t_Native_BAttr

	ffj_t_Native_Ext
)

var ffj_key_Native_Request = []byte("request")

var ffj_key_Native_Ver = []byte("ver")

var ffj_key_Native_API = []byte("api")

var ffj_key_Native_BAttr = []byte("battr")

var ffj_key_Native_Ext = []byte("ext")

func (uj *Native) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Native) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Nativebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Nativeno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Native_API, kn) {
						currentKey = ffj_t_Native_API
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_Native_BAttr, kn) {
						currentKey = ffj_t_Native_BAttr
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Native_Ext, kn) {
						currentKey = ffj_t_Native_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_Native_Request, kn) {
						currentKey = ffj_t_Native_Request
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_Native_Ver, kn) {
						currentKey = ffj_t_Native_Ver
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Native_Ext, kn) {
					currentKey = ffj_t_Native_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Native_BAttr, kn) {
					currentKey = ffj_t_Native_BAttr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Native_API, kn) {
					currentKey = ffj_t_Native_API
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Native_Ver, kn) {
					currentKey = ffj_t_Native_Ver
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Native_Request, kn) {
					currentKey = ffj_t_Native_Request
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Nativeno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Native_Request:
					goto handle_Request

				case ffj_t_Native_Ver:
					goto handle_Ver

				case ffj_t_Native_API:
					goto handle_API

				case ffj_t_Native_BAttr:
					goto handle_BAttr

				case ffj_t_Native_Ext:
					goto handle_Ext

				case ffj_t_Nativeno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Request:

	/* handler: uj.Request type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Request = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ver:

	/* handler: uj.Ver type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Ver = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_API:

	/* handler: uj.API type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.API = nil
		} else {

			uj.API = make([]int, 0)

			wantVal := true

			for {

				var tmp_uj__API int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__API type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__API = int(tval)

					}
				}

				uj.API = append(uj.API, tmp_uj__API)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BAttr:

	/* handler: uj.BAttr type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.BAttr = nil
		} else {

			uj.BAttr = make([]int, 0)

			wantVal := true

			for {

				var tmp_uj__BAttr int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__BAttr type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__BAttr = int(tval)

					}
				}

				uj.BAttr = append(uj.BAttr, tmp_uj__BAttr)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: uj.Ext type=json.RawMessage kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Ext = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.Ext == nil {
			uj.Ext = new(json.RawMessage)
		}

		err = uj.Ext.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
