package util

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"sort"
)

type OrderedMap[V any] struct {
	Lexical bool
	Order   []string
	Map     map[string]V
}

func NewOrderedMap[V any](lexical bool, capacity int) *OrderedMap[V] {
	return &OrderedMap[V]{Lexical: lexical, Order: make([]string, 0, capacity), Map: make(map[string]V, capacity)}
}

func (o *OrderedMap[V]) Append(k string, v V) {
	o.Order = append(o.Order, k)
	o.Map[k] = v
	if o.Lexical {
		sort.Strings(o.Order)
	}
}

func (o *OrderedMap[V]) Get(k string) (V, bool) {
	ret, ok := o.Map[k]
	return ret, ok
}

func (o *OrderedMap[V]) GetSimple(k string) V {
	return o.Map[k]
}

func (o OrderedMap[V]) MarshalYAML() (interface{}, error) {
	return o.Map, nil
}

func (o OrderedMap[V]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, key := range o.Order {
		n := xml.Name{Local: key}
		t := xml.StartElement{Name: n}

		v := o.Map[key]
		e := e.EncodeElement(v, t)
		if e != nil {
			return e
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return e.Flush()
}

func (o *OrderedMap[V]) UnmarshalJSON(b []byte) error {
	err := FromJSON(b, &o.Map)
	if err != nil {
		return err
	}

	index := make(map[string]int)
	for key := range o.Map {
		o.Order = append(o.Order, key)
		esc, _ := json.Marshal(key) // Escape the key
		index[key] = bytes.Index(b, esc)
	}

	if o.Lexical {
		sort.Slice(o.Order, func(i, j int) bool { return index[o.Order[i]] < index[o.Order[j]] })
	}
	return nil
}

func (o OrderedMap[V]) MarshalJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	buf.WriteByte('{')
	l := len(o.Order)
	for i, key := range o.Order {
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(km)
		buf.WriteByte(':')
		vm, err := json.Marshal(o.Map[key])
		if err != nil {
			return nil, err
		}
		buf.Write(vm)
		if i != l-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
