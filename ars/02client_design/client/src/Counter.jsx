import { useState } from "react"

function Counter(){
    const [count, setCount] = useState(0);
    const onIncrement = function(){
        setCount(count + 1);
    };
    const onDecrement = function(){
        setCount(count - 1);
    };
    return(
        <>
            <div>
                <button className="btn btn-warning" onClick ={onDecrement}>-</button>
                <span>{count}</span>
                <button className="btn btn-warning" onClick ={onIncrement}>+</button>
            </div>
        </>
    )
}

export default Counter;