import React from 'react'

export interface IFlight {
  id: number;
  mass: number;
  distance: number;
  excessMass: number;
  time: number;
}

export const Flight = ({props} : {props: IFlight}) => {

  return (
    <div className="flight" key={props.id}>
        <h3>Flight ID: {props.id} </h3><br></br>
        <p className="flightSpecs">
          Total Mass: {props.mass} <br></br>
          Take-off Distance: {props.distance} <br></br> 
          Take-off Time: {props.time} <br></br>
          Excess Mass: {props.excessMass}
        </p>
    </div>
  )
}

