import React, { useEffect, useState } from 'react';
import axios from 'axios';

import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';
import Button from '@mui/material/Button';

import './Slider.css';
import { baseURL } from './URL';

const Threshold = ( {onThresholdChange} ) => {
  const [thresValue, setThres] = useState(100);
  const [sliderValue, setSliderValue] = useState(null);

  useEffect(() => {     // gets the threshold value to show in the current threshold area. 
    axios.get(baseURL)
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
      axios.get(baseURL)    // has to get the state of isOn to be consistent. 
      .then((res) => {
        let isItOn = res.data.isOn;
        const settings = {
          "isOn": isItOn,
          "threshold": sliderValue
        }
        const headers = {
          "Content-Type": "application/json"
        }
        axios.put(baseURL, settings, headers)
        .then((res) => {
          setThres(sliderValue);
          console.log(res, `Threshold successfully changed to ${sliderValue}`);
          alert('Threshold has been updated!');
        })
        .catch((error) => {
          console.log(error);
          alert('Failed to send threshold');
        })
      })
      .catch((error) => {
        console.log("couldn't get isOn state while changing threshold");
        alert("couldn't get isOn state while changing threshold")
      })
      .then(() => {

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
            max={100} 
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