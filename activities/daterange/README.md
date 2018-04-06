# Date Range
This activity accepts a start date & end date. You can specify a unit offset for either date to calculate a start/end date range.

## Installation
### Flogo CLI
```bash
flogo install github.com/mellistibco/flogo-activities/activities/daterange
```

## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "format",
      "type": "string",
      "required": "false",
      "value": "01022006150405"
    },
    {
      "name": "startDate",
      "type": "object",
      "required": true
    },
    {
      "name": "endDate",
      "type": "object",
      "required": false
    }
  ],
  "output": [
    {
      "name": "newStartDate",
      "type": "string"
    },
    {
      "name": "newEndtDate",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| format         | False     | The date format. The default value is `01022006150405` |
| startDate      | True      | An object representing the start date in the format of MMDDYYYYHHMMSS, an offset and units         
| endDate        | False     | An object representing the end date in the format of MMDDYYYYHHMMSS, an offset and units. If not supplied, the end date is assumed to be NOW with no offset

The input object should be structured as follows:

```json
{
  "Date":   "04052018205400",
  "Offset": -5,
  "Units":  "minutes"
}
```

* To indicate that the date should be NOW, leave as an empty string or remove the `Date` field from the object.
* If you do not want to alter the date with an offset, specify 0 for the offset or remove the `Offset` field from the object

Valid unites include: **second or seconds, minute or minutes, days, months, years**

## Example
The below example will parse the supplied text.

```json
{
  "id": "daterange_2",
  "name": "Date Range",
  "description": "Calculate date range given a start date with an offset and an end date with an offset.",
  "activity": {
    "ref": "github.com/mellistibco/flogo-activities/activities/daterange",
    "input": {
    },
    "mappings": {
      "input": [
        {
          "type": "object",
          "value": {
            "Date": "04042018215500",
            "Offset": -5,
            "Units": "minutes"
          },
          "mapTo": "startDate"
        },
        {
          "type": "object",
          "value": {
            "Date": "",
            "Offset": 0
          },
          "mapTo": "endDate"
        }
      ]
    }
  }
}
```

The above sample will return a start date of 04042018215500-5 minutes and an end date of now.