# DataSet Calculator REST API 

Caluculate missing data in dataset by multi-resolver which will be selected by client's request.

# Feature
-   Caluculate missing data in dataset

## Run the app

    go run .

## Run the tests

    go test .

# REST API

The REST API to support above feature

## Calculate Data
- This end point will be receive reslover name and dataset for calculate missing data in dataset.
### Request

`POST /calculate`

- ResolverName - Resolver which use to solve this problem (In the ExampleReslover will use position+50 as result)
- Size - size of dataset include data which will be calculated
- Input - Array of Data in DataSet
- FindPosition - Position of data which will becalculated

#### Request Body
    {
        "ResolverName": "ExampleReslover",
        "Size": 7,
        "Input": [
            "1",
            "2",
            "x",
            "y",
            "3",
            "4",
            "5"
        ],
        "FindPosition": [
            2,
            4
        ]
    }

### Response

Response is depend on reslover which implement DataSetReslover interface.

    {
        "Size": 7,
        "Input": [
            "1",
            "2",
            "x",
            "y",
            "3",
            "4",
            "5"
        ],
        "FindPosition": [
            2,
            4
        ],
        "Result": {
            "2": "52",
            "4": "54"
        }
    }
