package cache

import "encoding/json"

func (c *ObjectCache[T]) unmarshal(data []byte, out any) error {
	if c.serializationType == SerializationTypeJson {
		return json.Unmarshal(data, out)
	}

	return nil
}

func (c *ObjectCache[T]) marshal(in any) ([]byte, error) {
	if c.serializationType == SerializationTypeJson {
		return json.Marshal(in)
	}

	return []byte{}, nil
}
