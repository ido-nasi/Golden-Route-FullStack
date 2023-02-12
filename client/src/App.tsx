import React from 'react';

import { Route, Routes } from "react-router-dom"
import Calculator from './CalculatorComponent/Calculator';
import WeatherApi from './WeatherComponents/WeatherApi';

function App() {
    return (
      <Routes>
        <Route path="/" element={<Calculator />} />
        <Route path="/weather" element={<WeatherApi />} />
      </Routes>
    );
}

export default App
