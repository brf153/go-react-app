import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "axios"
import { useEffect, useState } from 'react';

function App() {

  const [data, setData] = useState([])

  useEffect(()=>{
    fetchData()
  },[])

  const fetchData=async()=>{
    const response = await axios.get("http://localhost:4000/")
    // console.log(response)
    setData(response.data)
  }

  async function handleDelete(id){
    const response = await axios.delete(`http://localhost:4000/dish/${id}`)
    // console.log(response)
    fetchData()
  }

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
                  <p><Button variant="primary" onClick={()=>handleDelete(i.id)}>Delete Entry</Button></p>
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
