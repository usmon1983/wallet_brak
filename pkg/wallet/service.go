package wallet


import (
	"github.com/usmon1983/wallet/pkg/types"
	"github.com/google/uuid"
	"errors"
	"os"
	"log"
	"strconv"
	"fmt"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAmountMustBePositive = errors.New("amount must be greater than 0")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("balance is null")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite not found")
var ErrFileNotFound = errors.New("file not found")

type Service struct {
	nextAccountID int64 //Для генерации уникального номера аккаунта
	accounts []*types.Account
	payments []*types.Payment
	favorites []*types.Favorite
}

type Error string

func (e Error) Error() string  {
	return string(e)
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error)  {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account {
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)

	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}

func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error)  {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}
	
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}

	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}

	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment {
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, accounts := range s.accounts {
		if accounts.ID == accountID {
			account = accounts
			break
		} 
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil 
}

func (s *Service) Reject(paymentID string) error  {
	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return err
	}
	
	account, err := s.FindAccountByID(payment.AccountID)
	if err != nil {
		return nil
	}
	
	payment.Status = types.PaymentStatusFail
	account.Balance += payment.Amount
	return nil
}

func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error)  {
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}


func (s *Service) Repeat(paymentID string) (*types.Payment, error)  {
	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	repeatPayment, err := s.Pay(payment.AccountID, payment.Amount, payment.Category)
	if err != nil {
		return nil, err
	}

	return repeatPayment, nil
}

func (s *Service) FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	for _, favorite := range s.favorites {
		if favorite.ID == favoriteID {
			return favorite, nil
		}
	}
	return nil, ErrFavoriteNotFound
}

func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error)  {
	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}
	favoriteID := uuid.New().String()
	favorite := &types.Favorite {
		ID: favoriteID,
		AccountID: payment.AccountID,
		Name: name,
		Amount: payment.Amount,
		Category: payment.Category,
	}
	s.favorites = append(s.favorites, favorite)
	return favorite, nil
}

func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error)  {
	favoritePayment, err := s.FindFavoriteByID(favoriteID)
	if err != nil {
		return nil, err
	}
	payFavorite, err := s.Pay(favoritePayment.AccountID, favoritePayment.Amount, favoritePayment.Category)
	if err != nil {
		return nil, err
	}

	return payFavorite, nil
}

type testServiceUser struct {
	*Service
}

func mewTestServiceUser() *testServiceUser {
	return &testServiceUser{Service: &Service{}}
}

type testAccountUser struct {
	phone types.Phone
	balance types.Money
	payments []struct {
		amount types.Money
		category types.PaymentCategory
	}
}

var defaultTestAccountUser = testAccountUser {
	phone: "+992000000077",
	balance: 55_000_00,
	payments: []struct {
		amount types.Money
		category types.PaymentCategory	
	}{
		{amount: 5_000_00, category: "relax"},
	},
}

func (s *testServiceUser) addAccountUser(data testAccountUser) (*types.Account, []*types.Payment, error) {
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposit account, error = %v", err)
	}

	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments {
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}
	return account, payments, nil
}
func (s *Service) ExportToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	defer func () {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	data := ""
	for _, account := range s.accounts {
		id := strconv.Itoa(int(account.ID)) + ";"
		phone := string(account.Phone) + ";"
		balance := strconv.Itoa(int(account.Balance))

		data += id
		data += phone
		data += balance + "|"
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	return nil
}