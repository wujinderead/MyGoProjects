package tlser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"math/big"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestCert(t *testing.T) {
	// generate ca rsa key
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("generate rsa key err:", err)
		return
	}

	// print key in pem
	caKeyPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caKey)})
	fmt.Println("private key: ", string(caKeyPem))

	// generate self certificate
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		fmt.Println("failed to generate serial number:", err)
		return
	}
	fmt.Println("ca serial number: ", hex.EncodeToString(serialNumber.Bytes()))

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"SelfCA"},
		},
		NotBefore: time.Now().Add(-time.Hour * 24 * 365),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caCertBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, caKey.Public(), caKey)
	if err != nil {
		fmt.Println("Failed to create certificate:", err)
		return
	}

	selfCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caCertBytes})
	fmt.Println("self cert:", string(selfCertPem))

	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		fmt.Println("parse certificate err:", err)
		return
	}
	displayCert(caCert)

	serverKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("generate key err:", err)
		return
	}

	serverTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"MyServer"},
		},
		NotBefore: time.Now().Add(-time.Hour * 24 * 365),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		IPAddresses:           []net.IP{net.ParseIP("10.19.138.22"), net.ParseIP("10.19.138.135")},
		DNSNames:              []string{"dog.com", "leopard.com"},
	}

	serverCertBytes, err := x509.CreateCertificate(rand.Reader, &serverTemplate, caCert, serverKey.Public(), caKey)
	if err != nil {
		fmt.Println("server cert err:", err)
		return
	}
	serverCert, err := x509.ParseCertificate(serverCertBytes)
	if err != nil {
		fmt.Println("parse certificate err:", err)
		return
	}
	displayCert(serverCert)
}

func TestParseRsaKeyCert(t *testing.T) {
	// parse rsa key
	rsaKeyPem := readFile("../keys/rsa.ca.key")
	block, _ := pem.Decode([]byte(rsaKeyPem))
	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("parse key err:", err)
		return
	}
	fmt.Println("rsa.D:", hex.EncodeToString(rsaKey.D.Bytes()))
	fmt.Println()

	// parse ca certificate
	caCertPem := readFile("../keys/rsa.ca.crt")
	block, _ = pem.Decode([]byte(caCertPem))
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(caCert)
	fmt.Println()

	// parse client certificate request
	serverCsrPem := readFile("../keys/rsa.server.csr")
	block, _ = pem.Decode([]byte(serverCsrPem))
	serverCsr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		fmt.Println("parse cert requset err:", err)
		return
	}
	displayCertReq(serverCsr)
	fmt.Println()

	// parse client certificate
	serverCertPem := readFile("../keys/rsa.server.crt")
	block, _ = pem.Decode([]byte(serverCertPem))
	serverCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(serverCert)
	fmt.Println()

	fmt.Println("csr check signature: ", serverCsr.CheckSignature())
	fmt.Println("ca cert signed server cert: ", caCert.CheckSignature(
		serverCert.SignatureAlgorithm, serverCert.RawTBSCertificate, serverCert.Signature))
}

func TestParseEd25519KeyCert(t *testing.T) {
	// parse ed25519 key
	type pkcs8 struct {
		Version    int
		Algo       pkix.AlgorithmIdentifier
		PrivateKey []byte
		// optional attributes omitted.
	}
	ed25519CaKeyPem := readFile("../keys/ed25519.ca.key")
	block, _ := pem.Decode([]byte(ed25519CaKeyPem))
	var privKey pkcs8
	if _, err := asn1.Unmarshal(block.Bytes, &privKey); err != nil {
		fmt.Println("parse pkcs8 err:", err)
		return
	}
	fmt.Println("pkcs8 version:", privKey.Version)
	fmt.Println("pkcs8 algo:", privKey.Algo.Algorithm.String())
	fmt.Println("pkcs8 algo:", privKey.Algo.Parameters.FullBytes)
	ed25519Priv := ed25519.PrivateKey(privKey.PrivateKey)
	fmt.Println("ed25519 priv:", hex.EncodeToString(ed25519Priv))
	fmt.Println()

	// parse ca certificate
	caCertPem := readFile("../keys/ed25519.ca.crt")
	block, _ = pem.Decode([]byte(caCertPem))
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(caCert)
	fmt.Println()

	// parse client certificate request
	serverCsrPem := readFile("../keys/ed25519.server.csr")
	block, _ = pem.Decode([]byte(serverCsrPem))
	serverCsr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		fmt.Println("parse cert requset err:", err)
		return
	}
	displayCertReq(serverCsr)
	fmt.Println()

	// parse client certificate
	serverCertPem := readFile("../keys/ed25519.server.crt")
	block, _ = pem.Decode([]byte(serverCertPem))
	serverCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(serverCert)
	fmt.Println()
}

func TestParseEcdsaKeyCert(t *testing.T) {
	// parse rsa key
	caKeyPem := readFile("../keys/ec.ca.key")
	_, rest := pem.Decode([]byte(caKeyPem))
	block, _ := pem.Decode([]byte(rest))
	caKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("parse key err:", err)
		return
	}
	fmt.Println("ec.D:", hex.EncodeToString(caKey.D.Bytes()))
	fmt.Println()

	// parse ca certificate
	caCertPem := readFile("../keys/ec.ca.crt")
	block, _ = pem.Decode([]byte(caCertPem))
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(caCert)
	fmt.Println()

	// parse client certificate request
	serverCsrPem := readFile("../keys/ec.client.csr")
	block, _ = pem.Decode([]byte(serverCsrPem))
	serverCsr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		fmt.Println("parse cert requset err:", err)
		return
	}
	displayCertReq(serverCsr)
	fmt.Println()

	// parse client certificate
	serverCertPem := readFile("../keys/ec.client.crt")
	block, _ = pem.Decode([]byte(serverCertPem))
	serverCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(serverCert)
	fmt.Println()

	fmt.Println("check csr signature:", serverCsr.CheckSignature())
	fmt.Println("check ca signed server cert:", caCert.CheckSignature(
		caCert.SignatureAlgorithm, serverCert.RawTBSCertificate, serverCert.Signature))
}

func TestParseGoogleCert(t *testing.T) {
	// get google.com certificate
	googleCertPem := readFile("../keys/www.google.com.crt")
	block, _ := pem.Decode([]byte(googleCertPem))
	googleCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("get cert err:", err)
		return
	}
	displayCert(googleCert)
	fmt.Println()

	// get google CA certificate
	googleCaCertDer := readFile("../keys/GTSGIAG3.crt")
	googleCaCert, err := x509.ParseCertificate(googleCaCertDer)
	if err != nil {
		fmt.Println("parse der crl err:", err)
		return
	}
	displayCert(googleCaCert)
	fmt.Println()

	// verify signature
	fmt.Println("verify googleCaCert signed googleCert: ", googleCaCert.CheckSignature(
		googleCert.SignatureAlgorithm, googleCert.RawTBSCertificate, googleCert.Signature))

	// get certificate revocation list
	googleCaCrlDer := readFile("../keys/GTSGIAG3.crl")
	googleCaCrl, err := x509.ParseDERCRL(googleCaCrlDer)
	if err != nil {
		fmt.Println("parse der crl err:", err)
		return
	}
	displayCertificateList(googleCaCrl)
}

func displayCert(cert *x509.Certificate) {
	//fmt.Println("raw:", hex.EncodeToString(cert.Raw))
	//fmt.Println("raw tbs:", hex.EncodeToString(cert.RawTBSCertificate))
	//fmt.Println("raw pubk:", hex.EncodeToString(cert.RawSubjectPublicKeyInfo))
	//fmt.Println("raw sub:", hex.EncodeToString(cert.RawSubject))
	//fmt.Println("raw issue:", hex.EncodeToString(cert.RawIssuer))

	fmt.Println("signature:", hex.EncodeToString(cert.Signature))
	fmt.Println("sig algo:", cert.SignatureAlgorithm.String())

	fmt.Println("pubk algo:", cert.PublicKeyAlgorithm.String())
	fmt.Println("pubk type:", reflect.TypeOf(cert.PublicKey))

	fmt.Println("version:", cert.Version)
	fmt.Println("serial:", hex.EncodeToString(cert.SerialNumber.Bytes()))
	fmt.Println("issuer:", cert.Issuer.String())
	fmt.Println("subject:", cert.Subject.String())
	fmt.Println("time:", cert.NotBefore, cert.NotAfter)
	fmt.Println("key usage:", int(cert.KeyUsage))
	fmt.Println("extension:", cert.Extensions)
	fmt.Println("extra ext:", cert.ExtraExtensions)
	fmt.Println("un ext:", cert.UnhandledCriticalExtensions)
	fmt.Println("extra key use:", cert.ExtKeyUsage)
	fmt.Println("unknow key use:", cert.UnknownExtKeyUsage)
	fmt.Println("basic valid:", cert.BasicConstraintsValid)
	fmt.Println("isca:", cert.IsCA)
	fmt.Println("max pl:", cert.MaxPathLen)
	fmt.Println("max pl0:", cert.MaxPathLenZero)
	fmt.Println("sub key id:", hex.EncodeToString(cert.SubjectKeyId))
	fmt.Println("author key id:", hex.EncodeToString(cert.AuthorityKeyId))
	fmt.Println("ocsp:", cert.OCSPServer)
	fmt.Println("issuer url:", cert.IssuingCertificateURL)
	fmt.Println("dnss:", cert.DNSNames)
	fmt.Println("emails:", cert.EmailAddresses)
	fmt.Println("ips:", cert.IPAddresses)
	fmt.Println("uris:", cert.URIs)
	fmt.Println("pds:", cert.PermittedDNSDomains)
	fmt.Println("eds:", cert.ExcludedDNSDomains)
	fmt.Println("pr:", cert.PermittedIPRanges)
	fmt.Println("er:", cert.ExcludedIPRanges)
	fmt.Println("pe:", cert.PermittedEmailAddresses)
	fmt.Println("ee:", cert.ExcludedEmailAddresses)
	fmt.Println("pu:", cert.PermittedURIDomains)
	fmt.Println("eu:", cert.ExcludedURIDomains)
	fmt.Println("crldp:", cert.CRLDistributionPoints)
	fmt.Println("pi:", cert.PolicyIdentifiers)
}

func displayCertReq(csr *x509.CertificateRequest) {
	//fmt.Println("raw:", hex.EncodeToString(csr.Raw))
	//fmt.Println("raw tbs:", hex.EncodeToString(csr.RawTBSCertificateRequest))
	//fmt.Println("raw pubk:", hex.EncodeToString(csr.RawSubjectPublicKeyInfo))
	//fmt.Println("raw subject:", hex.EncodeToString(csr.RawSubject))
	fmt.Println("signature:", hex.EncodeToString(csr.Signature))
	fmt.Println("sig algo:", csr.SignatureAlgorithm.String())
	fmt.Println("pubk algo:", csr.PublicKeyAlgorithm.String())
	fmt.Println("pubk type:", reflect.TypeOf(csr.PublicKey))
	fmt.Println("version:", csr.Version)
	fmt.Println("subject:", csr.Subject.String())
	fmt.Println("extension:", csr.Extensions)
	fmt.Println("extra ext:", csr.ExtraExtensions)
	fmt.Println("dnss:", csr.DNSNames)
	fmt.Println("emails:", csr.EmailAddresses)
	fmt.Println("ips:", csr.IPAddresses)
	fmt.Println("uris:", csr.URIs)
}

func displayCertificateList(crl *pkix.CertificateList) {
	fmt.Println("sign algo:", crl.SignatureAlgorithm.Algorithm.String())
	fmt.Println("sign value:", hex.EncodeToString(crl.SignatureValue.Bytes))
	fmt.Println("extensions:", crl.TBSCertList.Extensions)
	fmt.Println("issuer:", crl.TBSCertList.Issuer.String())
	fmt.Println("next update:", crl.TBSCertList.NextUpdate)
	fmt.Println("this update:", crl.TBSCertList.ThisUpdate)
	fmt.Println("version:", crl.TBSCertList.Version)
}

func readFile(filename string) []byte {
	byter, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("read file err: " + err.Error())
	}
	return byter
}
