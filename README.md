# CurrencyCLI
CLI applicaltion  for returning if currency is supported using ISO 4217 code as argument.
Cheap Stocks, Inc is looking to build a cheap command line application centered around stocks that will be used by users in Africa. They are unsure of all the details now but they know they would like to support internationalization by allowing data to be displayed using supported currencies. The company has uploaded a CSV file of supported currencies to their server which can be accessed at: https://focusmobile-interview-materials.s3.eu-west-3.amazonaws.com/Cheap.Stocks.Internationalization.Currencies.csv.
The company would like you to build a command line application that takes in currency (using ISO 4217 code) as an argument and displays whether or not the currency is supported in the application.
The company will be making updates to the documents so your application should use the latest info to display the proper information to the user. Your assignment is to design and implement a functional MVP to the specifications listed.
Specifications
To guide you with building the MVP, here is a list of detailed specifications from Cheap Stocks, Inc. The command line application you design and implement should follow these guidelines.
CLI application that takes in an input flag corresponding to the input currency
Handle user input error

Running the command on the cmd should work
``` bash
 cd CurrencyCLI/ 
 sudo chmod a+x main

 ./main KES
