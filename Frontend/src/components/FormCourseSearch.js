import { useState , createRef} from 'react'
import styled from 'styled-components'
import {day} from '../datas/DayEnum'



export default function CourseSearchForm(props){
    const [searchTerm, setSearchTerm] = useState('');
    const [coursename, setCoursename] = useState('')
    const [subject, setSubject] = useState('I love hee')


    return (    
    <FormData>
        <FormRow>
            <ItemGrid justify='center' columngrid='1 / 3'>
                <SearchBar>
                    <SearchInput placeholder="Search by a course name" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)}/>
                </SearchBar>
            </ItemGrid>
            <ItemGrid columngrid='4'>
                <SearchButton onClick={() => props.search(searchTerm)}>Search</SearchButton>
            </ItemGrid>  
        </FormRow>
        <FormRow>
            <ItemGrid justify='center' columngrid='1 / 3'>
                <FormInput placeholder="Tutor Name" value={coursename} onChange={(e) => setCoursename(e.target.value)}/>
            </ItemGrid>
                <ItemGrid justify='center' columngrid='7'>
                    <Option value={subject} onChange={(e) => setSubject(e.target.value)}>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                    </Option>
               
            </ItemGrid>
        </FormRow>
        <FormRow>
            {day.map((idx, val) => (
                <CheckboxLable>
                    <Checkbox value={idx} type="checkbox"/>
                    day : {val}
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