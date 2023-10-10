import React from 'react';

import PowerSwitch from './PowerSwitch';
import Threshold from './Threshold'

const Controls = () => {
    return (
        <div className='controls-container'>
            <Threshold/>
            <PowerSwitch/>
        </div>
    );
}

export default Controls;