package crypto

/*
#cgo CFLAGS:-I../ecdsacmp/libcmp
#cgo LDFLAGS:-L../ecdsacmp/libcmp
#cgo CFLAGS:-DMCLBN_FP_UNIT_SIZE=6
#cgo LDFLAGS:-lcmp -lcrypto -lstdc++

typedef unsigned int (*ReadRandFunc)(void *, void *, unsigned int);
int wrapReadRandCgo(void *self, void *buf, unsigned int n);
#include <cmp_protocol.h>
#include <paillier_cryptosystem.h>
#include <algebraic_elements.h>
*/
import "C"
import (
	"fmt"
	"math/big"
	"unsafe"
)

func Hello() error {
	C.hello()
	return nil
}

func GeneratePaillierKeyPair(primeBits int, safe_prime_flag int, primeModule int) (N, p, q, phiN, lamda *big.Int, err error) {
	//fmt.Println("begin GeneratePaillierKeyPair")
	//C.hello()
	var privkey *C.paillier_private_key_t
	privkey = C.paillier_encryption_private_new()
	C.paillier_encryption_generate_private_ex(privkey, C.uint64_t(primeBits), C.int(safe_prime_flag))
	//fmt.Printf("buf %x \n", privkey)

	//uint8_t **bytes, uint64_t *byte_len, const paillier_private_key_t *priv, uint64_t paillier_modulus_bytes, int move_to_end
	var byte_len C.ulong

	C.paillier_struct_to_bytes(nil, &byte_len, nil, C.uint64_t(primeModule))
	//fmt.Printf("byte_len: %d \n", byte_len)
	buf := (*C.uchar)(C.malloc(byte_len))
	defer C.free(unsafe.Pointer(buf))
	var prevPtr *C.uchar
	prevPtr = buf
	C.paillier_struct_to_bytes((**C.uchar)(unsafe.Pointer(&buf)), &byte_len, privkey, C.uint64_t(primeModule))
	/*ps := (*[1536]C.uchar)(unsafe.Pointer(prevPtr))[:1536]
	fmt.Printf("%x \n", ps)*/
	convertedBytes := charToBytes(prevPtr, 6*primeBits)
	//fmt.Printf("%x \n", convertedBytes)
	/*privkeyHexstr := fmt.Sprintf("%x", convertedBytes)
	privkeyBytes, err := hex.DecodeString(privkeyHexstr)
	if err != nil {
		return
	}*/
	//fmt.Printf("privkeyBytes: %d \n", len(privkeyBytes))
	N = big.NewInt(0).SetBytes(convertedBytes[:primeModule])
	p = big.NewInt(0).SetBytes(convertedBytes[primeModule*2 : primeModule*3])
	q = big.NewInt(0).SetBytes(convertedBytes[primeModule*3 : primeModule*4])
	phiN = big.NewInt(0).SetBytes(convertedBytes[primeModule*4 : primeModule*5])
	lamda = big.NewInt(0).SetBytes(convertedBytes[primeModule*5 : primeModule*6])
	//fmt.Printf("p: %x \nq: %x \nN: %x \nphiN: %x \nlamda: %x \n", p.Bytes(), q.Bytes(), N.Bytes(), phiN.Bytes(), lamda.Bytes())
	//fmt.Printf("p is prime %v  q is prime %v \n", p.ProbablyPrime(30), q.ProbablyPrime(30))
	//fmt.Println("end GeneratePaillierKeyPair")
	return
}

func RandomPositiveBigInt(limit *big.Int) *big.Int {
	fmt.Printf("limit: %x \n", limit.Bytes())
	var limitScalar C.scalar_t
	var intLen uint64
	intLen = uint64(len(limit.Bytes()))
	fmt.Printf("intLen: %d \n", intLen)
	limitScalar = C.scalar_new()
	ptr := (*C.uint8_t)(C.CBytes(limit.Bytes()))
	defer C.free(unsafe.Pointer(ptr))
	C.scalar_from_bytes(limitScalar, (**C.uchar)(unsafe.Pointer(&ptr)), C.uint64_t(intLen), 0)
	buf := (*C.uchar)(C.malloc(C.size_t(intLen)))
	C.random_big_num((**C.uchar)(unsafe.Pointer(&buf)), C.uint64_t(intLen), limitScalar, 0)
	fmt.Printf("buf: %x \n", buf)
	randBitInt := big.NewInt(0).SetBytes(charToBytes(buf, len(limit.Bytes())))
	fmt.Printf("rand num: %x %d \n", randBitInt.Bytes(), randBitInt.Int64())
	return randBitInt
}

func charToBytes(src *C.uchar, sz int) []byte {
	return C.GoBytes(unsafe.Pointer(src), C.int(sz))
}

/*func charToBytes(src *C.uchar, sz C.int) []byte {
	return C.GoBytes(unsafe.Pointer(src), sz)
}*/
