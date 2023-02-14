# IAF Golden-Route Project

## Task 1
<p>
    <b>Question 5: </b> We can use testing modules such as the "testing" package for golang or "JUnit" for java to run our functions against known input and check if the expected output is actually generated.<br>
    Some edge cases that might happen is mistyped input, or an invalid input value. We can regex the input we get from the frontend to validate its type (which I have done). We can also run value checks on the input to deteremine its validity, like checking that the CargoMass value is not negative. 
</p>
<p>
<b>Question 6: </b> We can calculate the wind resistence.
 Moreover, we can implement complex mathmatical models that will enable us to calculate non constant accelaration, instead of constant one. Thus, making the model's conditions closer to reality.
</p>

## Task 4
<p>
<b>Question 4: </b> We might want to show the client if there's fog, which could cause blurry vision. 
</p>

## Task 5
### Overview
```mermaid
sequenceDiagram
    participant Frontend
    participant Backend
    participant RemoteWeatherAPI
    participant Database

    Frontend ->> Backend:  Post requests data, passing required fields for calculation in req body.

    Backend ->> Frontend: Evaluating Frontend's requests, calculating/requesting needed data, and sends it back.

    Backend ->> Database: Saves the Flight Records coming from Frontend via the Flight struct from Backend.

    Backend ->> RemoteWeatherAPI: requests weather data

    RemoteWeatherAPI ->> Backend: validate & parsing data from response, sends it back to client

```

### Weather Api Implementation UML
```mermaid
sequenceDiagram
    participant Frontend
    participant Backend
    participant RemoteWeatherAPI
    
    Frontend ->> Backend: Posts request with  needed fields
    Backend ->> RemoteWeatherAPI: Requests the weather data using the client's fields
    RemoteWeatherAPI ->> Backend: returns data in json format
    Backend ->> Frontend: returns selected fields from json

```

# Task 6
### Deployment
#### Step 1
Clone the project to your machine and cd into it: 
```bash
git clone "https://github.com/ido-nasi/Golden-Route-API.git"
cd Golden-Route-API
```

#### Step 2
Make sure you have the `docker-engine` and `docker-compose` installed<b>

Official installation guides:</b>
- docker engine: [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install)
- docker-compose: [https://docs.docker.com/compose/install/other/](https://docs.docker.com/compose/install/other/)

#### Step 3
After installing, from the root of the project, run:
```bash
docker-compose build
docker-compose up
```

The client will be running on `localhost:3000`