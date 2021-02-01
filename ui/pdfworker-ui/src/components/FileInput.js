import axios from 'axios'
import React, { Component } from 'react'
import { Button } from '@material-ui/core';
import { Input, InputLabel, Container } from '@material-ui/core';
class FileInput extends Component {
    constructor() {
        super();
    this.state = {
        file: null,
        numFileInputs: 1
    };
}

    onFileChange = event => {
        this.setState({ file: event.target.files[0] });
    }

    onFileUpload = () => {
        const formData = new FormData();
        formData.append(
            this.state.file.name,
            this.state.file,
            this.state.file.name
        );
        axios.post("http://localhost:2527/uploadFiles", formData);
    }

    renderFileInputs = () => {
        var i=0;
        var fileInputs = [];
        for(i=0; i< this.state.numFileInputs; i++) {
            console.log("adding " + i + "th input")
            fileInputs.push(<Input key={i} type='file'> </Input>);
        }
        return fileInputs;
    }

    addImage = () => {
        console.log("am i coming here?");
        this.setState({numFileInputs: numFileInputs + 1});
    }

    render() {

        return (
            
            <div>
                {this.renderFileInputs()}
                <Button variant='contained' color='primary' onClick={this.addImage()}> Add Image </Button>
            </div>
        )

    }

    
}

export default FileInput