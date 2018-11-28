package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

const (
	ecdsaPref = "ecdsa_cert"
	rsaPref   = "rsa_cert"
)

func publicKey(priv *rsa.PrivateKey) interface{} {

	return &priv.PublicKey

}

func pemBlockForKey(priv *rsa.PrivateKey) *pem.Block {

	return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)}

}

// GenerateCertFiles generates a certificate files
func GenerateCertFiles(hosts []string) error {

	if len(hosts) == 0 {
		return fmt.Errorf("No hosts")
	}
	priv, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	notBefore := time.Now().Add(-1 * time.Hour)

	notAfter := notBefore.Add(30 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	template.IsCA = true
	template.KeyUsage |= x509.KeyUsageCertSign

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	certOut, err := os.Create(".tmp/cert.pem")
	if err != nil {
		log.Fatalf("failed to open cert.pem for writing: %s", err)
	}
	defer certOut.Close()
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		log.Fatalf("failed to write cert.pem: %s", err)
	}
	log.Print("written cert.pem\n")

	keyOut, err := os.OpenFile(".tmp/key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {

		return fmt.Errorf("failed to open key.pem for writing: %s", err)
	}
	defer keyOut.Close()
	err = pem.Encode(keyOut, pemBlockForKey(priv))
	if err != nil {
		return fmt.Errorf("failed to write key.pem: %s", err)

	}
	log.Print("written key.pem\n")
	return nil
}
