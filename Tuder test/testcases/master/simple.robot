*** Settings ***
Library    SeleniumLibrary
Test Teardown    Close All Browsers
Resource  ../../testdata/test_data.robot
Resource    ../../keywords/ui/page/homepage.robot

*** Test Cases ***
Verify registration form with all input field
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify message after click submit form
    Close Browser

Verify registration form with empty input field at firstname
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    # Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify firstname display error message
    Close Browser

Verify registration form with empty input field at lastname
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    # Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify lastname display error message
    Close Browser

Verify registration form with empty input field at gender
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    # Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify gender display error message
    Close Browser

Verify registration form with empty input field at student id
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    # Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify student_id display error message
    Close Browser

Verify registration form with empty input field at mobile number
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    # Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify mobile display error message
    Close Browser

Verify registration form with empty input field at email
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    # Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify email display error message
    Close Browser

Verify registration form with empty input field at study time
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    # Click study time
    Click submit application
    Verify study time display error message
    Close Browser

Verify registration form with invalid input at student id
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    # Input and verify student id
    
    # input invalid student id
    Input and verify invalid student_id

    Input and verify mobile number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify study time display error message
    Close Browser

Verify registration form with invalid input field at mobile number
    Open browser student registration form

    Wait student registration load complete

    Verify firstname is empty

    Input and verify firstname
    Input and verify lastname
    Input and verify gender with male
    Input and verify student id
    # Input and verify mobile number
    Input and verify invalid mobile_number
    Input and verify email
    # Select and verify courses
    Click study time
    Click submit application
    Verify mobile display error message
    Close Browser
