import React from 'react';

import PowerSwitch from './PowerSwitch';
import Threshold from './Threshold'

const Controls = () => {
    return (
        <div className='controls-container'>
            <PowerSwitch/>
            <Threshold/>
        </div>
    );
}

export default Controls;