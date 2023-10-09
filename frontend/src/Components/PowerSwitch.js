import React, { useState } from 'react';
import ReactSwitch from 'react-switch';

const PowerSwitch = () => {
    const [checked, setChecked] = useState(true);

    const handleChange = val => {
        setChecked(val)
    }

    return (
        <div className="controls" id='power-switch'>
            
            <div id='power-title'>POWER</div>
            <ReactSwitch
                checked={checked}
                onChange={handleChange}
            />
        </div>
      );
}

export default PowerSwitch;