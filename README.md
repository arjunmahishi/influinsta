# influinsta
A tool that automates your way to becoming an Instagram influencer  

## Instagram client
This project uses a package called `goinsta` by @ahmdrz for interacting with instagram.

## How to use

- Build the binary from source
    ```bash
    $ go get -d
    $ go build -o influinsta
    ```
- Create a config file with following data
    ```json
    {
        "creds": {
            "username": "username",
            "password": "password"
        },
        "hashtags": ["hashtags", "you", "want", "to", "follow"],
        "likeThreshold": 300,
        "followThreshold": 300
    }
    ```
- Run either manually or create a cron to schedule it
    ```bash
    $ ./influinsta -config path-to/config.json
    ```
