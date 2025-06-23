package main

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"os"
)

type Codec struct {
}

// Encode encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	jsonData, err := json.Marshal(strs)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return base64.StdEncoding.EncodeToString(jsonData)
}

// Decode decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	jsonData, err := base64.StdEncoding.DecodeString(strs)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var items []string
	if err := json.Unmarshal(jsonData, &items); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return items
}

func EncodeDecode(strs []string) []string {
	var codec Codec
	return codec.Decode(codec.Encode(strs))
}
