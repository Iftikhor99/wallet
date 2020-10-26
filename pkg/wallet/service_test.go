package wallet

import (
	"log"
	"os"
	"fmt"
	"reflect"
	"testing"

	"github.com/Iftikhor99/wallet/pkg/types"
	"github.com/google/uuid"
)

var defaultTestAccount = testAccount{
	phone:   "+992000000001",
	balance: 10_000_00,
	payments: []struct {
		amount   types.Money
		category types.PaymentCategory
	}{
		{amount: 1_000_00, category: "auto"},
	},
}

type testAccount struct {
	phone    types.Phone
	balance  types.Money
	payments []struct {
		amount   types.Money
		category types.PaymentCategory
	}
}

type testService struct {
	*Service
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}

func TestService_FindPaymentByID_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// Tpo6yem HavtTu nnaTéx

	payment := payments[0]

	got, err := s.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}

	// CpaBHMBaem nnaTexu
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
		return
	}
}

func TestService_FindPaymentByID_fail(t *testing.T) {
	// co3paém cepsuc
	s := newTestService()
	_, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// TipoOyem HaWTM HeECyWeECTByWuMA nnaTéex
	_, err = s.FindPaymentByID(uuid.New().String())
	if err == nil {
		t.Error("FindPaymentByID(): must return error, returned nil")
		return
	}

	if err != ErrPaymentNotFound {
		t.Errorf("FindPaymentByID(): must return ErrPaymentNotFound, returned = %v", err)
		return
	}

}

func TestService_Reject_success(t *testing.T) {

	// co3paém cepsuc
	s := newTestService()

	_, payments, err := s.addAccount(defaultTestAccount)

	if err != nil {
		t.Error(err)
		return
	}

	// TipoOyem OTMeHMTb nnaTéx

	payment := payments[0]

	err = s.Reject(payment.ID)

	if err != nil {
		t.Errorf("Reject(): error = %v", err)
		return
	}

	savedPayment, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("Reject(): can't find payment by id, error = %v", err)
		return
	}
	if savedPayment.Status != types.PaymentStatusFail {
		t.Errorf("Reject(): status didn't changed, payment = %v", savedPayment)
		return
	}

	savedAccount, err := s.FindAccountByID(payment.AccountID)

	if err != nil {
		t.Errorf("Reject(): can't find account by id, error = %v", err)
		return
	}

	if savedAccount.Balance != defaultTestAccount.balance {
		t.Errorf("Reject(): balance didn't changed, account = %v", savedAccount)
		return
	}

}

func TestService_Repeat_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// Tpo6yem HavtTu nnaTéx

	payment := payments[0]

	payment, err = s.Repeat(payment.ID)

	if err != nil {
		t.Errorf("Repeat(): error = %v", err)
		return
	}

}

func TestService_PayFromFavorite_success(t *testing.T) {

	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	// Tpo6yem HavtTu nnaTéx

	payment := payments[0]

	favorite, err := s.FavoritePayment(payment.ID, "Tcell")

	payment, err = s.PayFromFavorite(favorite.ID)

	if err != nil {
		t.Errorf("PayFromFavorite(): error = %v", err)
		return
	}

}

func BenchmarkSumPayments_Success(b *testing.B) {
	s := newTestService()
	accountTest, err := s.RegisterAccount("+992000000001")
	//fmt.Println(accountTest)
	if err != nil {
		b.Error(err)
		return
	}
	err = s.Deposit(accountTest.ID, 200_000_00)
	if err != nil {
		log.Print("Аккаунт пользователя не найден")
		//return
	}
	//fmt.Println(accountTest.Balance)
	_, err = s.Pay(accountTest.ID, 1_000_00, "food")
	_, err = s.Pay(accountTest.ID, 2_000_00, "food")
	_, err = s.Pay(accountTest.ID, 3_000_00, "food")
	_, err = s.Pay(accountTest.ID, 4_000_00, "food")
	_, err = s.Pay(accountTest.ID, 5_000_00, "food")
	_, err = s.Pay(accountTest.ID, 6_000_00, "auto")
	//fmt.Println(newP)
	result:= types.Money(0)
	want := types.Money(2100000)
	for i := 0; i < b.N; i++ {
		result = s.SumPayments(int(accountTest.ID))
		if result != want {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
	}
	fmt.Println(result)
}

func BenchmarkFilterPayments(b *testing.B) {
	s := newTestService()
	accountTest, err := s.RegisterAccount("+992000000001")
	if err != nil {
		b.Error(err)
		return
	}
	err = s.Deposit(accountTest.ID, 200_000_00)
	if err != nil {
		fmt.Println("Аккаунт пользователя не найден")
		//return
	}

	_, err = s.Pay(accountTest.ID, 1_000_00, "food")
	_, err = s.Pay(accountTest.ID, 2_000_00, "food")
	_, err = s.Pay(accountTest.ID, 3_000_00, "food")
	_, err = s.Pay(accountTest.ID, 4_000_00, "food")
	_, err = s.Pay(accountTest.ID, 5_000_00, "food")
	_, err = s.Pay(accountTest.ID, 6_000_00, "auto")
	//fmt.Println(newP)
	//result := []types.Payment{} 
	//payments, _ := s.FilterPayments(accountTest.ID,2)
	want := 6
	//fmt.Printf("want %v", want)
	for i := 0; i < b.N; i++ {
		paymentsF, err := s.FilterPayments(accountTest.ID,2)
		if err != nil {
			b.Error(err)
			return
		}
		result := len(paymentsF)
		if result != want {
			b.Fatalf("invalid result, result %v, want %v", result, want)
		}
	}
	//fmt.Printf("result %v", result)
}

func BenchmarkFilterPaymentsByFn(b *testing.B) {
	s := newTestService()
	accountTest, err := s.RegisterAccount("+992000000001")
	if err != nil {
		b.Error(err)
		return
	}
	err = s.Deposit(accountTest.ID, 200_000_00)
	if err != nil {
		fmt.Println("Аккаунт пользователя не найден")
		//return
	}

	_, err = s.Pay(accountTest.ID, 1_000_00, "food")
	_, err = s.Pay(accountTest.ID, 2_000_00, "food")
	_, err = s.Pay(accountTest.ID, 3_000_00, "food")
	_, err = s.Pay(accountTest.ID, 4_000_00, "food")
	_, err = s.Pay(accountTest.ID, 5_000_00, "food")
	_, err = s.Pay(accountTest.ID, 6_000_00, "auto")
	//fmt.Println(newP)
	//result := []types.Payment{} 
	//payments, _ := s.FilterPayments(accountTest.ID,2)
	want := 6
	//fmt.Printf("want %v", want)
	for i := 0; i < b.N; i++ {
		paymentsF, err := s.FilterPaymentsByFn(filter, 2)
		if err != nil {
			b.Error(err)
			return
		}
		result := len(paymentsF)
		if result != want {
			b.Fatalf("invalid result, result %v, want %v", result, want)
		}
	}
	//fmt.Printf("result %v", result)
}

func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	// perucTpupyemM TaM nonb30BaTena
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	// MononHsem ero cyéT
	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposity account, error = %v", err)
	}

	// BeinonHaem nnaTexu
	// MOKeM CO3MaTb CNavc Cpa3y HYKHOM ONMHbI, NOCKONbKy 3HAaeM Ppa3sMep
	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments {
		// Torga 30€Ccb padoTaem npocto yepes index, a He Yepe3 append
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}

	return account, payments, nil
}

func TestService_ExportAccountHistory_Success(t *testing.T) {
	s := newTestService()
	var foundPayments []types.Payment
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	payment := payments[0]
	foundPayments = append(foundPayments, *payment)
	got, err := s.ExportAccountHistory(1)

	if err != nil {
		t.Errorf("ExportAccountHistory(): error = %v", err)
		return
	}

	if !reflect.DeepEqual(foundPayments, got) {
		t.Errorf("ExportAccountHistory(): wrong payment returned want %v, got %v", foundPayments, got)
		return
	}
}

func TestService_ExportAccountHistory_Fail(t *testing.T) {
	s := newTestService()
	var foundPayments []types.Payment
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	payment := payments[0]
	foundPayments = append(foundPayments, *payment)
	got, err := s.ExportAccountHistory(5)
	fmt.Println(got)
	if err == nil {
		t.Errorf("ExportAccountHistory(): error = %v", err)
		return
	}

	
}

func TestService_HistoryToFiles(t *testing.T) {
	type args struct {
		payments []types.Payment
		dir      string
		record1  int
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.HistoryToFiles(tt.args.payments, tt.args.dir, tt.args.record1); (err != nil) != tt.wantErr {
				t.Errorf("Service.HistoryToFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Import(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Import(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("Service.Import() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Export_Success(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	s := newTestService()
	
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	payment := payments[0]

	favorite, err := s.FavoritePayment(payment.ID, "Tcell")
	fmt.Println(favorite)
	err = s.Export(wd)
	if err != nil {
		t.Error(err)
		return
	}
}

// func TestService_Export_Fail(t *testing.T) {
	
// 	s := newTestService()
	
// 	_, payments, err := s.addAccount(defaultTestAccount)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	payment := payments[0]

// 	favorite, err := s.FavoritePayment(payment.ID, "Tcell")
// 	fmt.Println(favorite)
// 	err = s.Export("c")
// 	fmt.Println(err)
// 	if err == nil {
// 		t.Error("Export(): must return error, returned nil")
// 		return
// 	}
// }

func TestService_Import_Success(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	s := newTestService()
	
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	payment := payments[0]

	favorite, err := s.FavoritePayment(payment.ID, "Tcell")
	fmt.Println(favorite)
	err = s.Import(wd)
	if err != nil {
		t.Error(err)
		return
	}
}


// func BenchmarkSumPaymentsWithProgress(b *testing.B) {
// 	s := newTestService()
	
// 	//fmt.Printf("want %v", want)
// 	for i := 0; i < b.N; i++ {
// 		paymentsF := s.SumPaymentsWithProgress()
// 		want := 45
// 		result := len(paymentsF)
// 		if result != want {
// 			b.Fatalf("invalid result, result %v, want %v", result, want)
// 		}
// 	}
// 	//fmt.Printf("result %v", result)
// }

// func BenchmarkFilterPaymentsNew(b *testing.B) {
// 	s := newTestService()
	
// 	//fmt.Println(newP)
// 	//result := []types.Payment{} 
// 	//payments, _ := s.FilterPayments(accountTest.ID,2)

// 	accountTest, err := s.RegisterAccount("+992000000001")
// 	if err != nil {
// 		fmt.Print(err)
		
// 	}
// 	err = s.Deposit(accountTest.ID, 2_000_000_000_000_000)
// 	if err != nil {
// 		fmt.Println("Аккаунт пользователя не найден")
// 		//return
// 	}
// 	//data := make([]int, 10_000_001)
// 	for j := 1; j < 1_000_001; j++ {
// 		_, _ = s.Pay(1, types.Money(j), "food")
				
// 	}
		
// 	want := 1_000_000
// 	//fmt.Printf("want %v", want)
// 	for i := 0; i < b.N; i++ {
// 		paymentsF, err := s.FilterPaymentsNew(accountTest.ID,100000)
// 	//	fmt.Print(paymentsF)
// 		if err != nil {
// 			b.Error(err)
// 			return
// 		}
// 		result := len(paymentsF)
// 		if result != want {
// 			b.Fatalf("invalid result, result %v, want %v", result, want)
// 		}
// 	}
// 	//fmt.Printf("result %v", result)
// }

// func BenchmarkSumPaymentsWithProgress(b *testing.B) {
// 	s := newTestService()
	
// 	//fmt.Println(newP)
// 	//result := []types.Payment{} 
// 	//payments, _ := s.FilterPayments(accountTest.ID,2)

// 	accountTest, err := s.RegisterAccount("+992000000001")
// 	if err != nil {
// 		fmt.Print(err)
		
// 	}
// 	err = s.Deposit(accountTest.ID, 2_000_000_000_000_000)
// 	if err != nil {
// 		fmt.Println("Аккаунт пользователя не найден")
// 		//return
// 	}
// 	data := make([]int, 1_000_000)
// 	for j := 1; j < 1_000_522; j++ {
// 		_, _ = s.Pay(1, types.Money(1), "food")
				
// 	}
// 	log.Print(len(data))
	
// 	want := Progress{}
// 	want.Part = 49999500000
// 	want.Result = 500000500000
// 	//chanel := make(chan Progress, 1)
// 	//fmt.Printf("want %v", want)
// 	for i := 0; i < b.N; i++ {
// 		paymentsF := s.SumPaymentsWithProgress()
// 	//	fmt.Print(paymentsF)
		
// 	//	chanel = paymentsF
// 		// for j := range paymentsF {
// 		// 	log.Print(j)
// 		// }
// 		resultChanel := <- paymentsF
// 		//defer close(paymentsF)
// //		log.Printf("resultChanel %v", resultChanel.Result)
// 		result := resultChanel
// 		if result != want {
// 			b.Fatalf("invalid result, result %v, want %v", result, want)
// 		}
// 	}
// 	//fmt.Printf("result %v", result)
// }