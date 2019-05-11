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
        "followThreshold": 300,
        "actions": [
            {
                "name": "action-name",
                "args": ["arg1", "arg2"]
            }
        ]
    }
    ```
    See [list of available actions](##List-of-available-actions).

- Run either manually or schedule it as a cron
    ```bash
    $ ./influinsta -config path-to/config.json
    ```

## List of available actions
- `reshare-video` - Looks for videos for the given `hashtags`, selects the best one, and reposts it 
- `random-comments` [WIP] - Looks for videos/images for the given `hashtags` and randomly comments on the posts
- `random-follow` [WIP] - Looks for videos/images for the given `hashtags` and randomly follows the followers of that `hashtag`

## In the (near) future
- Make the automation flow configurable through config.json
- Find a way to post videos
- Autonomously comment on posts
- Follow revelant users
- Auto-reply on DM

## Instagram client
This project uses a package called [goinsta](https://github.com/ahmdrz/goinsta) by [@ahmdrz](https://github.com/ahmdrz) for interacting with instagram.

## Legal
This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial tool. Use at your own risk.