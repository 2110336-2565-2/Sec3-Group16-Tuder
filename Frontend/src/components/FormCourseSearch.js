import { useState , createRef} from 'react'
import styled from 'styled-components'
import {day} from '../datas/DayEnum'
import  {searchCourseHandler}  from '../handlers/searchCourseHandler';
import { useDataContext } from '../pages/Courses';
import { useQuery } from 'react-query';

export default function CourseSearchForm(props){
    const [searchTerm, setSearchTerm] = useState('');
    const [coursename, setCoursename] = useState('')
    const [subject, setSubject] = useState('')
    const [topic, setTopic] = useState('')
    const [tutor_name, setTutorName] = useState('')
    const [days, setDays] = useState([0,0,0,0,0,0,0])


    const {data, setData} = useDataContext();


    const {initData, isLoading, error, refetch} = useQuery(
        'filter_courses',() =>
        {
            const searchData = {
                searchTerm,
                coursename,
                subject,
                topic,
                tutor_name,
                days
            }
            searchCourseHandler(searchData).then((res) => {
                setData({data:res.data})
            })
        },
        {
            enabled: false,
            refetchOnWindowFocus: false,
        }
        );

    const submitHandler = () => {
        refetch();
    }


    if (isLoading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error: {error.message}</div>
    }
    
    

    return (    
    <FormData>
        <FormRow>
            <ItemGrid justify='center' columngrid='1 / 3'>
                <SearchBar>
                    <SearchInput key="searchBar" placeholder="Search by a course name" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)}/>
                </SearchBar>
            </ItemGrid>
            <ItemGrid columngrid='4'>
                <SearchButton onClick={submitHandler} >Search</SearchButton>
            </ItemGrid>  
        </FormRow>
        <FormRow>
            <ItemGrid justify='center' columngrid='1 / 3'>
                <FormInput key="topic" placeholder="Topic Name" value={topic} onChange={(e) => setTopic(e.target.value)}/>
            </ItemGrid>
        </FormRow>
        <FormRow>
        
            <ItemGrid justify='center' columngrid='1 / 3'>
                    <FormInput key="tutor" placeholder="Tutor Name" value={coursename} onChange={(e) => setCoursename(e.target.value)}/>
            </ItemGrid>
            <ItemGrid  columngrid='4'>
                    <FormInput key="subject" placeholder="Subject" value={subject} onChange={(e) => setSubject(e.target.value)}/>
            </ItemGrid>
        </FormRow>
        <FormRow>
            {day.map((idx, val) => (
                <CheckboxLable key={val}>
                    <Checkbox value={idx} type="checkbox" onChange={e => {
                        const newDays = [...days];
                        newDays[val] = e.target.checked ? 1 : 0;
                        setDays(newDays);
                        }
                    }/>
                    day : {idx[val+1]}
                </CheckboxLable>
            ))
            }
            
        </FormRow>

        </FormData>
        )
}


const FormRow = styled.div`
    display: flex;
    justify-content: center;
    gap: 20px;
`
const FormData = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
    gap: 20px;
`

const FormInput = styled.input`
    width: 250px;
    height: 30px;
    border-radius: 16px;
    border: 1px solid black;
    padding: 10px;
    font-size: 20px;
`

const CheckboxLable = styled.label`
    display: flex;
    justify-content: center;
    align-items: center;
    width: 150px;
    height: 30px;
    border-radius: 16px;
    border: 1px solid black;
    padding: 10px;
    font-size: 20px;
`  
const Checkbox = styled.input`
    width: 20px;
    height: 20px;
    border-radius: 16px;
    border: 1px solid black;
    padding: 10px;
    font-size: 20px;
`

const Option = styled.select`
    width: 250px;
    height: 50px;
    border-radius: 16px;
    border: 1px solid black;
    padding: 10px;
    font-size: 20px;
`



const ItemGrid = styled.div`
    grid-column: ${(props) => {
        return props.columngrid
    }};
    justify-self: ${(props) => {
        return props.justify
    }}
    
`

const SearchBar = styled.div`
    display: flex;
    justify-content: center;
`

const SearchInput = styled.input`
    width: 500px;
    height: 30px;
    border-radius: 16px;
    border: 1px solid black;
    padding: 10px;
    font-size: 20px;
`

const SearchButton = styled.button`
    width: 100px;
    height: 50px;
    border-radius: 16px;
    border: 1px solid black;
    background-color: #FFC300;
    padding: 10px;
    font-size: 20px;
`