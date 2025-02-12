//import { useState } from 'react'
//import reactLogo from './assets/react.svg'
//import viteLogo from '/vite.svg'
//import './App.css'
import FlightCreate from './flights/FlightCreate'
import FlightEdit from './flights/FlightEdit'
import FlightList from './flights/FlightList'
import { BrowserRouter } from 'react-router-dom';
import { Routes } from 'react-router-dom';
import { Route } from 'react-router-dom';



function App() {

  return (
    <>
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<FlightList />} />
            <Route path="/flights/list" element={<FlightList />} />
            <Route path="/flights/create" element={<FlightCreate />} />
            <Route path="/flights/edit/:id" element={<FlightEdit />} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  )
}

export default App
