import { useState } from "react";
import { useNavigate } from "react-router-dom";
import PageHeader from "../header/Pageheader";
import axios from "axios";

function FlightCreate() {
    const[flight, setFlight] = useState({id:"45", number:"AI 121", airline_name:"Air India", source:"Trichy", destination:"Mysore", capacity:"150", price:7000})
    const OnBoxChange = (event) => {
        const newFlight = {...flight};
        newFlight[event.target.id] = event.target.vaue ;
        setFlight(newFlight);

    }

    const navigate = useNavigate();
    const OnCreate  = async () => {
        try {
            const baseUrl = "http://localhost:8080"
            const response = await axios.post(`${baseUrl}/flights` , {...flight, capacity:parseInt(flight.capacity), price:parseFloat(flight.price)});
            alert(response.data.message)
            navigate(('/flight/list'));
        }
        catch(error){
            alert("Server Error");
        }
    };

    return (
        <>
            <PageHeader PageNumber={2} />
            <h3><a className="btn btn-success" href="/flights/list">Go back</a>New Flight</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label" >Flight Number :</label>
                    <input type="text" className="form-control" id="number" placeholder="Please enter flight number" value={flight.number} onChange = {OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline name :</label>
                    <input type="text" className="form-control" id="airline_name" placeholder="Please enter airline name" value={flight.airline_name} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source :</label>
                    <input type="text" className="form-control" id="source" placeholder="Please enter the source" value={flight.source} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination :</label>
                    <input type="text" className="form-control" id="destination" placeholder="Please enter the destination" value={flight.destination} onChange = {OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity ( No of Seats ) :</label>
                    <input type="text" className="form-control" id="capacity" placeholder="Please enter the capacity" value={flight.capacity} onChange = {OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price :</label>
                    <input type="text" className="form-control" id="price" placeholder="Please enter the price" value={flight.price} onChange = {OnBoxChange}/>
                </div>
                <button className="btn btn-success" onClick={OnCreate}>Create Flight</button>
            </div>
        </>
    );
}

export default FlightCreate;