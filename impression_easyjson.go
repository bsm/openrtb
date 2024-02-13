// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

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

func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV3(in *jlexer.Lexer, out *Impression) {
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
		case "id":
			out.ID = string(in.String())
		case "banner":
			if in.IsNull() {
				in.Skip()
				out.Banner = nil
			} else {
				if out.Banner == nil {
					out.Banner = new(Banner)
				}
				easyjson7ebaa60bDecodeGithubComBsmOpenrtbV31(in, out.Banner)
			}
		case "video":
			if in.IsNull() {
				in.Skip()
				out.Video = nil
			} else {
				if out.Video == nil {
					out.Video = new(Video)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Video).UnmarshalJSON(data))
				}
			}
		case "audio":
			if in.IsNull() {
				in.Skip()
				out.Audio = nil
			} else {
				if out.Audio == nil {
					out.Audio = new(Audio)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Audio).UnmarshalJSON(data))
				}
			}
		case "native":
			if in.IsNull() {
				in.Skip()
				out.Native = nil
			} else {
				if out.Native == nil {
					out.Native = new(Native)
				}
				easyjson7ebaa60bDecodeGithubComBsmOpenrtbV32(in, out.Native)
			}
		case "pmp":
			if in.IsNull() {
				in.Skip()
				out.PMP = nil
			} else {
				if out.PMP == nil {
					out.PMP = new(PMP)
				}
				easyjson7ebaa60bDecodeGithubComBsmOpenrtbV33(in, out.PMP)
			}
		case "displaymanager":
			out.DisplayManager = string(in.String())
		case "displaymanagerver":
			out.DisplayManagerVersion = string(in.String())
		case "instl":
			out.Interstitial = int(in.Int())
		case "tagid":
			out.TagID = string(in.String())
		case "bidfloor":
			out.BidFloor = float64(in.Float64())
		case "bidfloorcur":
			out.BidFloorCurrency = string(in.String())
		case "secure":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Secure).UnmarshalJSON(data))
			}
		case "qty":
			if in.IsNull() {
				in.Skip()
				out.Quantity = nil
			} else {
				if out.Quantity == nil {
					out.Quantity = new(Quantity)
				}
				easyjson7ebaa60bDecodeGithubComBsmOpenrtbV34(in, out.Quantity)
			}
		case "exp":
			out.Exp = int(in.Int())
		case "iframebuster":
			if in.IsNull() {
				in.Skip()
				out.IFrameBusters = nil
			} else {
				in.Delim('[')
				if out.IFrameBusters == nil {
					if !in.IsDelim(']') {
						out.IFrameBusters = make([]string, 0, 4)
					} else {
						out.IFrameBusters = []string{}
					}
				} else {
					out.IFrameBusters = (out.IFrameBusters)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.IFrameBusters = append(out.IFrameBusters, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV3(out *jwriter.Writer, in Impression) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Banner != nil {
		const prefix string = ",\"banner\":"
		out.RawString(prefix)
		easyjson7ebaa60bEncodeGithubComBsmOpenrtbV31(out, *in.Banner)
	}
	if in.Video != nil {
		const prefix string = ",\"video\":"
		out.RawString(prefix)
		out.Raw((*in.Video).MarshalJSON())
	}
	if in.Audio != nil {
		const prefix string = ",\"audio\":"
		out.RawString(prefix)
		out.Raw((*in.Audio).MarshalJSON())
	}
	if in.Native != nil {
		const prefix string = ",\"native\":"
		out.RawString(prefix)
		easyjson7ebaa60bEncodeGithubComBsmOpenrtbV32(out, *in.Native)
	}
	if in.PMP != nil {
		const prefix string = ",\"pmp\":"
		out.RawString(prefix)
		easyjson7ebaa60bEncodeGithubComBsmOpenrtbV33(out, *in.PMP)
	}
	if in.DisplayManager != "" {
		const prefix string = ",\"displaymanager\":"
		out.RawString(prefix)
		out.String(string(in.DisplayManager))
	}
	if in.DisplayManagerVersion != "" {
		const prefix string = ",\"displaymanagerver\":"
		out.RawString(prefix)
		out.String(string(in.DisplayManagerVersion))
	}
	if in.Interstitial != 0 {
		const prefix string = ",\"instl\":"
		out.RawString(prefix)
		out.Int(int(in.Interstitial))
	}
	if in.TagID != "" {
		const prefix string = ",\"tagid\":"
		out.RawString(prefix)
		out.String(string(in.TagID))
	}
	if in.BidFloor != 0 {
		const prefix string = ",\"bidfloor\":"
		out.RawString(prefix)
		out.Float64(float64(in.BidFloor))
	}
	if in.BidFloorCurrency != "" {
		const prefix string = ",\"bidfloorcur\":"
		out.RawString(prefix)
		out.String(string(in.BidFloorCurrency))
	}
	if in.Secure != 0 {
		const prefix string = ",\"secure\":"
		out.RawString(prefix)
		out.Int(int(in.Secure))
	}
	if in.Quantity != nil {
		const prefix string = ",\"qty\":"
		out.RawString(prefix)
		easyjson7ebaa60bEncodeGithubComBsmOpenrtbV34(out, *in.Quantity)
	}
	if in.Exp != 0 {
		const prefix string = ",\"exp\":"
		out.RawString(prefix)
		out.Int(int(in.Exp))
	}
	if len(in.IFrameBusters) != 0 {
		const prefix string = ",\"iframebuster\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.IFrameBusters {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Impression) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7ebaa60bEncodeGithubComBsmOpenrtbV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Impression) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7ebaa60bEncodeGithubComBsmOpenrtbV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Impression) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7ebaa60bDecodeGithubComBsmOpenrtbV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Impression) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7ebaa60bDecodeGithubComBsmOpenrtbV3(l, v)
}
func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV34(in *jlexer.Lexer, out *Quantity) {
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
		case "multiplier":
			out.Multiplier = float64(in.Float64())
		case "sourcetype":
			out.SourceType = MeasurementSourceType(in.Int())
		case "vendor":
			out.Vendor = string(in.String())
		case "ext":
			if in.IsNull() {
				in.Skip()
				out.Ext = nil
			} else {
				if out.Ext == nil {
					out.Ext = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Ext).UnmarshalJSON(data))
				}
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV34(out *jwriter.Writer, in Quantity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"multiplier\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Multiplier))
	}
	if in.SourceType != 0 {
		const prefix string = ",\"sourcetype\":"
		out.RawString(prefix)
		out.Int(int(in.SourceType))
	}
	if in.Vendor != "" {
		const prefix string = ",\"vendor\":"
		out.RawString(prefix)
		out.String(string(in.Vendor))
	}
	if in.Ext != nil {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((*in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV33(in *jlexer.Lexer, out *PMP) {
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
		case "private_auction":
			out.Private = int(in.Int())
		case "deals":
			if in.IsNull() {
				in.Skip()
				out.Deals = nil
			} else {
				in.Delim('[')
				if out.Deals == nil {
					if !in.IsDelim(']') {
						out.Deals = make([]Deal, 0, 0)
					} else {
						out.Deals = []Deal{}
					}
				} else {
					out.Deals = (out.Deals)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Deal
					if data := in.Raw(); in.Ok() {
						in.AddError((v4).UnmarshalJSON(data))
					}
					out.Deals = append(out.Deals, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV33(out *jwriter.Writer, in PMP) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Private != 0 {
		const prefix string = ",\"private_auction\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Private))
	}
	if len(in.Deals) != 0 {
		const prefix string = ",\"deals\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Deals {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Raw((v6).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV32(in *jlexer.Lexer, out *Native) {
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
		case "request":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Request).UnmarshalJSON(data))
			}
		case "ver":
			out.Version = string(in.String())
		case "api":
			if in.IsNull() {
				in.Skip()
				out.APIs = nil
			} else {
				in.Delim('[')
				if out.APIs == nil {
					if !in.IsDelim(']') {
						out.APIs = make([]APIFramework, 0, 8)
					} else {
						out.APIs = []APIFramework{}
					}
				} else {
					out.APIs = (out.APIs)[:0]
				}
				for !in.IsDelim(']') {
					var v7 APIFramework
					v7 = APIFramework(in.Int())
					out.APIs = append(out.APIs, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "battr":
			if in.IsNull() {
				in.Skip()
				out.BlockedAttrs = nil
			} else {
				in.Delim('[')
				if out.BlockedAttrs == nil {
					if !in.IsDelim(']') {
						out.BlockedAttrs = make([]CreativeAttribute, 0, 8)
					} else {
						out.BlockedAttrs = []CreativeAttribute{}
					}
				} else {
					out.BlockedAttrs = (out.BlockedAttrs)[:0]
				}
				for !in.IsDelim(']') {
					var v8 CreativeAttribute
					v8 = CreativeAttribute(in.Int())
					out.BlockedAttrs = append(out.BlockedAttrs, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV32(out *jwriter.Writer, in Native) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		out.RawString(prefix[1:])
		out.Raw((in.Request).MarshalJSON())
	}
	if in.Version != "" {
		const prefix string = ",\"ver\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	if len(in.APIs) != 0 {
		const prefix string = ",\"api\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v9, v10 := range in.APIs {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlockedAttrs) != 0 {
		const prefix string = ",\"battr\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.BlockedAttrs {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v12))
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV31(in *jlexer.Lexer, out *Banner) {
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
		case "w":
			out.Width = int(in.Int())
		case "h":
			out.Height = int(in.Int())
		case "format":
			if in.IsNull() {
				in.Skip()
				out.Formats = nil
			} else {
				in.Delim('[')
				if out.Formats == nil {
					if !in.IsDelim(']') {
						out.Formats = make([]Format, 0, 1)
					} else {
						out.Formats = []Format{}
					}
				} else {
					out.Formats = (out.Formats)[:0]
				}
				for !in.IsDelim(']') {
					var v13 Format
					easyjson7ebaa60bDecodeGithubComBsmOpenrtbV35(in, &v13)
					out.Formats = append(out.Formats, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wmax":
			out.WidthMax = int(in.Int())
		case "hmax":
			out.HeightMax = int(in.Int())
		case "wmin":
			out.WidthMin = int(in.Int())
		case "hmin":
			out.HeightMin = int(in.Int())
		case "id":
			out.ID = string(in.String())
		case "btype":
			if in.IsNull() {
				in.Skip()
				out.BlockedTypes = nil
			} else {
				in.Delim('[')
				if out.BlockedTypes == nil {
					if !in.IsDelim(']') {
						out.BlockedTypes = make([]BannerType, 0, 8)
					} else {
						out.BlockedTypes = []BannerType{}
					}
				} else {
					out.BlockedTypes = (out.BlockedTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v14 BannerType
					v14 = BannerType(in.Int())
					out.BlockedTypes = append(out.BlockedTypes, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "battr":
			if in.IsNull() {
				in.Skip()
				out.BlockedAttrs = nil
			} else {
				in.Delim('[')
				if out.BlockedAttrs == nil {
					if !in.IsDelim(']') {
						out.BlockedAttrs = make([]CreativeAttribute, 0, 8)
					} else {
						out.BlockedAttrs = []CreativeAttribute{}
					}
				} else {
					out.BlockedAttrs = (out.BlockedAttrs)[:0]
				}
				for !in.IsDelim(']') {
					var v15 CreativeAttribute
					v15 = CreativeAttribute(in.Int())
					out.BlockedAttrs = append(out.BlockedAttrs, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pos":
			out.Position = AdPosition(in.Int())
		case "mimes":
			if in.IsNull() {
				in.Skip()
				out.MIMEs = nil
			} else {
				in.Delim('[')
				if out.MIMEs == nil {
					if !in.IsDelim(']') {
						out.MIMEs = make([]string, 0, 4)
					} else {
						out.MIMEs = []string{}
					}
				} else {
					out.MIMEs = (out.MIMEs)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.MIMEs = append(out.MIMEs, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "topframe":
			out.TopFrame = int(in.Int())
		case "expdir":
			if in.IsNull() {
				in.Skip()
				out.ExpDirs = nil
			} else {
				in.Delim('[')
				if out.ExpDirs == nil {
					if !in.IsDelim(']') {
						out.ExpDirs = make([]ExpDir, 0, 8)
					} else {
						out.ExpDirs = []ExpDir{}
					}
				} else {
					out.ExpDirs = (out.ExpDirs)[:0]
				}
				for !in.IsDelim(']') {
					var v17 ExpDir
					v17 = ExpDir(in.Int())
					out.ExpDirs = append(out.ExpDirs, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "api":
			if in.IsNull() {
				in.Skip()
				out.APIs = nil
			} else {
				in.Delim('[')
				if out.APIs == nil {
					if !in.IsDelim(']') {
						out.APIs = make([]APIFramework, 0, 8)
					} else {
						out.APIs = []APIFramework{}
					}
				} else {
					out.APIs = (out.APIs)[:0]
				}
				for !in.IsDelim(']') {
					var v18 APIFramework
					v18 = APIFramework(in.Int())
					out.APIs = append(out.APIs, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "vcm":
			out.VCM = int(in.Int())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV31(out *jwriter.Writer, in Banner) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Width != 0 {
		const prefix string = ",\"w\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"h\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Height))
	}
	if len(in.Formats) != 0 {
		const prefix string = ",\"format\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.Formats {
				if v19 > 0 {
					out.RawByte(',')
				}
				easyjson7ebaa60bEncodeGithubComBsmOpenrtbV35(out, v20)
			}
			out.RawByte(']')
		}
	}
	if in.WidthMax != 0 {
		const prefix string = ",\"wmax\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.WidthMax))
	}
	if in.HeightMax != 0 {
		const prefix string = ",\"hmax\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.HeightMax))
	}
	if in.WidthMin != 0 {
		const prefix string = ",\"wmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.WidthMin))
	}
	if in.HeightMin != 0 {
		const prefix string = ",\"hmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.HeightMin))
	}
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if len(in.BlockedTypes) != 0 {
		const prefix string = ",\"btype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.BlockedTypes {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v22))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlockedAttrs) != 0 {
		const prefix string = ",\"battr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v23, v24 := range in.BlockedAttrs {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v24))
			}
			out.RawByte(']')
		}
	}
	if in.Position != 0 {
		const prefix string = ",\"pos\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Position))
	}
	if len(in.MIMEs) != 0 {
		const prefix string = ",\"mimes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v25, v26 := range in.MIMEs {
				if v25 > 0 {
					out.RawByte(',')
				}
				out.String(string(v26))
			}
			out.RawByte(']')
		}
	}
	if in.TopFrame != 0 {
		const prefix string = ",\"topframe\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.TopFrame))
	}
	if len(in.ExpDirs) != 0 {
		const prefix string = ",\"expdir\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v27, v28 := range in.ExpDirs {
				if v27 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v28))
			}
			out.RawByte(']')
		}
	}
	if len(in.APIs) != 0 {
		const prefix string = ",\"api\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v29, v30 := range in.APIs {
				if v29 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v30))
			}
			out.RawByte(']')
		}
	}
	if in.VCM != 0 {
		const prefix string = ",\"vcm\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.VCM))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson7ebaa60bDecodeGithubComBsmOpenrtbV35(in *jlexer.Lexer, out *Format) {
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
		case "w":
			out.Width = int(in.Int())
		case "h":
			out.Height = int(in.Int())
		case "wratio":
			out.WidthRatio = int(in.Int())
		case "hration":
			out.HeightRatio = int(in.Int())
		case "wmin":
			out.WidthMin = int(in.Int())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjson7ebaa60bEncodeGithubComBsmOpenrtbV35(out *jwriter.Writer, in Format) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Width != 0 {
		const prefix string = ",\"w\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.Width))
	}
	if in.Height != 0 {
		const prefix string = ",\"h\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Height))
	}
	if in.WidthRatio != 0 {
		const prefix string = ",\"wratio\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.WidthRatio))
	}
	if in.HeightRatio != 0 {
		const prefix string = ",\"hration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.HeightRatio))
	}
	if in.WidthMin != 0 {
		const prefix string = ",\"wmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.WidthMin))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
