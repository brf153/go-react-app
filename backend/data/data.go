package data

type DataStruct struct{
	id string 
	dish string 
	ingredients string 
	calories string
	fat string
}

var Data = []DataStruct{
    {
      id: "1",
      dish: "Marghetti Pizza",
      ingredients: "Tomato, Cheese, Bread",
      calories: "400",
      fat: "2",
    },
    {
      id: "2",
      dish: "Vegas Zoodles",
      ingredients: "Zirchini, Avacado, Olives",
      calories: "220",
      fat: "3",
    },
    {
      id: "3",
      dish: "Marghetti Pizza",
      ingredients: "Almonds, Refined Sugar, Honey",
      calories: "300",
      fat: "2",
    },
}