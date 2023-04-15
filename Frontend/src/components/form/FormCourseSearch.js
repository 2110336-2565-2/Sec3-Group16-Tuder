import React, { useState, createRef } from "react";
import styled from "styled-components";
import { day } from "../../datas/DayEnum";
import { searchCourseHandler } from "../../handlers/searchCourseHandler";
import { useDataContext } from "../../pages/Courses";
import { useQuery } from "react-query";
import { toast } from "react-hot-toast";
import { WindowsFilled } from "@ant-design/icons";

export default function CourseSearchForm(props) {
  const [coursename, setCoursename] = useState("");
  const [subject, setSubject] = useState("");
  const [topic, setTopic] = useState("");
  const [tutor_name, setTutorName] = useState("");
  const [days, setDays] = useState([
    false,
    false,
    false,
    false,
    false,
    false,
    false,
  ]);

  const { data, setData } = useDataContext();

  const { initData, isLoading, error, refetch } = useQuery(
    "filter_courses",
    () => {
      const searchData = {
        title: coursename,
        subject: subject,
        topic: topic,
        tutor_name: tutor_name,
        days: days,
      };
      searchCourseHandler(searchData).then((res) => {
        setData(res.data.result);
      });
    },
    {
      enabled: false,
      refetchOnWindowFocus: false,
    }
  );

  const submitHandler = (e) => {
    e.preventDefault();
      refetch();
    window.scrollTo(0, document.getElementById("course").offsetTop-100);
  };

  if (isLoading) {
    toast.promise(initData, {
      loading: "Loading",
      success: "Success",
      error: "Error",
    });
    return <div>Loading...</div>;
  }

  if (error) {
    toast.error("Error");
    return <div>Error: {error.message}</div>;
  }

  return (
    <FormData action="#">
      <FormRow gap="20px">
        <ItemGrid justify="center" columngrid="1 / 3">
          <SearchBar>
            <SearchInput
              boxsize="350px"
              key="searchBar"
              placeholder="Search by a Course Name"
              value={coursename}
              onChange={(e) => setCoursename(e.target.value)}
            />
          </SearchBar>
        </ItemGrid>
        <ItemGrid justify="center" columngrid="1 / 3">
          <SearchInput
            boxsize="350px"
            key="topic"
            placeholder="Topic Name"
            value={topic}
            onChange={(e) => setTopic(e.target.value)}
          />
        </ItemGrid>
      </FormRow>
      <FormRow gap="20px">
        <ItemGrid justify="center" columngrid="1 / 3">
          <SearchInput
            boxsize="350px"
            key="tutor"
            placeholder="Tutor Name"
            value={tutor_name}
            onChange={(e) => setTutorName(e.target.value)}
          />
        </ItemGrid>
        <ItemGrid columngrid="4">
          <SearchInput
            boxsize="350px"
            key="subject"
            placeholder="Subject Name"
            value={subject}
            onChange={(e) => setSubject(e.target.value)}
          />
        </ItemGrid>
      </FormRow>
      <FormRow gap="0px">
        {day.map((idx, val) => (
          <CheckboxLable key={val}>
            <Checkbox
              value={idx}
              type="checkbox"
              onChange={(e) => {
                const newDays = [...days];
                newDays[val] = e.target.checked ? true : false;
                setDays(newDays);
              }}
            />
            {idx[val + 1]}
          </CheckboxLable>
        ))}
      </FormRow>
      <FormRow>
        <ItemGrid columngrid="4">
          <SearchButton type="submit" onClick={submitHandler}>
            Search
          </SearchButton>
        </ItemGrid>
      </FormRow>
    </FormData>
  );
}

const FormRow = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: ${(props) => {
    return props.gap;
  }};
`;
const FormData = styled.form`
  display: flex;
  justify-content: center;
  flex-direction: column;
  gap: 20px;
`;

const FormInput = styled.input`
  width: 250px;
  height: 30px;
  border-radius: 16px;
  border: 1px solid black;
  padding: 10px;
  font-size: 20px;
`;

const CheckboxLable = styled.label`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 150px;
  height: 30px;
  padding: 0px;
  font-size: 16px;
  gap: 5px;
`;
const Checkbox = styled.input`
  width: 20px;
  height: 20px;
  border-radius: 16px;
  border: 1px solid black;
  padding: 0px;
  font-size: 20px;
  accent-color: #eb7b42;
`;

const ItemGrid = styled.div`
  grid-column: ${(props) => {
    return props.columngrid;
  }};
  justify-self: ${(props) => {
    return props.justify;
  }};
`;

const SearchBar = styled.div`
  display: flex;
  justify-content: center;
`;

const SearchInput = styled.input`
  width: ${(props) => {
    return props.boxsize;
  }};
  height: 30px;
  border: 1px solid #bebebe;
  border-radius: 5px;
  padding: 10px;
  font-size: 16px;
`;

const SearchButton = styled.button`
  width: 100px;
  padding: 10px 10px;
  border: 0px solid;
  border-radius: 6px;
  color: white;
  font-size: 14px;
  background-color: #eb7b42;
  cursor: pointer;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
    rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
  &:hover {
    background-color: rgb(240, 123, 36);
  }
`;
