# User Event Aggregator

A simple application for user event aggregation. This appliation will allow you to add events (like, comments, post) against the user and get the total number for of events for each user

## Table of Contents

- [Features](#features)
- [How To Run The App ? ](#HowToRunTheApp?)
- [Important Points To Remember](#ImportantPoitsToRemember)

 ## Features

Daily Summary Reports: This APP provides daily summary reports that break down user activity, including the number of posts made , the total number of likes and comments received. 

Real-Time Updates: As new events are added to the dataset, the application updates the summary reports in real-time to reflect the latest user activity.

## How To Run The App ?
  - ### Run Locally :
      - Run the following command in the root of the project to build the app `go build main.go` and once the build is successful , run `./main -i input.json -o output.json`
      - For update action you can run the same command i.e `./main -i input.json -o output.json` with the udated data. The output.json file will be updated with both latest and earlier summary

## Important Points To Note
  - Update the input data in the file `input.json` present in the root of the directory. 
  - You dont need to worry about the `output.json` , if app wont be able to find output.json file in the directory of project then it will create one by its own.
  - This application using inmemory for dataset , so once the execution will be ended all the data inside the in-memory will be lost.
  - After every cli command i.e `./main -i input.json -o output.json` the app will terminate means the in-memory data will be lost, but it doesnt affect the `output.json file` .
    `output.json` will contains all output. 

