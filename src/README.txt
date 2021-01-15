Starting the app:
- IMPORTANT: Be sure that port 5000 is open, and available on your computer.
- To start the app, simply type go run main.go in the terminal.
- The app will be running on localhost:5000

Usage:
- The app only supports two request methods. GET, and POST.
- Doing a GET on localhost:5000 will return the current batch, and the batch history if at least one batch has been dispatched.
- In order to post transactions to a batch, you will need to do a POST to localhost:5000 with the following JSON body:
{
    "amount": <some integer value without angle brackets>
}

Notes:
- I was torn between returning the current batch and batch history in JSON format, and just returning a "text" version of the history. 
No format was specified, so I just returned it something human readable for the sake of this test. If this was API/service were to be consumed by something, 
I would have returned it in JSON.
- Outside of the unit tests, I did my own testing for concurrency using a shell script that spammed curl commands. This script is present in the folder.
- There are log statements in the code for debugging purposes that have been commented out.