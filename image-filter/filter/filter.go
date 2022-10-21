package filter

type Filter func(pixel []byte)

var Gray Filter = func(pixel []byte) {
	avg := (int32(pixel[0]) + int32(pixel[1]) + int32(pixel[2])) / 3
	pixel[0] = byte(avg)
	pixel[1] = byte(avg)
	pixel[2] = byte(avg)
}

var Thres Filter = func(pixel []byte) {
	color := byte(0)
	if ((int32(pixel[0]) + int32(pixel[1]) + int32(pixel[2])) / 3) > 127 {
		color = 0xff
	}
	pixel[0] = color
	pixel[1] = color
	pixel[2] = color
}
