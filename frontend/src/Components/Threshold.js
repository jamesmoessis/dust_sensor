import React from 'react';
import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';

const Threshold = () => {
  return (
    <div className='control' id='threshold-slider'>

        <Box sx={{ width: 300 }}>
            <h2>Threshold</h2>
            <Slider 
            defaultValue={0} 
            aria-label="Default" 
            valueLabelDisplay="auto" 
            min={0} 
            max={1} 
            step={0.1}
            color='secondary'
            />
        </Box>
    </div>
  );
}

export default Threshold;