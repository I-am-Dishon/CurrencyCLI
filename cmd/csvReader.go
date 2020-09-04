package cmd

import (
    "os"
    "io"
    //"fmt"
    "log"
    "bufio"
    "encoding/csv"
)

/** 
* Currency
*/
type Currency struct {
    Country string   
    currency  string   
    isocode   string 
}

func readcsv() []Currency {
    csvFile, _ := os.Open("./Cheap.Stocks.Internationalization.Currencies.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var currency []Currency
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        currency = append(currency, Currency{
            Country: line[0],   
            currency:  line[1],   
            isocode:   line[2], 
        })
    }
    //fmt.Println(currency[9])
    return currency
}