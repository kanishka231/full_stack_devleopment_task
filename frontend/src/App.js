
import './App.css';
import {useState} from 'react';

function App() {
  const [selectedCell, setSelectedCell] = useState([]);
  const [path, setPath] = useState([]);
  const handleCellClick = (row,col) => {
    if(selectedCell.length<2){
      setSelectedCell((prev)=>[...prev,{row,col}]);
    }else{
      setSelectedCell([{row,col}]);
    }
  }

  const isCellSelected = (row,col) => {
     if(selectedCell.length===0){
       return;
     }
     const [start,end] = selectedCell;
     return(
      (start && start.row === row && start.col === col) || 
      (end && end.row === row && end.col === col)
     );
  }
  const gridSize = 20;
 
  
  const findShortestPath = async () =>{
    if(selectedCell.length<2){
      alert('Please select start and end points');
      return;
    }
    try{
      const response = await fetch('http://localhost:5000/shortest-path',{
        method:'POST',
        headers:{
          'Content-Type':'application/json'
        },
        body:JSON.stringify({
          start:selectedCell[0],
          end:selectedCell[1],
        }),
      });
      if (!response.ok){
        throw new Error('Something went wrong');
      }
      const data = await response.json();
      setPath(data.path);
      console.log(data);
      }catch(err){
        console.error(err,"error");
        alert('Something went wrong');
        }

    }

    
  const isCellInPath = (row,col) =>{
    return selectedCell.some((point)=>point.row === row && point.col === col);
  }

  return (
   <div className='app'>
    <button 
    className='find-path-button'
    onClick={findShortestPath}>Find Shortest Path</button>
     <div className="grid-container">
      {Array.from({length:gridSize}).map(( _ ,row)=>(
        <div key= {row} className='grid-row'> 
        {Array.from({length:gridSize}).map((_,col)=>(
          <div key={`${row}-${col}`}
          onClick={()=>handleCellClick(row,col)}
          className={`grid-cell ${isCellSelected(row,col) ? "selected" : isCellInPath(row,col)?"path":""}`}>
            {`${row},${col}`}
          </div>
      ))}
      </div>
      ))}
    </div>
   </div>
  );
}

export default App;
