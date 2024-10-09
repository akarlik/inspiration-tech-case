package repositories

import (
	"fmt"
	"inspiration-tech-case/configuration"
	"inspiration-tech-case/internal/entities"
	"inspiration-tech-case/internal/utils"
	"strconv"
	"github.com/patrickmn/go-cache"
)

func NewAccountRepository(DB *cache.Cache, config configuration.Config) AccountRepository {
	return &AccountRepositoryImpl{DB: DB, config: config}
}

type AccountRepositoryImpl struct {
	DB     *cache.Cache
	config configuration.Config
}

func (repository *AccountRepositoryImpl) Seed() error {
	err := repository.initializeAccounts()
	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountRepositoryImpl) initializeAccounts() error {
	filename := repository.config.Get("DATASOURCE.SEED")
	accounts, err := getAccounts(filename)
	if err != nil {
		return err
	}
	for _, account := range accounts {
		repository.DB.Set(strconv.FormatInt(account.AccountId, 10), account, cache.NoExpiration)
	}

	return nil
}

func getAccounts(filename string) ([]entities.Account, error) {
	records, err := utils.ReadAccountsFromCSV(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %v", err)
	}

	var accounts []entities.Account
	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}
		accountId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid accountId at line %d: %v", i+1, err)
		}
		balance, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid balance at line %d: %v", i+1, err)
		}

		accounts = append(accounts, entities.Account{
			AccountId: accountId,
			Balance:   balance,
		})
	}

	return accounts, nil
}

func (repository *AccountRepositoryImpl) GetAccountByID(accountId int64) entities.Account {
	var account entities.Account
	accountData, ok := repository.DB.Get(strconv.Itoa(int(accountId)))
	if !ok {
		return account
	}
	account, ok = accountData.(entities.Account)
	if !ok {
		return account
	}
	return account
}
