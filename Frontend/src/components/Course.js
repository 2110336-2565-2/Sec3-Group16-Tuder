import styled from 'styled-components';

export default function Course(props){
    return (
        <>
            <CardHeader>
                <CardImg src={props.img}/>
                <CardContent >
                    {props.coursename}
                </CardContent>
                <CardContent>    
                    {props.tutor}
                </CardContent>
                <CardContent>
                    {props.topic}
                </CardContent>  
                <CardContent>
                    {props.subject}
                </CardContent>  
                <CardContent>
                    Time: {props.time} hr
                </CardContent>  
                <CardContent>
                    Price/hr: {props.price}
                </CardContent>    
            </CardHeader>
        </>
    )
}

const CardImg = styled.img`
    width: 500px;
    height: auto;
    object-fit: cover;
    border-radius: 10px;
`

const CardHeader = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
`

const CardContent = styled.div`
    display: flex;
    font-size: 20px;
`