import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { MyButton, ButtonWithProps } from './MyStuff'

function App() {
  const [count, setCount] = useState(0)
  const userName = "whee";
  const userData = { name: "Jane", surname: "Doe"};
  

  // rendering lists
  // const products = [
  //   { title: 'Cabbage', id: 1 },
  //   { title: 'Garlic', id: 2 },
  //   { title: 'Apple', id: 3 },
  // ];

  // first we turn the list into react objects
  // const listItems = products.map(product =>
  //   <li key={product.id}>
  //     {product.title}
  //   </li>
  // );

  const listItems = [
    <li key="1">Whee</li>,
    <li key="2">Bree</li>,
    <li key="3">Glee</li>
  ]

  //console.log("Type of listItems: ", typeof listItems)
  //console.log("Type of react object: ", typeof <button> whee </button>)

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      <p>
        This is my humble sentence :)
      </p>
      {
        /* 
        This is a comment btw. 
        React components must always start with Capital letters, HTMP tags with lower case 
        ... className seems to have no effect in component tags.
        */
      }
      <MyButton className="blue-edge" />
      <MyButton></MyButton>
      <h3>
        Merry christmas {userName}. May the stars shine down your path (or whatever)...
      </h3>
      {/* Passing style attribute as javascript object (first set of {} escapes from html into js, second set creates a js object) */}
      <p style={{"background-color" : "lightblue"}}> Btw we know your real name is {userData.name} (wink*) </p>
      
      <p>Rendering list of items as "ul"</p>
      <ul>{listItems}</ul>

      <p>Rendering list of items as "ol"</p>
      <ol>{listItems}</ol>
      <ButtonWithProps name="Light grey button" color="lightgrey"></ButtonWithProps>
      <ButtonWithProps name="Light green button" color="lightgreen"></ButtonWithProps>
    </>
  )
}

export default App
