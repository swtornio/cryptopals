# Write a function that takes two equal-length buffers and produces their XOR combination.

# If your function works properly, then when you feed it the string:

# 1c0111001f010100061a024b53535009181c
# ... after hex decoding, and when XOR'd against:

# 686974207468652062756c6c277320657965
# ... should produce:

# 746865206b696420646f6e277420706c6179

def xor_byte_strings(input1, input2):
	"""XOR two byte strings
	"""
	return bytes([b1 ^ b2 for b1, b2 in zip(input1, input2)])

def main():
	string1 = "1c0111001f010100061a024b53535009181c"
	string2 = "686974207468652062756c6c277320657965"
	expected_xor = "746865206b696420646f6e277420706c6179"

	string1_bytes = bytes.fromhex(string1)
	string2_bytes = bytes.fromhex(string2)
	print(xor_byte_strings(string1_bytes, string2_bytes).hex())

	print("Should be\n746865206b696420646f6e277420706c6179")


# If challenge1.py is run (instead of imported)
# call the main() function
if __name__=='__main__':
	main()