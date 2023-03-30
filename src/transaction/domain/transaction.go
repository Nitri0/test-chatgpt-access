package domain

import (
	"fmt"
	"time"
)

type Checker struct{}

type LevelPriority int

const LOW_PRIORITY LevelPriority = 1
const HIGH_PRIORITY LevelPriority = 2

type TransactionStatus int

const WAITING_AUTORIZE_STATE TransactionStatus = 1
const AUTORIZE_REMOTELY_STATE TransactionStatus = 2
const AUTORIZE_LOCALY_STATE TransactionStatus = 3

type CheckpointStatus struct {
	Status TransactionStatus
	date   time.Time
}

type Transaction struct {
	ID                int
	PosNumber         int
	StoreNumber       int
	Priority          LevelPriority
	TransactionNumber int
	Anulations        []Anulation
	ChangePrice       []ChangePrice
	Status            []CheckpointStatus
	CreateAt          time.Time
}

func (t *Transaction) SetStatus(idAnulation int, trxStatus TransactionStatus) error {
	lastCheckPoint := t.Status[len(t.Status)-1]
	if lastCheckPoint.Status == trxStatus {
		return fmt.Errorf("Transaction is already %v", trxStatus)
	}

	checkpoint := CheckpointStatus{
		Status: trxStatus,
		date:   time.Now().UTC(),
	}
	t.Status = append(t.Status, checkpoint)
	return nil
}

func (t *Transaction) AddAnulation() error {

	return nil
}

type TransactionRepository interface {
	GetTransaction() (Transaction, error)
	CreateTransaction() (Transaction, error)
	UpdateTransaction() (Transaction, error)
}
