import React from 'react'
import {Button} from '@material-ui/core'

const CustomButton = (props) => {
    console.log("props from custom button" + props.samplekey)
    return <Button>Click here</Button>
}

export default CustomButton