import styled from 'styled-components';

export default function Course(props){
    return (
        <>
            <CardHeader>
                <CardContent >
                    {props.coursename}
                </CardContent>
                <CardContent>    
                    {props.tutor}
                </CardContent>
                <CardContent>
                    {props.content}
                </CardContent>    
            </CardHeader>
        </>
    )
}



const CardHeader = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
`

const CardContent = styled.div`
    display: flex;
    font-size: 20px;
`