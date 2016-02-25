package bencode

import (
	"fmt"
	"unicode/utf8"
)

func EncodeInt(x int) string {
	return fmt.Sprintf("i%de", x)
}

func EncodeList(x []string) string {
	tmp := "l"
	for i := range x {
		tmp = fmt.Sprintf("%s%s", tmp, EncodeByteString(x[i]))
	}
	tmp = fmt.Sprintf("%se", tmp)
	return tmp
}

func EncodeDictionary(key string, value string) string {
	if value[0] == 'l' || value[0] == 'd' {
		return fmt.Sprintf("d%s%se", EncodeByteString(key), value)
	}
	// TODO(ian): Allow for detection of integers
	// this only supports strings and list/dicts right now
	return fmt.Sprintf("d%s%se", EncodeByteString(key), EncodeByteString(value))
}

func EncodeByteString(x string) string {
	return fmt.Sprintf("%d:%s", utf8.RuneCountInString(x), x)
}

func EncodePeerList(x []string, interval int) string {
	peerList := EncodeList(x)
	// intDictionary := EncodeDictionary("interval", EncodeInt(interval))
	peers := EncodeDictionary("peers", peerList)
	return peers
}
