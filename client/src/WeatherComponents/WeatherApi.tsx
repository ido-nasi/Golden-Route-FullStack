import React, { useState } from 'react'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import DatePicker, { DateObject, Value } from "react-multi-date-picker"
import InputIcon from "react-multi-date-picker/components/input_icon"
import "react-multi-date-picker/styles/backgrounds/bg-dark.css"
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
    }
    if (data.avgTemp !== undefined) {
      setAvgTemp(data.avgTemp);
      setValidHours([]);
    }
  }

  return (

    <div className='app'>
        <h1>IAF Weather Forecast</h1>
        
        <div className="getUserInput">
          <h3>Latitude: </h3>
          <input placeholder="value between -90 to 90" onChange={(e) => setLatitude(e.target.value)}/>
        </div>
        
        <div className="getUserInput">
          <h3>Longitude: </h3>
          <input placeholder="value between -180 to 180" onChange={(e) => setLongitude(e.target.value)} />
        </div>
        <ToastContainer/>
        
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
            <div className='container'>
              <h3>It is possible to carry out the mission</h3>
              <ul>{validHours.map((hour) => (<li key={hour}>{hour}</li>))}</ul>
          </div>
          ) : (
            <div>
              <h3>Couldn't find matching hour</h3>
              <h3>Average temperature: {avgTemp}</h3>
            </div>
          )
        }

      </div>
  )
}

export default WeatherApi;