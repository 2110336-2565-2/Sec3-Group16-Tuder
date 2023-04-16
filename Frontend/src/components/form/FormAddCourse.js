import React, { useState } from "react";
import styled from "styled-components";
import { getTutorID } from "../../utils/jwtGet";
import { useNavigate, useParams, useLocation } from "react-router-dom";
import { toast } from "react-hot-toast";
import { addCourseHandler } from "../../handlers/addCourseHandler";
import FileUploader from "../global/FileUploader";

export default function FormAddcourse() {

    const navigate = useNavigate();
    const location = useLocation();

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
        console.log(formData);
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        const data = {
            tutor_id: getTutorID(),
            title: formData.title,
            subject: formData.subject,
            topic: formData.topic,
            estimated_time: parseInt(formData.estimated_time),
            description: formData.description,
            price_per_hour: parseInt(formData.price_per_hour),
            level: formData.level,
            course_picture: formData.course_picture_url.split(",")[1],
        };
        console.log(data)
        addCourseHandler(data)
            .then((res) => {
                if (res.data.success) {
                    toast.success("Create course successfully");
                    navigate("/");
                }
            })
            .catch((err) => {
                console.log(err);
                toast.error(err.response.data.message);
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

    return (
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
                />
            </TopicSection>
            <TopicSection>
                <Topic>Subject</Topic>
                <RequestTitle
                type="text"
                name="subject"
                value={subject}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>Topic</Topic>
                <RequestTitle
                type="text"
                name="topic"
                value={topic}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>estimated_time</Topic>
                <RequestTitle
                type="number"
                name="estimated_time"
                value={estimated_time}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>Description</Topic>
                <RequestTitle
                type="text"
                name="description"
                value={description}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>Price per hour</Topic>
                <RequestTitle
                type="number"
                name="price_per_hour"
                value={price_per_hour}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>Level</Topic>
                <RequestTitle
                type="text"
                name="level"
                value={level}
                onChange={handleChange}
                />
            </TopicSection>
            <TopicSection>
                <Topic>Picture</Topic>
                
                <RequestPicture
                name="course_picture_url"
                src={formData.course_picture_url ? formData.course_picture_url : defaultImgUrl}
                onClick={() => setIsFileUploaderOpen(true)}
                />
            </TopicSection>
            <ButtonSection>
                <CancelButton onClick={() => navigate("/")}>
                Cancel
                </CancelButton>
                <SubmitButton onClick={handleSubmit}>Submit</SubmitButton>
            </ButtonSection>
        </AddCourseRequest>
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
  width: 40px;
  height: 40px;
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
