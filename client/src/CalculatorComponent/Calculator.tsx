import React from 'react'
import { useState } from 'react'
import './Calculator.css'
import { Flight, IFlight } from './Flight'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import { Link } from 'react-router-dom'

const BACKEND_URL = "http://localhost:4000/api/v1" // remains localhost in production

function Calculator() {
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
      }
      else {
        toast.error("Cargo Mass Value Must Be a Number");
      }
    }
  
    return (
      <div className='app'>
        <h1>IAF Flight Calculator</h1>
        <Link to="/weather">Go to weather calculator</Link>
        <div className='getUserInput input'>
          <input
            placeholder='Enter cargo mass'
            value={mass}
            onChange={(e) => setMass(e.target.value)}
          />
          <button onClick={calculateData}>submit</button>
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

export default Calculator;