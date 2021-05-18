package myshawdowsockets


type cipher struct {
	encodePassword *password
	decodePassword *password
}

func (cipher *cipher) Encode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

func (cipher *cipher) Decode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

func NewCipher(encodePassword *password) *cipher {
	decodePassword := &password{}
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}
	return &cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}