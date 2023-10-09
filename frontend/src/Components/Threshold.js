import React from 'react';

const Threshold = () => {
    return (
        <form action='' className='controls' id='thres-sel'>
            <label for="selection">Select Threshold: </label>
            <br></br>
            <select id='selection' name='threshold'>
                <option value='0.1'>0.1</option>
                <option value='0.2'>0.2</option>
                <option value='0.3'>0.3</option>
                <option value='0.4'>0.4</option>
                <option value='0.5'>0.5</option>
                <option value='0.6'>0.6</option>
                <option value='0.7'>0.7</option>
                <option value='0.8'>0.8</option>
                <option value='0.9'>0.9</option>
            </select>
            <br></br>
            <input type='submit' value='Set'></input>
            <br></br>
            <label for="man-selection">Or Manual Selection:</label>
            <br></br>
            <input type='text' id='man-selection'></input>
            <input type='submit' value='Manual Set'></input>
        </form>
    );
}

export default Threshold;