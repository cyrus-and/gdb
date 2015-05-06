package gdb

import (
	"encoding/json"
	"reflect"
	"testing"
)

func checkParseRecord(t *testing.T, input, expected string) {
	inputObj := parseRecord(input)
	expectedObj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(inputObj, expectedObj) {
		t.Fatal(inputObj, "!=", expectedObj)
	}
}

func TestParseRecord(t *testing.T) {
	// async record
	checkParseRecord(t, `*foo`, `{"type":"exec","class":"foo"}`)
	checkParseRecord(t, `+bar`, `{"type":"status","class":"bar"}`)
	checkParseRecord(t, `=baz`, `{"type":"notify","class":"baz"}`)
	checkParseRecord(t, `*foo,a="x"`, `{"type":"exec","class":"foo","payload":{"a":"x"}}`)
	checkParseRecord(t, `*foo,a="x",b="y"`, `{"type":"exec","class":"foo","payload":{"a":"x","b":"y"}}`)
	checkParseRecord(t, `+bar,a="x"`, `{"type":"status","class":"bar","payload":{"a":"x"}}`)
	checkParseRecord(t, `+bar,a="x",b="y"`, `{"type":"status","class":"bar","payload":{"a":"x","b":"y"}}`)
	checkParseRecord(t, `=baz,a="x"`, `{"type":"notify","class":"baz","payload":{"a":"x"}}`)
	checkParseRecord(t, `=baz,a="x",b="y"`, `{"type":"notify","class":"baz","payload":{"a":"x","b":"y"}}`)

	// stream record
	checkParseRecord(t, `~"foo"`, `{"type":"console","payload":"foo"}`)
	checkParseRecord(t, `@"bar"`, `{"type":"target","payload":"bar"}`)
	checkParseRecord(t, `&"baz"`, `{"type":"log","payload":"baz"}`)

	// result record
	checkParseRecord(t, `123^foo`, `{"sequence":"123","class":"foo"}`)
	checkParseRecord(t, `456^bar,a="x"`, `{"sequence":"456","class":"bar","payload":{"a":"x"}}`)
	checkParseRecord(t, `789^baz,a="x",b="y"`, `{"sequence":"789","class":"baz","payload":{"a":"x","b":"y"}}`)

	// complex record payload
	checkParseRecord(t, `000^foo,a={"b"="x","c"="y"}`, `{"sequence":"000","class":"foo","payload":{"a":{"b":"x","c":"y"}}}`)
	checkParseRecord(t, `000^foo,a=["b"="x","c"="y"]`, `{"sequence":"000","class":"foo","payload":{"a":[{"b":"x"},{"c":"y"}]}}`)
	checkParseRecord(t, `000^foo,a=["x","y"]`, `{"sequence":"000","class":"foo","payload":{"a":["x","y"]}}`)
	checkParseRecord(t, `000^foo,a=[[["x"]]]`, `{"sequence":"000","class":"foo","payload":{"a":[[["x"]]]}}`)
	checkParseRecord(t, `000^foo,a={b={c={x="y"}}}`, `{"sequence":"000","class":"foo","payload":{"a":{"b":{"c":{"x":"y"}}}}}`)
	checkParseRecord(t, `000^foo,a={}`, `{"sequence":"000","class":"foo","payload":{"a":{}}}`)
	checkParseRecord(t, `000^foo,a=[]`, `{"sequence":"000","class":"foo","payload":{"a":[]}}`)

	// escape sequences and strings in general
	checkParseRecord(t, `~""`, `{"type":"console","payload":""}`)
	checkParseRecord(t, `~"\b\f\n\r\t\""`, `{"type":"console","payload":"\b\f\n\r\t\""}`)
	checkParseRecord(t, `@""`, `{"type":"target","payload":""}`)
	checkParseRecord(t, `@"\b\f\n\r\t\""`, `{"type":"target","payload":"\b\f\n\r\t\""}`)
	checkParseRecord(t, `&""`, `{"type":"log","payload":""}`)
	checkParseRecord(t, `&"\b\f\n\r\t\""`, `{"type":"log","payload":"\b\f\n\r\t\""}`)
	checkParseRecord(t, `000^foo,a=""`, `{"sequence":"000","class":"foo","payload":{"a":""}}`)
	checkParseRecord(t, `000^bar,a="\b\f\n\r\t\""`, `{"sequence":"000","class":"bar","payload":{"a":"\b\f\n\r\t\""}}`)
}
