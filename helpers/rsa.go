package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSA() (public, private interface{}) {
	bitSize := 4096

	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)
	// Encode private key to PKCS#1 ASN.1 PEM.
	privPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	//if err := ioutil.WriteFile("sapi"+".rsa", privPEM, 0700); err != nil {
	//	panic(err)
	//}
	//
	//// Write public key to file.
	//if err := ioutil.WriteFile("sapi"+".rsa.pub", pubPEM, 0755); err != nil {
	//	panic(err)
	//}
	return string(pubPEM), string(privPEM)
}
