import "./MyStuff.css"
import { useState } from 'react';

function MyButton() {
    // hooks (functions starting with use... may be called only on top of components)
    const [count, setCount] = useState(0);
    function handleClick() {
      //alert('You clicked me!');
      console.log(`Count is ${count}`);
      setCount(count + 1);
      // here the increment is not yet registred ... weird
      console.log(`Click! And new count is ${count} (according to count variable, but actually it should be ${count + 1})`);
    }

    return (
      <button
        onClick={handleClick}
        className="thick-red-border">
        Button clicked {count} times
      </button>
    );
  }

  function ButtonWithProps({name, color}) {
    return <button style={{"background-color" : color}}> {name} </button>;
  }

  export {
    MyButton,
    ButtonWithProps,
  };