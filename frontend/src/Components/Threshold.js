import React, { useEffect, useState } from 'react';
import axios from 'axios';

import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';
import Button from '@mui/material/Button';

import './Slider.css';

const Threshold = () => {
  
  const [thresValue, setThres] = useState(null);

  useEffect(() => {     // gets the threshold value to show in the current threshold area. 
    axios
    .get("http://localhost:8080/api/settings")
    .then((res) => {
      setThres(res.data.threshold);
    })
    .catch((error) => {
      console.log("failed to collect threshold from API", error);
    })
  }, [])

  return (
    <div className='control' id='threshold-slider'>
      <h2>THRESHOLD</h2>
        <Box sx={{ width: 300 }} id="slider-container"> 
          <h3>Current Threshold: {thresValue}</h3>
            <Slider id="slider"
            defaultValue={0} 
            aria-label="Default" 
            valueLabelDisplay="auto" 
            min={0} 
            max={100} 
            step={1}
            color='secondary'
            />
        </Box>
        <Button variant="contained" color="success" id="set-btn">
          Set
        </Button>
    </div>
  );
}

export default Threshold;