## Rainin balls on an island

## Authors
 Diego Armando Gracia Hinojosa
 Eduardo Alonso
 
## USE CASE

The program is just to displa the information so the user can read it and undestand it.
![Untitled Diagram](https://user-images.githubusercontent.com/17838196/70180409-8ed59b80-16a5-11ea-91a6-f8d5299d9af9.png)

## How to use it

1.- **choose the number of balls:** the user can be able to choose the number of balls falling into the island 
the value of the movement is base on the height.

2.- **Observe the movement of the balls:** the user can be able to watch how te balls interact on the map by a map drawn in ascii

3.- **View how the balls get stuck** the user can watch it the balls stop moving in the island

4.- **Watch the velocity of the balls** the user can be able to track the velocity of the balls falling

5.- **view the landing of the balls**  the user can watch where the balls lands and where it ends

6.- **Know if the ball gets to the sea**


## Architecture

The prgoram is just one clas that you run on Go.

You dont need anithyng else to run it.

## Multithreading

we have a function in charge of starting the multithreading that creating a goroutines 
and we have a variable that works letting the progaram know if it its over.




## process

![Untitled Diagram (1)](https://user-images.githubusercontent.com/17838196/70181646-ed9c1480-16a7-11ea-9552-45d6c44408c9.png)

**Size of the map** you can choose the size of the map in the terminal

** create the ball multithread:** each ball should have a individual thread

**create a falling ball** They fall randombly and by the geight you choose

*balls display** display teh balls in each area








