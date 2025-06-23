package main

import (
	"bytes"
	"strings"
)

type Codec struct {
}

// ### Escaping approach

const delimiter string = "/:"

func (c Codec) Encode(strs []string) string {
	var buf bytes.Buffer
	for _, str := range strs {
		str = strings.ReplaceAll(str, "/", "//")
		buf.WriteString(str)
		buf.WriteString(delimiter)
	}

	return buf.String()
}

func (c Codec) Decode(strs string) []string {
	var result []string
	var str string
	for i := 0; i < len(strs); i++ {
		if strs[i] == '/' && strs[i+1] == ':' {
			result = append(result, str)
			str = ""
			i++
		} else if strs[i] == '/' && strs[i+1] == '/' {
			str += "/"
			i++
		} else {
			str += string(strs[i])
		}
	}
	return result
}

// ### JSON Approach

// Encode encodes a list of strings to a single string.
//func (codec *Codec) Encode(strs []string) string {
//	jsonData, err := json.Marshal(strs)
//	if err != nil {
//		slog.Error(err.Error())
//		os.Exit(1)
//	}
//
//	return base64.StdEncoding.EncodeToString(jsonData)
//}

// Decode decodes a single string to a list of strings.
//func (codec *Codec) Decode(strs string) []string {
//	jsonData, err := base64.StdEncoding.DecodeString(strs)
//	if err != nil {
//		slog.Error(err.Error())
//		os.Exit(1)
//	}
//
//	var items []string
//	if err := json.Unmarshal(jsonData, &items); err != nil {
//		slog.Error(err.Error())
//		os.Exit(1)
//	}
//
//	return items
//}

func EncodeDecode(strs []string) []string {
	var codec Codec
	return codec.Decode(codec.Encode(strs))
}
