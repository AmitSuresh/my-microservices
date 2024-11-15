import { useState, useEffect } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import axios from "axios";

// Define the structure of the ping response
interface PingResponse {
  success_message: string;
}

function App() {
  const [count, setCount] = useState(0);
  const [pingResponse, setPingResponse] = useState<string | null>(null);

  // Function to send a ping request to the Go Gin server
  const sendPingRequest = async () => {
    try {
      const response = await axios.get<PingResponse>("/api/ping"); // Proxy server routes this to mTLS Go Gin server
      setPingResponse(response.data.success_message);
    } catch (error) {
      console.error("Error making mTLS request:", error);
      setPingResponse("Error connecting to server");
    }
  };

  // Send ping request on component mount
  useEffect(() => {
    sendPingRequest();
  }, []);

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
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
        <p>Ping response: {pingResponse}</p>{" "}
        {/* Display the ping response here */}
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
