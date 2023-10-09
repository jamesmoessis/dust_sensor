import './App.css';
import React from 'react';

import Controls from './Components/Controls';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h2>Dust Sensor Interface</h2>
      </header>
      <body className='App-body'>
        <Controls/>
      </body>
    </div>
  );
}

export default App;