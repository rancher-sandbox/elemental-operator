/*
 ** Copyright 2019 Bloomberg Finance L.P.
 **
 ** Licensed under the Apache License, Version 2.0 (the "License");
 ** you may not use this file except in compliance with the License.
 ** You may obtain a copy of the License at
 **
 **     http://www.apache.org/licenses/LICENSE-2.0
 **
 ** Unless required by applicable law or agreed to in writing, software
 ** distributed under the License is distributed on an "AS IS" BASIS,
 ** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 ** See the License for the specific language governing permissions and
 ** limitations under the License.
 */

package tpm

import (
	"crypto/sha256"
	"encoding/pem"
	"fmt"

	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/go-attestation/attest"
)

// AttestationData is used to generate challanges from EKs
type AttestationData struct {
	EK []byte
	AK *attest.AttestationParameters
}

// Challenge represent the struct returned from the ws server,
// used to resolve the TPM challenge.
type Challenge struct {
	EC *attest.EncryptedCredential
}

// ChallengeResponse represent the struct returned to the ws server
// as a challenge response.
type ChallengeResponse struct {
	Secret []byte
}

// DecodePubHash returns the public key from an attestation EK
func DecodePubHash(ek *attest.EK) (string, error) {
	data, err := pubBytes(ek)
	if err != nil {
		return "", err
	}
	pubHash := sha256.Sum256(data)
	hashEncoded := fmt.Sprintf("%x", pubHash)
	return hashEncoded, nil
}

func encodeEK(ek *attest.EK) ([]byte, error) {
	if ek.Certificate != nil {
		return pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: ek.Certificate.Raw,
		}), nil
	}

	data, err := pubBytes(ek)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: data,
	}), nil
}

func pubBytes(ek *attest.EK) ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(ek.Public)
	if err != nil {
		return nil, fmt.Errorf("error marshaling ec public key: %v", err)
	}
	return data, nil
}
