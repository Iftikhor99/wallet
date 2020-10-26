package main

import (
	"sync"
	//	"io/ioutil"
	"path/filepath"
	"github.com/Iftikhor99/wallet/pkg/types"
	//	"strings"
	//	"strconv"
	//	"io"
	"fmt"
	"log"
	"os"

	"github.com/Iftikhor99/wallet/pkg/wallet"
)

//Progress for
type Progress struct {
	Part int
	Result types.Money
}

func main() {
	svc := &wallet.Service{}
	accountTest1, err1 := svc.RegisterAccount("+992000000001")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	err1 = svc.Deposit(accountTest1.ID, 100_000_00)
	if err1 != nil {
		switch err1 {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}
	fmt.Println(accountTest1.Balance)

	// newP1, ee21 := svc.Pay(accountTest1.ID, 6_000_00, "car")
	// fmt.Println(newP1)
	// fmt.Println(ee21)

	accountTest, err := svc.RegisterAccount("+992000000002")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(accountTest.ID, 200_000_00)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}
	fmt.Println(accountTest.Balance)

	// data := make([]int, 10)
	// var newP *types.Payment
	// for i := range data {
	// 	newP, err = svc.Pay(accountTest.ID, types.Money(i), "food")
	// 	data[i] = i
	// }

	// newP, ee2 := svc.Pay(accountTest.ID, 1_000_00, "food")
	// newP, ee2 = svc.Pay(accountTest.ID, 2_000_00, "food")
	// newP, ee2 = svc.Pay(accountTest.ID, 3_000_00, "food")
	// newP, ee2 = svc.Pay(accountTest.ID, 4_000_00, "food")
	// newP, ee2 = svc.Pay(accountTest.ID, 5_000_00, "food")
	// newP, ee2 = svc.Pay(accountTest.ID, 6_000_00, "auto")

	fmt.Println(accountTest.Balance)
//	fmt.Println(newP)
	//fmt.Println(ee2)
	//fmt.Println(svc.pay)
	// newP2, ee3 := svc.FindPaymentByID(newP.ID)
	// fmt.Println(newP2)
	// fmt.Println(ee3)

	// //ee4 := svc.Reject(newP.ID)
	// //fmt.Println(account.Balance)
	// //fmt.Println(ee4)

	// newP3, ee5 := svc.Repeat(newP.ID)
	// fmt.Println(ee5)
	// fmt.Println(newP3.Amount)
	// fmt.Println(account.Balance)

	// fav, errFv := svc.FavoritePayment(newP.ID, "Tcell")
	// fmt.Println(errFv)
	// fmt.Println(fav)

	// newP4, eeFv2 := svc.PayFromFavorite("fav.ID")
	// fmt.Println(eeFv2)
	// fmt.Println(newP4)

	// fmt.Println(account.Balance)

	abs, err := filepath.Abs("data/readme.txt")

	if err != nil {
		log.Print(err)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(wd)
	log.Print(abs)

	// err = ioutil.WriteFile("c:/projects/wallet/data/readme1.txt", []byte(accountTest.ID), 0600)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// err = svc.ImportFromFile("c:/projects/wallet/data/accounts.dump")
	// if err != nil {
	//  	log.Print(err)
	//  	return
	//  }

	// err = svc.Export(wd)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// err = svc.Import(wd)
	// if err != nil {
	//  	log.Print(err)
	//  	return
	// }
	
	// paymentsFound, err := svc.ExportAccountHistory(newP1.AccountID)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// err = svc.HistoryToFiles(paymentsFound, wd, 2)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	//svc.Simple()
	// pay, err8 := svc.FilterPayments(2,4) 
	// log.Print(pay)
	// log.Print(err8)
	// log.Print(len(pay))

	// done := make(chan struct{},1)
	// log.Print(len(done))
	// done <- struct{}{}
	// <- done
	// log.Print("done")	
	//println(data[500])
	//data := make([]int, 1_00)
	// for i:= range data {
	// 	data[i] =i
	// }
	
	// for i := 1; i< 101; i++ {
	// 	_, _ = svc.Pay(1, types.Money(i), "food")
						
 	// }
	// foundPayments, _ := svc.ExportAccountHistoryWithoutID()
	// parts := 10

	// size := len(foundPayments)/parts
	// channels := make([]<-chan Progress, parts) 
	// for i := 0; i < parts; i++ {
	// 	ch := make(chan Progress)
	// 	channels[i] = ch
	// 	go func(ch chan<- Progress, foundPayments []types.Payment){
	// 		defer close(ch)
	// 		sum := Progress{}
	// 		for j, v := range foundPayments{
	// 			sum.Part += j
	// 			sum.Result += v.Amount
	// 		}
	// 		ch<- sum
	// 	}(ch, foundPayments[i*size:(i+1)*size])
		
	// }

	// total := Progress{}
	// for value := range 	merge(channels) {
	// 	total.Part += value.Part
	// 	total.Result += value.Result
	// }
	// log.Print(total)
	total := payProces()
	log.Print(total)	
	// svc.SumPaymentsWithProgress()

}

func payProces() Progress {
	svc := &wallet.Service{}
	accountTest1, err1 := svc.RegisterAccount("+992000000001")
	if err1 != nil {
		fmt.Println(err1)
		
	}

	err1 = svc.Deposit(accountTest1.ID, 1_000_000_000_000)
	if err1 != nil {
		switch err1 {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		
	}
	fmt.Println(accountTest1.Balance)


	for i := 1; i< 1_000_001; i++ {
		_, _ = svc.Pay(1, types.Money(i), "food")
						
 	}
	foundPayments, _ := svc.ExportAccountHistoryWithoutID()
	parts := 100_000

	size := len(foundPayments)/parts
	channels := make([]<-chan Progress, parts) 
	for i := 0; i < parts; i++ {
		ch := make(chan Progress)
		channels[i] = ch
		go func(ch chan<- Progress, foundPayments []types.Payment){
			defer close(ch)
			sum := Progress{}
			for j, v := range foundPayments{
				sum.Part += j
				sum.Result += v.Amount
			}
			ch<- sum
		}(ch, foundPayments[i*size:(i+1)*size])
		
	}

	total := Progress{}
	for value := range 	merge(channels) {
		total.Part += value.Part
		total.Result += value.Result
	}
	log.Print(total)
	return total
}

func merge(channels []<-chan Progress) <-chan Progress {
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	merged := make(chan Progress)

	for _, ch := range channels {
		go func(ch <- chan Progress) {
			defer wg.Done()
			for val := range ch {
				merged <- val
			}
		}(ch)
	}

	go func() {
		defer close(merged)
		wg.Wait()
	}()
	return merged
}

//Concurrently for
func Concurrently() int64 {

	wg := sync.WaitGroup{}

	wg.Add(2) // cKonbKOo ropyTMH péM

	mu := sync.Mutex{}
	sum := int64(0)

	go func() {
		defer wg.Done() // cooOwaem, 4TO 3aKkoHUunN
		val := int64(0)
		for i := 0; i < 1000; i++ {
			val++
		}
		mu.Lock()
		defer mu.Unlock()
		sum += val // TOMbKO B KOHUE 3anvcbiBaeM CYMMY
	}()

	go func() {
		defer wg.Done() // coo6waem, 4TO 3aKoH4UNN
		val := int64(0)
		for i := 0; i < 1000; i++ {
			val++
		}
		mu.Lock()
		defer mu.Unlock()
		sum += val // TOMbKO B KOHWE 3anucbiBaeM CyMMy

	}()

	wg.Wait()
	return sum
}
