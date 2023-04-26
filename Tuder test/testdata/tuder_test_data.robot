*** Variables ***

${firstname}    john
${lastname}     doe

# username is concatenated by johndoe and date time in format dd-mm-yyyy-hh-mm-ss
${username}     johndoe11
${username_not_exist}   johndoeNotExist
${username_existed}     student1
${username_tutor}  johndoeTutor



${email}        iamnumbeR1@somemail.com
${password}     P@ssw0rd
${password_Incorrect}     P@ssw0rd1
${password_alphanumeric_char}    P@ssw0rd1ß∂≈¥
${password_no_uppercase}     p@ssw0rd
${password_no_lowercase}     P@SSW0RD
${password_no_number}        P@ssword
${password_similar_to_username}    johndoe11A
${password_similar_to_email}       iamnumbeR1.com
${password_similar_to_firstname}   johnjohn1A
${password_similar_to_lastname}    Misterdoe1
${invalid_password_too_short}   P@ssw0r
${invalid_password_too_long}    p@ssw0rd901234567

# ${invalid_email}    invalidemail
# ${invalid_email_no_at}    invalidemail.com
# ${invalid_email_no_dot}   invalidemail@com

${address}    1234 Main St
${contact_number}    1234567890
${gender_male}    male
${gender_female}    female
${birthdate}    11-03-2002
${expected_birthdate}    2002-03-11
${role_student}    student
${role_tutor}    tutor
# for tutor only
# ${citizen_id}    1234567890123
${description}    I am a tutor
${card_holder_name}    John Doe
${card_number}    4162 0260 0572 2393
${card_expiry}    12 / 2025
${card_cvv}    123

# error messages
${error_username_used}    the username has already been used
${error_password_mismatch}     the password isn't match
${error_password_too_short}    Password length must be not lower that 8 chars
${error_password_too_long}     Password length must be not greater that 16 chars

# the password must contains only lowercase+uppercase+number+nonalphanumeric
${error_password_invalid_char}    Password must contains only the password must contains only abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890(~!@#$%^&*_-+=`|\\(){}[]:;\"'<>,.?/)
${error_password_no_lowercase}    Password must contains at least 1 chars from abcdefghijklmnopqrstuvwxyz
${error_password_no_uppercase}    Password must contains at least 1 chars from ABCDEFGHIJKLMNOPQRSTUVWXYZ
${error_password_no_number}    Password must contains at least 1 chars from 1234567890
# prompt The password is too similar to the %s
${error_password_is_similar_to_username}    The password is too similar to the ${username}
${error_password_is_similar_to_email}    The password is too similar to the ${email}
${error_password_is_similar_to_firstname}   The password is too similar to the ${firstname}
${error_password_is_similar_to_lastname}    The password is too similar to the ${lastname}

# success messages
${success_register}    Sign up successfully
