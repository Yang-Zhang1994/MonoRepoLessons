import { useState } from "react";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);
  const [name, setName] = useState("");

  return (
    <div>
      <h1>Hello, I'm Yang!</h1>
      <p>Welcome to my React Vite App 🚀</p>

      <button onClick={() => setCount(count + 1)}>
        Clicked {count} times
      </button>

      <input
        type="text"
        placeholder="Enter your name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <p>Your name: {name}</p>
    </div>
  );
}

export default App;
