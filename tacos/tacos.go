package main

import( "fmt"
        "io/ioutil"
        "net/http"
        "encoding/json"
)

type Component struct {
  Name string
  Recipe string
  Url string
}

type TacoRecipe struct {
  Condiment Component
  BaseLayer Component `json:"base_layer"`
  MixIn Component
  Seasoning Component
  Shell Component
}

func main() {
  var taco_recipe TacoRecipe

  url := "http://taco-randomizer.herokuapp.com/random"

  response, err := http.Get(url)
  if err != nil {
    fmt.Println("API call did not work!", err)
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Could not read the body of the response")
  }

  json.Unmarshal([]byte(body), &taco_recipe)
  printTacoRecipe(taco_recipe)
  printComponent(taco_recipe.Seasoning)
}

func printTacoRecipe(recipe TacoRecipe) {
  fmt.Printf("Taco Shell: %s\n", recipe.Shell.Name)
  fmt.Printf("Taco Base Layer: %s\n", recipe.BaseLayer.Name)
  fmt.Printf("Taco MixIn: %s\n", recipe.MixIn.Name)
  fmt.Printf("Taco Condiment: %s\n", recipe.Condiment.Name)
  fmt.Printf("Taco Seasoning: %s\n", recipe.Seasoning.Name)
}

func printComponent(component Component) {
  fmt.Printf("\n\nComponent Information:\n")
  fmt.Printf("Name: %s\n", component.Name)
  fmt.Printf("Recipe: %s\n", component.Recipe)
  fmt.Printf("URL: %s\n", component.Url)
}
