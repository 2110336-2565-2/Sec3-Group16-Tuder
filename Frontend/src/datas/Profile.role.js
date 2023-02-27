export const tutorFields = [
  {
    label: "First Name",
    type: "text",
    id: "first_name",
    name: "first_name",
    width: "",
  },
  {
    label: "Last Name",
    type: "text",
    id: "last_name",
    name: "last_name",
    width: "",
  },
  { 
    label: "Email", 
    type: "email", 
    id: "email", 
    name: "email", 
    width: "100%" 
  },
  {
    label: "Address",
    type: "text",
    id: "address",
    name: "address",
    width: "100%",
  },
  {
    label: "Contact Number",
    type: "text",
    id: "phone",
    name: "phone",
    width: "100%",
  },
  {
    label: "Gender",
    type: "select",
    id: "gender",
    name: "gender",
    isRequired: true,
    options: [
      { value: "", label: "Please select", disabled: true, hidden: true },
      { value: "male", label: "male" },
      { value: "female", label: "female" },
    ],
    width: "",
  },
  {
    label: "Birth Date",
    type: "date",
    id: "birth_date",
    name: "birth_date",
    width: "",
  },
  {
    label: "Description",
    type: "textArea",
    id: "description",
    name: "description",
    width: "100%",
  },
  {
    label: "Available Time",
    type: "time",
    id: "schedule",
    name: "schedule",
    width: "100%",
  }
];

export const dummyStudent = {
  role: "student",
  profile_picture_URL:
    "https://i.pinimg.com/originals/68/12/46/681246b04458bbb03a7120b68c0bad01.jpg",
  first_name: "Mehrab",
  last_name: "Jaidee",
  email: "mehrab@gmail.com",
  address: "123/456 Bangkoontien Nontaburi 11110",
  phone: "0812345678",
  gender: "female",
  birth_date: "2010-10-10",
  school: "มหาวิทยาลัยนนทบุรี",
};

export const dummyTutor = {
  role: "tutor",
  profile_picture_URL:
    "https://i.pinimg.com/originals/68/12/46/681246b04458bbb03a7120b68c0bad01.jpg",
  first_name: "Mehrab",
  last_name: "Jaidee",
  email: "mehrab@gmail.com",
  address: "123/456 Bangkoontien Nontaburi 11110",
  phone: "0812345678",
  gender: "female",
  birth_date: "2010-10-10",
  citizen_id: "1234567890123",
  description:
    "Experienced tutor with a passion for education and helping students succeed. With a Master's degree in Education and 5 years of teaching experience, I have a proven track record of improving student performance and confidence in subjects such as Math, English, and Science. I customize my lessons to meet each student's individual needs and learning style, creating a supportive and engaging environment for maximum success. Outside of the classroom, I enjoy reading, hiking, and trying new recipes. Let's work together to reach your educational goals!",
  bankName: "Kasikorn",
  accountNumber: "1234567890",
  accountName: "Mehrab Jaidee",
  schedule: [
    {
      day: "Sunday",
      timeSlot: [
        {
          from: "10:00",
          to: "12:00",
        },
        {
          from: "14:00",
          to: "15:00",
        },
      ],
    },
    {
      day: "Monday",
      timeSlot: [
        {
          from: "10:00",
          to: "12:00",
        },
        {
          from: "14:00",
          to: "15:00",
        },
      ],
    },
    {
      day: "Tuesday",
      timeSlot: [
        {
          from: "10:00",
          to: "12:00",
        },
      ],
    },
  ],
};