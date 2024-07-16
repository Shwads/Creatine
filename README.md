# Creatine

When you're tired our from ```cURL```ing, creatine will help you recover.

## Accepted File Tags
Example request file:

```
request:
    file: false         // optional - defaults to true
    console: false      // optional - defualts to false
    verbose: false      // optional - defaults to true

    url: https://examplelink.com
    method: GET

    headers:
        Authorization: 1234
        ItemList:
            - item 1
            - item 2

    body: |
        {
            "some": "multi line",
            "request": "body content",
        }

request:
    ...
    
```
