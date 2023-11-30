import React from 'react';

import PowerSwitch from './PowerSwitch';
import Threshold from './Threshold'

const Controls = ( {threshold, onThresholdChange, onPowerChange} ) => {
    return (
        <div className='controls-container'>
            <Threshold onThresholdChange={onThresholdChange}/>
            <PowerSwitch />
        </div>
    );
}

export default Controls;