# Assignment 

## Question 1
Design API to find value x,y,z in data set [1, X, 8, 17, Y, Z, 78, 113] ***I mock X, Y, and Z to the dataset -> [1, 2, 8, 17, 22, 37, 78, 113]***

***
## RESTful API

### API documentation

#### GET */get-xyz-by-position*
Description:
Get X, Y and Z by position

Query: -

Response:

```json
{
  "X": 2,
  "Y": 22,
  "Z": 37
}
```
___

#### GET */get-xyz-by-remove-know-data*
Description:
Get X, Y and Z by remove know data

Query: -

Response:

```json
{
  "X": 2,
  "Y": 22,
  "Z": 37
}
```
***
## Building
    $ go build assignment/question_1