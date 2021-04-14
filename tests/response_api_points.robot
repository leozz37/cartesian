*** Settings ***
Library         RequestsLibrary
Library         Collections
Library         Process
Library         JSONLibrary

*** Variable ***
${HOST}         http://localhost:8080      

*** Keywords ***
Setup
    Create Session      requestAuth         url=${HOST}

*** Test Case ***
Check Status Code /api/points
    Setup
    &{params}=          Create Dictionary       x=-94    y=89   distance=318
    ${response}=        GET On Session          requestAuth     /api/points     params=${params}
    ${msg_response}=    Get Value From Json     ${response.json()}      $.message.x
    ${coordiante}=      Evaluate                "".join(${msg_response})
    Should Contain      -94     ${coordiante}
