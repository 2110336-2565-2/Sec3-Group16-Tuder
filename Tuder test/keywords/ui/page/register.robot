*** Setting ***
Resource    ../../../keywords/ui/common/commonKeywords.robot
Resource    ../../../testdata/environment.robot
Resource    ../../../testdata/tuder_test_data.robot
Library    SeleniumLibrary
Library    String

*** Keywords ***
Open browser Register page
    Open Browser   ${WEB_URL}/sign-up    ${WEB_BROWSER}
    Maximize Browser Window
    # Sleep    2

Wait until register page is loaded
    # wait until a div with the text "Sign in" is visible
    Wait Until Element Is Visible    //*[.='create account']    10

Verify Firstname is empty
    # get the value of the input field with name "username"
    ${tfirstname}    Get Value    //input[@name='firstname']
    Should Be Equal    ${tfirstname}    ${EMPTY}

Verify Lastname is empty
    # get the value of the input field with name "password"
    ${tlastname}    Get Value    //input[@name='lastname']
    Should Be Equal    ${tlastname}    ${EMPTY}

Verify Username is empty
    # get the value of the input field with name "username"
    ${tusername}    Get Value    //input[@name='username']
    Should Be Equal    ${tusername}    ${EMPTY}

Verify Email is empty
    # get the value of the input field with name "username"
    ${temail}    Get Value    //input[@name='email']
    Should Be Equal    ${temail}    ${EMPTY}

Verify Password is empty
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${EMPTY}

Verify Confirm Password is empty 
    # get the value of the input field with name "password"
    ${tconfirm_password}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirm_password}    ${EMPTY}

Verify Address is empty
    # get the value of the input field with name "password"
    ${taddress}    Get Value    //input[@name='address']
    Should Be Equal    ${taddress}    ${EMPTY}

Verify Contact Number is empty
    # get the value of the input field with name "password"
    ${tcontact_number}    Get Value    //input[@name='contactnumber']
    Should Be Equal    ${tcontact_number}    ${EMPTY}

Verify Gender is empty
    # get the value of the select field with name "gender"
    ${tgender}    Get Value    //select[@name='gender']
    Should Be Equal    ${tgender}    ${EMPTY}

Verify Birth Date is empty
    # get the value of the input field with name "password"
    ${tbirth_date}    Get Value    //input[@name='birthdate']
    Should Be Equal    ${tbirth_date}    ${EMPTY}

Verify Role is empty
    # get the value of the select field with name "role"
    ${trole}    Get Value    //select[@name='role']
    Should Be Equal    ${trole}    ${EMPTY}

Verify Empty for student
    Verify Firstname is empty
    Verify Lastname is empty
    Verify Email is empty
    Verify Password is empty
    Verify Confirm Password is empty
    Verify Address is empty
    Verify Contact Number is empty
    Verify Gender is empty
    Verify Birth Date is empty
    Verify Role is empty

Verify Empty for tutor
    Verify Citizen ID is empty
    Verify Description is empty
    Verify Card Holder Name is empty
    Verify Card Number is empty
    Verify Expire Date is empty
    Verify CVV is empty

Verify Citizen ID is not visible
    # an input field with name "password" should not be visible
    Element Should Not Be Visible    //input[@name='citizenid']

Verify Citizen ID is visible
    # an input field with name "password" should not be visible
    Element Should Be Visible    //input[@name='citizenid']

Verify Citizen ID is empty
    # get the value of the input field with name "password"
    ${tcitizen_id}    Get Value    //input[@name='citizen_id']
    Should Be Equal    ${tcitizen_id}    ${EMPTY}

Verify Description is empty
    # get the value of the input field with name "password"
    ${tdescription}    Get Value    //input[@name='description']
    Should Be Equal    ${tdescription}    ${EMPTY}

Verify Card Holder Name is empty
    # get the value of the input field with name "password"
    ${tcard_holder_name}    Get Value    //input[@name='card_holder name']
    Should Be Equal    ${tcard_holder_name}    ${EMPTY}

Verify Card Number is empty
    # get the value of the input field with name "password"
    ${tcard_number}    Get Value    //input[@name='card_number']
    Should Be Equal    ${tcard_number}    ${EMPTY}

Verify Expire Date is empty
    # get the value of the input field with name "password"
    ${texpire_date}    Get Value    //input[@name='expiry_date']
    Should Be Equal    ${texpire_date}    ${EMPTY}

Verify CVV is empty
    # get the value of the input field with name "password"
    ${tcvv}    Get Value    //input[@name='cvv']
    Should Be Equal    ${tcvv}    ${EMPTY}

Input and Verify Firstname
    # input the firstname into the input field with name "firstname"
    Input Text    //input[@name='firstname']    ${firstname}
    # get the value of the input field with name "firstname"
    ${tfirstname}    Get Value    //input[@name='firstname']
    Should Be Equal    ${tfirstname}    ${firstname}

Input and Verify Lastname
    # input the lastname into the input field with name "lastname"
    Input Text    //input[@name='lastname']    ${lastname}
    # get the value of the input field with name "lastname"
    ${tlastname}    Get Value    //input[@name='lastname']
    Should Be Equal    ${tlastname}    ${lastname}

Input and Verify Username
    # input the username into the input field with name "username"
    # this optional
    [Arguments]   ${date}=${EMPTY}    ${username}=${username}
    Input Text    //input[@name='username']    ${username}${date}
    # get the value of the input field with name "username"
    ${tusername}    Get Value    //input[@name='username']
    Should Be Equal    ${tusername}    ${username}${date}

Input and Verify Email
    # input the email into the input field with name "email"
    Input Text    //input[@name='email']    ${email}
    # get the value of the input field with name "email"
    ${temail}    Get Value    //input[@name='email']
    Should Be Equal    ${temail}    ${email}

# Input and Verify Email Invalid
#     # input the email into the input field with name "email"
#     Input Text    //input[@name='email']    ${invalid_email}
#     # get the value of the input field with name "email"
#     ${temail}    Get Value    //input[@name='email']
#     Should Be Equal    ${temail}    ${invalid_email}

# Input and Verify Email no at
#     # input the email into the input field with name "email"
#     Input Text    //input[@name='email']    ${invalid_email_no_at}
#     # get the value of the input field with name "email"
#     ${temail}    Get Value    //input[@name='email']
#     Should Be Equal    ${temail}    ${invalid_email_no_at}

# Input and Verify Email no dot
#     # input the email into the input field with name "email"
#     Input Text    //input[@name='email']    ${invalid_email_no_dot}
#     # get the value of the input field with name "email"
#     ${temail}    Get Value    //input[@name='email']
#     Should Be Equal    ${temail}    ${invalid_email_no_dot}

Input and Verify Password
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password}

Input and Verify Password mismatch
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password}missmatch
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password}missmatch

Input and Verify Password alphanumeric char
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_alphanumeric_char}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_alphanumeric_char}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_alphanumeric_char}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_alphanumeric_char}

Input and Verify Password no uppercase
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_no_uppercase}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_no_uppercase}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_no_uppercase}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_no_uppercase}

Input and Verify Password no lowercase
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_no_lowercase}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_no_lowercase}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_no_lowercase}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_no_lowercase}


Input and Verify Password no number
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_no_number}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_no_number}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_no_number}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_no_number}


Input and Verify Password similar to username
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_similar_to_username}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_similar_to_username}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_similar_to_username}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_similar_to_username}
    

Input and Verify Password similar to email
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_similar_to_email}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_similar_to_email}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_similar_to_email}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_similar_to_email}


Input and Verify Password similar to firstname
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_similar_to_firstname}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_similar_to_firstname}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_similar_to_firstname}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_similar_to_firstname}


Input and Verify Password similar to lastname
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_similar_to_lastname}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${password_similar_to_lastname}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${password_similar_to_lastname}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${password_similar_to_lastname}


Input and Verify Password too short
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${invalid_password_too_short}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${invalid_password_too_short}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${invalid_password_too_short}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${invalid_password_too_short}


Input and Verify Password too long
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${invalid_password_too_long}
    # get the value of the input field with name "password"
    ${tpassword}    Get Value    //input[@name='password']
    Should Be Equal    ${tpassword}    ${invalid_password_too_long}
    # also input the confirm password into the input field with name "confirmpassword"
    Input Text    //input[@name='confirmpassword']    ${invalid_password_too_long}
    # get the value of the input field with name "confirmpassword"
    ${tconfirmpassword}    Get Value    //input[@name='confirmpassword']
    Should Be Equal    ${tconfirmpassword}    ${invalid_password_too_long}


Input and Verify Address
    # input the address into the input field with name "address"
    Input Text    //input[@name='address']    ${address}
    # get the value of the input field with name "address"
    ${taddress}    Get Value    //input[@name='address']
    Should Be Equal    ${taddress}    ${address}

Input and Verify Contact Number
    # input the contact number into the input field with name "contact_number"
    Input Text    //input[@name='contactnumber']    ${contactnumber}
    # get the value of the input field with name "contact_number"
    ${tcontact_number}    Get Value    //input[@name='contactnumber']
    Should Be Equal    ${tcontact_number}    ${contactnumber}

Select and Verify Gender Male
    # select the gender as male on the select box with name "gender"
    Click Element    //option[@value='male']
    # get the value of the select field with name "gender"
    ${tgender}    Get Value    //select[@name='gender']
    Should Be Equal    ${tgender}    ${gender_male}

Select and Verify Gender Female
    # select the gender as male on the select box with name "gender"
    Click Element    //option[@value='female']
    # get the value of the select field with name "gender"
    ${tgender}    Get Value    //select[@name='gender']
    Should Be Equal    ${tgender}    ${gender_female}

Input and Verify Birthdate
    # input the birthdate into the input field with name "birthdate"
    Input Text    //input[@name='birthdate']    ${birthdate}
    # get the value of the input field with name "birthdate"
    ${tbirthdate}    Get Value    //input[@name='birthdate']
    # the birthdate is in the format dd-mm-yyyy
    Should Be Equal    ${tbirthdate}    ${expected_birthdate}

Select and Verify Role student
    # select the role as student on the select box with name "role"
    Click Element    //option[@value='student']
    # get the value of the select field with name "role"
    ${trole}    Get Value    //select[@name='role']
    Should Be Equal    ${trole}    ${role_student}

Select and Verify Role tutor
    # select the role as tutor on the select box with name "role"
    Click Element    //option[@value='tutor']
    # get the value of the select field with name "role"
    ${trole}    Get Value    //select[@name='role']
    Should Be Equal    ${trole}    ${role_tutor}

Input and Verify Citizen ID
    # input the citizen id into the input field with name "citizen_id"
    ${citizen_id}   Generate Random String    13    [NUMBERS]
    Input Text    //input[@name='citizen_id']    ${citizen_id}
    # get the value of the input field with name "citizen_id"
    ${tcitizen_id}    Get Value    //input[@name='citizen_id']
    Should Be Equal    ${tcitizen_id}    ${citizen_id}

Input and Verify Description
    # input the description into the input field with name "description"
    Input Text    //input[@name='description']    ${description}
    # get the value of the input field with name "description"
    ${tdescription}    Get Value    //input[@name='description']
    Should Be Equal    ${tdescription}    ${description}

Input and Verify Card Holder Name
    # input the card holder name into the input field with name "card_holder_name"
    Input Text    //input[@name='card_holder name']    ${card_holder_name}
    # get the value of the input field with name "card_holder_name"
    ${tcard_holder_name}    Get Value    //input[@name='card_holder name']
    Should Be Equal    ${tcard_holder_name}    ${card_holder_name}

Input and Verify Card Number
    # input the card number into the input field with name "card_number"
    Input Text    //input[@name='card_number']    ${card_number}
    # get the value of the input field with name "card_number"
    ${tcard_number}    Get Value    //input[@name='card_number']
    Should Be Equal    ${tcard_number}    ${card_number}

Input and Verify Expiration Date
    # input the expiration date into the input field with name "expiration_date"
    Input Text    //input[@name='expiry_date']    ${card_expiry}
    # get the value of the input field with name "expiration_date"
    ${texpiration_date}    Get Value    //input[@name='expiry_date']
    Should Be Equal    ${texpiration_date}    ${card_expiry}

Input and Verify CVV
    # input the cvv into the input field with name "cvv"
    Input Text    //input[@name='cvv']    ${card_cvv}
    # get the value of the input field with name "cvv"
    ${tcvv}    Get Value    //input[@name='cvv']
    Should Be Equal    ${tcvv}    ${card_cvv}

Verify Register Failed Password Missmatch
    # wait until a div with the div with role=status and contain "the password isn't match"
    Wait Until Element Contains    //div[@role='status']    ${error_password_mismatch}    5

Verify Register Failed Password Too Short
    # wait until a div with the div with role=status and contain "the password is too short"
    Wait Until Element Contains    //div[@role='status']    ${error_password_too_short}    5

Verify Register Failed Password Too Long
    # wait until a div with the div with role=status and contain "the password is too long"
    Wait Until Element Contains    //div[@role='status']    ${error_password_too_long}    5

Verify Register Failed Password Invalid Char
    # wait until a div with the div with role=status and contain "the password contains invalid character"
    Wait Until Element Contains    //div[@role='status']    ${error_password_invalid_char}    5

Verify Register Failed Password No Lowercase
    # wait until a div with the div with role=status and contain "the password doesn't contain lowercase"
    Wait Until Element Contains    //div[@role='status']    ${error_password_no_lowercase}    5

Verify Register Failed Password No Uppercase
    # wait until a div with the div with role=status and contain "the password doesn't contain uppercase"
    Wait Until Element Contains    //div[@role='status']    ${error_password_no_uppercase}    5

Verify Register Failed Password No Number
    # wait until a div with the div with role=status and contain "the password doesn't contain number"
    Wait Until Element Contains    //div[@role='status']    ${error_password_no_number}    5

Verify Register Failed Password is Similar to Username
    # wait until a div with the div with role=status and contain "the password is similar to username"
    Wait Until Element Contains    //div[@role='status']    ${error_password_is_similar_to_username}    5

Verify Register Failed Password is Similar to Email
    # wait until a div with the div with role=status and contain "the password is similar to email"
    Wait Until Element Contains    //div[@role='status']    ${error_password_is_similar_to_email}    5

Verify Register Failed Password is Similar to First Name
    # wait until a div with the div with role=status and contain "the password is similar to first name"
    Wait Until Element Contains    //div[@role='status']    ${error_password_is_similar_to_firstname}    5

Verify Register Failed Password is Similar to Last Name
    # wait until a div with the div with role=status and contain "the password is similar to last name"
    Wait Until Element Contains    //div[@role='status']    ${error_password_is_similar_to_lastname}    5

Verify Register Failed Username Used
    # wait until a div with the div with role=status and contain "the username is used"
    Wait Until Element Contains    //div[@role='status']    ${error_username_used}    5

Verify Register Success
    # wait until a div with the div with role=status and contain "register success"
    Wait Until Element Contains    //div[@role='status']    ${success_register}    5

Click Register Button
    # click the button with the text "register"
    Click Button    //button[.='Sign Up']

Verify Form is Not Submitted
    # the browser should display label on a field
    Sleep    5
    Page Should Contain Element    //div[.="create account"]
    