package common

import (
	"bytes"
	"encoding/gob"
)

func ObjectToByte(obj interface{}) []byte {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)
	if err := e.Encode(obj); err != nil {
		panic(err)
	}

	// buf := &bytes.Buffer{}
	// binary.Write(buf, binary.BigEndian, obj)
	return buf.Bytes()
}

func ByteToObject(src []byte, dest *Message) error {
	buf := bytes.NewBuffer(src)
	d := gob.NewDecoder(buf)
	var err error = nil
	if err = d.Decode(&dest); err != nil {
		panic(err)
	}
	// buf := &bytes.Buffer{}
	// err := binary.Read(buf, binary.BigEndian, &dest)
	return err
}
