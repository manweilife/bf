package bfc_types

import "github.com/manweilife/bf/lib"

type Digest = lib.Base58

type ObjectDigest = Digest

type TransactionDigest = Digest

type TransactionEffectsDigest = Digest

type TransactionEventsDigest = Digest

type CheckpointDigest = Digest

type CertificateDigest = Digest

type CheckpointContentsDigest = Digest

func NewDigest(str string) (*Digest, error) {
	return lib.NewBase58(str)
}
