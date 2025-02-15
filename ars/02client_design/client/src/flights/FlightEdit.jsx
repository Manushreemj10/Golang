import { useState } from "react";
import PageHeader from "../header/Pageheader";
import { useParams } from "react-router-dom";

function FlightEdit() {
    /*const[flight, setFlight] = useState({id:"45", number:"AI 121", airline_name:"Air India", source:"Trichy", destination:"Mysore", capacity:"150", price:7000})
    const OnBoxChange = (event) => {
        const newFlight = {...flight};
        newFlight[event.target.id] = event.target.vaue ;
        setFlight(newFlight);

    }
    const param = useParams();*/

    return (
        <>
            <PageHeader PageNumber={1}/>
            <h3><a className="btn btn-success" href="/flights/list">Go back</a>Edit Flight Ticket Price</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label">Flight Number :</label>
                    <div className="form-control" id="number" >???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline name :</label>
                    <div className="form-control" id="airlinr_name" >???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source :</label>
                    <div className="form-control" id="source">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination :</label>
                    <div className="form-control" id="destination">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity ( No of Seats ) :</label>
                    <div className="form-control" id="capacity">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price :</label>
                    <input type="text" className="form-control" id="price" placeholder="Please enter the price" />
                </div>
                <button className="btn btn-success">Update Price</button>
            </div>
        </>
    );
}

export default FlightEdit;