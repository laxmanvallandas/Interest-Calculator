# assignment
This is a micro-service written in go to calculate montly installments plan and this service listens on port 8080
### Request payload:
```json
{
  "loanAmount": "5000",
  "nominalRate": "5.0",
  "duration": 24,
  "startDate": "2018-01-01T00:00:01Z"
}
```

### Sample Response:
```json
{
	"borrowerPayments": [{
			"borrowerPaymentAmount": "219.36",
			"date": "2018-01-01T00:00:00Z",
			"initialOutstandingPrincipal": "5000.00",
			"interest": "20.83",
			"principal": "198.53",
			"remainingOutstandingPrincipal": "4801.47"
		},
		{
			"borrowerPaymentAmount": "219.36",
			"date": "2018-02-01T00:00:00Z",
			"initialOutstandingPrincipal": "4801.47",
			"interest": "20.01",
			"principal": "199.35",
			"remainingOutstandingPrincipal": "4602.12"
		},
      ...
		{
			"borrowerPaymentAmount": "219.28",
			"date": "2019-12-01T00:00:00Z",
			"initialOutstandingPrincipal": "218.37",
			"interest": "0.91",
			"principal": "218.37",
			"remainingOutstandingPrincipal": "0"
		}
	]
}
```
### Formulaes:
Interest:
Interest = (Rate * Days in Month * Initial Outstanding Principal) / Days in
Year e.g. first installment = (0.05 * 30 * 5000.00) / 360 = 20.83 € (with
rounding)

Principal:
Principal = Annuity - Interest e.g. first principal = 219.36 - 20.83 = 198.53 €

BorrowerAmount:
Borrower Payment Amount(Annuity) = Principal + Interest e.g. first borrower
payment = 198.53 + 20.83 = 219.36 €

### How to build

For amd build
  ``` make ```

For arm64 build
  ```ARCH=arm64 make``` for arm

Creating docker image
  ``` make docker ```
### How to use

- cd assignment

- With executable, run ./assignment

- With docker image, run ``` docker run -it -p 8080:8080 assignment:<version> ```
   eg : ``` docker run -it -p 8080:8080 assignment:0.1.0 ```

- curl to REST endpoint http://localhost:8080/generate-plan to know view the plan
