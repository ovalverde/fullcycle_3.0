package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {

	err := t.IsAmountValid()

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) IsAmountValid() error {

	if t.Amount < 1 {
		return errors.New("the amount must b greater then 1")

	}

	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")

	}
	return nil

}

func (t *Transaction) SetCreditCard(card CreditCard) {

	t.CreditCard = card

}
