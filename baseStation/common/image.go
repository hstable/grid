package common

func IsImage(img []byte) bool {
	var h [2]byte
	copy(h[:], img[:2])
	//255216是jpg;7173是gif;6677是BMP,13780是PNG
	switch h {
	case [2]byte{255, 216}, [2]byte{71, 73}, [2]byte{66, 77}, [2]byte{137, 80}:
		return true
	}
	return false
}
