import { useState } from "react"

function Counter(){
    const [count, setCount] = useState(1);
    const onIncrement = function(){
        setCount(count + 1);
    };
    return(
        <>
            <div>
                <button className="btn btn-warning">-</button>
                <span>{count}</span>
                <button className="btn btn-warning" onclick ={onIncrement}>+</button>
            </div>
        </>
    )
}

export default Counter;