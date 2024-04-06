package hasher

// Djb2 is a hashing function that calculates the hash value for the given byte slice key
// using the DJB2 algorithm.
//
// The DJB2 algorithm, developed by Daniel J. Bernstein, is a simple and efficient hash
// function. It iterates through each byte of the key, updating the hash value
// based on the ASCII value of the characters and a prime constant.
//
// - key: The byte slice representing the key to be hashed.
//
// - The hash value calculated for the given key as an integer.
func Djb2(key []byte) int {
	// Initial prime value
	hash := 5381
	for _, c := range key {
		char_code := int(c)

		// (hash<<5) means hash*(2^5)
		hash = ((hash << 5) + hash) + char_code
	}
	return hash
}
