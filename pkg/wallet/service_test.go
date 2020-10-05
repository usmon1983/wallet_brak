package wallet

import (
	"testing"
	"github.com/usmon1983/wallet/pkg/types"
)

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