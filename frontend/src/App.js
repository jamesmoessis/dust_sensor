import './App.css';
import React from 'react';

import Controls from './Components/Controls';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h2>Dust Sensor Interface</h2>
      </header>
      <section className='App-body'>
        <Controls/>
      </section>
    </div>
  );
}

export default App;