package ecc

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
)

const (
	SN_secp112r1               = "secp112r1"
	SN_secp112r2               = "secp112r2"
	SN_secp128r1               = "secp128r1"
	SN_secp128r2               = "secp128r2"
	SN_secp160k1               = "secp160k1"
	SN_secp160r1               = "secp160r1"
	SN_secp160r2               = "secp160r2"
	SN_secp192k1               = "secp192k1"
	SN_secp224k1               = "secp224k1"
	SN_secp224r1               = "secp224r1"
	SN_secp256k1               = "secp256k1"
	SN_secp384r1               = "secp384r1"
	SN_secp521r1               = "secp521r1"
	SN_X9_62_prime192v1        = "prime192v1"
	SN_X9_62_prime192v2        = "prime192v2"
	SN_X9_62_prime192v3        = "prime192v3"
	SN_X9_62_prime239v1        = "prime239v1"
	SN_X9_62_prime239v2        = "prime239v2"
	SN_X9_62_prime239v3        = "prime239v3"
	SN_X9_62_prime256v1        = "prime256v1"
	SN_sect113r1               = "sect113r1"
	SN_sect113r2               = "sect113r2"
	SN_sect131r1               = "sect131r1"
	SN_sect131r2               = "sect131r2"
	SN_sect163k1               = "sect163k1"
	SN_sect163r1               = "sect163r1"
	SN_sect163r2               = "sect163r2"
	SN_sect193r1               = "sect193r1"
	SN_sect193r2               = "sect193r2"
	SN_sect233k1               = "sect233k1"
	SN_sect233r1               = "sect233r1"
	SN_sect239k1               = "sect239k1"
	SN_sect283k1               = "sect283k1"
	SN_sect283r1               = "sect283r1"
	SN_sect409k1               = "sect409k1"
	SN_sect409r1               = "sect409r1"
	SN_sect571k1               = "sect571k1"
	SN_sect571r1               = "sect571r1"
	SN_X9_62_c2pnb163v1        = "c2pnb163v1"
	SN_X9_62_c2pnb163v2        = "c2pnb163v2"
	SN_X9_62_c2pnb163v3        = "c2pnb163v3"
	SN_X9_62_c2pnb176v1        = "c2pnb176v1"
	SN_X9_62_c2tnb191v1        = "c2tnb191v1"
	SN_X9_62_c2tnb191v2        = "c2tnb191v2"
	SN_X9_62_c2tnb191v3        = "c2tnb191v3"
	SN_X9_62_c2pnb208w1        = "c2pnb208w1"
	SN_X9_62_c2tnb239v1        = "c2tnb239v1"
	SN_X9_62_c2tnb239v2        = "c2tnb239v2"
	SN_X9_62_c2tnb239v3        = "c2tnb239v3"
	SN_X9_62_c2pnb272w1        = "c2pnb272w1"
	SN_X9_62_c2pnb304w1        = "c2pnb304w1"
	SN_X9_62_c2tnb359v1        = "c2tnb359v1"
	SN_X9_62_c2pnb368w1        = "c2pnb368w1"
	SN_X9_62_c2tnb431r1        = "c2tnb431r1"
	SN_wap_wsg_idm_ecid_wtls1  = "wap-wsg-idm-ecid-wtls1"
	SN_wap_wsg_idm_ecid_wtls3  = "wap-wsg-idm-ecid-wtls3"
	SN_wap_wsg_idm_ecid_wtls4  = "wap-wsg-idm-ecid-wtls4"
	SN_wap_wsg_idm_ecid_wtls5  = "wap-wsg-idm-ecid-wtls5"
	SN_wap_wsg_idm_ecid_wtls6  = "wap-wsg-idm-ecid-wtls6"
	SN_wap_wsg_idm_ecid_wtls7  = "wap-wsg-idm-ecid-wtls7"
	SN_wap_wsg_idm_ecid_wtls8  = "wap-wsg-idm-ecid-wtls8"
	SN_wap_wsg_idm_ecid_wtls9  = "wap-wsg-idm-ecid-wtls9"
	SN_wap_wsg_idm_ecid_wtls10 = "wap-wsg-idm-ecid-wtls10"
	SN_wap_wsg_idm_ecid_wtls11 = "wap-wsg-idm-ecid-wtls11"
	SN_wap_wsg_idm_ecid_wtls12 = "wap-wsg-idm-ecid-wtls12"
	SN_brainpoolP160r1         = "brainpoolP160r1"
	SN_brainpoolP160t1         = "brainpoolP160t1"
	SN_brainpoolP192r1         = "brainpoolP192r1"
	SN_brainpoolP192t1         = "brainpoolP192t1"
	SN_brainpoolP224r1         = "brainpoolP224r1"
	SN_brainpoolP224t1         = "brainpoolP224t1"
	SN_brainpoolP256r1         = "brainpoolP256r1"
	SN_brainpoolP256t1         = "brainpoolP256t1"
	SN_brainpoolP320r1         = "brainpoolP320r1"
	SN_brainpoolP320t1         = "brainpoolP320t1"
	SN_brainpoolP384r1         = "brainpoolP384r1"
	SN_brainpoolP384t1         = "brainpoolP384t1"
	SN_brainpoolP512r1         = "brainpoolP512r1"
	SN_brainpoolP512t1         = "brainpoolP512t1"
)

type EcCurveSpec struct {
	Name  string
	Curve *EcCurve
	Desc  string
}

var initEc sync.Once
var EcCurveSpecs = make(map[string]*EcCurveSpec)
var FpCurveNames = make([]string, 0)
var F2mCurveNames = make([]string, 0)
var Aneg3CurveNames = make([]string, 0)  // Fp curve with a=-3
var KoblitzCurveNames = make([]string, 0) // Fp curve with a=0
var TrivialCurveNames = make([]string, 0) // Fp curve with no special parameters

func initEcCurves() {
	EcCurveSpecs[SN_secp112r1] = &EcCurveSpec{SN_secp112r1, _EC_SECG_PRIME_112R1,
		"SECG/WTLS curve over a 112 bit prime field"}
	EcCurveSpecs[SN_secp112r2] = &EcCurveSpec{SN_secp112r2, _EC_SECG_PRIME_112R2,
		"SECG curve over a 112 bit prime field"}
	EcCurveSpecs[SN_secp128r1] = &EcCurveSpec{SN_secp128r1, _EC_SECG_PRIME_128R1,
		"SECG curve over a 128 bit prime field"}
	EcCurveSpecs[SN_secp128r2] = &EcCurveSpec{SN_secp128r2, _EC_SECG_PRIME_128R2,
		"SECG curve over a 128 bit prime field"}
	EcCurveSpecs[SN_secp160k1] = &EcCurveSpec{SN_secp160k1, _EC_SECG_PRIME_160K1,
		"SECG curve over a 160 bit prime field"}
	EcCurveSpecs[SN_secp160r1] = &EcCurveSpec{SN_secp160r1, _EC_SECG_PRIME_160R1,
		"SECG curve over a 160 bit prime field"}
	EcCurveSpecs[SN_secp160r2] = &EcCurveSpec{SN_secp160r2, _EC_SECG_PRIME_160R2,
		"SECG/WTLS curve over a 160 bit prime field"}
	/* SECG secp192r1 is the same as X9.62 prime192v1 and hence omitted */
	EcCurveSpecs[SN_secp192k1] = &EcCurveSpec{SN_secp192k1, _EC_SECG_PRIME_192K1,
		"SECG curve over a 192 bit prime field"}
	EcCurveSpecs[SN_secp224k1] = &EcCurveSpec{SN_secp224k1, _EC_SECG_PRIME_224K1,
		"SECG curve over a 224 bit prime field"}

	EcCurveSpecs[SN_secp224r1] = &EcCurveSpec{SN_secp224r1, _EC_NIST_PRIME_224,
		"NIST/SECG curve over a 224 bit prime field"}

	EcCurveSpecs[SN_secp224r1] = &EcCurveSpec{SN_secp224r1, _EC_NIST_PRIME_224,
		"NIST/SECG curve over a 224 bit prime field"}

	EcCurveSpecs[SN_secp256k1] = &EcCurveSpec{SN_secp256k1, _EC_SECG_PRIME_256K1,
		"SECG curve over a 256 bit prime field"}
	/* SECG secp256r1 is the same as X9.62 prime256v1 and hence omitted */
	EcCurveSpecs[SN_secp384r1] = &EcCurveSpec{SN_secp384r1, _EC_NIST_PRIME_384,
		"NIST/SECG curve over a 384 bit prime field"}

	EcCurveSpecs[SN_secp521r1] = &EcCurveSpec{SN_secp521r1, _EC_NIST_PRIME_521,
		"NIST/SECG curve over a 521 bit prime field"}

	/* X9.62 curves */
	EcCurveSpecs[SN_X9_62_prime192v1] = &EcCurveSpec{SN_X9_62_prime192v1, _EC_NIST_PRIME_192,
		"NIST/X9.62/SECG curve over a 192 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime192v2] = &EcCurveSpec{SN_X9_62_prime192v2, _EC_X9_62_PRIME_192V2,
		"X9.62 curve over a 192 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime192v3] = &EcCurveSpec{SN_X9_62_prime192v3, _EC_X9_62_PRIME_192V3,
		"X9.62 curve over a 192 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime239v1] = &EcCurveSpec{SN_X9_62_prime239v1, _EC_X9_62_PRIME_239V1,
		"X9.62 curve over a 239 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime239v2] = &EcCurveSpec{SN_X9_62_prime239v2, _EC_X9_62_PRIME_239V2,
		"X9.62 curve over a 239 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime239v3] = &EcCurveSpec{SN_X9_62_prime239v3, _EC_X9_62_PRIME_239V3,
		"X9.62 curve over a 239 bit prime field"}
	EcCurveSpecs[SN_X9_62_prime256v1] = &EcCurveSpec{SN_X9_62_prime256v1, _EC_X9_62_PRIME_256V1,
		"X9.62/SECG curve over a 256 bit prime field"}

	/* characteristic two field curves */
	/* NIST/SECG curves */
	EcCurveSpecs[SN_sect113r1] = &EcCurveSpec{SN_sect113r1, _EC_SECG_CHAR2_113R1,
		"SECG curve over a 113 bit binary field"}
	EcCurveSpecs[SN_sect113r2] = &EcCurveSpec{SN_sect113r2, _EC_SECG_CHAR2_113R2,
		"SECG curve over a 113 bit binary field"}
	EcCurveSpecs[SN_sect131r1] = &EcCurveSpec{SN_sect131r1, _EC_SECG_CHAR2_131R1,
		"SECG/WTLS curve over a 131 bit binary field"}
	EcCurveSpecs[SN_sect131r2] = &EcCurveSpec{SN_sect131r2, _EC_SECG_CHAR2_131R2,
		"SECG curve over a 131 bit binary field"}
	EcCurveSpecs[SN_sect163k1] = &EcCurveSpec{SN_sect163k1, _EC_NIST_CHAR2_163K,
		"NIST/SECG/WTLS curve over a 163 bit binary field"}
	EcCurveSpecs[SN_sect163r1] = &EcCurveSpec{SN_sect163r1, _EC_SECG_CHAR2_163R1,
		"SECG curve over a 163 bit binary field"}
	EcCurveSpecs[SN_sect163r2] = &EcCurveSpec{SN_sect163r2, _EC_NIST_CHAR2_163B,
		"NIST/SECG curve over a 163 bit binary field"}
	EcCurveSpecs[SN_sect193r1] = &EcCurveSpec{SN_sect193r1, _EC_SECG_CHAR2_193R1,
		"SECG curve over a 193 bit binary field"}
	EcCurveSpecs[SN_sect193r2] = &EcCurveSpec{SN_sect193r2, _EC_SECG_CHAR2_193R2,
		"SECG curve over a 193 bit binary field"}
	EcCurveSpecs[SN_sect233k1] = &EcCurveSpec{SN_sect233k1, _EC_NIST_CHAR2_233K,
		"NIST/SECG/WTLS curve over a 233 bit binary field"}
	EcCurveSpecs[SN_sect233r1] = &EcCurveSpec{SN_sect233r1, _EC_NIST_CHAR2_233B,
		"NIST/SECG/WTLS curve over a 233 bit binary field"}
	EcCurveSpecs[SN_sect239k1] = &EcCurveSpec{SN_sect239k1, _EC_SECG_CHAR2_239K1,
		"SECG curve over a 239 bit binary field"}
	EcCurveSpecs[SN_sect283k1] = &EcCurveSpec{SN_sect283k1, _EC_NIST_CHAR2_283K,
		"NIST/SECG curve over a 283 bit binary field"}
	EcCurveSpecs[SN_sect283r1] = &EcCurveSpec{SN_sect283r1, _EC_NIST_CHAR2_283B,
		"NIST/SECG curve over a 283 bit binary field"}
	EcCurveSpecs[SN_sect409k1] = &EcCurveSpec{SN_sect409k1, _EC_NIST_CHAR2_409K,
		"NIST/SECG curve over a 409 bit binary field"}
	EcCurveSpecs[SN_sect409r1] = &EcCurveSpec{SN_sect409r1, _EC_NIST_CHAR2_409B,
		"NIST/SECG curve over a 409 bit binary field"}
	EcCurveSpecs[SN_sect571k1] = &EcCurveSpec{SN_sect571k1, _EC_NIST_CHAR2_571K,
		"NIST/SECG curve over a 571 bit binary field"}
	EcCurveSpecs[SN_sect571r1] = &EcCurveSpec{SN_sect571r1, _EC_NIST_CHAR2_571B,
		"NIST/SECG curve over a 571 bit binary field"}
	/* X9.62 curves */
	EcCurveSpecs[SN_X9_62_c2pnb163v1] = &EcCurveSpec{SN_X9_62_c2pnb163v1, _EC_X9_62_CHAR2_163V1,
		"X9.62 curve over a 163 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb163v2] = &EcCurveSpec{SN_X9_62_c2pnb163v2, _EC_X9_62_CHAR2_163V2,
		"X9.62 curve over a 163 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb163v3] = &EcCurveSpec{SN_X9_62_c2pnb163v3, _EC_X9_62_CHAR2_163V3,
		"X9.62 curve over a 163 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb176v1] = &EcCurveSpec{SN_X9_62_c2pnb176v1, _EC_X9_62_CHAR2_176V1,
		"X9.62 curve over a 176 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb191v1] = &EcCurveSpec{SN_X9_62_c2tnb191v1, _EC_X9_62_CHAR2_191V1,
		"X9.62 curve over a 191 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb191v2] = &EcCurveSpec{SN_X9_62_c2tnb191v2, _EC_X9_62_CHAR2_191V2,
		"X9.62 curve over a 191 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb191v3] = &EcCurveSpec{SN_X9_62_c2tnb191v3, _EC_X9_62_CHAR2_191V3,
		"X9.62 curve over a 191 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb208w1] = &EcCurveSpec{SN_X9_62_c2pnb208w1, _EC_X9_62_CHAR2_208W1,
		"X9.62 curve over a 208 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb239v1] = &EcCurveSpec{SN_X9_62_c2tnb239v1, _EC_X9_62_CHAR2_239V1,
		"X9.62 curve over a 239 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb239v2] = &EcCurveSpec{SN_X9_62_c2tnb239v2, _EC_X9_62_CHAR2_239V2,
		"X9.62 curve over a 239 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb239v3] = &EcCurveSpec{SN_X9_62_c2tnb239v3, _EC_X9_62_CHAR2_239V3,
		"X9.62 curve over a 239 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb272w1] = &EcCurveSpec{SN_X9_62_c2pnb272w1, _EC_X9_62_CHAR2_272W1,
		"X9.62 curve over a 272 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb304w1] = &EcCurveSpec{SN_X9_62_c2pnb304w1, _EC_X9_62_CHAR2_304W1,
		"X9.62 curve over a 304 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb359v1] = &EcCurveSpec{SN_X9_62_c2tnb359v1, _EC_X9_62_CHAR2_359V1,
		"X9.62 curve over a 359 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2pnb368w1] = &EcCurveSpec{SN_X9_62_c2pnb368w1, _EC_X9_62_CHAR2_368W1,
		"X9.62 curve over a 368 bit binary field"}
	EcCurveSpecs[SN_X9_62_c2tnb431r1] = &EcCurveSpec{SN_X9_62_c2tnb431r1, _EC_X9_62_CHAR2_431R1,
		"X9.62 curve over a 431 bit binary field"}
	/*
	 * the WAP/WTLS curves [unlike SECG, spec has its own OIDs for curves
	 * from X9.62]
	 */
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls1] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls1, _EC_WTLS_1,
		"WTLS curve over a 113 bit binary field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls3] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls3, _EC_NIST_CHAR2_163K,
		"NIST/SECG/WTLS curve over a 163 bit binary field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls4] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls4, _EC_SECG_CHAR2_113R1,
		"SECG curve over a 113 bit binary field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls5] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls5, _EC_X9_62_CHAR2_163V1,
		"X9.62 curve over a 163 bit binary field"}

	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls6] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls6, _EC_SECG_PRIME_112R1,
		"SECG/WTLS curve over a 112 bit prime field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls7] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls7, _EC_SECG_PRIME_160R2,
		"SECG/WTLS curve over a 160 bit prime field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls8] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls8, _EC_WTLS_8,
		"WTLS curve over a 112 bit prime field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls9] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls9, _EC_WTLS_9,
		"WTLS curve over a 160 bit prime field"}

	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls10] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls10, _EC_NIST_CHAR2_233K,
		"NIST/SECG/WTLS curve over a 233 bit binary field"}
	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls11] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls11, _EC_NIST_CHAR2_233B,
		"NIST/SECG/WTLS curve over a 233 bit binary field"}

	EcCurveSpecs[SN_wap_wsg_idm_ecid_wtls12] = &EcCurveSpec{SN_wap_wsg_idm_ecid_wtls12, _EC_WTLS_12,
		"WTLS curve over a 224 bit prime field"}

	/* brainpool curves */
	EcCurveSpecs[SN_brainpoolP160r1] = &EcCurveSpec{SN_brainpoolP160r1, _EC_brainpoolP160r1,
		"RFC 5639 curve over a 160 bit prime field"}
	EcCurveSpecs[SN_brainpoolP160t1] = &EcCurveSpec{SN_brainpoolP160t1, _EC_brainpoolP160t1,
		"RFC 5639 curve over a 160 bit prime field"}
	EcCurveSpecs[SN_brainpoolP192r1] = &EcCurveSpec{SN_brainpoolP192r1, _EC_brainpoolP192r1,
		"RFC 5639 curve over a 192 bit prime field"}
	EcCurveSpecs[SN_brainpoolP192t1] = &EcCurveSpec{SN_brainpoolP192t1, _EC_brainpoolP192t1,
		"RFC 5639 curve over a 192 bit prime field"}
	EcCurveSpecs[SN_brainpoolP224r1] = &EcCurveSpec{SN_brainpoolP224r1, _EC_brainpoolP224r1,
		"RFC 5639 curve over a 224 bit prime field"}
	EcCurveSpecs[SN_brainpoolP224t1] = &EcCurveSpec{SN_brainpoolP224t1, _EC_brainpoolP224t1,
		"RFC 5639 curve over a 224 bit prime field"}
	EcCurveSpecs[SN_brainpoolP256r1] = &EcCurveSpec{SN_brainpoolP256r1, _EC_brainpoolP256r1,
		"RFC 5639 curve over a 256 bit prime field"}
	EcCurveSpecs[SN_brainpoolP256t1] = &EcCurveSpec{SN_brainpoolP256t1, _EC_brainpoolP256t1,
		"RFC 5639 curve over a 256 bit prime field"}
	EcCurveSpecs[SN_brainpoolP320r1] = &EcCurveSpec{SN_brainpoolP320r1, _EC_brainpoolP320r1,
		"RFC 5639 curve over a 320 bit prime field"}
	EcCurveSpecs[SN_brainpoolP320t1] = &EcCurveSpec{SN_brainpoolP320t1, _EC_brainpoolP320t1,
		"RFC 5639 curve over a 320 bit prime field"}
	EcCurveSpecs[SN_brainpoolP384r1] = &EcCurveSpec{SN_brainpoolP384r1, _EC_brainpoolP384r1,
		"RFC 5639 curve over a 384 bit prime field"}
	EcCurveSpecs[SN_brainpoolP384t1] = &EcCurveSpec{SN_brainpoolP384t1, _EC_brainpoolP384t1,
		"RFC 5639 curve over a 384 bit prime field"}
	EcCurveSpecs[SN_brainpoolP512r1] = &EcCurveSpec{SN_brainpoolP512r1, _EC_brainpoolP512r1,
		"RFC 5639 curve over a 512 bit prime field"}
	EcCurveSpecs[SN_brainpoolP512t1] = &EcCurveSpec{SN_brainpoolP512t1, _EC_brainpoolP512t1,
		"RFC 5639 curve over a 512 bit prime field"}

	for name, curve := range EcCurveSpecs {
		three := new(big.Int).SetInt64(3)
		if curve.Curve.head.fieldType == NID_X9_62_prime_field {
			FpCurveNames = append(FpCurveNames, name)
			if new(big.Int).Add(curve.Curve.A, three).Cmp(curve.Curve.P) == 0 {
				Aneg3CurveNames = append(Aneg3CurveNames, name)
				continue
			}
			if curve.Curve.A.Cmp(Zero) == 0 {
				KoblitzCurveNames = append(KoblitzCurveNames, name)
				continue
			}
			TrivialCurveNames = append(TrivialCurveNames, name)
		}

		if curve.Curve.head.fieldType == NID_X9_62_characteristic_two_field {
			F2mCurveNames = append(F2mCurveNames, name)
		}
 	}
}

func GetEcCurveSpec(name string) (*EcCurveSpec, error) {
	initEc.Do(initEcCurves)
	spec := EcCurveSpecs[name]
	if spec == nil {
		return nil, errors.New(fmt.Sprintf("curve '%s' not exists", name))
	}
	return spec, nil
}

func GetEcCurve(name string) (*EcCurve, error) {
	initEc.Do(initEcCurves)
	spec := EcCurveSpecs[name]
	if spec == nil {
		return nil, errors.New(fmt.Sprintf("curve '%s' not exists", name))
	}
	return spec.Curve, nil
}

func GetFpCurve(name string) (*FpCurve, error) {
	initEc.Do(initEcCurves)
	spec := EcCurveSpecs[name]
	if spec == nil {
		return nil, errors.New(fmt.Sprintf("curve '%s' not exists", name))
	}
	if spec.Curve.head.fieldType != NID_X9_62_prime_field {
		return nil, errors.New(fmt.Sprintf("curve '%s' is not for prime field", name))
	}
	return (*FpCurve)(spec.Curve), nil
}

func GetF2mCurve(name string) (*F2mCurve, error) {
	initEc.Do(initEcCurves)
	spec := EcCurveSpecs[name]
	if spec == nil {
		return nil, errors.New(fmt.Sprintf("curve '%s' not exists", name))
	}
	if spec.Curve.head.fieldType != NID_X9_62_characteristic_two_field {
		return nil, errors.New(fmt.Sprintf("curve '%s' is not for binary field", name))
	}
	return &F2mCurve{spec.Curve, bn_gf2m_poly2arr(spec.Curve.P.Bits())}, nil
}
