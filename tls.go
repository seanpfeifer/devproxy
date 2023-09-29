package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"time"
)

func generateTLSConfig(hostIP string) (*tls.Config, error) {
	ip := net.ParseIP(hostIP)
	if ip == nil {
		return nil, fmt.Errorf("parsing IP address: %s", hostIP)
	}

	curve := elliptic.P256()

	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generating ECDSA key: %w", err)
	}

	c := &x509.Certificate{
		SerialNumber: big.NewInt(0),
		NotBefore:    time.Now().Add(-24 * time.Hour),
		NotAfter:     time.Now().AddDate(3, 0, 0), // Valid for 3 years from app start
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
		IsCA:                  true, // We are our own CA
		MaxPathLenZero:        true,
		IPAddresses:           []net.IP{ip},
	}

	certData, err := x509.CreateCertificate(rand.Reader, c, c, key.Public(), key)
	if err != nil {
		return nil, fmt.Errorf("creating certificate: %w", err)
	}

	cert, err := certDataToCertificate(certData, key)
	if err != nil {
		return nil, fmt.Errorf("converting data to certificate: %w", err)
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func certDataToCertificate(certData []byte, key *ecdsa.PrivateKey) (tls.Certificate, error) {
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certData,
	})

	keyData, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return tls.Certificate{}, err
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "ECDSA PRIVATE KEY",
		Bytes: keyData,
	})

	return tls.X509KeyPair(certPEM, keyPEM)
}
