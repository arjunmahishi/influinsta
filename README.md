# influinsta [![Build Status](https://travis-ci.com/arjunmahishi/influinsta.svg?branch=master)](https://travis-ci.com/arjunmahishi/influinsta)
A tool that automates your way to becoming an Instagram influencer  

## What it does
It automates certain `actions` on Instagram. An `action` could be anything tedious like following a large number of people, commenting on random posts for a given hashtag, repost images/videos etc. These actions are being done manually by people (which have proven to increase their follower count). But these tasks are tedious and should not be done by human beings. So, this tool helps automate those actions and save a lot of time while making it more efficient.

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
- `reshare-video` - Looks for videos for the given `hashtags`, selects the best one and reposts it 
- `reshare-image` - Looks for images for the given `hashtags`, selects the best one and reposts it
- `random-comments` - Looks for videos/images for the given `hashtags` and randomly comments on the posts
- `random-follow` - Looks for videos/images for the given `hashtags` and randomly follows the followers of that `hashtag`

## In the (near) future
- Find a way to post videos
- Auto-reply on DM

## Instagram client
This project uses a package called [goinsta](https://github.com/ahmdrz/goinsta/v2) by [@ahmdrz](https://github.com/ahmdrz) for interacting with instagram.

## Legal
This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial tool. Use at your own risk.