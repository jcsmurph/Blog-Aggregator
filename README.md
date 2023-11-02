# Blog Aggregator 

This applications links to your favorite blogs and returns their content.

# Supported Features
- User Creation
- Create new Blog feeds that link to your favorite blogs
- Follow feeds that are created by other users
- Unfollow feeds that no longer interest you

# Supported Feeds
-XML

#Prerequisites
- Postgres <https://www.postgresql.org/download/>
- Postgres Client <https://www.pgadmin.org/download/pgadmin-4-container/>
- Goose for Database Migrations <https://github.com/pressly/goose>

# To Run
- Clone the repository locally
- Create a database in postgres
- On the command line, run goose postgres postgres://postgres:USERNAME@HOST:PORT/DATABASE?sslmode=disable up to create the necessary tables
- Use an API client to create a user and get an API token
- Use the API token to create a feed to your favorite blog
- Use the API to follow other blogs that have been created by other users

# Supported APIs
- Create a User (POST localhost:8080/v1/users)
<br/> 
Body:
```
{
"name": Username
}

```
- Get User Information (GET localhost:8080/v1/users)
  <br/>
Body:
```
{
"name": Username
}

```

- Create a feed (POST localhost:8080/v1/feeds)
  <br/>
Header:
```
Authorization: ApiKey <key>
```

- Follow a created feed (POST localhost:8080/v1/feed_follows)
  <br/>
Header:
```
Authorization: ApiKey <key>
```

- Unfollow a feed (DELETE localhost:8080/v1/feed_follows/{feedFollowID})
<br/>
Header:
```
Authorization: ApiKey <key>
```

- Unfollow a feed (DELETE localhost:8080/v1/feed_follows/{feedFollowID})
<br/>
Header:
```
Authorization: ApiKey <key>
```

- Get posts from a feed (GET localhost:8080/v1/posts?limit=)
  <br/>
Header:
```
Authorization: ApiKey <key>
```


