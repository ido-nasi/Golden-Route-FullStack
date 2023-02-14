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
        <h3 className="flightId">Flight ID: {props.id} </h3><br></br>
        <p className="flightSpecs">
          <b>Total Mass:</b> {props.mass} [kg]<br></br>
          <b>Take-off Distance:</b> {props.distance} [m]<br></br> 
          <b>Take-off Time:</b> {props.time} [seconds]<br></br>
          <b>Excess Mass:</b> {props.excessMass} [kg]
        </p>
    </div>
  )
}

