# Server Endpoints

*All endpoints for API calls are prefixed with `/api`*

## POST /register

Registers and creates a new user

**Parameters**
|Name|Description|Required|
|---|---|---|
|username|The name of the user to be created|Yes
|password|The password for the user to be created|Yes|

**Response**
|Name|Description|
|---|---|
|error_code|Code for error (if any)|
|error|Description of error (if any)|

## POST /login

Logs in and returns a new session token for an existing user

**Parameters**
|Name|Description|Required|
|---|---|---|
|username|The name of the user|Yes
|password|The password for the user|Yes|

**Response**
|Name|Description|
|---|---|
|token|JSON Web Token|
|error_code|Code for error (if any)|
|error|Description of error (if any)|
