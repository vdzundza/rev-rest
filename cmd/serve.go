// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"

	log "log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

//var subs = new(Subscriptions)

var Activated = false

var NoActiveSubs = `{"data":[]}`

var ActiveSubs = `{
	"data": [
		{
			"application_fee_percent": null,
			"billing": "charge_automatically",
			"billing_cycle_anchor": 1532274447,
			"cancel_at_period_end": false,
			"canceled_at": null,
			"created": 1532274447,
			"current_period_end": 1534952847,
			"current_period_start": 1532274447,
			"customer": "cus_DHIKWYVit9bTTu",
			"days_until_due": null,
			"discount": null,
			"ended_at": null,
			"id": "sub_DHINURIYa4fqEZ",
			"items": {
				"data": [
					{
						"created": 1532274448,
						"id": "si_DHINtgAtF8cKZ4",
						"metadata": {},
						"object": "subscription_item",
						"plan": {
							"active": true,
							"aggregate_usage": "max",
							"amount": 100,
							"billing_scheme": "per_unit",
							"created": 1532273553,
							"currency": "usd",
							"id": "plan_DHI84gZdxoYTeR",
							"interval": "month",
							"interval_count": 1,
							"livemode": false,
							"metadata": {},
							"nickname": "Assets metered",
							"object": "plan",
							"product": "prod_DHI7DFn0wEKjND",
							"tiers": null,
							"tiers_mode": null,
							"transform_usage": null,
							"trial_period_days": null,
							"usage_type": "metered"
						},
						"subscription": "sub_DHINURIYa4fqEZ"
					},
					{
						"created": 1532274448,
						"id": "si_DHINlgL3uNLr7t",
						"metadata": {},
						"object": "subscription_item",
						"plan": {
							"active": true,
							"aggregate_usage": "sum",
							"amount": 1,
							"billing_scheme": "per_unit",
							"created": 1532272971,
							"currency": "usd",
							"id": "plan_DHHyrKHGGdXSSq",
							"interval": "month",
							"interval_count": 1,
							"livemode": false,
							"metadata": {},
							"nickname": "SRL metered",
							"object": "plan",
							"product": "prod_DHHwuNcI5MP3Q2",
							"tiers": null,
							"tiers_mode": null,
							"transform_usage": {
								"divide_by": 1000,
								"round": "up"
							},
							"trial_period_days": null,
							"usage_type": "metered"
						},
						"subscription": "sub_DHINURIYa4fqEZ"
					},
					{
						"created": 1532342455,
						"id": "si_DHaei3wbJqNwJ8",
						"metadata": {},
						"object": "subscription_item",
						"plan": {
							"active": true,
							"aggregate_usage": null,
							"amount": 5000,
							"billing_scheme": "per_unit",
							"created": 1532335804,
							"currency": "usd",
							"id": "plan_DHYrBaL8yCERf2",
							"interval": "month",
							"interval_count": 1,
							"livemode": false,
							"metadata": {},
							"nickname": "Basic Support",
							"object": "plan",
							"product": "prod_DHYqMK9szsT0mG",
							"tiers": null,
							"tiers_mode": null,
							"transform_usage": null,
							"trial_period_days": null,
							"usage_type": "licensed"
						},
						"quantity": 2,
						"subscription": "sub_DHINURIYa4fqEZ"
					}
				],
				"has_more": false,
				"object": "list",
				"total_count": 3,
				"url": "/v1/subscription_items?subscription=sub_DHINURIYa4fqEZ"
			},
			"livemode": false,
			"metadata": {},
			"object": "subscription",
			"plan": null,
			"quantity": null,
			"start": 1532342454,
			"status": "active",
			"tax_percent": null,
			"trial_end": null,
			"trial_start": null
		}
	],
	"has_more": false,
	"object": "list",
	"total_count": 1,
	"url": "/v1/customers/cus_DHIKWYVit9bTTu/subscriptions"
}`

// Get SUbscriptions
func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetSubscriptions called")

	var byt []byte
	fmt.Println(Activated)
	if Activated == false {
		byt = []byte(NoActiveSubs)
	} else {
		byt = []byte(ActiveSubs)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(dat)
	// json.NewEncoder(w).Encode(subs)
}

func GetStripeKey(w http.ResponseWriter, r *http.Request) {
	byt := []byte(`
	{
		"key": "pk_test_4DmGPmpaGOHr92X4uL2IxoKc"
	}
	`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(dat)
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	byt := []byte(`
	{
	
	}
	`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	Activated = true
	json.NewEncoder(w).Encode(dat)
}

func GetBillingInfo(w http.ResponseWriter, r *http.Request) {
	byt := []byte(`
	{
		"data": {
		 "card":  {
			 "name": "America",
			 "expires": "2012-04-23T18:25:43.511Z",
			 "holder_name": "Bruno Guez",
			 "postal_code": "333-545-45"
		 },
		 "contact": {
			 "name": "Bruno Guez",
			 "email": "email@email.com",
			 "number": "7777-777-77"
		 },
		 "address": {
			 "name": "",
			 "state": "",
			 "address_line": "",
			 "zip": "",
			 "city": "",
			 "country": "UA"
		 }
		}
	 }
	`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	Activated = true
	json.NewEncoder(w).Encode(dat)
}

func UpdateBillingInfo(w http.ResponseWriter, r *http.Request) {
	byt := []byte(`
	{
	
	}
	`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(dat)
}

func ListInvoices(w http.ResponseWriter, r *http.Request) {
	// sc.Init("sk_test_kPDzN6RqoH937Y6aCkIuSSF5", nil)
	byt := []byte(`
	{
		"data": [{
			"date": "2017-10-05T13:46:06Z",
			"invoice": "Davin",
			"amount": 758.26,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-12-26T12:56:08Z",
			"invoice": "Skippie",
			"amount": 1533.46,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-03-29T17:56:50Z",
			"invoice": "Philly",
			"amount": 821.5,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-09-13T11:19:16Z",
			"invoice": "Ezra",
			"amount": 1889.32,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-03-27T18:52:28Z",
			"invoice": "Vail",
			"amount": 410.54,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-07-29T11:29:03Z",
			"invoice": "Clare",
			"amount": 571.28,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-09-11T02:45:37Z",
			"invoice": "Morissa",
			"amount": 468.14,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-09-20T00:43:36Z",
			"invoice": "Hewitt",
			"amount": 652.44,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-09-25T22:21:30Z",
			"invoice": "Carrissa",
			"amount": 1895.0,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-10-07T09:20:56Z",
			"invoice": "Elva",
			"amount": 65.76,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-11-29T04:46:23Z",
			"invoice": "Lynnette",
			"amount": 524.57,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-02-10T19:59:38Z",
			"invoice": "Emmett",
			"amount": 329.62,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-10-21T12:00:18Z",
			"invoice": "Janessa",
			"amount": 928.08,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-04-20T02:02:35Z",
			"invoice": "Iorgo",
			"amount": 334.49,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-06-27T02:23:07Z",
			"invoice": "Robbie",
			"amount": 1382.74,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-02-11T18:15:34Z",
			"invoice": "Othelia",
			"amount": 1586.2,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-01-31T22:48:25Z",
			"invoice": "Lind",
			"amount": 1191.64,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-05-31T19:17:51Z",
			"invoice": "Susette",
			"amount": 838.81,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-04-25T17:39:30Z",
			"invoice": "Babbie",
			"amount": 1570.09,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-03-19T18:36:30Z",
			"invoice": "Cherie",
			"amount": 180.28,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-02-25T15:15:05Z",
			"invoice": "Dawn",
			"amount": 1116.49,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-05-12T07:52:21Z",
			"invoice": "Claretta",
			"amount": 1564.91,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-02-25T07:59:11Z",
			"invoice": "Dina",
			"amount": 1558.18,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-09-17T02:03:34Z",
			"invoice": "Maritsa",
			"amount": 1218.38,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-03-16T16:10:57Z",
			"invoice": "Mame",
			"amount": 1020.03,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-04-03T09:01:00Z",
			"invoice": "Baily",
			"amount": 615.34,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-04-30T09:42:17Z",
			"invoice": "Caesar",
			"amount": 830.42,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-03-22T15:05:39Z",
			"invoice": "Hetty",
			"amount": 1548.74,
			"status": "paid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2018-04-24T10:07:26Z",
			"invoice": "Caressa",
			"amount": 86.06,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}, {
			"date": "2017-11-07T16:41:37Z",
			"invoice": "Markos",
			"amount": 246.23,
			"status": "unpaid",
			"link": "https://dashboard.stripe.com/emails/receipts/invrc_1D1BDqIeGD42qwSiQKZLqnx3/pdf"
		}]
		}
	`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	Activated = true
	json.NewEncoder(w).Encode(dat)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called 2")
		router := mux.NewRouter()
		allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS", "PATCH", "PUT"})
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedHeaders := handlers.AllowedHeaders([]string{
			"Organization-Id",
			"Authorization",
			"content-type",
			"X-Requested-With",
			"Access-Control-Allow-Origin",
			"Access-Control-Request-Headers",
			"Access-Control-Allow-Methods",
		})

		router.HandleFunc("/api/payments/subscriptions/", GetSubscriptions).Methods("GET")
		router.HandleFunc("/api/payments/stripe_key/", GetStripeKey).Methods("GET")
		router.HandleFunc("/api/payments/card/", AddCard).Methods("POST")
		router.HandleFunc("/api/payments/billing/history/", ListInvoices).Methods("GET")
		router.HandleFunc("/api/payments/billing/", GetBillingInfo).Methods("GET")
		router.HandleFunc("/api/payments/billing/", UpdateBillingInfo).Methods("PATCH")
		log.Fatal(http.ListenAndServe(":7777", handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router)))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
