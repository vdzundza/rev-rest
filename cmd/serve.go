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

	"log"
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
