package main


import (
    "fmt"
	"os"
	"bufio"
  "io"
  "log"
  "strings"
  "sync"
  "strconv"

)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type CityCounter struct {
  Avg float64
  Count  int
}

type countingMap struct {
  counters map[string]CityCounter
  mutex    sync.RWMutex
}
var mpHolder = countingMap{counters : make(map[string]CityCounter) }

func main() {


    f, err := os.Open("brc/measurements.txt")
 if err != nil {
   fmt.Println("cannot able to read the file", err)
   return
 }
defer f.Close()  

//  var wg sync.WaitGroup
reader := bufio.NewReader(f)

for {
    line, _, err := reader.ReadLine()
    if err != nil {
        if err == io.EOF {
            break
        }
        log.Fatalf("a real error happened here: %v\n", err)
    }
    var currentLineString string = string(line)

    splittedString := strings.Split(currentLineString, ";")


    var cityName string = splittedString[0]

    var temperature float64
    temperature, err = strconv.ParseFloat(splittedString[1], 64)

    if err != nil {
      panic(err)
  }

    _, ok := mpHolder.counters[cityName]
if ok {
  if entry, ok := mpHolder.counters[cityName]; ok {
    entry.Avg += temperature
  entry.Count++  
  mpHolder.counters[cityName] = entry
}
  
} else {

  mpHolder.counters[cityName] = CityCounter{
    Avg : temperature,
    Count : 1,
  }
}
}
calculateCityWiseTemp()
calculateOveralltemp()
}

func calculateOveralltemp () {
  var tmp float64 = 0;
  var cnt int = 0;
  for _, val := range mpHolder.counters {
    cnt += val.Count
    tmp += val.Avg
}
fmt.Println("Overall", tmp / float64(cnt))
}


func calculateCityWiseTemp () {
  m := make(map[string]float64)
  for city, val := range mpHolder.counters {

    _, ok := m[city]
    if(ok){
      // m[vl] = vl / 
    } else {
      // fmt.Println("m", m)
      // fmt.Println("city", city)
      // fmt.Println("vall", val.Avg /float64(val.Count))
      m[city] = val.Avg /float64(val.Count)
    }
}
fmt.Println(m)
}