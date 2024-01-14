# dating-app

Users

user_id (Primary Key)
username (Unique)
password (Hashed and Salted)
email (Unique)
date_of_birth
gender
orientation
profile_picture
bio
location (can be more specific with city, state, country columns)
created_at
last_login

Preferences

preference_id (Primary Key)
user_id (Foreign Key)
preferred_gender
preferred_age_range_start
preferred_age_range_end
preferred_location_range (e.g., within X miles)

User_Interests

interest_id (Primary Key)
user_id (Foreign Key)
interest

Messages

message_id (Primary Key)
sender_id (Foreign Key referencing Users)
receiver_id (Foreign Key referencing Users)
content
timestamp


Matches

match_id (Primary Key)
user1_id (Foreign Key referencing Users)
user2_id (Foreign Key referencing Users)
matched_on (timestamp)


Photos (for users who upload multiple photos)

photo_id (Primary Key)
user_id (Foreign Key)
photo_url
upload_date


Swipes

swipe_id (Primary Key)
swiper_id (Foreign Key referencing Users)
swiped_user_id (Foreign Key referencing Users)
swipe_type (Like, Dislike)
timestamp


## install go

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
``` 

## start server 

```bash
cd backend
go build .
./main
```


### start server with tls
```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost
```
