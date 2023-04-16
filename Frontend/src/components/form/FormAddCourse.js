import React, { useState } from "react";
import styled from "styled-components";
import { getTutorID } from "../../utils/jwtGet";
import { useNavigate, useParams, useLocation } from "react-router-dom";
import { toast } from "react-hot-toast";
import { addCourseHandler } from "../../handlers/addCourseHandler";
import FileUploader from "../global/FileUploader";

export default function FormAddcourse() {

    const navigate = useNavigate();

    const [formData, setFormData] = useState({
        title: "",
        subject: "",
        topic: "",
        estimated_time: "",
        description: "",
        price_per_hour: "",
        level: "",
        course_picture_url: "",
    });

    const [defaultImgUrl, setdefaultImgUrl] = useState("/images/Union.png");
    const [isFileUploaderOpen, setIsFileUploaderOpen] = useState(false);

    const handleChange = (event) => {
        setFormData({
        ...formData,
        [event.target.name]: event.target.value,
        });
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        const data = {
            tutor_id: getTutorID(),
            title: formData.title,
            subject: formData.subject,
            topic: formData.topic,
            estimate_time: parseInt(formData.estimated_time),
            description: formData.description,
            price_per_hour: parseInt(formData.price_per_hour),
            level: formData.level,
            course_picture: formData.course_picture_url.split(",")[1],
        };
        addCourseHandler(data)
            .then((res) => {
                if (res.data.success) {
                    toast.success("Create course successfully");
                    navigate("/courses");
                }else{
                    toast.error("Fail to create course");
                }
            })
            .catch((err) => {
                console.log(err);
                toast.error("Something went wrong");
            });
    };

    var title = formData.title;
    var topic = formData.topic;
    var subject = formData.subject;
    var estimated_time = formData.estimated_time;
    var description = formData.description;
    var price_per_hour = formData.price_per_hour;
    var level = formData.level;
    var course_picture_url = formData.course_picture_url;
    const showPic = ()=>{
        if (formData.course_picture_url){
            return (
                <RequestPicture type="normal"
                name="course_picture_url"
                src={formData.course_picture_url}
                onClick={() => setIsFileUploaderOpen(true)}
                />
            )
        }else{
            return(
                <RequestPicture type="default"
                name="course_picture_url"
                src={defaultImgUrl}
                onClick={() => setIsFileUploaderOpen(true)}
                />
            )
        }
    }
    return (
        <form onSubmit={handleSubmit}>
            <AddCourseRequest>
                <FileUploader
                    isOpen={isFileUploaderOpen}
                    setIsOpen={setIsFileUploaderOpen}
                    handleChange={handleChange}
                    name="course_picture_url"
                />
                <Title>New Course</Title>
                <TopicSection>
                    <Topic>Title</Topic>
                    <RequestTitle
                    type="text"
                    name="title"
                    value={title}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>Subject</Topic>
                    <RequestTitle
                    type="text"
                    name="subject"
                    value={subject}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>Topic</Topic>
                    <RequestTitle
                    type="text"
                    name="topic"
                    value={topic}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>estimated_time</Topic>
                    <RequestTitle
                    type="number"
                    name="estimated_time"
                    value={estimated_time}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>Description</Topic>
                    <RequestTitle
                    type="text"
                    name="description"
                    value={description}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>Price per hour</Topic>
                    <RequestTitle
                    type="number"
                    name="price_per_hour"
                    value={price_per_hour}
                    onChange={handleChange}
                    required
                    />
                </TopicSection>
                <TopicSection>
                    <Topic>Level</Topic>
                    <RequestSelect
                    type="text"
                    name="level"
                    value={level}
                    onChange={handleChange}
                    required
                    >
                        <RequestOption value="">Please Select Level</RequestOption>
                        <RequestOption value="Grade1">Grade1</RequestOption>
                        <RequestOption value="Grade2">Grade2</RequestOption>
                        <RequestOption value="Grade3">Grade3</RequestOption>
                        <RequestOption value="Grade4">Grade4</RequestOption>
                        <RequestOption value="Grade5">Grade5</RequestOption>
                        <RequestOption value="Grade6">Grade6</RequestOption>
                        <RequestOption value="Grade7">Grade7</RequestOption>
                        <RequestOption value="Grade8">Grade8</RequestOption>
                        <RequestOption value="Grade9">Grade9</RequestOption>
                        <RequestOption value="Grade10">Grade10</RequestOption>
                        <RequestOption value="Grade11">Grade11</RequestOption>
                        <RequestOption value="Grade12">Grade12</RequestOption>

                    </RequestSelect>
                </TopicSection>
                <PicSection>
                    <Topic>Picture</Topic>
                    {showPic()}
                </PicSection>
                <ButtonSection>
                    <CancelButton type="button" onClick={() => navigate("/courses")}>
                    Cancel
                    </CancelButton>
                    <SubmitButton type="submit">Submit</SubmitButton>
                </ButtonSection>
            </AddCourseRequest>
        </form>
    );
}

const AddCourseRequest = styled.div`
  box-sizing: border-box;
  min-width: 800px;
  max-width: 700px;
  margin-top: 50px;
  margin-bottom: 50px;
  margin-left: auto;
  margin-right: auto;
  padding: 50px;
  justify-content: center;

  background: #ffffff;
  border: 1px solid #858585;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
`;

const Title = styled.div`
  width: 100%;
  text-align: center;
  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 700;
  font-size: 30px;
  line-height: 44px;
  padding-bottom: 20px;
`;

const TopicSection = styled.div`
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: 76px;
  left: 220px;
  top: 311px;
`;

const PicSection = styled.div`
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  left: 220px;
  top: 311px;
`;

const Topic = styled.div`
  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
`;

const RequestTitle = styled.input`
  border: 2px solid #858585;
  border-radius: 5px;

  width: 100%;
  height: 100%;
  font-size: 20px;
  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
`;

const RequestPicture = styled.img`
  width: ${(props)=>props.type==="normal"?"500px":"40px"};
  height: ${(props)=>props.type==="normal"?"400px":"40px"};
`;

const ButtonSection = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  gap: 20px;
  width: 100%;
  margin-top: 20px;
`;

const CancelButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #ff7008;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ff7008;
  background-color: #ffffff;

  &:hover {
    background-color: #ff7008;
    color: #ffffff;
  }
`;

const SubmitButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #ff7008;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ffffff;
  background-color: #ff7008;

  &:hover {
    background-color: #ff9009;
    border: 2px solid #ff9008;
    color: #ffffff;
  }
`;

const RequestSelect = styled.select`
  border: 2px solid #858585;
  border-radius: 5px;

  width: 100%;
  height: 100%;
  font-size: 20px;
  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
`;
const RequestOption = styled.option`
    color: black;
    background: white;
`