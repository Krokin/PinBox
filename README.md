# PinBox

Web application for creating and storing temporary notes.
 
    POST /create - Creates a note with the specified title and text and enters it into the database.
    GET /pin/{id} - Returns a note with the specified id.
    / - Returns all notes from the database that have not expired.



### POST struct:
Json

    {
        "title":"text",
        "content":"text"
    }

## Response to the request:

### JSON in the format:

    {
        [
            {
                "id": "int",
                "title": "string"
                "content": "string",
                "created": "time.Time" 
                "expires": "time.Time"
            }
        ]
    }
