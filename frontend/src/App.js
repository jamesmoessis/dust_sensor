import './App.css';
import React, { useState } from 'react';

import Controls from './Components/Controls';

// export const baseURL = "http://localhost:8080/api/settings";
export const baseURL = "https://jfzlkq6spl.execute-api.ap-southeast-2.amazonaws.com/api/settings";

function App() {

  const [threshold, setThreshold] = useState(null);

  const handleThresholdChange = (newThreshold) => {
    setThreshold(newThreshold);
  };

  return (
    <div className="App">
      <header className="App-header">
        <h2>Dust Sensor Interface</h2>
      </header>
      <section className='App-body'>
        <Controls threshold={threshold} onThresholdChange={handleThresholdChange}/>
      </section>
    </div>
  );
}

export default App;