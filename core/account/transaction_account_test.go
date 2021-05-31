package account

import (
	"testing"

	accountpb "github.com/dappley/go-dappley/core/account/pb"
	"github.com/stretchr/testify/assert"
)

func TestTransactionAccount_ToProto(t *testing.T) {
	pubKeyHash := newUserPubKeyHash([]byte("hash"))
	address := pubKeyHash.GenerateAddress()
	transactionAccount := &TransactionAccount{pubKeyHash: pubKeyHash, address: address}

	expected := &accountpb.TransactionAccount{
		Address: &accountpb.Address{
			Address: address.address,
		},
		PubKeyHash: pubKeyHash,
	}
	assert.Equal(t, expected, transactionAccount.ToProto())
}

func TestTransactionAccount_FromProto(t *testing.T) {
	pubKeyHash := newUserPubKeyHash([]byte("hash"))
	address := pubKeyHash.GenerateAddress()

	transactionAccount := &TransactionAccount{}
	transactionAccountProto := &accountpb.TransactionAccount{
		Address: &accountpb.Address{
			Address: address.address,
		},
		PubKeyHash: pubKeyHash,
	}
	transactionAccount.FromProto(transactionAccountProto)

	expected := &TransactionAccount{address: address, pubKeyHash: pubKeyHash}
	assert.Equal(t, expected, transactionAccount)
}

func TestTransactionAccount_IsValid(t *testing.T) {
	transactionAccount := NewContractTransactionAccount()
	assert.True(t, transactionAccount.IsValid())
	transactionAccount.address.address = "address000000000000000000000000011"
	assert.False(t, transactionAccount.IsValid())
}

func TestNewTransactionAccountByPubKey(t *testing.T) {
	pubKeyBytes := []byte("address1000000000000000000000000")
	transactionAccount := NewTransactionAccountByPubKey(pubKeyBytes)

	assert.NotNil(t, transactionAccount)
	assert.NotNil(t, transactionAccount.pubKeyHash)
	assert.NotNil(t, transactionAccount.address)
	assert.Equal(t, PubKeyHash([]byte{0x5a, 0xad, 0xec, 0x2c, 0x21, 0x3b, 0x67, 0xfa, 0x96, 0xe5, 0xa8, 0xb9, 0xb4, 0x99, 0xf3, 0x26, 0x41, 0xf7, 0xff, 0x36, 0x8a}), transactionAccount.pubKeyHash)
	assert.Equal(t, NewAddress("dVGuSWFXE91Ay36n9HnCzpu8AfckEgvnnR"), transactionAccount.address)
}

func TestNewContractAccountByPubKeyHash(t *testing.T) {
	pubKeyBytes := PubKeyHash([]byte{versionUser, 0xb1, 0x34, 0x4c, 0x17, 0x67, 0x4c, 0x18, 0xd1, 0xa2, 0xdc, 0xea, 0x9f, 0x17, 0x16, 0xe0, 0x49, 0xf4, 0xa0, 0x5e, 0x6c})
	transactionAccount := NewContractAccountByPubKeyHash(pubKeyBytes)

	assert.NotNil(t, transactionAccount)
	assert.NotNil(t, transactionAccount.pubKeyHash)
	assert.NotNil(t, transactionAccount.address)
	assert.Equal(t, pubKeyBytes, transactionAccount.pubKeyHash)
	assert.Equal(t, NewAddress("dVaFsQL9He4Xn4CEUh1TCNtfEhHNHKX3hs"), transactionAccount.address)
}

func TestGeneratePubKeyHashByAddress(t *testing.T) {
	address1 := NewAddress("dZSj3ehsCXKzbTAxfgZU6hokbNFe7Unsuy")
	address2 := NewAddress("invalid000000000000000000000000000")
	address3 := NewAddress("tooshort")
	hash1, success1 := generatePubKeyHashByAddress(address1)
	hash2, success2 := generatePubKeyHashByAddress(address2)
	hash3, success3 := generatePubKeyHashByAddress(address3)

	expectedHash1 := PubKeyHash([]byte{0x5a, 0xdb, 0xa8, 0x28, 0x9b, 0xe2, 0xa9, 0xf, 0x21, 0x1f, 0xf5, 0x0, 0x5f, 0x2a, 0x8e, 0x1e, 0xe8, 0x90, 0x62, 0x5c, 0x2})

	assert.True(t, success1)
	assert.Equal(t, expectedHash1, hash1)

	assert.False(t, success2)
	assert.Nil(t, hash2)

	assert.False(t, success3)
	assert.Nil(t, hash3)
}

func TestNewTransactionAccountByAddress(t *testing.T) {
	address := NewAddress("dZSj3ehsCXKzbTAxfgZU6hokbNFe7Unsuy")
	transactionAccount := NewTransactionAccountByAddress(address)
	expectedHash := PubKeyHash([]byte{0x5a, 0xdb, 0xa8, 0x28, 0x9b, 0xe2, 0xa9, 0xf, 0x21, 0x1f, 0xf5, 0x0, 0x5f, 0x2a, 0x8e, 0x1e, 0xe8, 0x90, 0x62, 0x5c, 0x2})

	assert.NotNil(t, transactionAccount)
	assert.NotNil(t, transactionAccount.pubKeyHash)
	assert.NotNil(t, transactionAccount.address)
	assert.Equal(t, expectedHash, transactionAccount.pubKeyHash)
	assert.Equal(t, address, transactionAccount.address)
}

func TestChecksum(t *testing.T) {
	pubKeyBytes1 := []byte{versionUser, 0xb1, 0x34, 0x4c, 0x17, 0x67, 0x4c, 0x18, 0xd1, 0xa2, 0xdc, 0xea, 0x9f, 0x17, 0x16, 0xe0, 0x49, 0xf4, 0xa0, 0x5e, 0x6c}
	pubKeyBytes2 := []byte{versionUser, 0xb0, 0x34, 0x4c, 0x17, 0x67, 0x4c, 0x18, 0xd1, 0xa2, 0xdc, 0xea, 0x9f, 0x17, 0x16, 0xe0, 0x49, 0xf4, 0xa0, 0x5e, 0x6c}

	checksum1 := Checksum(pubKeyBytes1)
	checksum2 := Checksum(pubKeyBytes2)

	assert.Equal(t, []byte{0x8d, 0xc6, 0x1e, 0x9a}, checksum1)
	assert.Equal(t, []byte{0x84, 0xbd, 0xf4, 0xa2}, checksum2)
}

func TestGetAddressPayloadLength(t *testing.T) {
	assert.Equal(t, 25, GetAddressPayloadLength())
}