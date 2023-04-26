*** Setting ***
Resource    ../../../keywords/ui/common/commonKeywords.robot
Resource    ../../../testdata/environment.robot
Resource    ../../../testdata/tuder_test_data.robot
Library    SeleniumLibrary
Library    DateTime

*** Keywords ***
Open browser sign in page
    Open Browser   ${WEB_URL}/sign-in    ${WEB_BROWSER}
    # Open Browser    ${WEB_URL}    ${WEB_BROWSER}
    Maximize Browser Window
    Sleep    2

Wait until sign in page is loaded
    # wait until a div with the text "Sign in" is visible
    Wait Until Element Is Visible    //button[.='Sign in']    30

Verify Username is empty
    # get the value of the input field with name "username"
    ${username}    Get Value    //input[@name='username']
    Should Be Equal    ${username}    ${EMPTY}

Verify Password is empty
    # get the value of the input field with name "password"
    ${password}    Get Value    //input[@name='password']
    Should Be Equal    ${password}    ${EMPTY}

# get date as an argument
Input and Verify Username sign-in 
    # input the username into the input field with name "username"
    # also concatenate the username with the date time with the format "ddMMyyyyHHmmss"
    [Arguments]    ${date}
    Input Text    //input[@name='username']   ${username}${date}
    # get the value of the input field with name "username"
    ${username}    Get Value    //input[@name='username']
    Should Be Equal    ${username}    ${username}${date}

Input and Verify Password sign-in
    # input the password into the input field with name "password"
    Input Text    //input[@name='password']    ${password}
    # get the value of the input field with name "password"
    ${password}    Get Value    //input[@name='password']
    Should Be Equal    ${password}    ${password}

Input and Verify Invalid Username
    # input the invalid username into the input field with name "username"
    Input Text    //input[@name='username']    ${username_not_exist}
    # get the value of the input field with name "username"
    ${username}    Get Value    //input[@name='username']
    Should Be Equal    ${username}    ${username_not_exist}

Input and Verify Invalid Password
    # input the invalid password into the input field with name "password"
    Input Text    //input[@name='password']    ${password_Incorrect}
    # get the value of the input field with name "password"
    ${password}    Get Value    //input[@name='password']
    Should Be Equal    ${password}    ${password_Incorrect}

Click Sign In Button
    # click the button with the text "Sign in"
    Click Button    //button[.='Sign in']
    Sleep    2

Verify Sign In Failed by Username
    # wait until a div with the div with role=status and contain "the username isn't match"
    Wait Until Element Contains    //div[@role='status']    the username isn't match    30

Verify Sign In Failed by Password
    # wait until a div with the div with role=status and contain "the password isn't match"
    Wait Until Element Contains    //div[@role='status']    the password isn't match    5

Verify Sign In Success
    # wait until a div contain the text "Welcome to Tudor" is visible
    # Wait Until Element Contains    //div[@role='status']    Singed in    5
    Wait Until Element Is Visible  //a[.="My Classes"]