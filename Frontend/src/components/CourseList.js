import styled from 'styled-components';
import Course from '../components/Course';
import {useQuery} from 'react-query';
import { fetchCourseHandler } from '../handlers/fetchCourseHandler';
import { useDataContext } from '../pages/Courses';

// import { useEffect , useState} from 'react';



export default function CourseList(){ 
    const {data, setData} = useDataContext();

    // useEffect(() => {
    //     // update component
    //     // console.log(data)
    // }, [data])

    const {initData, isLoading, error} = useQuery(
        'courses',() =>
        {
            fetchCourseHandler().then((res) => {
                setData({data:res.data});
            })
        },
        {
            refetchOnWindowFocus: false,
        }
        );


    if (isLoading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error: {error.message}</div>
    }
    

   

    return (
        <>
            <CourseListPage>
                { data.data.map(item => (  
                    <CourseListcontent key={item.course_id}>
                        <Course  coursename={item.title} topic={item.topic} tutor={item.tutor_name} subject={item.subject} time={item.estimate_time} price={item.price_per_hour} img={item.course_picture_url}/>       
                    </CourseListcontent>
                ))}
            </CourseListPage>
        </>
    )
}

const CourseListPage = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: center;
`

const CourseListcontent = styled.div`
    display: flex;
    border: 1px solid black;
    border-radius: 16px;
    padding: 10px;
`