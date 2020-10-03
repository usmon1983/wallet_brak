package wallet


import (
	"reflect"
	"github.com/usmon1983/wallet/pkg/types"
	"testing"
)

func TestService_RegisterAccount_unsuccess(t *testing.T)  {
	vc := Service{}

	accounts := []types.Account {
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
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(3)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_success(t *testing.T)  {
	s := &Service{}

	phone := types.Phone("+9920000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): can't register account, error = %v", err)
		return
	}

	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Reject(): can't deposit account, error = %v", err)
		return
	}

	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): can't create payment, error = %v", err)
		return
	}

	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v", err)
		return
	}
}

func TestService_FindPaymentByID_success(t *testing.T)  {
	s := &Service{}

	phone := types.Phone("+9920000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): can't register account, error = %v", err)
		return
	}

	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Reject(): can't deposit account, error = %v", err)
		return
	}

	payment, err := s.Pay(account.ID, 1000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): can't create payment, error = %v", err)
		return
	}

	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	} 

	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
	}
}