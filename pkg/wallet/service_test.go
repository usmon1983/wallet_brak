package wallet

import (
	"testing"
	"github.com/usmon1983/wallet/pkg/types"
	//"fmt"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_RegisterAccount_unsuccess(t *testing.T) {
	vc := Service{}

	accounts := []types.Account{
		{ID: 1, Phone: "+992000000001", Balance: 2_000_00},
		{ID: 2, Phone: "+992000000002", Balance: 3_000_00},
		{ID: 3, Phone: "+992000000003", Balance: 4_000_00},
		{ID: 4, Phone: "+992000000004", Balance: 5_000_00},
		{ID: 5, Phone: "+992000000005", Balance: 6_000_00},
		{ID: 6, Phone: "+992000000006", Balance: 7_000_00},
	}
	result, err := vc.RegisterAccount("+992000000007")
	for _, account := range accounts {
		if account.Phone == result.Phone {
			t.Errorf("invalid result, expected: %v, actual: %v", err, result)
			break
		}
	}
}

func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	vc := Service{}
	vc.RegisterAccount("+9920000001")

	account, err := vc.FindAccountByID(3)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_success(t *testing.T) {
	vc := Service{}
	vc.RegisterAccount("+9920000001")

	account, err := vc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := vc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := vc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Reject(pay.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Reject_fail(t *testing.T) {
	vc := Service{}
	vc.RegisterAccount("+9920000001")

	account, err := vc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := vc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := vc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	editPayID := pay.ID + "edit :)"
	err = vc.Reject(editPayID)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Repeat_success(t *testing.T) {
	vc := Service{}
	vc.RegisterAccount("+9920000001")

	account, err := vc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := vc.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := vc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err = vc.Repeat(pay.ID)
	if err != nil {
		t.Errorf("Repeat(): Error(): can't pay for an account(%v): %v", pay.ID, err)
	}
}

func TestService_FavoritePayment_success(t *testing.T) {
	vc := Service{}

	account, err := vc.RegisterAccount("+9920000001")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := vc.Pay(account.ID, 100_00, "babilon-m")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	favoritePay, err := vc.FavoritePayment(payment.ID, "MyFavoritePayment")
	if err != nil {
		t.Errorf("FavoritePayment(): Error(): can't create favorite payment(%v): %v", favoritePay.Name, err)
	}
}

func TestService_PayFromFavorite_success(t *testing.T) {
	vc := Service{}
	
	account, err := vc.RegisterAccount("+9920000001")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = vc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := vc.Pay(account.ID, 100_00, "babilon-m")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	favoritePay, err := vc.FavoritePayment(payment.ID, "MyFavoritePayment")
	if err != nil {
		t.Errorf("FavoritePayment(): Error(): can't create favorite payment(%v): %v", favoritePay.ID, err)
	}

	payFromFavorite, err := vc.PayFromFavorite(favoritePay.ID)
	if err != nil {
		t.Errorf("PayFromFavorite(): Error(): can't create payment from favorite(%v): %v", payFromFavorite, err)
	}
}

func TestService_ExportToFile_success(t *testing.T) {
	var vc Service
	vc.RegisterAccount("+992000000077")
	vc.RegisterAccount("+992000000088")
	vc.RegisterAccount("+992000000099")

	err := vc.ExportToFile("export.txt")
	if err != nil {
		t.Errorf("ExportToFile return not nil error, err = %v", err)
	}
}

func TestService_ImportFromFile_success(t *testing.T) {
	var vc Service
	err := vc.ImportFromFile("export.txt")
	if err != nil {
		t.Errorf("ImportFromFile return not nil error, err = %v", err)
	}
}

func TestService_Export_success(t *testing.T) {
	var vc Service
	vc.RegisterAccount("+992000000100")
	vc.RegisterAccount("+992000000101")
	vc.RegisterAccount("+992000000102")

	err := vc.Export("/data")
	if err != nil {
		t.Errorf("Export return not nil error, err = %v", err)
	}
}

func TestService_Import_success(t *testing.T) {
	var vc Service
	err := vc.Import("/data")
	if err != nil {
		t.Errorf("Import return not nil error, err = %v", err)
	}
}

func TestService_ExportAccountHistory_success(t *testing.T){
	var vc Service
	account, err := vc.RegisterAccount("+992000000077")
	if err != nil {
		t.Errorf("RegisterAccount returned not nil error, account = %v", account)
	}
	
	err = vc.Deposit(1, 100_00)
	if err != nil {
		t.Errorf("Deposit returned not nil error, error = %v", err)
	}
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")

	payments, err := vc.ExportAccountHistory(account.ID)
	if err != nil {
		t.Errorf("ExportAccountHistory returned not nil error, error = %v", err)
	}

	err = vc.HistoryToFiles(payments, "/data", 2)
	if err != nil {
		t.Errorf("HistoryToFiles returned not nil error, error = %v", err)
	}
}
func BenchmarkSumPayments(b *testing.B) {
	var vc Service
	account, err := vc.RegisterAccount("+992000000077")
	if err != nil {
		b.Errorf("method RegisterAccount returned not nil error, account = %v", account)
	}
	
	err = vc.Deposit(1, 100_00)
	if err != nil {
		b.Errorf("method Deposit returned not nil error, error = %v", err)
	}
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")
	_, err = vc.Pay(1, 10, "Service center")

	want := types.Money(50)
	got := vc.SumPayments(2)
	if want != got {
		b.Errorf("want = %v got = %v", want, got)
	}
}