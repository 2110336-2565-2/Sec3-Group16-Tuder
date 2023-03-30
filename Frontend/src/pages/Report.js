import React, { useState } from 'react';
import FormT from '../components/form/FormStyle.js';
import { toast } from 'react-hot-toast';
import submitIssueReportHandler from '../handlers/submitIssueReportHandler';

export default function Report(){

    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [contact, setContact] = useState('');

    async function submitHandler(e){
        e.preventDefault();
        let date = Date.now()
        const reportData = {
            title,
            description,
            contact,
            date
        }
        try{
            await submitIssueReportHandler(reportData)
            toast.success('Thank you for your attention.')
        } catch (error){
            toast.error(error.message)
        }
    }

    return (
        <form onSubmit={submitHandler} >
            <FormT.Div FormW='350px'>
                <FormT.Header>Problem report</FormT.Header>
                <FormT.Content>If you have any problem or question, let me know here</FormT.Content>
                <FormT.Content>
                    <FormT.TextInput BoxSize='200px' name='title' type='text' placeholder='Title' value={title} onChange={(e) => setTitle(e.target.value)}/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput  BoxSize='200px' name='description' type='description' placeholder='Description' value={description} onChange={(e) => setDescription(e.target.value)}/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.TextInput  BoxSize='200px' name='contact' type='contact' placeholder='contact' value={contact} onChange={(e) => setContact(e.target.value)}/>
                </FormT.Content>
                <FormT.Content>
                    <FormT.Button type='submit'>Submit</FormT.Button>
                </FormT.Content>
            </FormT.Div>
        </form>
    )
}