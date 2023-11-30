import React, { useState, useEffect } from 'react';
import axios from 'axios';

import Box from '@mui/material/Box'
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import Switch from '@mui/material/Switch';
import Typography from '@mui/material/Typography';

import { baseURL } from './URL';
import '../App.css';

const PowerSwitch = () => {

  const AntSwitch = styled(Switch)(({ theme }) => ({
    width: 28,
    height: 16,
    padding: 0,
    display: 'flex',
    '&:active': {
      '& .MuiSwitch-thumb': {
        width: 15,
      },
      '& .MuiSwitch-switchBase.Mui-checked': {
        transform: 'translateX(9px)',
      },
    },
    '& .MuiSwitch-switchBase': {
      padding: 2,
      '&.Mui-checked': {
        transform: 'translateX(12px)',
        color: '#fff',
        '& + .MuiSwitch-track': {
          opacity: 1,
          backgroundColor: theme.palette.mode === 'dark' ? '#177ddc' : '#1890ff',
        },
      },
    },
    '& .MuiSwitch-thumb': {
      boxShadow: '0 2px 4px 0 rgb(0 35 11 / 20%)',
      width: 12,
      height: 12,
      borderRadius: 6,
      transition: theme.transitions.create(['width'], {
        duration: 200,
      }),
    },
    '& .MuiSwitch-track': {
      borderRadius: 16 / 2,
      opacity: 1,
      backgroundColor:
        theme.palette.mode === 'dark' ? 'rgba(255,255,255,.35)' : 'rgba(0,0,0,.25)',
      boxSizing: 'border-box',
    },
  }));

  const [isOn, setIsOn] = useState(null);

  const togglePower = (event) => {
    axios.get(baseURL)  // get the threshold again because that's much much simpler and easier than sharing the state
    .then((res) => {
      console.log("power on?: ", event.target.checked);
      let isItOn = event.target.checked;
      setIsOn(event.target.checked);
      console.log(isItOn, "is it on???");
      const powerSettings = {
        "isOn": isItOn,
        "threshold": res.data.threshold
      }
      const headers = {
        "Content-Type": "application/json"
      }
      axios.put(baseURL, powerSettings, headers)
      .then((res) => {
        console.log(res, `Power on successfully changed to ${isItOn}`);
      })
      .catch((error) => {
        console.log(error);
        alert("Power toggle failed");
      })
    })
  }

  useEffect(() => {
    axios.get(baseURL)
    .then((res) => {
      setIsOn(res.data.isOn);
      console.log(isOn, "is it on?");
      console.log('initial isOn state collected', res.data)
    })
    .catch((error) => {
      console.log("failed to collect state from API", error);
      alert("couldn't collect state from API")
    })
  }, [])

  /* Has to render the whole power switch separately whether its state is on or off */

  if (isOn) {
    return (
      <div className='control' id='power'>

        <h2>POWER</h2>
        <Box sx={{ width: 300 }}>
          <Stack direction="row" spacing={1} alignItems="center" justifyContent="center">
            <Typography color='white'>Off</Typography>
            <AntSwitch defaultChecked inputProps={{ 'aria-label': 'ant design' }} 
                onChange={togglePower}/>
            <Typography color='white'>On</Typography>
          </Stack>
        </Box>
      
      </div>
    );
  } else {
    return (
      <div className='control' id='power'>

        <h2>POWER</h2>
        <Box sx={{ width: 300 }}>
          <Stack direction="row" spacing={1} alignItems="center" justifyContent="center">
            <Typography color='white'>Off</Typography>
            <AntSwitch inputProps={{ 'aria-label': 'ant design' }} 
                onChange={togglePower}/>
            <Typography color='white'>On</Typography>
          </Stack>
        </Box>
      
      </div>
    );
  }

}

export default PowerSwitch;