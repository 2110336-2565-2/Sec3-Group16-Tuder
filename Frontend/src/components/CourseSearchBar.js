import styled from "styled-components"
import {useState} from 'react'

export default function CourseSearchBar(props){
    const [searchTerm, setSearchTerm] = useState('')

    return (   
        <SearchSection>
            <ItemGrid justify='center' columngrid='1 / 3'>
                <SearchBar>
                    <SearchInput placeholder="Search for a course" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)}/>
                </SearchBar>
            </ItemGrid>
            <ItemGrid columngrid='4'>
                <SearchButton onClick={() => props.search(searchTerm)}>Search</SearchButton>
            </ItemGrid>
        </SearchSection>    
    )
}

const SearchSection = styled.div`
    display: flex;
    justify-content: center;
`

const ItemGrid = styled.div`
    grid-column: ${(props) => {
        return props.columngrid
    }};
    justify-self: ${(props) => {
        return props.justify
    }}
`;


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