package wallet


import (
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T)  {
	vc := Service{}
	result, err := vc.RegisterAccount("+992000000005")

	if err != nil {
		t.Errorf("invalid result, expected: %v, actual: %v", err, result)
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