package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"time"
	"unsafe"
)

func main() {
	//tlsDialBaidu()
	tlsClientBaidu()
}

// tls dial baidu
func tlsDialBaidu() {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		fmt.Println("system cert poll err:", err)
		return
	}
	clientConfig := &tls.Config{
		RootCAs: certPool,
	}
	conn, err := tls.Dial("tcp", "www.baidu.com:443", clientConfig)
	defer toClose3(conn)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	displayConn(conn)
}

// tcp dial baidu and create tls client
func tlsClientBaidu() {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		fmt.Println("system cert poll err:", err)
		return
	}
	clientConfig := &tls.Config{
		RootCAs:            certPool,
		ClientSessionCache: tls.NewLRUClientSessionCache(20),
		ServerName:         "www.baidu.com",
	}
	tcpconn, err := net.Dial("tcp", "www.baidu.com:443")
	if err != nil {
		fmt.Println("tcp dial err:", err)
		return
	}
	conn := tls.Client(tcpconn, clientConfig)
	fmt.Println("handshaked?", conn.ConnectionState().HandshakeComplete)
	err = conn.Handshake()
	if err != nil {
		fmt.Println("handshake err:", err)
		return
	}
	fmt.Println("handshaked?", conn.ConnectionState().HandshakeComplete)
	displayConn(conn)
	fmt.Println()

	// display session information
	session, ok := clientConfig.ClientSessionCache.Get(clientConfig.ServerName)
	if ok {
		displayClientSessionState(session)
		fmt.Println()
	}
	toClose3(conn)
}

var suites = map[uint16]string{
	// TLS 1.0 - 1.2 cipher suites.
	0x0005: "TLS_RSA_WITH_RC4_128_SHA",
	0x000a: "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
	0x002f: "TLS_RSA_WITH_AES_128_CBC_SHA",
	0x0035: "TLS_RSA_WITH_AES_256_CBC_SHA",
	0x003c: "TLS_RSA_WITH_AES_128_CBC_SHA256",
	0x009c: "TLS_RSA_WITH_AES_128_GCM_SHA256",
	0x009d: "TLS_RSA_WITH_AES_256_GCM_SHA384",
	0xc007: "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
	0xc009: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
	0xc00a: "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
	0xc011: "TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	0xc012: "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
	0xc013: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
	0xc014: "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
	0xc023: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
	0xc027: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
	0xc02f: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	0xc02b: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
	0xc030: "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	0xc02c: "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
	0xcca8: "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
	0xcca9: "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",

	// TLS 1.3 cipher suites.
	0x1301: "TLS_AES_128_GCM_SHA256",
	0x1302: "TLS_AES_256_GCM_SHA384",
	0x1303: "TLS_CHACHA20_POLY1305_SHA256",

	// TLS_FALLBACK_SCSV isn't a standard cipher suite but an indicator
	// that the client is doing version fallback. See RFC 7507.
	0x5600: "TLS_FALLBACK_SCSV",
}

func toClose3(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}
}

func displayConn(conn *tls.Conn) {
	fmt.Println("conn local:", conn.LocalAddr().String())
	fmt.Println("conn remote:", conn.RemoteAddr().String())
	state := conn.ConnectionState()
	fmt.Println("version:", state.Version)
	fmt.Println("handshake complete:", state.HandshakeComplete)
	fmt.Println("dis resume:", state.DidResume)
	fmt.Println("cipher suite:", suites[state.CipherSuite], state.CipherSuite)
	fmt.Println("negotiated protocol:", state.NegotiatedProtocol)
	fmt.Println("negotiated protocol mutual:", state.NegotiatedProtocolIsMutual)
	fmt.Println("server name:", state.ServerName)
	for i, cert := range state.PeerCertificates {
		fmt.Println(i, "sub:", cert.Subject.CommonName, ", issuer:", cert.Issuer.CommonName)
	}
	for i, chain := range state.VerifiedChains {
		fmt.Println("chain:", i)
		for j, cert := range chain {
			fmt.Println(j, "sub:", cert.Subject.CommonName, ", issuer:", cert.Issuer.CommonName)
		}
	}
	for i, ts := range state.SignedCertificateTimestamps {
		fmt.Println(i, "sct:", hex.EncodeToString(ts))
	}
	fmt.Println("OCSPResponse:", hex.EncodeToString(state.OCSPResponse))
	fmt.Println("TLSUnique:", hex.EncodeToString(state.TLSUnique))
}

type ClientSessionState struct {
	sessionTicket      []uint8               // Encrypted ticket used for session resumption with server
	vers               uint16                // SSL/TLS version negotiated for the session
	cipherSuite        uint16                // Ciphersuite negotiated for the session
	masterSecret       []byte                // Full handshake MasterSecret, or TLS 1.3 resumption_master_secret
	serverCertificates []*x509.Certificate   // Certificate chain presented by the server
	verifiedChains     [][]*x509.Certificate // Certificate chains we built for verification
	receivedAt         time.Time             // When the session ticket was received from the server

	// TLS 1.3 fields.
	nonce  []byte    // Ticket nonce sent by the server, to derive PSK
	useBy  time.Time // Expiration of the ticket lifetime as set by the server
	ageAdd uint32    // Random obfuscation factor for sending the ticket age
}

func displayClientSessionState(state *tls.ClientSessionState) {
	mysession := (*ClientSessionState)(unsafe.Pointer(state))
	fmt.Println("sessionTicket     :", mysession.sessionTicket)
	fmt.Println("version           :", mysession.vers)
	fmt.Println("cipherSuite       :", suites[mysession.cipherSuite])
	fmt.Println("masterSecret      :", hex.EncodeToString(mysession.masterSecret))
	for i, cert := range mysession.serverCertificates {
		fmt.Println(i, "sub:", cert.Subject.CommonName, ", issuer:", cert.Issuer.CommonName)
	}
	for i, chain := range mysession.verifiedChains {
		fmt.Println("chain:", i)
		for j, cert := range chain {
			fmt.Println(j, "sub:", cert.Subject.CommonName, ", issuer:", cert.Issuer.CommonName)
		}
	}
	fmt.Println("receivedAt        :", mysession.receivedAt)
	fmt.Println("nonce             :", hex.EncodeToString(mysession.nonce))
	fmt.Println("useBy             :", mysession.useBy)
	fmt.Println("ageAdd            :", mysession.ageAdd)
}
