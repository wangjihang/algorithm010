package reverse_bits

func reverseBits(num uint32) uint32 {
	var res uint32
	for power := 31; num > 0; num >>= 1 {
		res += (num & 1) << power // 最低位
		power--
	}
	return res
}
