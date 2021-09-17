package main

import (
    "fmt"
    "log"
)

//walletFacade фасад кошлеька - содержит структуры с данными для каждого кошлеька
type walletFacade struct {
    account      *account
    wallet       *wallet
    securityCode *securityCode
    notification *notification
    ledger       *ledger
}

//account аккаунт
type account struct {
    name string
}

//ledger учетная книга
type ledger struct {
}

//notification уведомления
type notification struct {
}

//securityCode код безопасности
type securityCode struct {
    code int
}

//walelt кошелек
type wallet struct {
    balance int
}

func main() {
    fmt.Println()
    walletFacade := newWalletFacade("abc", 1234)
    fmt.Println()

    err := walletFacade.addMoneyToWallet("abc", 1234, 10)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }

    fmt.Println()
    err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }
}

//newWalletFacade возвращает элемент типа *walletFacade
func newWalletFacade(accountID string, code int) *walletFacade {
    fmt.Println("Starting create account")
    walletFacacde := &walletFacade{
        account:      newAccount(accountID),
        securityCode: newSecurityCode(code),
        wallet:       newWallet(),
        notification: &notification{},
        ledger:       &ledger{},
    }
    fmt.Println("Account created")
    return walletFacacde
}

//addMoneyToWallet добавить денег на кошелек
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting add money to wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }
    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    w.wallet.creditBalance(amount)
    w.notification.sendWalletCreditNotification()
    w.ledger.makeEntry(accountID, "credit", amount)
    return nil
}
//deductMoneyFromWallet снять деньги с кошелька
func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting debit money from wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    err = w.wallet.debitBalance(amount)
    if err != nil {
        return err
    }
    w.notification.sendWalletDebitNotification()
    w.ledger.makeEntry(accountID, "credit", amount)
    return nil
}

//newAccount возвращает *account, создали аккаут
func newAccount(accountName string) *account {
    return &account{
        name: accountName,
    }
}

//checkAccount проверка аккаунта по имени
func (a *account) checkAccount(accountName string) error {
    if a.name != accountName {
        return fmt.Errorf("Account Name is incorrect")
    }
    fmt.Println("Account Verified")
    return nil
}

//makeEntry вход
func (s *ledger) makeEntry(accountID, txnType string, amount int) {
    fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
    return
}

//sendWalletCreditNotification уведомление при получении средств счета
func (n *notification) sendWalletCreditNotification() {
    fmt.Println("Sending wallet credit notification")
}

//sendWalletCreditNotification уведомление при пополнении счета
func (n *notification) sendWalletDebitNotification() {
    fmt.Println("Sending wallet debit notification")
}

//newSecurittyCode новый код безопасности
func newSecurityCode(code int) *securityCode {
    return &securityCode{
        code: code,
    }
}

//checkCode проверка кода безопасности
func (s *securityCode) checkCode(incomingCode int) error {
    if s.code != incomingCode {
        return fmt.Errorf("Security Code is incorrect")
    }
    fmt.Println("SecurityCode Verified")
    return nil
}

//newWallet создание кошелька возвращет *wallet
func newWallet() *wallet {
    return &wallet{
        balance: 0,
    }
}

//creditBalance пополнение счета
func (w *wallet) creditBalance(amount int) {
    w.balance += amount
    fmt.Println("Wallet balance added successfully")
    return
}

//debitBalance снятие со счета
func (w *wallet) debitBalance(amount int) error {
    if w.balance < amount {
        return fmt.Errorf("Balance is not sufficient")
    }
    fmt.Println("Wallet balance is Sufficient")
    w.balance = w.balance - amount
    return nil
}