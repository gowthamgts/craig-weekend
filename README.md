# Daniel Craig's "Ladies and gentlemen, the weekend" mastodon bot

This bot is based on the original twitter handle ["ladies and gentlemen, the weekend 😌"](https://twitter.com/CraigWeekend) which will publish Daniel Craig video every weekend. I missed seeing this on mastodon, hence spent some time to build this. You can follow [`@CraigWeekend@techhub.social`](https://techhub.social/@CraigWeekend) mastodon account to get the toot.

## How it works

Nothing fancy, just a cron job that runs at 12:00 AM UTC on every saturday.

## Installation

```sh
# clone the repo
git clone https://github.com/gowthamgts/craig-weekend

# change directory
cd craig-weekend

# build the binary
go build main.go

# test the binary - this will throw an error stating you are missing some environment variables
./craig-weekend

# add the following entry to your crontab by issuing `crontab -e`. don't forget to update the environment variables
0 0 * * SAT MASTODON_SERVER="" MASTODON_CLIENT_ID="" MASTODON_CLIENT_SECRET="" MASTODON_USERNAME="" MASTODON_PASSWORD="" <build directory>/craig-weekend
```

## Environment variables

The following environment variables are needed for running the bot.

| Environment Variable Name | Description                                          |
| ------------------------- | ---------------------------------------------------- |
| MASTODON_SERVER           | the mastodon server url. ex: https://mastodon.social |
| MASTODON_CLIENT_ID        | your application's client id                         |
| MASTODON_CLIENT_SECRET    | your application's client secret                     |
| MASTODON_USERNAME         | your mastodon account username                       |
| MASTODON_PASSWORD         | your mastodon account password                       |

## License

GNU GPL v3
