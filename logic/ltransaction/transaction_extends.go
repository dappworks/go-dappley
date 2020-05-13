// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either pubKeyHash 3 of the License, or
// (at your option) any later pubKeyHash.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//

package ltransaction

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/core/account"
	"github.com/dappley/go-dappley/core/block"
	"github.com/dappley/go-dappley/core/scState"
	"github.com/dappley/go-dappley/core/transaction"
	"github.com/dappley/go-dappley/core/transactionbase"
	"github.com/dappley/go-dappley/core/utxo"
	"github.com/dappley/go-dappley/logic/lutxo"
	"github.com/dappley/go-dappley/util"
	logger "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	scheduleFuncName = "dapp_schedule"
)

// Normal transaction
type TxNormal struct {
	*transaction.Transaction
}

// TxContract contains contract value
type TxContract struct {
	*transaction.Transaction
	Address account.Address
}

// Coinbase transaction, rewards to miner
type TxCoinbase struct {
	*transaction.Transaction
}

// GasReward transaction, rewards to miner during vm execution
type TxGasReward struct {
	*transaction.Transaction
}

// GasChange transaction, change value to from user
type TxGasChange struct {
	*transaction.Transaction
}

// Reward transaction, step reward
type TxReward struct {
	*transaction.Transaction
}

// Returns decorator of transaction
func NewTxDecorator(tx *transaction.Transaction) TxDecorator {
	// old data adapter
	adaptedTx := transaction.NewTxAdapter(tx)
	tx = adaptedTx.Transaction
	switch tx.Type {
	case transaction.TxTypeNormal:
		return &TxNormal{tx}
	case transaction.TxTypeContract:
		return NewTxContract(tx)
	case transaction.TxTypeCoinbase:
		return &TxCoinbase{tx}
	case transaction.TxTypeGasReward:
		return &TxGasReward{tx}
	case transaction.TxTypeGasChange:
		return &TxGasChange{tx}
	case transaction.TxTypeReward:
		return &TxReward{tx}
	}
	return nil
}

func (tx *TxNormal) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return tx.Transaction.Sign(privKey, prevUtxos)
}

func (tx *TxNormal) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	prevUtxos, err := lutxo.FindVinUtxosInUtxoPool(utxoIndex, tx.Transaction)
	if err != nil {
		logger.WithError(err).WithFields(logger.Fields{
			"txid": hex.EncodeToString(tx.ID),
		}).Warn("Verify: cannot find vin while verifying normal tx")
		return err
	}
	return tx.Transaction.Verify(prevUtxos)
}

func (tx *TxContract) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return tx.Transaction.Sign(privKey, prevUtxos)
}

func (tx *TxContract) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	prevUtxos, err := lutxo.FindVinUtxosInUtxoPool(utxoIndex, tx.Transaction)
	if err != nil {
		return err
	}
	err = tx.verifyInEstimate(utxoIndex, prevUtxos)
	if err != nil {
		return err
	}
	totalBalance, err := tx.GetTotalBalance(prevUtxos)
	if err != nil {
		return err
	}
	return tx.VerifyGas(totalBalance)
}

func (tx *TxCoinbase) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return nil
}

func (tx *TxCoinbase) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	//TODO coinbase vout check need add tip
	if tx.Vout[0].Value.Cmp(transaction.Subsidy) < 0 {
		return errors.New("Transaction: subsidy check failed")
	}
	bh := binary.BigEndian.Uint64(tx.Vin[0].Signature)
	if blockHeight != bh {
		return fmt.Errorf("Transaction: block height check failed expected=%v actual=%v", blockHeight, bh)
	}
	return nil
}

func (tx *TxGasReward) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return nil
}

func (tx *TxGasReward) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	return nil
}

func (tx *TxGasChange) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return nil
}

func (tx *TxGasChange) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	return nil
}

func (tx *TxReward) Sign(privKey ecdsa.PrivateKey, prevUtxos []*utxo.UTXO) error {
	return nil
}

func (tx *TxReward) Verify(utxoIndex *lutxo.UTXOIndex, blockHeight uint64) error {
	return nil
}

func NewTxContract(tx *transaction.Transaction) *TxContract {
	adaptedTx := transaction.NewTxAdapter(tx)
	if adaptedTx.IsContract() {
		address := tx.Vout[transaction.ContractTxouputIndex].GetAddress()
		return &TxContract{tx, address}
	}
	return nil
}

// IsScheduleContract returns if the contract contains 'dapp_schedule'
func (ctx *TxContract) IsScheduleContract() bool {
	if !strings.Contains(ctx.GetContract(), scheduleFuncName) {
		return true
	}
	return false
}

//GetContract returns the smart contract code in a transaction
func (ctx *TxContract) GetContract() string {
	return ctx.Vout[transaction.ContractTxouputIndex].Contract
}

//GetContractPubKeyHash returns the smart contract pubkeyhash in a transaction
func (ctx *TxContract) GetContractPubKeyHash() account.PubKeyHash {
	return ctx.Vout[transaction.ContractTxouputIndex].PubKeyHash
}

// GasCountOfTxBase calculate the actual amount for a tx with data
func (ctx *TxContract) GasCountOfTxBase() (*common.Amount, error) {
	txGas := transaction.MinGasCountPerTransaction
	if dataLen := ctx.DataLen(); dataLen > 0 {
		dataGas := common.NewAmount(uint64(dataLen)).Mul(transaction.GasCountPerByte)
		baseGas := txGas.Add(dataGas)
		txGas = baseGas
	}
	return txGas, nil
}

// DataLen return the length of payload
func (ctx *TxContract) DataLen() int {
	return len([]byte(ctx.GetContract()))
}

// VerifyGas verifies if the transaction has the correct GasLimit and GasPrice
func (ctx *TxContract) VerifyGas(totalBalance *common.Amount) error {
	baseGas, err := ctx.GasCountOfTxBase()
	if err == nil {
		if ctx.GasLimit.Cmp(baseGas) < 0 {
			logger.WithFields(logger.Fields{
				"limit":       ctx.GasLimit,
				"acceptedGas": baseGas,
			}).Warn("Failed to check GasLimit >= txBaseGas.")
			// GasLimit is smaller than based tx gas, won't giveback the tx
			return transaction.ErrOutOfGasLimit
		}
	}

	limitedFee := ctx.GasLimit.Mul(ctx.GasPrice)
	if totalBalance.Cmp(limitedFee) < 0 {
		return transaction.ErrInsufficientBalance
	}
	return nil
}

//GetContractAddress gets the smart contract's address if a transaction deploys a smart contract
func (tx *TxContract) GetContractAddress() account.Address {
	return tx.Address
}

// VerifyInEstimate returns whether the current tx in estimate mode is valid.
func (tx *TxContract) VerifyInEstimate(utxoIndex *lutxo.UTXOIndex) error {
	prevUtxos, err := lutxo.FindVinUtxosInUtxoPool(utxoIndex, tx.Transaction)
	if err != nil {
		return err
	}
	return tx.verifyInEstimate(utxoIndex, prevUtxos)
}

func (tx *TxContract) verifyInEstimate(utxoIndex *lutxo.UTXOIndex, prevUtxos []*utxo.UTXO) error {
	if tx.IsScheduleContract() && !tx.IsContractDeployed(utxoIndex) {
		return errors.New("Transaction: contract state check failed")
	}
	err := tx.Transaction.Verify(prevUtxos)
	return err
}

// IsContractDeployed returns if the current contract is deployed
func (tx *TxContract) IsContractDeployed(utxoIndex *lutxo.UTXOIndex) bool {
	pubkeyhash := tx.GetContractPubKeyHash()
	if pubkeyhash == nil {
		return false
	}

	contractUtxoTx := utxoIndex.GetAllUTXOsByPubKeyHash(pubkeyhash)
	return contractUtxoTx.Size() > 0
}

//Execute executes the smart contract the transaction points to. it doesnt do anything if is a contract deploy transaction
func (tx *TxContract) Execute(prevUtxos []*utxo.UTXO,
	isContractDeployed bool,
	utxoIndex *lutxo.UTXOIndex,
	scStorage *scState.ScState,
	rewards map[string]string,
	engine ScEngine,
	currblkHeight uint64,
	parentBlk *block.Block) (uint64, []*transaction.Transaction, error) {

	if engine == nil {
		return 0, nil, nil
	}
	if !isContractDeployed {
		return 0, nil, nil
	}

	vout := tx.Vout[transaction.ContractTxouputIndex]

	function, args := util.DecodeScInput(vout.Contract)
	if function == "" {
		return 0, nil, ErrUnsupportedSourceType
	}
	if err := engine.SetExecutionLimits(tx.GasLimit.Uint64(), 0); err != nil {
		return 0, nil, ErrInvalidGasLimit
	}

	totalArgs := util.PrepareArgs(args)
	address := vout.GetAddress()
	logger.WithFields(logger.Fields{
		"contract_address": address.String(),
		"invoked_function": function,
		"arguments":        totalArgs,
	}).Debug("Transaction: is executing the smart contract...")

	createContractUtxo := utxoIndex.GetContractCreateUTXOByPubKeyHash([]byte(vout.PubKeyHash))
	if createContractUtxo == nil {
		return 0, nil, ErrLoadError
	}
	engine.ImportSourceCode(createContractUtxo.Contract)
	engine.ImportLocalStorage(scStorage)
	engine.ImportContractAddr(address)
	engine.ImportSourceTXID(tx.ID)
	engine.ImportRewardStorage(rewards)
	engine.ImportTransaction(tx.Transaction)
	engine.ImportContractCreateUTXO(createContractUtxo)
	engine.ImportPrevUtxos(prevUtxos)
	engine.ImportCurrBlockHeight(currblkHeight)
	engine.ImportSeed(parentBlk.GetTimestamp())
	engine.ImportUtxoIndex(utxoIndex)
	_, err := engine.Execute(function, totalArgs)
	gasCount := engine.ExecutionInstructions()
	// record base gas
	baseGas, _ := tx.GasCountOfTxBase()
	gasCount += baseGas.Uint64()
	if err != nil {
		return gasCount, nil, err
	}
	return gasCount, engine.GetGeneratedTXs(), err
}

// Execute contract and return the generated transactions
func (tx *TxContract) CollectContractOutput(utxoIndex *lutxo.UTXOIndex, prevUtxos []*utxo.UTXO, isContractDeployed bool, scStorage *scState.ScState,
	engine ScEngine, currBlkHeight uint64, parentBlk *block.Block, minerAddr account.Address, rewards map[string]string, count int) (generatedTxs []*transaction.Transaction, err error) {
	gasCount, generatedTxs, err := tx.Execute(prevUtxos, isContractDeployed, utxoIndex, scStorage, rewards, engine, currBlkHeight, parentBlk)
	if err != nil {
		logger.WithError(err).Error("BlockProducer: executeSmartContract error.")
	}
	// record gas used
	if !tx.GasPrice.IsZero() {
		if gasCount > 0 {
			minerTA := account.NewTransactionAccountByAddress(minerAddr)
			grtx, err := NewGasRewardTx(minerTA, currBlkHeight, common.NewAmount(gasCount), tx.GasPrice, count)
			if err == nil {
				generatedTxs = append(generatedTxs, &grtx)
			}
		}
		gctx, err := NewGasChangeTx(tx.GetDefaultFromTransactionAccount(), currBlkHeight, common.NewAmount(gasCount), tx.GasLimit, tx.GasPrice, count)
		if err == nil {
			generatedTxs = append(generatedTxs, &gctx)
		}
	}
	return generatedTxs, nil
}

//NewRewardTx creates a new transaction that gives reward to addresses according to the input rewards
func NewRewardTx(blockHeight uint64, rewards map[string]string) transaction.Transaction {

	bh := make([]byte, 8)
	binary.BigEndian.PutUint64(bh, uint64(blockHeight))

	txin := transactionbase.TXInput{nil, -1, bh, transaction.RewardTxData}
	txOutputs := []transactionbase.TXOutput{}
	for address, amount := range rewards {
		amt, err := common.NewAmountFromString(amount)
		if err != nil {
			logger.WithError(err).WithFields(logger.Fields{
				"address": address,
				"amount":  amount,
			}).Warn("Transaction: failed to parse reward amount")
		}
		acc := account.NewTransactionAccountByAddress(account.NewAddress(address))
		txOutputs = append(txOutputs, *transactionbase.NewTXOutput(amt, acc))
	}
	tx := transaction.Transaction{nil, []transactionbase.TXInput{txin}, txOutputs, common.NewAmount(0), common.NewAmount(0), common.NewAmount(0), time.Now().UnixNano() / 1e6, transaction.TxTypeReward}

	tx.ID = tx.Hash()

	return tx
}

// NewGasRewardTx returns a reward to miner, earned for contract execution gas fee
func NewGasRewardTx(to *account.TransactionAccount, blockHeight uint64, actualGasCount *common.Amount, gasPrice *common.Amount, uniqueNum int) (transaction.Transaction, error) {
	fee := actualGasCount.Mul(gasPrice)
	txin := transactionbase.TXInput{nil, -1, getUniqueByte(blockHeight, uniqueNum), transaction.GasRewardData}
	txout := transactionbase.NewTXOutput(fee, to)
	tx := transaction.Transaction{nil, []transactionbase.TXInput{txin}, []transactionbase.TXOutput{*txout}, common.NewAmount(0), common.NewAmount(0), common.NewAmount(0), time.Now().UnixNano() / 1e6, transaction.TxTypeGasReward}
	tx.ID = tx.Hash()
	return tx, nil
}

// NewGasChangeTx returns a change to contract invoker, pay for the change of unused gas
func NewGasChangeTx(to *account.TransactionAccount, blockHeight uint64, actualGasCount *common.Amount, gasLimit *common.Amount, gasPrice *common.Amount, uniqueNum int) (transaction.Transaction, error) {
	if gasLimit.Cmp(actualGasCount) <= 0 {
		return transaction.Transaction{}, transaction.ErrNoGasChange
	}
	change, err := gasLimit.Sub(actualGasCount)

	if err != nil {
		return transaction.Transaction{}, err
	}
	changeValue := change.Mul(gasPrice)

	txin := transactionbase.TXInput{nil, -1, getUniqueByte(blockHeight, uniqueNum), transaction.GasChangeData}
	txout := transactionbase.NewTXOutput(changeValue, to)
	tx := transaction.Transaction{nil, []transactionbase.TXInput{txin}, []transactionbase.TXOutput{*txout}, common.NewAmount(0), common.NewAmount(0), common.NewAmount(0), time.Now().UnixNano() / 1e6, transaction.TxTypeGasChange}

	tx.ID = tx.Hash()
	return tx, nil
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to account.Address, data string, blockHeight uint64, tip *common.Amount) transaction.Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	bh := make([]byte, 8)
	binary.BigEndian.PutUint64(bh, uint64(blockHeight))
	toAccount := account.NewTransactionAccountByAddress(to)
	txin := transactionbase.TXInput{nil, -1, bh, []byte(data)}
	txout := transactionbase.NewTXOutput(transaction.Subsidy.Add(tip), toAccount)
	tx := transaction.Transaction{nil, []transactionbase.TXInput{txin}, []transactionbase.TXOutput{*txout}, common.NewAmount(0), common.NewAmount(0), common.NewAmount(0), time.Now().UnixNano() / 1e6, transaction.TxTypeCoinbase}
	tx.ID = tx.Hash()

	return tx
}

// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(utxos []*utxo.UTXO, sendTxParam transaction.SendTxParam) (transaction.Transaction, error) {
	fromAccount := account.NewTransactionAccountByAddress(sendTxParam.From)
	toAccount := account.NewTransactionAccountByAddress(sendTxParam.To)
	sum := transaction.CalculateUtxoSum(utxos)
	change, err := transaction.CalculateChange(sum, sendTxParam.Amount, sendTxParam.Tip, sendTxParam.GasLimit, sendTxParam.GasPrice)
	if err != nil {
		return transaction.Transaction{}, err
	}
	txType := transaction.TxTypeNormal
	if sendTxParam.Contract != "" {
		txType = transaction.TxTypeContract
	}
	tx := transaction.Transaction{
		nil,
		prepareInputLists(utxos, sendTxParam.SenderKeyPair.GetPublicKey(), nil),
		prepareOutputLists(fromAccount, toAccount, sendTxParam.Amount, change, sendTxParam.Contract),
		sendTxParam.Tip,
		sendTxParam.GasLimit,
		sendTxParam.GasPrice,
		time.Now().UnixNano() / 1e6,
		txType,
	}
	tx.ID = tx.Hash()

	err = tx.Sign(sendTxParam.SenderKeyPair.GetPrivateKey(), utxos)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return tx, nil
}

func NewSmartContractDestoryTX(utxos []*utxo.UTXO, contractAddr account.Address, sourceTXID []byte) transaction.Transaction {
	sum := transaction.CalculateUtxoSum(utxos)
	tips := common.NewAmount(0)
	gasLimit := common.NewAmount(0)
	gasPrice := common.NewAmount(0)

	tx, _ := NewContractTransferTX(utxos, contractAddr, account.NewAddress(transaction.SCDestroyAddress), sum, tips, gasLimit, gasPrice, sourceTXID)
	return tx
}

func NewContractTransferTX(utxos []*utxo.UTXO, contractAddr, toAddr account.Address, amount, tip *common.Amount, gasLimit *common.Amount, gasPrice *common.Amount, sourceTXID []byte) (transaction.Transaction, error) {
	contractAccount := account.NewTransactionAccountByAddress(contractAddr)
	toAccount := account.NewTransactionAccountByAddress(toAddr)
	if !contractAccount.IsValid() {
		return transaction.Transaction{}, account.ErrInvalidAddress
	}
	if isContract, err := contractAccount.GetPubKeyHash().IsContract(); !isContract {
		return transaction.Transaction{}, err
	}

	sum := transaction.CalculateUtxoSum(utxos)
	change, err := transaction.CalculateChange(sum, amount, tip, gasLimit, gasPrice)
	if err != nil {
		return transaction.Transaction{}, err
	}

	// Intentionally set PubKeyHash as PubKey (to recognize it is from contract) and sourceTXID as signature in Vin
	tx := transaction.Transaction{
		nil,
		prepareInputLists(utxos, contractAccount.GetPubKeyHash(), sourceTXID),
		prepareOutputLists(contractAccount, toAccount, amount, change, ""),
		tip,
		gasLimit,
		gasPrice,
		time.Now().UnixNano() / 1e6,
		transaction.TxTypeNormal,
	}
	tx.ID = tx.Hash()

	return tx, nil
}

//prepareInputLists prepares a list of txinputs for a new transaction
func prepareInputLists(utxos []*utxo.UTXO, publicKey []byte, signature []byte) []transactionbase.TXInput {
	var inputs []transactionbase.TXInput

	// Build a list of inputs
	for _, utxo := range utxos {
		input := transactionbase.TXInput{utxo.Txid, utxo.TxIndex, signature, publicKey}
		inputs = append(inputs, input)
	}

	return inputs
}

//prepareOutPutLists prepares a list of txoutputs for a new transaction
func prepareOutputLists(from, to *account.TransactionAccount, amount *common.Amount, change *common.Amount, contract string) []transactionbase.TXOutput {
	var outputs []transactionbase.TXOutput
	toAddr := to

	if toAddr.GetAddress().String() == "" {
		toAddr = account.NewContractTransactionAccount()
	}

	if contract != "" {
		outputs = append(outputs, *transactionbase.NewContractTXOutput(toAddr, contract))
	}

	outputs = append(outputs, *transactionbase.NewTXOutput(amount, toAddr))
	if !change.IsZero() {
		outputs = append(outputs, *transactionbase.NewTXOutput(change, from))
	}
	return outputs
}

func (tx *TxContract) GetTotalBalance(prevUtxos []*utxo.UTXO) (*common.Amount, error) {
	totalPrev := transaction.CalculateUtxoSum(prevUtxos)
	totalVoutValue, _ := tx.CalculateTotalVoutValue()
	totalBalance, err := totalPrev.Sub(totalVoutValue)
	if err != nil {
		return nil, transaction.ErrInsufficientBalance
	}
	totalBalance, _ = totalBalance.Sub(tx.Tip)
	return totalBalance, nil
}

func getUniqueByte(height uint64, uniqueNum int) []byte {
	bh := make([]byte, 8)
	binary.BigEndian.PutUint64(bh, uint64(height))
	bUnique := make([]byte, 2)
	binary.BigEndian.PutUint16(bUnique, uint16(uniqueNum))
	bh[0] = bUnique[0]
	bh[1] = bUnique[1]
	return bh
}
