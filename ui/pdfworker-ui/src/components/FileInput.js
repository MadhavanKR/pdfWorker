import axios from 'axios'
import React, { Component } from 'react'
import { Button } from '@material-ui/core';
import { Input, InputLabel, Container, Card, FlatButton } from '@material-ui/core';
import {v4 as uuidv4 } from 'uuid'
import CustomButton from './CustomButton'
class FileInput extends Component {
    constructor() {
        super();
        this.state = {
            files: [],
            fileInputRefs: []
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
        console.log("rendering file inputs")
        var i = 0;
        var fileInputs = [];
        console.log("size = " + this.state.fileInputRefs.length)
        for(i=0; i<this.state.fileInputRefs.length; i++)  {
            var id = this.state.fileInputRefs[i].id;
            console.log("iterating number: " + i);
            console.log(id);
            fileInputs.push(
                <Card variant='outlined' key={this.state.fileInputRefs[i].id}>
                    <Button label="Choose File">
                        <Input type='file'> </Input>
                    </Button>
                    {<Button samplekey={this.state.fileInputRefs[i].id} onClick={() => this.removeImage(id)}> Remove File</Button> }
                </Card>);
        }
        return fileInputs;
    }

    removeImage = (index) => {
        console.log("removing image");
        console.log("index = " + index);
        var tempFileInputRefs = [];
        for(var i=0; i<this.state.fileInputRefs.length; i++)  {
            if(this.state.fileInputRefs[i].id != index) {
                tempFileInputRefs.push(this.state.fileInputRefs);
            } else {
                console.log("removing id: " + index)
            }
        }
        console.log("after deleting, length of fileInputRefs: " + tempFileInputRefs.length)
        this.setState({fileInputRefs: tempFileInputRefs});
    }

    addImage = () => {
        console.log("adding image");
        console.log("am i coming here?");
        var fileInputRef = {
            "id": uuidv4()
        }
        console.log(fileInputRef)
        this.setState({ fileInputRefs: this.state.fileInputRefs.concat(fileInputRef) });
    }

    render() {

        return (

            <div>
                {this.renderFileInputs()}
                <Button variant='contained' color='primary' onClick={this.addImage}> Add Image </Button>
            </div>
        )

    }


}

export default FileInput