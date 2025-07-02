#Gator
##To run gator you need these installed:
- postgres (used to store posts from the feeds)
- go (to build the project)

To install gator you can use:
```
go install github.com/Dass33/gator

```

Config file for gator has to be set up in home directory and named .gatorconfig.json
Example of config file:
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"USER"}
```

##Commands available
- login
- register
    - adds new user
- reset
    - resets the database
- users
    - prints all registered users
- agg
    - starts colleting posts from loged user feeds
- addfeed
    - adds new feed and subscribs loged user to it
- feeds
    - prints all feeds that were added
- follow
- following
    - prints all the feeds user follows
- unfollow
- browse
    - prints loged users stored posts
