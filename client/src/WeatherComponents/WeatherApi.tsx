import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import DatePicker, { Value } from "react-multi-date-picker"
import InputIcon from "react-multi-date-picker/components/input_icon"
import "react-multi-date-picker/styles/backgrounds/bg-dark.css"
import "react-multi-date-picker/styles/colors/red.css"
import "./Weather.css" 

const BACKEND_URL = "http://localhost:4000/api/v1" // remains localhost in production

function WeatherApi() {
  const [validHours, setValidHours] = useState([])
  const [longitude, setLongitude] = useState('')
  const [latitude, setLatitude] = useState('')
  const [date, setDate] = useState<Value>(new Date())
  const [avgTemp, setAvgTemp] = useState('')

  const getWeatherFromBackend = async () => {
    //never happens, but the compiler will not let to convert to date without it
    if (date === null) {
      console.log(date)
      return;
    }

    // have to create another object to be able to jsonify it properly and not turn into a number
    let dateJSONable = new Date(date.toString())    
    const response = await fetch(`${BACKEND_URL}/weatherApi`, {
      method: 'POST',
        mode: 'cors',
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            latitude: latitude,
            longitude: longitude,
            date: dateJSONable
        })
    });

    const data = await response.json();

    if (data.message !== undefined) {
      toast.error(data.message);
      return;
    }
    if (data.hours !== undefined) {
      setValidHours(data.hours);
      setAvgTemp('');
      toast.success("Updated possible hours for take-off")
    }
    if (data.avgTemp !== undefined) {
      setAvgTemp(data.avgTemp);
      setValidHours([]);
      toast.success("Updated the average temperature")
    }
  }

  return (

    <div className='app'>
        <h1>IAF Weather Forecast</h1>
        <Link to="/" style={{ textDecoration: 'none', color: "#bd4715", fontFamily: "Roboto Slab, serif", fontSize: "2.2rem"  }}>Go to flight calculator</Link>

        <div className="getUserInput">
          <h3>Latitude: </h3>
          <input 
            style={{fontSize: "1.5rem"}}
            placeholder="value between -90 to 90" 
            value={latitude}
            onChange={(e) => setLatitude(e.target.value)}
          />
        </div>
        
        <div className="getUserInput">
          <h3>Longitude: </h3>
          <input 
            style={{fontSize: "1.5rem"}}
            placeholder="value between -180 to 180" 
            value={longitude}
            onChange={(e) => setLongitude(e.target.value)} 
          />


        </div>
        <ToastContainer/>
        
        <h3>Pick a date: </h3>
        <DatePicker
          value={date}
          render={<InputIcon />}
          onChange={(e) => {setDate(e)}}
          className="bg-dark"
        />
        
        <div className="submitButton">
          <button onClick={getWeatherFromBackend}>Check Weather!</button>
        </div>
        
        {
          validHours?.length > 0 ? (
            <div style={{fontSize: "1.8rem"}}>
              <h3 style={{marginBottom:"1rem", fontSize: "2rem"}}>
                It is possible to carry out the mission on:</h3>
              <ul className="hourList" style={{display: "flex", alignItems: "center", flexDirection: "column"}}>
                {validHours.map((hour) => (<li style={{color:"#e3c3a8", font: "0.5rem"}} key={hour}>{hour}</li>))}</ul>
          </div>
          ) : (
            <div className="container1">
              <h3>Couldn't find matching hour</h3>
              <h3>Average temperature: {avgTemp} °C</h3>
            </div>
          )
        }

      </div>
  )
}

export default WeatherApi;