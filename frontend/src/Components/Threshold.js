import React, { useEffect, useState } from 'react';
import axios from 'axios';

import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';
import Button from '@mui/material/Button';

import './Slider.css';

const Threshold = ( {onThresholdChange} ) => {
  
  const [thresValue, setThres] = useState(null);
  const [sliderValue, setSliderValue] = useState(null);

  useEffect(() => {     // gets the threshold value to show in the current threshold area. 
    axios.get("http://localhost:8080/api/settings")
    .then((res) => {
      setThres(res.data.threshold);
      console.log('threshold is collected', res.data)
      onThresholdChange(res.data.threshold);
    })
    .catch((error) => {
      console.log("failed to collect threshold from API", error);
    })
  }, [onThresholdChange])

  const getSliderVal = (event, val) => {

    setSliderValue(val)    // can use 'val' or 'event.target.value'

    // console.log(val, "val");
    // console.log(event.target.value, "event.target.val")

  }

  const putThres = () => {

    const settings = {
      "isOn": false,
      "threshold": sliderValue
    }

    const headers = {
      "Content-Type": "application/json"
    }

    /* For debugging purposes, window does not currently automatically reload when it sets the threshold.
        This makes it so you can still see console logs */

    axios.put("http://localhost:8080/api/settings", settings, headers)
    .then((res) => {
      console.log(res, `Threshold successfully changed to ${sliderValue}`);
      // alert('Threshold has been updated!');
      alert('Threshold has been updated! Refresh page to see the difference. (debug mode)');
    })
    // .then( () => {
    //    window.location.reload();
    // }) 
    .catch((error) => {
      console.log(error);
    })
  }
    
  return (
    <div className='control' id='threshold-slider'>
      <h2>THRESHOLD</h2>
        <Box sx={{ width: 300 }} id="slider-container"> 
          <h3>Current: {thresValue}</h3>
            <Slider id="slider"
            key={`slider-${thresValue}`}  // allows to change default value after it's fetched. https://stackoverflow.com/questions/62711040/a-component-is-changing-the-default-value-state-of-an-uncontrolled-slider-after
            defaultValue={thresValue}
            aria-label="Default" 
            valueLabelDisplay="auto" 
            min={0} 
            max={300} 
            step={1}
            color='secondary'
            onChange={getSliderVal}
            />
        </Box>
        <Button variant="contained" color="primary" id="set-btn" onClick={putThres}>
            Set
        </Button>
    </div>
  );
}

export default Threshold;