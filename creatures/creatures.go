package main

import ("encoding/json"
        "fmt")

type Creature struct {
  Type string
  Description string
  Furry bool
 }

func main() {
  // parse json with attributes that match the fields in the Creature structure
  parseCat()

  // Parse json that has an attribute that is not in the Creature Struct
//parseComplexCat()

}

func parseCat() {
  example_json := `{"type": "cat",
                    "description": "A four legged furry creature that meows.",
                    "furry": true }`

  var cat Creature
  // store json in the `cat` variable, which has data type Creature
  // convert string to a byte array before trying to Unmarshal!
  json.Unmarshal([]byte(example_json), &cat)

  // Access attributes of the structure with . notation
  fmt.Println("Type:", cat.Type)
  fmt.Println("Description:", cat.Description)
  fmt.Println("Furry:", cat.Furry)
}

func parseComplexCat() {
  // Contains extra attribute not in the Creature structure
  example_json := `{"type": "cat",                                              
                    "description": "A four legged furry creature that meows.",
                    "furry": true,
                    "purry": true }`

  var cat Creature
  // No errors parsing with extra attribute!
  json.Unmarshal([]byte(example_json), &cat)
  fmt.Println("Type:", cat.Type)
  fmt.Println("Description:", cat.Description)
  fmt.Println("Furry:", cat.Furry)
}
