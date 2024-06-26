// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package TimeRules

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules(in *jlexer.Lexer, out *YearLimitsObj) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "min":
			out.Min = int64(in.Int64())
		case "max":
			out.Max = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules(out *jwriter.Writer, in YearLimitsObj) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Min != 0 {
		const prefix string = ",\"min\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Min))
	}
	if in.Max != 0 {
		const prefix string = ",\"max\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Max))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v YearLimitsObj) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v YearLimitsObj) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *YearLimitsObj) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *YearLimitsObj) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules(l, v)
}
func easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules1(in *jlexer.Lexer, out *TimeConfigurationRulesObj) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "inf":
			(out.INF).UnmarshalEasyJSON(in)
		case "format":
			(out.FormatsDef).UnmarshalEasyJSON(in)
		case "year":
			(out.Year).UnmarshalEasyJSON(in)
		case "maxHour":
			out.MaxHour = uint16(in.Uint16())
		case "maxMin":
			out.MaxMin = uint16(in.Uint16())
		case "month":
			if in.IsNull() {
				in.Skip()
				out.Month = nil
			} else {
				in.Delim('[')
				if out.Month == nil {
					if !in.IsDelim(']') {
						out.Month = make([]MonthObj, 0, 1)
					} else {
						out.Month = []MonthObj{}
					}
				} else {
					out.Month = (out.Month)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MonthObj
					(v1).UnmarshalEasyJSON(in)
					out.Month = append(out.Month, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "daysInYear":
			out.DaysInYear = uint64(in.Uint64())
		case "timezones":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Timezones = make(map[string]int64)
				} else {
					out.Timezones = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int64
					v2 = int64(in.Int64())
					(out.Timezones)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules1(out *jwriter.Writer, in TimeConfigurationRulesObj) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if true {
		const prefix string = ",\"inf\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.INF).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.FormatsDef).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"year\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Year).MarshalEasyJSON(out)
	}
	if in.MaxHour != 0 {
		const prefix string = ",\"maxHour\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint16(uint16(in.MaxHour))
	}
	if in.MaxMin != 0 {
		const prefix string = ",\"maxMin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint16(uint16(in.MaxMin))
	}
	if len(in.Month) != 0 {
		const prefix string = ",\"month\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Month {
				if v3 > 0 {
					out.RawByte(',')
				}
				(v4).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.DaysInYear != 0 {
		const prefix string = ",\"daysInYear\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.DaysInYear))
	}
	if len(in.Timezones) != 0 {
		const prefix string = ",\"timezones\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Timezones {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.Int64(int64(v5Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TimeConfigurationRulesObj) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TimeConfigurationRulesObj) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TimeConfigurationRulesObj) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TimeConfigurationRulesObj) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules1(l, v)
}
func easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules2(in *jlexer.Lexer, out *SystemInfoObj) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ver":
			out.Ver = string(in.String())
		case "creator":
			out.Creator = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules2(out *jwriter.Writer, in SystemInfoObj) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ver != "" {
		const prefix string = ",\"ver\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Ver))
	}
	if in.Creator != "" {
		const prefix string = ",\"creator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Creator))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SystemInfoObj) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SystemInfoObj) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SystemInfoObj) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SystemInfoObj) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules2(l, v)
}
func easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules3(in *jlexer.Lexer, out *MonthObj) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fullName":
			out.FullName = string(in.String())
		case "shortName":
			out.ShortName = string(in.String())
		case "days":
			out.Days = uint16(in.Uint16())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules3(out *jwriter.Writer, in MonthObj) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FullName != "" {
		const prefix string = ",\"fullName\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FullName))
	}
	if in.ShortName != "" {
		const prefix string = ",\"shortName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ShortName))
	}
	if in.Days != 0 {
		const prefix string = ",\"days\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint16(uint16(in.Days))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MonthObj) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MonthObj) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MonthObj) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MonthObj) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules3(l, v)
}
func easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules4(in *jlexer.Lexer, out *FormatsInfoObj) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			out.Date = string(in.String())
		case "time":
			out.Time = string(in.String())
		case "full":
			out.Full = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules4(out *jwriter.Writer, in FormatsInfoObj) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Date != "" {
		const prefix string = ",\"date\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Date))
	}
	if in.Time != "" {
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Time))
	}
	if in.Full != "" {
		const prefix string = ",\"full\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Full))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FormatsInfoObj) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FormatsInfoObj) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComBookshelfWriterTimeRules4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FormatsInfoObj) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FormatsInfoObj) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComBookshelfWriterTimeRules4(l, v)
}
