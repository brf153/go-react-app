import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import ListGroup from 'react-bootstrap/ListGroup';
import axios from "./axios"
import { useEffect, useRef, useState } from 'react';

function App() {

  const [data, setData] = useState([])
  const [visible, setVisible] = useState(false)
  const [id,setId] = useState()

  useEffect(()=>{
    fetchData()
  },[])

  const ingredientRef= useRef()

  const handleVisibilty=(id)=>{
    if (visible){
      setId("")
      setVisible(false)
      ingredientRef.current.style.display= "none"
    }
    else{
      setId(id)
      setVisible(true)
      ingredientRef.current.style.display= "flex"
    }
  }

  const fetchData=async()=>{
    const response = await axios.get("/")
    // console.log(response)
    setData(response.data)
  }

  async function handleDelete(id){
    const response = await axios.delete(`/dish/${id}`)
    // console.log(response)
    fetchData()
  }

  const handleSubmit=async(id)=>{
    const response = await axios.post(`/dish/ingredient/${id}`)
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
                  <p><Button variant="primary" onClick={()=>handleVisibilty(i.id)}>Change Ingredients</Button></p>
                  <p><Button variant="primary">Change Entry</Button></p>
                </div>
              </ListGroup.Item>
                
            ) 
          })
        }
     
      </ListGroup>
      <div ref={ingredientRef} class="card" style={{width:"18rem", position:"absolute"}}>
        <div class="card-body">
          <h5 class="card-title">Ingredients for #{id}</h5>
          <form onSubmit={()=>handleSubmit(id)}>
            <div class="mb-3">
              <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" />
              <div id="emailHelp" class="form-text">Write whatever you eat!</div>
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default App;
