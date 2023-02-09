import React from 'react';
import { useEffect, useState } from 'react'
import './App.css'
import { Flight, IFlight } from './components/Flight'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import { Toast } from 'react-toastify/dist/components';


const BACKEND_URL = "http://localhost:4000" // change for {backend} in production
const API_URL = "https://api.open-meteo.com/v1/forecast?"



function App() {
  const [mass, setMass] = useState('');
  const [flights, setFlights] = useState<IFlight[]>([]);

  const calculateData = async () => {
    const response = await fetch(`${BACKEND_URL}/calculate`, {
      method: 'POST',
      mode: 'cors',
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
          mass: mass
      })
    });

    const data = await response.json();
    
    if (data.flightData !== undefined) {
      setFlights(flights.length === 0 ? [data.flightData] : flights.concat(data.flightData));
      // setMass('');
    }
    else {
      toast.error("Cargo Mass Value Must Be a Number");
    }
  }

  return (
    <div className='app'>
      <h1>IAF Flight Calculator</h1>
      <div className='getUserInput input'>
        <input
          placeholder='Enter cargo mass'
          value={mass}
          onChange={(e) => setMass(e.target.value)}
        />
        <button onClick={calculateData} /> 
      </div>
      <ToastContainer/>

      <div>
        <h2> Previous Flights</h2>  
      </div>
      
        {/* rendering the flights on the screen */}
        {
          flights?.length > 0 
          ? (
            <div className='container'>
              {flights.slice(0).reverse().map((flight) => ( 
                <Flight props={flight} key={flight.id}/>
              ))}
            </div>
           ) : (
            <div className='empty'>
              <h3 className="flightId">No Flight Records</h3>
            </div>
          )
        }
    
    </div>
  );
    
  
}

export default App
