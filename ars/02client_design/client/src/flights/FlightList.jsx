import PageHeader from "../header/Pageheader";

function FlightList() {
    return (
        <>
            <PageHeader PageNumber={1} />
            <h3>List of Flights</h3>
            <div className="container">
                <table className="table table-info table-striped">
                    <thead className="table-warning">
                        <tr>
                            <th scope="col">Flight Number</th>
                            <th scope="col">Airline Name</th>
                            <th scope="col">Source</th>
                            <th scope="col">Destination</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <th scope="row">AI 845</th>
                            <td>Air India</td>
                            <td>Mumbai</td>
                            <td>Abu dhabi</td>
                            <td><a href="/flights/edit/12345" className="btn btn-warning">Edit Price</a>
                                <button className="btn btn-danger">Delete</button></td>
                        </tr>
                        <tr>
                            <th scope="row">6E 151</th>
                            <td>Indigo</td>
                            <td>Hyderabad</td>
                            <td>Banglore</td>
                            <td><a href="/flights/edit/164532" className="btn btn-warning">Edit Price</a>
                                <button className="btn btn-danger">Delete</button></td>
                        </tr>
                        <tr>
                            <th scope="row">AI 171</th>
                            <td>Air India</td>
                            <td>Ahmedabad</td>
                            <td>Newyark</td>
                            <td><a href="/flights/edit/26745" className="btn btn-warning">Edit Price</a>
                                <button className="btn btn-danger">Delete</button></td>
                        </tr>
                        <tr>
                            <th scope="row">6E 7101</th>
                            <td>Indigo</td>
                            <td>Hyderabad</td>
                            <td>Manglore</td>
                            <td><a href="/flights/edit/45732" className="btn btn-warning">Edit Price</a>
                                <button className="btn btn-danger">Delete</button></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default FlightList;