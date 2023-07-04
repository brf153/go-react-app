import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import ListGroup from 'react-bootstrap/ListGroup';

function App() {
  const data =[
    {
      dish: "Marghetti Pizza",
      ingredients: "Tomato, Cheese, Bread",
      calories: "400",
      fat: "2"
    },
    {
      dish: "Vegas Zoodles",
      ingredients: "Zirchini, Avacado, Olives",
      calories: "220",
      fat: "3"
    },
    {
      dish: "Marghetti Pizza",
      ingredients: "Almonds, Refined Sugar, Honey",
      calories: "300",
      fat: "2"
    },
  ]
  return (
    <div className="App">
      <Button variant="primary">Track Today's Calorie</Button>
      <ListGroup>
        {
          data.map((i)=>{
            return(
              <ListGroup.Item>
                <div className='listgroup-item'>
                  <p>Dish: {i.dish}</p>
                  <p>Ingredients: {i.ingredients}</p>
                  <p>Calories: {i.calories}</p>
                  <p>Fat: {i.fat}</p>
                  <p><Button variant="primary">Delete Entry</Button></p>
                  <p><Button variant="primary">Change Ingredients</Button></p>
                  <p><Button variant="primary">Change Entry</Button></p>
                </div>
              </ListGroup.Item>
                
            ) 
          })
        }
     
      </ListGroup>
    </div>
  );
}

export default App;
