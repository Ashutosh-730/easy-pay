import { useState } from 'react'
import './App.css'

function App() {
  const [responseData, setResponseData] = useState("");

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080/customer/profile');
      const data = await response.json();
      setResponseData(responseData === "" ? data : "");
    } catch (error) {
      console.error('Error:', error);
    }
  };


  return (
    <>
      <div>
        <h1>Backend Response:  {responseData["message"]}</h1>
        <button onClick={fetchData}>Fetch Data</button>
        <p>Click the button to fetch data</p>
      </div>
    </>
  )
}

export default App
