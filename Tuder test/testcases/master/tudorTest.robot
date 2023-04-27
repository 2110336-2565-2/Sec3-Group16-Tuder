*** Settings ***
Library    SeleniumLibrary
Test Teardown    Close All Browsers
Resource  ../../testdata/test_data.robot
Resource    ../../keywords/ui/page/sign-in.robot
Resource    ../../keywords/ui/page/register.robot

*** Test Cases ***

Test Sign In
    Open browser Sign In page
    Wait until sign in page is loaded
    Input and Verify Username    ${EMPTY}    ${username_existed}
    Input and Verify Password sign-in
    Click Sign In Button
    Verify Sign In Success
    Close Browser



TC1-1 Register Fail Existed Username
    # ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    # print ${date} to console
    Open browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${EMPTY}    ${username_existed}
    Input and Verify Email
    Input and Verify Password
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    Verify Register Failed Username Used
    Close Browser

TC1-2 Register Failed password missmatch
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    # Input and Verify Password missmatch
    Input and Verify Password mismatch 
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    Verify Register Failed Password Missmatch
    Close Browser

TC1-3 Register Failed password too short
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password too short
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    Verify Register Failed Password Too Short
    Close Browser

TC1-4 Register Failed password too long
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password too long
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    Verify Register Failed Password Too Long
    Close Browser

TC1-5 Register Failed password contains invalid char
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password alphanumeric char
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Form is Not Submitted
    Close Browser

TC1-6 Register Failed password not contain lowercase
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password no lowercase
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Register Failed Password No Lowercase
    Close Browser

TC1-7 Register Failed password not contain uppercase
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password no uppercase
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Register Failed Password No Uppercase
    Close Browser

TC1-8 Register Failed password not contain number
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password no number
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Register Failed Password No Number
    Close Browser

TC1-9.1 Register Failed password is similar to email
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password similar to email
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Register Failed Password is Similar to Email
    Close Browser

TC1-9.2 Register Failed password is similar to firstname
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    Open Browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password similar to email
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    # Verify Register Failed Password Invalid Char
    Verify Register Failed Password is Similar to Email
    Close Browser



TC1-11 Register Success as Student
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    # print ${date} to console
    Open browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role student
    Click Register Button
    Verify Register Success
    Wait until sign in page is loaded
    Input and Verify Username   ${date}
    Input and Verify Password sign-in
    Click Sign In Button
    Verify Sign In Success
    Close Browser

TC1-12 Register Success as Tutor
    ${date}=    Get Current Date    result_format=%Y%m%d%H%M%S
    # print ${date} to console
    Sleep  1
    Open browser Register page
    Wait until register page is loaded
    Verify Empty for student
    Input and Verify Firstname
    Input and Verify Lastname
    Input and Verify Username   ${date}
    Input and Verify Email
    Input and Verify Password
    Input and Verify Address
    Input and Verify Contact Number
    Select and Verify Gender Male
    Input and Verify Birthdate
    Select and Verify Role tutor

    Verify Empty for tutor
    Input and Verify Citizen ID
    Input and Verify Description
    Input and Verify Card Holder Name
    Input and Verify Card Number
    Input and Verify Expiration Date
    Input and Verify CVV


    Click Register Button
    Verify Register Success
    Wait until sign in page is loaded
    Input and Verify Username   ${date}
    Input and Verify Password sign-in
    Click Sign In Button
    Verify Sign In Success
    Close Browser