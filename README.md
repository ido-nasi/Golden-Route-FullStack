# IAF Golden-Route Project

## Task 1
<p>
    <b>Question 5: </b> We can use testing modules such as the "testing" package for golang or "JUnit" for java to run our functions against known input and check if the expected output is actually generated.<br>
    Some edge cases that might happen is mistyped input, or an invalid input value. We can regex the input we get from the frontend to validate its type (which I have done). We can also run value checks on the input to deteremine its validity, like checking that the CargoMass value is not negative. 
</p>
<p>
<b>Question 6: </b> We should calculate the wind resistence and direction, as it could change our take off distance.
 Moreover, we can implement complex mathmatical models that will enable us to calculate non constant accelaration, instead of constant one. Thus, making the calculator's conditions closer to reality.
</p>

## Task 4
<p>
<b>Question 4: </b> We might want to show the client if there's fog, which could cause blurry vision. 
</p>

## Task 5
### Calculator UML
```mermaid
sequenceDiagram
    participant Frontend
    participant Backend
    participant Database

    Frontend ->> Backend: Post request to the backend with the user selected data.

    Backend ->> Frontend: Calculates the metrics and sends it back.

    Backend ->> Database: Saves the Flight Records in Postgresql DB.
```

### Weather Api Implementation UML
```mermaid
sequenceDiagram
    participant Frontend
    participant Backend
    participant RemoteWeatherAPI
    
    Frontend ->> Backend: Post request to the backend with the user selected data.
    Backend ->> RemoteWeatherAPI: Requests the weather data using the client's data
    RemoteWeatherAPI ->> Backend: returns data in json format
    Backend ->> Frontend: returns selected fields from json

```

### Database tables
- Flight - saved in the database. 

`FlightId` is a sequence, auto incremented. 
```mermaid
classDiagram
    class Flight {
        +uint FlightId
        +float64 CargoMass
        +float64 TakeOffDistance
        +float64 ExcessCargoMass
        +float64 TakeOffTime
    }
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