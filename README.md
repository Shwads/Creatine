# Creatine

When you're tired out from ```cURL```ing, creatine will help you recover.

## Accepted File Tags
Example request file.yml:

```
request:
    file: false                         // optional - defaults to true (write output to file)
    console: false                      // optional - defualts to false (print output to console)
    verbose: false                      // optional - defaults to true 

    url: https://examplelink.com
    method: GET

    headers:                            // optional - defaults to empty
        Authorization: 1234
        ItemList:
            - item 1
            - item 2

    body: |                             // optional - defaults to empty. '|' indicates the request parser should preserve newline characters
        {
            "some": "multi line",
            "request": "body content",
        }

request:
    url: https://examplelink.com
    method: GET

    body: >                             // optional - defaults to empty. '>' indicates newline characters will not be preserved
        {
            "some": "multi line",
            "request": "body content",
        }
    
```
