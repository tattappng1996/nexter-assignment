This is a simple cashier-system using save in local (./save/cash register.json) instead of connecting complex DB

For running code in project 
-> go run main.go

Example 1
GET localhost:8080/exam-one

Example 2

API [GetCashRegister]
GET localhost:8080/exam-two/v1/cash-register
- CASE: If you want to check current cash in cash register


API [AddCashToCashRegister]
POST localhost:8080/exam-two/v1/cash-register
Body: 
{
    "cash_register": [
        {
            "max_quantity": 10,
            "quantity": 20,
            "value": 1000,
            "index": 0
        },
        {
            "max_quantity": 20,
            "quantity": 20,
            "value": 500,
            "index": 1
        },
        {
            "max_quantity": 15,
            "quantity": 20,
            "value": 100,
            "index": 2
        },
        {
            "max_quantity": 20,
            "quantity": 20,
            "value": 50,
            "index": 3
        },
        {
            "max_quantity": 30,
            "quantity": 20,
            "value": 20,
            "index": 4
        },
        {
            "max_quantity": 20,
            "quantity": 20,
            "value": 10,
            "index": 5
        },
        {
            "max_quantity": 20,
            "quantity": 20,
            "value": 5,
            "index": 6
        },
        {
            "max_quantity": 20,
            "quantity": 20,
            "value": 1,
            "index": 7
        },
        {
            "max_quantity": 50,
            "quantity": 20,
            "value": 0.25,
            "index": 8
        }
    ]
}
- CASE: If you want to add cash to cash register
- CASE: after customer is paid you can add bank_notes or coins from your "total_product_price"


API [CustomerPaid]
POST localhost:8080/exam-two/v1/customer-paid
Body: 
{
    "total_paid": 2000,
    "total_product_price": 850
}
- CASE: If you want to test customer paid

For more unit-test (Makefile)
- make [command]

-------------------------------------
Hope you enjoy this test if you got any problem
Just let me know
tattapong.kun@gmail.com
+66896361559