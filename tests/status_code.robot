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
Check Status Code /points
    Setup
    ${response}=        GET On Session      requestAuth     /points
    Status Should Be    200                 ${response}

*** Test Case ***
Check Status Code /api/points
    &{params}=    Create Dictionary     x=-94    y=89   distance=318
    ${response}=        GET On Session      requestAuth     /points     params=${params}
    Status Should Be    200                 ${response}
