package data

type DataStruct struct{
	  ID          string `json:"id"`
    Dish        string `json:"dish"`
    Ingredients string `json:"ingredients"`
    Calories    string `json:"calories"`
    Fat         string `json:"fat"`
}

var Data = []DataStruct{
    {
      ID: "1",
      Dish: "Marghetti Pizza",
      Ingredients: "Tomato, Cheese, Bread",
      Calories: "400",
      Fat: "2",
    },
    {
      ID: "2",
      Dish: "Vegas Zoodles",
      Ingredients: "Zirchini, Avacado, Olives",
      Calories: "220",
      Fat: "3",
    },
    {
      ID: "3",
      Dish: "Marghetti Pizza",
      Ingredients: "Almonds, Refined Sugar, Honey",
      Calories: "300",
      Fat: "2",
    },
}