package main_test

import (
	"log"
	"sync"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MainTest struct {
	suite.Suite
	ApiClient *resty.Client
}

func (suite *MainTest) SetupTest() {
	suite.ApiClient = resty.New()
}

func (suite *MainTest) Test_Checkout_Pay() {
	log.Println("User 1 checkout")
	resp, _ := suite.ApiClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"user_1", "address":"address1"}`)).
		Post("http://localhost:8080/api/v1/order/checkout")
	log.Println(resp)

	assert.Equal(suite.T(), 200, resp.StatusCode())

	log.Println("\nUser 2 checkout")
	resp, _ = suite.ApiClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"user_2", "address":"address1"}`)).
		Post("http://localhost:8080/api/v1/order/checkout")
	log.Println(resp)

	assert.Equal(suite.T(), 200, resp.StatusCode())

	log.Println("\nUser 3 checkout")
	resp, _ = suite.ApiClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"user_3", "address":"address1"}`)).
		Post("http://localhost:8080/api/v1/order/checkout")
	log.Println(resp)

	assert.Equal(suite.T(), 200, resp.StatusCode())

	log.Println("\nUser 4 checkout")
	resp, _ = suite.ApiClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"user_4", "address":"address1"}`)).
		Post("http://localhost:8080/api/v1/order/checkout")
	log.Println(resp)

	assert.Equal(suite.T(), 200, resp.StatusCode())

	log.Println("\nUser 5 checkout")
	resp, _ = suite.ApiClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"user_5", "address":"address1"}`)).
		Post("http://localhost:8080/api/v1/order/checkout")
	log.Println(resp)

	assert.Equal(suite.T(), 200, resp.StatusCode())

	// Test concurrect payment

	log.Println("\n\n ==============Concurrency test started =======================")
	var waitgroup sync.WaitGroup
	waitgroup.Add(5)

	go func() {
		resp, _ := suite.ApiClient.R().
			SetHeader("Content-Type", "application/json").
			Put("http://localhost:8080/api/v1/order/pay/1")
		log.Println(resp.String() + " - user 1")
		waitgroup.Done()
	}()
	go func() {
		resp, _ := suite.ApiClient.R().
			SetHeader("Content-Type", "application/json").
			Put("http://localhost:8080/api/v1/order/pay/2")
		log.Println(resp.String() + " - user 2")
		waitgroup.Done()
	}()
	go func() {
		resp, _ := suite.ApiClient.R().
			SetHeader("Content-Type", "application/json").
			Put("http://localhost:8080/api/v1/order/pay/3")
		log.Println(resp.String() + " - user 3")
		waitgroup.Done()
	}()
	go func() {
		resp, _ := suite.ApiClient.R().
			SetHeader("Content-Type", "application/json").
			Put("http://localhost:8080/api/v1/order/pay/4")
		log.Println(resp.String() + " - user 4")
		waitgroup.Done()
	}()
	go func() {
		resp, _ := suite.ApiClient.R().
			SetHeader("Content-Type", "application/json").
			Put("http://localhost:8080/api/v1/order/pay/5")
		log.Println(resp.String() + " - user 5")
		waitgroup.Done()
	}()

	waitgroup.Wait()

	log.Println("\n\n ==============Concurrency test finished =======================")
}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTest))
}
