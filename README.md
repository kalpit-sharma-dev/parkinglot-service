# parkinglot-service
ParkingLot Service

## Problem Statement

I own a parking lot that can hold up to 'n' cars at any given point in time. Each slot is
given a number starting at 1 increasing with increasing distance from the entry point
in steps of one. I want to create an automated ticketing system that allows my
customers to use my parking lot without human intervention.  

When a car enters my parking lot, I want to have a ticket issued to the driver. The
ticket issuing process includes us documenting the registration number (number
plate) and the colour of the car and allocating an available parking slot to the car
before actually handing over a ticket to the driver (we assume that our customers are
nice enough to always park in the slots allocated to them). The customer should be
allocated a parking slot which is nearest to the entry. At the exit the customer returns
the ticket which then marks the slot they were using as being available.  

Due to government regulation, the system should provide me with the ability to find out:  
* Registration numbers of all cars of a particular colour.
* Slot number in which a car with a given registration number is parked.
* Slot numbers of all slots where a car of a particular colour is parked.

We interact with the system over the API or any other approach which produce a specific
output. Please take a look at the example below, which includes all the commands
you need to support - they're self explanatory.

#### Events ():
create_parking_lot 6
park KA-01-AB-1234 White  
park KA-01-CD-9999 White  
park KA-01-EF-0001 Black  
park KA-01-GH-7777 Red  
park KA-01-IJ-2701 Blue  
park KA-01-KL-3141 Black  
leave 4  
status  
park KA-01-M-333 White  
park KL-12-NN-9999 White  
registration_numbers_for_cars_with_colour White  
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-KL-3141  
slot_number_for_registration_number MH-04-AY-1111
todays_earning
monthly_earning
find_average_car_parking_time
find_peak_hours

