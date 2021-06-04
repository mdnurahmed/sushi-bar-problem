# sushi-bar-problem
This is a concurrency problem from the book [The Little Book of Semaphores](https://greenteapress.com/wp/semaphores/)

## Problem Description 

Imagine a sushi bar with 5 seats. If you arrive while there is an empty seat, you can take a seat immediately. But if you arrive when all 5 seats are full, that means that all of them are dining together, and you will have to wait for the entire party to leave before you sit down.

## My Approach 

A customer can do these things in a restaurent in serial - 

- Waits to be seated
- Sits and eats
- Leaves

In a resaurent 2 things can happen 

- A customer enters the restaurent 
- A customer leaves the restaurent 

When table becomes full we no longer accepts customers . We do this by making the channel nil. When all the customer leaves restaurent starts accpeting customers.

## Why this won't reach deadlock 

Cause at anytime there is at least the customer generator go routine is running which never blocks . 

## How to run 
using go 
```
go run sushi.go
```
or using docker 
```
docker-compose -f sushi.yaml up --build
docker run -it sushi-bar-problem_sushi /bin/sh
go run sushi.go
exit
docker-compose -f sushi.yaml down
```