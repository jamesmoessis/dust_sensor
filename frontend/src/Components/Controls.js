import React from 'react';

import PowerSwitch from './PowerSwitch';
import Threshold from './Threshold'

const Controls = ( {threshold, onThresholdChange} ) => {
    return (
        <div className='controls-container'>
            <Threshold onThresholdChange={onThresholdChange}/>
            <PowerSwitch threshold={threshold}/>
        </div>
    );
}

export default Controls;