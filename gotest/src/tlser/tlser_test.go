package tlser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"testing"
	"time"
	"reflect"
	"net"
)

/*
openssl genrsa -out ca.key 2048

#CA self certificate
openssl req -x509 -new -nodes -key ca.key -days 50000 -out ca.crt -subj "/CN=SelfCA"

openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=clienter" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 50000
 */


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
		NotBefore: time.Now().Add(-time.Hour*24*365),
		NotAfter:  time.Now().Add(time.Hour*24*365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
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
		NotBefore: time.Now().Add(-time.Hour*24*365),
		NotAfter:  time.Now().Add(time.Hour*24*365),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: false,
		IPAddresses: []net.IP{net.ParseIP("10.19.138.22"), net.ParseIP("10.19.138.135")},
		DNSNames: []string{"dog.com", "leopard.com"},
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

func TestParse(t *testing.T) {
	// parse rsa key
	rsaKeyPem := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAsKS8zjC7W/nL4jzhmYsD31/9zGt2ykU2fvHVkhL/f3h1+1bI
KotR49WldczGRtBrtsHQFP8hfGS2fn8BCBxZqjyb+UgPPcltFqc7Et1r52y1vTbU
VAf5G64ML0VGWbZjfD4pPmB1yG1915BWGDPwZn9JsidKMqH+A8ynFQOYTgJHrFzE
dMBM9TWaDq0fXgOmOfG9LygrxKJAzdpqARu6sQp9060QnH/OVIM+uIqAJ22mCiQ/
kdpI7kK4rilF56IJBLJ8i8WY+6YZzaMPpNXzrcNnFgwX4uoyLPCtydb9nvK44ZOz
ISiQ5DEu1IKHPQmMBqGZA8AXRFrnHEjL0Re1uQIDAQABAoIBADxrzQ02TBAU7LFx
I7XSgDua6QRQSey8KfzYGbaCexSODsUvFP7Acv1cqeEWb0fvqLh1qQhVkI2tIWM5
bA/rKpx5aNym0lfPG1phT2qPhIY/gBa9t3ka1RGrwg01Q/AR3Au2c4MbmY46LY5b
l6dltLhKl9mxaMbS9EE2cnxUo3ci7NHVzgXYLlG6HDsnTFdFFdKtgml7o+/2q5K5
eJOU32EaGEvl+bNVM8HuBBNdpfuppfKB5NOgwWXktvyWbvkqSa5tFhVwEnAuE3tX
5OgYKZvUGqZWIUJhEq2G8SB8zU8qM26M+P198ph6Ft6B4F3HZSVA2a89hmZD4w+F
aiiJcyUCgYEA4Ip52f/RQUmLt4eK7XJ9NvWRzj9iFiaXITx2MT02KBjDciBuAS6O
e/CiY/qUAcsR5eQIynWObo+GPiivjj+DR0pLKYxXJblBCuRHO7niPwhPkRB+VM5+
GvZp7iTxNN8BNWOHgehXOw1jG0PaN+UeRExonwpHnZhdmdJf/a2xo3sCgYEAyWRX
1Z6UcC2eAjcQBTF0rPXKY3/qSsr/JizxBpXnV9nkEJlyFsfF0ogJTb9e4+/xKWGp
84Lxc+c8sPTlCQczBHU5d6wqyekkb9bDRyIh1YNMgBrPCAnsbUoU4ns6KCETRoY7
J/pZ/w3ryX0e8+h/bnTS5/Pt96WuseqPJrYr+1sCgYEA3GfRGCW998IDfVF1E4LQ
WkROrQ8WbvvOWXeJ5Oda45z7LGmc0Vgr0IjyPgVXhzMYDHr8Dg+6kdgcQ6OYP58k
c7P/d3ckjAj+SXyuV3gtFwZHY/O2rfRLYJgEfxiQE//apddeyiuQhIytfbPq3fbu
8Me34nUquw02w6j3RIFc30UCgYEAgDNPMp5DUfHIDxLsMNIduuiwUIyiIcB9kdi0
CgQtA6Ch7OsxVE0RogaVHZgAGMuqUjRokqo9eBGwcdlDX27kzCavUX4YsvWmC0fE
gai5rwhpD3eBaVWf2qZ5Cv90swzzD0btq3JUDefXvCjZJl1PmYnmpF+Ekcw3m+x7
+iGnd9cCgYAEH2aRvzON+je2KxosPe7DpAAY4b6hRjM6U6NUePKE6mgC3O8IS2n3
PHHftaDoCr8AFbxYnEZFsbo/wqWA7lpLcOWZFdyCKL2J3XKjusMtbxtJRicVT4E1
JyMOp/DkPVOZf9c4jgB+31YlLiRXthXt+Yfl7So/mglVI5cc1Xz3Kw==
-----END RSA PRIVATE KEY-----`
	block, _ := pem.Decode([]byte(rsaKeyPem))
	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("parse key err:", err)
		return
	}
	fmt.Println("rsa.D:", hex.EncodeToString(rsaKey.D.Bytes()))
	fmt.Println()

	// parse ca certificate
	caCertPem := `-----BEGIN CERTIFICATE-----
MIIC9TCCAd2gAwIBAgIJANJjZ2NFbNHxMA0GCSqGSIb3DQEBCwUAMBAxDjAMBgNV
BAMMBW93bmNhMCAXDTE5MDUwNzA5MTkyN1oYDzIxNTYwMzI5MDkxOTI3WjAQMQ4w
DAYDVQQDDAVvd25jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALCk
vM4wu1v5y+I84ZmLA99f/cxrdspFNn7x1ZIS/394dftWyCqLUePVpXXMxkbQa7bB
0BT/IXxktn5/AQgcWao8m/lIDz3JbRanOxLda+dstb021FQH+RuuDC9FRlm2Y3w+
KT5gdchtfdeQVhgz8GZ/SbInSjKh/gPMpxUDmE4CR6xcxHTATPU1mg6tH14Dpjnx
vS8oK8SiQM3aagEburEKfdOtEJx/zlSDPriKgCdtpgokP5HaSO5CuK4pReeiCQSy
fIvFmPumGc2jD6TV863DZxYMF+LqMizwrcnW/Z7yuOGTsyEokOQxLtSChz0JjAah
mQPAF0Ra5xxIy9EXtbkCAwEAAaNQME4wHQYDVR0OBBYEFJKDCrDWrKScs74zKYMq
pandiZXyMB8GA1UdIwQYMBaAFJKDCrDWrKScs74zKYMqpandiZXyMAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQELBQADggEBADzbQ8qDjWKmvcGuRQ0UgxwonOcLWiim
USu7yBTapfmMA2xza0JMxiHIQaQ/dMyq/iUM4IdBIXrMu0HjLAMcvvfFoSr4Hwq6
FY8HzjYiHNmNV8Hg6xcWSBIrgEqHcdJXOGX9eWy3qMdiA9LDHFy49fUH6dB5ELKu
qVYmAiFUWuxvmS2cZCEd5Mt/N5OnPKIBdvEkoIQ0WYY0rusguk3PIPPIJqtWPUS8
ceSKY3s8Q8BIBXIh3GYHUXNlUrcv6EnWTnEnvsgt7SzfEYSUChIFY4ytGhxzJS+c
QymlEEelUyWPTXR7FcWOHAyw7PhFUUz1z22F5FGIYqOBwiE6iDc8b/g=
-----END CERTIFICATE-----`
	block, _ = pem.Decode([]byte(caCertPem))
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(caCert)
	fmt.Println()

	// parse client certificate request
	serverCsrPem := `-----BEGIN CERTIFICATE REQUEST-----
MIIDOzCCAiMCAQAwFjEUMBIGA1UEAwwLa3ViZS1tYXN0ZXIwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDI9LKEMo9BrFtHUgpxWMeg3c7N4cIewd2xAgYx
9qAc+DOPsyJ7a2R3VHp+JBiMIj+qVSWFqZ4jNFWuSbqllGiSXHH5TwTQgXv+heGS
UdOxTB0vj7yg3epBMYhMpEjumj3Ii0Kr7LySexgUhPXxcfzmcCEwrzO0eDORJeGh
D1Jt2fhsP50UbujPy1h92Y1Da00az+k7kDxLnxRRv0EQG21WILKa0sprv3iHEOzZ
LiZJy7GAXUvTBV6CDtwZVW0uZqLMp8kIpQa8PZOpC/y52DuvSV2FbUYJ/bRjPf1H
Gn189SjR+DNBFAuVXV7l+3XV9fXkB517ukHodEHBWsKo0KUpAgMBAAGggd8wgdwG
CSqGSIb3DQEJDjGBzjCByzAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DCBsAYDVR0R
BIGoMIGlggtrdWJlLW1hc3RlcoIKa3ViZXJuZXRlc4ISa3ViZXJuZXRlcy5kZWZh
dWx0ghZrdWJlcm5ldGVzLmRlZmF1bHQuc3ZjgiRrdWJlcm5ldGVzLmRlZmF1bHQu
c3ZjLmNsdXN0ZXIubG9jYWyCLGt1YmVybmV0ZXMuZGVmYXVsdC5zdmMuY2x1c3Rl
ci5sb2NhbC1jbHVzdGVyhwQKE4qHhwQK/gABMA0GCSqGSIb3DQEBCwUAA4IBAQAn
sKDhYgStGBglJYqO7JC5M/i5GRElrNXi2L2zcOBkM47drmG01kWxHEoh+jEbGy5x
hBbx7GtCGFNJm6iLQlwvKsSUxlNM6rg0kOeZI1blOPRZflZGHPyim1jBvahVTXZx
Fe/lUF9IQ9Pw7wWwd557arf7HTC5i5Yywsz7+SZSh1zDuijIcC0QiC4SdQy76brU
zF8iwRgzqET6ffGp3C1UWMCFSkW9i0cvJy+mNhp4qo/Z9AxUCkBEJuKM8QCznfVT
aesJEsoVmNDoK9wFrVQQVa3g+HbsALfQwPOt2aqJDz9ai3IDdwxVZi6J5xLHdiCW
CIp/IKrSdK2vqEvT3Wqr
-----END CERTIFICATE REQUEST-----`
	block, _ = pem.Decode([]byte(serverCsrPem))
	serverCsr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		fmt.Println("parse cert requset err:", err)
		return
	}
	displayCertReq(serverCsr)
	fmt.Println()

	// parse client certificate
	serverCertPem := `-----BEGIN CERTIFICATE-----
MIIDfDCCAmSgAwIBAgIJAOEFvANJ57uvMA0GCSqGSIb3DQEBCwUAMBIxEDAOBgNV
BAMMB2t1YmUtY2EwIBcNMTcwNTIxMDMyOTA4WhgPMjE1NDA0MTMwMzI5MDhaMBYx
FDASBgNVBAMMC2t1YmUtbWFzdGVyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAyPSyhDKPQaxbR1IKcVjHoN3OzeHCHsHdsQIGMfagHPgzj7Mie2tkd1R6
fiQYjCI/qlUlhameIzRVrkm6pZRoklxx+U8E0IF7/oXhklHTsUwdL4+8oN3qQTGI
TKRI7po9yItCq+y8knsYFIT18XH85nAhMK8ztHgzkSXhoQ9Sbdn4bD+dFG7oz8tY
fdmNQ2tNGs/pO5A8S58UUb9BEBttViCymtLKa794hxDs2S4mScuxgF1L0wVegg7c
GVVtLmaizKfJCKUGvD2TqQv8udg7r0ldhW1GCf20Yz39Rxp9fPUo0fgzQRQLlV1e
5ft11fX15Aede7pB6HRBwVrCqNClKQIDAQABo4HOMIHLMAkGA1UdEwQCMAAwCwYD
VR0PBAQDAgXgMIGwBgNVHREEgagwgaWCC2t1YmUtbWFzdGVyggprdWJlcm5ldGVz
ghJrdWJlcm5ldGVzLmRlZmF1bHSCFmt1YmVybmV0ZXMuZGVmYXVsdC5zdmOCJGt1
YmVybmV0ZXMuZGVmYXVsdC5zdmMuY2x1c3Rlci5sb2NhbIIsa3ViZXJuZXRlcy5k
ZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsLWNsdXN0ZXKHBAoTioeHBAr+AAEwDQYJ
KoZIhvcNAQELBQADggEBADS3rboBP9DuVhFZde80gd/099/85Kjsa10TxKa0Khpu
PYz8ULLOQzjr+uvl7E87Fuur7No/Yi7gYT1MfSAttsPWsU4P5A2ZQQorCgC8x+Eh
9WvuOQb5vt3/8dgaeFbWwdsRQCNTsx2bKg95M4CVzOstMOoCno3eCM6xGD8yiq7b
pA6f5APzSQG5G9scMv0r0WbKxM+mc4QPSJe7/Hl8yndkfu4xTwDZvRvkuaqbLV5f
E/hbDkK1mep7quUe2ncfiJbu2TxwKyVGBbWS6SuPCiPFkWxr5VcrA0Sm85QGTVaS
HzyzS8cgWfrEfO+HNGA92hS0gdJVRJX9uaSwGmDpLuI=
-----END CERTIFICATE-----`
	block, _ = pem.Decode([]byte(serverCertPem))
	serverCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("parse cert err:", err)
		return
	}
	displayCert(serverCert)
	fmt.Println()
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
	fmt.Println("url:", cert.IssuingCertificateURL)
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
	fmt.Println("poi:", cert.PolicyIdentifiers)
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
