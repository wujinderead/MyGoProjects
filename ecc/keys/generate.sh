#!/bin/bash

# generate rsa key
openssl genrsa -out rsa.ca.key 2048
openssl req -x509 -new -nodes -key rsa.ca.key -days 50000 -out rsa.ca.crt -subj "/CN=CaRsa"

openssl genrsa -out rsa.client.key 2048
openssl req -new -key rsa.client.key -subj "/CN=ClientRsa" -out rsa.client.csr
openssl x509 -req -in rsa.client.csr -CA rsa.ca.crt -CAkey rsa.ca.key -CAcreateserial -out rsa.client.crt -days 50000

openssl genrsa -out rsa.server.key 2048
openssl req -new -key rsa.server.key -subj "/CN=ServerRsa" -out rsa.server.csr -config openssl.conf
openssl x509 -req -in rsa.server.csr -CA rsa.ca.crt -CAkey rsa.ca.key -CAcreateserial -out rsa.server.crt -days 50000 -extensions v3_req -extfile openssl.conf


# ed25519 generate key and cert
# openssl version >= 1.1.1
openssl genpkey -algorithm ed25519 -out ed25519.ca.key
openssl req -x509 -new -nodes -key ed25519.ca.key -days 50000 -out ed25519.ca.crt -subj "/CN=Ca25519"

openssl genpkey -algorithm ed25519 -out ed25519.server.key
openssl req -new -key ed25519.server.key -out ed25519.server.csr -subj "/CN=Server25519" -config openssl.conf
openssl x509 -req -in ed25519.server.csr -CA ed25519.ca.crt -CAkey ed25519.ca.key -CAcreateserial -out ed25519.server.crt -days 730 -extensions v3_req -extfile openssl.conf


# generate ec key
openssl ecparam -genkey -name secp384r1 -out ec.ca.key
openssl req -x509 -new -nodes -key ec.ca.key -days 50000 -out ec.ca.crt -subj "/CN=CaEc"

openssl ecparam -genkey -name secp384r1 -out ec.client.key
openssl req -new -key ec.client.key -subj "/CN=ClientEc" -out ec.client.csr
openssl x509 -req -in ec.client.csr -CA ec.ca.crt -CAkey ec.ca.key -CAcreateserial -out ec.client.crt -days 50000

openssl ecparam -genkey -name secp384r1 -out ec.server.key
openssl req -new -key ec.server.key -subj "/CN=ServerEc" -out ec.server.csr -config openssl.conf
openssl x509 -req -in ec.server.csr -CA ec.ca.crt -CAkey ec.ca.key -CAcreateserial -out ec.server.crt -days 50000 -extensions v3_req -extfile openssl.conf
