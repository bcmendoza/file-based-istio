package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

func ErrorExit(msg string, a ...interface{}) {
	fmt.Printf(msg, a)
	fmt.Println()
	os.Exit(1)
}

func MarshallJson(w proto.Message) []byte {
	buffer := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{}).Marshal(buffer, w)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

func MarshallYaml(w proto.Message) []byte {
	b, _ := yaml.JSONToYAML([]byte(MarshallJson(w)))
	return b
}

func SanitizeName(name string) string {
	return strings.ReplaceAll(name, "|", "_.")
}
