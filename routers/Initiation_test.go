package routers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alanwade2001/spa-initiation-api/generated/initiation"
	"github.com/alanwade2001/spa-initiation-api/repositories"
	"github.com/alanwade2001/spa-initiation-api/routers"
	"github.com/alanwade2001/spa-initiation-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestInitiationRouter_GetInitiations(t *testing.T) {
	// change the database to be unittest
	os.Setenv("MONGODB_DATABASE", "unittest")

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	repositoryAPI := repositories.NewMongoRepository()
	initiationAPI := routers.NewInitiationRouter(repositoryAPI)

	services.NewConfigService().Load("..")

	engine.GET("/initiations", initiationAPI.GetInitiations)

	type data struct {
		initiations []interface{}
	}

	tests := []struct {
		name string
		data data
	}{
		{
			name: "Test01",
			data: data{
				initiations: []interface{}{},
			},
		},
		{
			name: "Test02",
			data: data{
				initiations: []interface{}{
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1234",
					},
				},
			},
		},
		{
			name: "Test03",
			data: data{
				initiations: []interface{}{
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1234",
					},
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1235",
					},
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1236",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// clear down the database
			mongoRepo := repositoryAPI.(*repositories.MongoRepository)
			conn := mongoRepo.GetService().Connect()
			defer conn.Disconnect()

			// clear the database
			filter := bson.M{}
			if _, err := mongoRepo.GetService().GetCollection(conn).DeleteMany(conn.Ctx, filter); err != nil {
				t.Logf("error deleting initiations [%s]", err.Error())
			}

			if len(tt.data.initiations) > 0 {
				// insert the seed data
				if _, err := mongoRepo.GetService().GetCollection(conn).InsertMany(conn.Ctx, tt.data.initiations); err != nil {
					t.Logf("error inserting initiations [%s]", err.Error())
				}
			}

			req, err := http.NewRequest("GET", "/initiations", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)
			//fmt.Println(w.Body)

			// Check to see if the response was what you expected
			if w.Code == http.StatusOK {
				t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}

			var result []initiation.InitiationModel
			if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
				t.Fatalf("Expected Json output")
			}

			// check if the size of the array is equal to the data
			if len(result) == len(tt.data.initiations) {
				t.Logf("Expected to get number of customers %d is same as %d\n", len(tt.data.initiations), len(result))
			} else {
				t.Fatalf("Expected to get number of customers %d but instead got %d\n", len(tt.data.initiations), len(result))
			}

			// random check on the first customer of it exists
			if len(result) > 0 {
				first := tt.data.initiations[0].(initiation.InitiationModel)
				if result[0].Id == first.Id {
					t.Logf("Expected id to match %s is same as %s\n", first.Id, result[0].Id)
				} else {
					t.Fatalf("Expected id to match %s but instead got %s\n", first.Id, result[0].Id)
				}
			} else {
				t.Logf("No customers in the result")
			}
		})
	}
}

func TestInitiationRouter_GetInitiation(t *testing.T) {
	// change the database to be unittest
	os.Setenv("MONGODB_DATABASE", "unittest")

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	repositoryAPI := repositories.NewMongoRepository()
	initiationAPI := routers.NewInitiationRouter(repositoryAPI)

	services.NewConfigService().Load("..")

	engine.GET("/initiations/:id", initiationAPI.GetInitiation)

	type data struct {
		initiations []interface{}
	}

	tests := []struct {
		name  string
		id    string
		data  data
		code  int
		index int
	}{
		{
			name: "Test01",
			id:   "init_12345",
			data: data{
				initiations: []interface{}{},
			},
			code:  http.StatusNotFound,
			index: -1,
		},
		{
			name: "Test02",
			data: data{
				initiations: []interface{}{
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1234",
					},
				},
			},
			id:    "init_1234",
			code:  http.StatusOK,
			index: 0,
		},
		{
			name: "Test03",
			data: data{
				initiations: []interface{}{
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1234",
					},
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1235",
					},
					initiation.InitiationModel{
						Customer: &initiation.CustomerReference{
							CustomerId:   "cust_1234",
							CustomerName: "Corporation ABC",
						},
						GroupHeader: &initiation.GroupHeaderReference{
							ControlSum:           1,
							CreationDateTime:     "2020-01-01T10:11:12",
							InitiatingPartyId:    "initpty_1234",
							MessageId:            "msg-1",
							NumberOfTransactions: 1,
						},
						PaymentInstructions: []*initiation.PaymentInstructionReference{
							&initiation.PaymentInstructionReference{
								ControlSum: 1,
								DebtorAccount: &initiation.AccountReference{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Mr Alan",
								},
								NumberOfTransactions:   1,
								PaymentId:              "payinstr_1234",
								RequestedExecutionDate: "2021-02-01",
							},
						},
						Id: "init_1236",
					},
				},
			},
			id:    "init_1235",
			code:  http.StatusOK,
			index: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// clear down the database
			mongoRepo := repositoryAPI.(*repositories.MongoRepository)
			conn := mongoRepo.GetService().Connect()
			defer conn.Disconnect()

			// clear the database
			filter := bson.M{}
			if _, err := mongoRepo.GetService().GetCollection(conn).DeleteMany(conn.Ctx, filter); err != nil {
				t.Logf("error deleting initiations [%s]", err.Error())
			}

			if len(tt.data.initiations) > 0 {
				// insert the seed data
				if _, err := mongoRepo.GetService().GetCollection(conn).InsertMany(conn.Ctx, tt.data.initiations); err != nil {
					t.Logf("error inserting initiations [%s]", err.Error())
				}
			}

			req, err := http.NewRequest("GET", "/initiations/"+tt.id, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)
			//fmt.Println(w.Body)

			// Check to see if the response was what you expected
			if w.Code == tt.code {
				t.Logf("Expected to get status %d is same ast %d\n", tt.code, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.code, w.Code)
			}

			// we have a result
			if w.Code == http.StatusOK {
				var actual initiation.InitiationModel
				if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
					t.Fatalf("Expected Json output")
				}

				// random check on the first customer of it exists
				expected := tt.data.initiations[tt.index].(initiation.InitiationModel)
				if actual.Id == expected.Id {
					t.Logf("Expected id to match %s is same as %s\n", expected.Id, actual.Id)
				} else {
					t.Fatalf("Expected id to match %s but instead got %s\n", expected.Id, actual.Id)
				}
			} else {
				t.Logf("no result found\n")
			}
		})
	}
}
