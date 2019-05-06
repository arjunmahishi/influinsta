# influinsta
A tool that automates your way to becoming an Instagram influencer  

## What it does
For now, it searches for a given set of hashtags, selects the top post based on parameters like number of likes, number of views etc, and reposts the image with credits to the user. While doing this, it likes all the posts it has encountered (sticking to the likeThreshold), follows the user whose post it reposts. 

This current behavior can be modified in the main function. 

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
- Run either manually or schedule it as a cron
    ```bash
    $ ./influinsta -config path-to/config.json
    ```

## In the (near) future
- Make the automation flow configurable through config.json
- Find a way to post videos
- Autonomously comment on posts
- Follow revelant users
- Auto-reply on DM

## Instagram client
This project uses a package called [goinsta](https://github.com/ahmdrz/goinsta) by [@ahmdrz](https://github.com/ahmdrz) for interacting with instagram.
